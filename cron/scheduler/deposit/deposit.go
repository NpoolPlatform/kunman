package deposit

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/cron/scheduler/deposit/finish"
	"github.com/NpoolPlatform/kunman/cron/scheduler/deposit/transfer"
	"github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "deposit"

var running sync.Map

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)

	user.Initialize(ctx, cancel, &running)
	finish.Initialize(ctx, cancel, &running)
	transfer.Initialize(ctx, cancel, &running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	transfer.Finalize(ctx)
	finish.Finalize(ctx)
	user.Finalize(ctx)
}
