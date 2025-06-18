package transferring

import (
	"context"
	"sync"
	"time"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/cron/scheduler/base"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/transferring/executor"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/transferring/notif"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/transferring/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/transferring/sentinel"
)

const subsystem = "benefitpowerrentaltransferring"

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
