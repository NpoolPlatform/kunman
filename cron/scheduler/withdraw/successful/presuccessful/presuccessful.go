package presuccessful

import (
	"context"
	"sync"
	"time"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/successful/presuccessful/executor"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/successful/presuccessful/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/successful/presuccessful/reward"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/successful/presuccessful/sentinel"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "withdrawsuccessfulpresuccessful"

var h *base.Handler

func Initialize(ctx context.Context, cancel context.CancelFunc, running *sync.Map) {
	_h, err := base.NewHandler(
		ctx,
		cancel,
		base.WithSubsystem(subsystem),
		base.WithScanInterval(time.Minute),
		base.WithScanner(sentinel.NewSentinel()),
		base.WithExec(executor.NewExecutor()),
		base.WithExecutorNumber(1),
		base.WithPersistenter(persistent.NewPersistent()),
		base.WithRewarder(reward.NewReward()),
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
