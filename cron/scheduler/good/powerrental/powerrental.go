package powerrental

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/checkhashrate"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/creategooduser"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good/powerrental/wait"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "goodpowerrental"

func Initialize(ctx context.Context, cancel context.CancelFunc, running *sync.Map) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)
	wait.Initialize(ctx, cancel, running)
	creategooduser.Initialize(ctx, cancel, running)
	checkhashrate.Initialize(ctx, cancel, running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	wait.Finalize(ctx)
	creategooduser.Finalize(ctx)
	checkhashrate.Finalize(ctx)
}
