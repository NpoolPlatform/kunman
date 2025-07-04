package base

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/executor"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base/notif"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base/retry"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base/reward"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/framework/action"
	"github.com/NpoolPlatform/kunman/framework/logger"
	redis2 "github.com/NpoolPlatform/kunman/framework/redis"
	"github.com/NpoolPlatform/kunman/framework/watcher"
)

type idDesc struct {
	id        string
	subsystem string
	start     time.Time
	value     interface{}
}

type syncMap struct {
	*sync.Map
	count      int
	concurrent int
	subsystem  string
}

func (s *syncMap) Store(key, value interface{}) (bool, bool) {
	desc := &idDesc{
		id:        key.(string),
		subsystem: s.subsystem,
		start:     time.Now(),
		value:     value,
	}
	_desc, loaded := s.LoadOrStore(key, desc)
	if !loaded {
		if s.count >= s.concurrent && s.concurrent < math.MaxInt {
			s.Map.Delete(key)
			return false, true
		}
		s.count++
		return false, false
	}
	if time.Now().After(_desc.(*idDesc).start.Add(1 * time.Minute)) {
		desc.start = time.Now()
		s.Map.Store(key, desc)
		logger.Sugar().Warnw(
			"Store",
			"Ent", value,
			"Key", key,
			"StoreEnt", _desc.(*idDesc).value,
			"ID", _desc.(*idDesc).id,
			"StoreSubsystem", _desc.(*idDesc).subsystem,
			"Start", _desc.(*idDesc).start,
			"Subsystem", s.subsystem,
			"Count", s.count,
			"State", "Processing",
		)
	}
	return true, false
}

func (s *syncMap) Delete(key interface{}) {
	desc, ok := s.LoadAndDelete(key)
	if !ok {
		return
	}
	if time.Since(desc.(*idDesc).start).Seconds() > 10 {
		logger.Sugar().Warnw(
			"Delete",
			"ID", desc.(*idDesc).id,
			"StoreSubsystem", desc.(*idDesc).subsystem,
			"Start", desc.(*idDesc).start,
			"Elapsed", time.Since(desc.(*idDesc).start),
		)
	}
	s.count--
}

type Handler struct {
	persistent        chan interface{}
	reward            chan interface{}
	notif             chan interface{}
	done              chan interface{}
	w                 *watcher.Watcher
	sentinel          sentinel.Sentinel
	scanner           sentinel.Scanner
	executors         []executor.Executor
	execer            executor.Exec
	executorNumber    int
	executorIndex     int
	persistenter      persistent.Persistent
	persistentor      persistent.Persistenter
	rewarder          reward.Reward
	rewardor          reward.Rewarder
	notifier          notif.Notif
	notify            notif.Notify
	subsystem         string
	scanInterval      time.Duration
	running           *syncMap
	runningConcurrent int
	locked            bool
	cancel            context.CancelFunc
}

func (h *Handler) lockKey() string {
	return fmt.Sprintf("%v:%v", "scheduler", h.subsystem)
}

func NewHandler(ctx context.Context, cancel context.CancelFunc, options ...func(*Handler)) (*Handler, error) {
	h := &Handler{
		executorNumber: 1,
	}
	for _, opt := range options {
		opt(h)
	}
	if b := config.SupportSubsystem(h.subsystem); !b {
		return nil, nil
	}
	if h.running == nil {
		return nil, fmt.Errorf("invalid running map")
	}
	if h.runningConcurrent > 0 {
		h.running.concurrent = h.runningConcurrent
	}

	h.persistent = make(chan interface{})
	h.reward = make(chan interface{})
	h.notif = make(chan interface{})
	h.done = make(chan interface{})
	ctx, h.cancel = context.WithCancel(ctx)

	h.sentinel = sentinel.NewSentinel(ctx, cancel, h.scanner, h.scanInterval, h.subsystem)
	for i := 0; i < h.executorNumber; i++ {
		h.executors = append(h.executors, executor.NewExecutor(ctx, cancel, h.persistent, h.notif, h.done, h.execer, h.subsystem))
	}
	h.persistenter = persistent.NewPersistent(ctx, cancel, h.reward, h.notif, h.done, h.persistentor, h.subsystem)
	h.notifier = notif.NewNotif(ctx, cancel, h.notify, h.subsystem)
	h.rewarder = reward.NewReward(ctx, cancel, h.notif, h.done, h.rewardor, h.subsystem)

	h.w = watcher.NewWatcher()
	if err := redis2.TryLock(h.lockKey(), 0); err == nil {
		h.locked = true
	}

	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", h.subsystem,
		"Locked", h.locked,
	)
	return h, nil
}

func WithSubsystem(subsystem string) func(*Handler) {
	return func(h *Handler) {
		h.subsystem = subsystem
	}
}

func WithScanner(scanner sentinel.Scanner) func(*Handler) {
	return func(h *Handler) {
		h.scanner = scanner
	}
}

func WithScanInterval(scanInterval time.Duration) func(*Handler) {
	return func(h *Handler) {
		h.scanInterval = scanInterval
	}
}

func WithExec(exec executor.Exec) func(*Handler) {
	return func(h *Handler) {
		h.execer = exec
	}
}

func WithExecutorNumber(n int) func(*Handler) {
	return func(h *Handler) {
		h.executorNumber = n
	}
}

func WithPersistenter(persistenter persistent.Persistenter) func(*Handler) {
	return func(h *Handler) {
		h.persistentor = persistenter
	}
}

func WithRewarder(rewarder reward.Rewarder) func(*Handler) {
	return func(h *Handler) {
		h.rewardor = rewarder
	}
}

func WithNotify(notify notif.Notify) func(*Handler) {
	return func(h *Handler) {
		h.notify = notify
	}
}

func WithRunningMap(m *sync.Map) func(*Handler) {
	return func(h *Handler) {
		h.running = &syncMap{
			Map:        m,
			concurrent: 3,
			subsystem:  h.subsystem,
		}
	}
}

func WithRunningConcurrent(concurrent int) func(*Handler) {
	return func(h *Handler) {
		h.runningConcurrent = concurrent
	}
}

func (h *Handler) Run(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(h.subsystem); !b {
		return
	}
	action.Watch(ctx, cancel, h.run, h.paniced)
}

func (h *Handler) execEnt(ctx context.Context, ent interface{}) {
	h.executors[h.executorIndex].Feed(ctx, ent)
	h.executorIndex++
	h.executorIndex %= len(h.executors)
}

func (h *Handler) handler(ctx context.Context) bool {
	select {
	case ent := <-h.sentinel.Exec():
		if loaded, overflow := h.running.Store(h.scanner.ObjectID(ent), ent); loaded || overflow {
			if overflow {
				// Here is a bit strange, but let's use sentinel exec firstly
				retry.Retry(h.scanner.ObjectID(ent), ent, h.sentinel.Exec())
			}
			return false
		}
		h.execEnt(ctx, ent)
		return false
	case ent := <-h.persistent:
		h.persistenter.Feed(ctx, ent)
		return false
	case ent := <-h.reward:
		h.rewarder.Feed(ctx, ent)
		return false
	case ent := <-h.notif:
		h.notifier.Feed(ctx, ent)
		return false
	case ent := <-h.done:
		h.running.Delete(h.scanner.ObjectID(ent))
		return false
	case <-h.w.CloseChan():
		logger.Sugar().Infow(
			"handler",
			"State", "Close",
			"Subsystem", h.subsystem,
			"Error", ctx.Err(),
		)
		close(h.w.ClosedChan())
		return true
	}
}

func (h *Handler) retryLock(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Minute):
			logger.Sugar().Infow(
				"RetryLock",
				"Subsystem", h.subsystem,
			)
			if err := redis2.TryLock(h.lockKey(), 0); err == nil {
				h.locked = true
				return
			}
		}
	}
}

func (h *Handler) run(ctx context.Context) {
	if !h.locked {
		h.retryLock(ctx)
	}
	if !h.locked {
		close(h.w.ClosedChan())
		return
	}
	for {
		if b := h.handler(ctx); b {
			break
		}
	}
}

func (h *Handler) paniced(ctx context.Context) {
	logger.Sugar().Errorw(
		"Paniced",
		"Subsystem", h.subsystem,
	)
	close(h.w.ClosedChan())
}

func (h *Handler) Trigger(cond interface{}) {
	h.sentinel.Trigger(cond)
}

func (h *Handler) Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(h.subsystem); !b {
		return
	}
	if h.locked {
		_ = redis2.Unlock(h.lockKey())
	}
	h.cancel()
	h.sentinel.Finalize(ctx)
	if h.w != nil {
		h.w.Shutdown(ctx)
	}
	for _, e := range h.executors {
		e.Finalize(ctx)
	}
	h.persistenter.Finalize(ctx)
	h.rewarder.Finalize(ctx)
	h.notifier.Finalize(ctx)
	logger.Sugar().Infow(
		"Finalize",
		"Subsystem", h.subsystem,
	)
}
