package reward

import (
	"context"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	"github.com/NpoolPlatform/kunman/framework/action"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/watcher"
)

type Reward interface {
	Feed(context.Context, interface{})
	Finalize(context.Context)
}

type Rewarder interface {
	Update(context.Context, interface{}, chan interface{}, chan interface{}) error
}

type handler struct {
	feeder    chan interface{}
	notif     chan interface{}
	done      chan interface{}
	w         *watcher.Watcher
	rewarder  Rewarder
	subsystem string
}

func NewReward(ctx context.Context, cancel context.CancelFunc, notif, done chan interface{}, rewarder Rewarder, subsystem string) Reward {
	p := &handler{
		feeder:    make(chan interface{}),
		notif:     notif,
		done:      done,
		w:         watcher.NewWatcher(),
		rewarder:  rewarder,
		subsystem: subsystem,
	}
	go action.Watch(ctx, cancel, p.run, p.paniced)
	return p
}

func (p *handler) handler(ctx context.Context) bool {
	select {
	case ent := <-p.feeder:
		if err := p.rewarder.Update(ctx, ent, p.notif, p.done); err != nil {
			logger.Sugar().Errorw(
				"handler",
				"State", "Update",
				"Subsystem", p.subsystem,
				"Ent", ent,
				"Error", err,
			)
		}
		return false
	case <-p.w.CloseChan():
		close(p.w.ClosedChan())
		return true
	}
}

func (p *handler) run(ctx context.Context) {
	for {
		if b := p.handler(ctx); b {
			break
		}
	}
}

func (p *handler) paniced(ctx context.Context) { //nolint
	close(p.w.ClosedChan())
}

func (p *handler) Finalize(ctx context.Context) {
	if p.w != nil {
		p.w.Shutdown(ctx)
	}
}

func (p *handler) Feed(ctx context.Context, ent interface{}) {
	cancelablefeed.CancelableFeed(ctx, ent, p.feeder)
}
