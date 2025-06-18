package collector

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/cron/scheduler/payment/collector/finish"
	"github.com/NpoolPlatform/kunman/cron/scheduler/payment/collector/transfer"
)

const subsystem = "paymentcollector"

var running sync.Map

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)

	finish.Initialize(ctx, cancel, &running)
	transfer.Initialize(ctx, cancel, &running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	transfer.Finalize(ctx)
	finish.Finalize(ctx)
}
