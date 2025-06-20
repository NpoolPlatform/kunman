package obselete

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/transfer/bookkeeping"
	"github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/transfer/unlockaccount"
	"github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/unlockbalance"
	"github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/wait"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "paymentobselete"

var running sync.Map

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)

	wait.Initialize(ctx, cancel, &running)
	unlockbalance.Initialize(ctx, cancel, &running)
	bookkeeping.Initialize(ctx, cancel, &running)
	unlockaccount.Initialize(ctx, cancel, &running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	unlockaccount.Finalize(ctx)
	bookkeeping.Finalize(ctx)
	unlockbalance.Finalize(ctx)
	wait.Finalize(ctx)
}
