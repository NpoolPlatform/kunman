package good

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "good"

var running sync.Map

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)

	powerrental.Initialize(ctx, cancel, &running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	powerrental.Finalize(ctx)
}
