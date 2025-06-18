package txqueue

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/created"
	"github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/transferring"
	"github.com/NpoolPlatform/kunman/cron/scheduler/txqueue/wait"
)

const subsystem = "txqueue"

var running sync.Map

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)

	created.Initialize(ctx, cancel, &running)
	wait.Initialize(ctx, cancel, &running)
	transferring.Initialize(ctx, cancel, &running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	transferring.Finalize(ctx)
	wait.Finalize(ctx)
	created.Finalize(ctx)
}
