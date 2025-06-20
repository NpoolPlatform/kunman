package order

import (
	"context"

	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	fee "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee"
	powerrental "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental"
	powerrentalsimulate "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/simulate"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "order"

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)
	powerrental.Initialize(ctx, cancel)
	fee.Initialize(ctx, cancel)
	powerrentalsimulate.Initialize(ctx, cancel)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	powerrentalsimulate.Finalize(ctx)
	fee.Finalize(ctx)
	powerrental.Finalize(ctx)
}
