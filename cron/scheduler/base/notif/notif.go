package notif

import (
	"context"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	"github.com/NpoolPlatform/kunman/framework/action"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/watcher"
)

type Notif interface {
	Feed(context.Context, interface{})
	Finalize(context.Context)
}

type Notify interface {
	Notify(context.Context, interface{}, chan interface{}) error
}

type handler struct {
	feeder    chan interface{}
	w         *watcher.Watcher
	notify    Notify
	subsystem string
}

func NewNotif(ctx context.Context, cancel context.CancelFunc, notify Notify, subsystem string) Notif {
	p := &handler{
		feeder:    make(chan interface{}),
		w:         watcher.NewWatcher(),
		notify:    notify,
		subsystem: subsystem,
	}
	go action.Watch(ctx, cancel, p.run, p.paniced)
	return p
}

func (p *handler) handler(ctx context.Context) bool {
	select {
	case ent := <-p.feeder:
		if p.notify == nil {
			return false
		}
		if err := p.notify.Notify(ctx, ent, p.feeder); err != nil {
			logger.Sugar().Infow(
				"handler",
				"State", "Notify",
				"Subsystem", p.subsystem,
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
