package good

import (
	"context"
	"sync"
	"time"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/good/executor"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/good/notif"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/good/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/good/sentinel"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "benefitpowerrentalbookkeepinggood"

var h *base.Handler

func Initialize(ctx context.Context, cancel context.CancelFunc, running *sync.Map) {
	_h, err := base.NewHandler(
		ctx,
		cancel,
		base.WithSubsystem(subsystem),
		base.WithScanInterval(1*time.Minute),
		base.WithScanner(sentinel.NewSentinel()),
		base.WithExec(executor.NewExecutor()),
		base.WithNotify(notif.NewNotif()),
		base.WithPersistenter(persistent.NewPersistent()),
		base.WithRunningMap(running),
	)
	if err != nil || _h == nil {
		logger.Sugar().Errorw(
			"Initialize",
			"Subsystem", subsystem,
			"Error", err,
		)
		return
	}

	h = _h
	go h.Run(ctx, cancel)
}

func Finalize(ctx context.Context) {
	if h != nil {
		h.Finalize(ctx)
	}
}
