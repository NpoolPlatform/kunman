package powerrental

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/done"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/fail"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/transferring"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/wait"
	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
)

const subsystem = "benefitpowerrental"

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
	fail.Initialize(ctx, cancel, &running)
	done.Initialize(ctx, cancel, &running)
	transferring.Initialize(ctx, cancel, &running)
	bookkeeping.Initialize(ctx, cancel, &running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	bookkeeping.Finalize(ctx)
	transferring.Finalize(ctx)
	done.Finalize(ctx)
	fail.Finalize(ctx)
	wait.Finalize(ctx)
}
