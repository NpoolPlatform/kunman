package user

import (
	"context"
	"sync"
	"time"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base"
	"github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user/executor"
	"github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user/notif"
	"github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user/persistent"
	"github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user/reward"
	"github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user/sentinel"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "deposituser"

var h *base.Handler

func Initialize(ctx context.Context, cancel context.CancelFunc, running *sync.Map) {
	_h, err := base.NewHandler(
		ctx,
		cancel,
		base.WithSubsystem(subsystem),
		base.WithScanInterval(10*time.Minute),
		base.WithScanner(sentinel.NewSentinel()),
		base.WithNotify(notif.NewNotif()),
		base.WithExec(executor.NewExecutor()),
		base.WithExecutorNumber(20),
		base.WithRunningConcurrent(20),
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
