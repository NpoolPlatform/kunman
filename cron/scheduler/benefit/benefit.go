package benefit

import (
	"context"

	powerrental "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental"
	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "benefit"

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)
	powerrental.Initialize(ctx, cancel)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	powerrental.Finalize(ctx)
}
