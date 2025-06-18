package couponwithdraw

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/cron/scheduler/couponwithdraw/approved"
	"github.com/NpoolPlatform/kunman/cron/scheduler/couponwithdraw/reviewing"
)

const subsystem = "couponwithdraw"

var running sync.Map

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)

	approved.Initialize(ctx, cancel, &running)
	reviewing.Initialize(ctx, cancel, &running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	reviewing.Finalize(ctx)
	approved.Finalize(ctx)
}
