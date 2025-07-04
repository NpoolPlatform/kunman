package benefit

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base"
	"github.com/NpoolPlatform/kunman/cron/scheduler/notif/benefit/powerrental/executor"
	"github.com/NpoolPlatform/kunman/cron/scheduler/notif/benefit/powerrental/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/notif/benefit/powerrental/sentinel"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "notifbenefitpowerrental"

var (
	h       *base.Handler
	running sync.Map
)

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	_h, err := base.NewHandler(
		ctx,
		cancel,
		base.WithSubsystem(subsystem),
		base.WithScanInterval(1*time.Minute),
		base.WithScanner(sentinel.NewSentinel()),
		base.WithExec(executor.NewExecutor()),
		base.WithRunningConcurrent(math.MaxInt),
		base.WithPersistenter(persistent.NewPersistent()),
		base.WithRunningMap(&running),
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
