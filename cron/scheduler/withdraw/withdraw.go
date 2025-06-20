package withdraw

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/approved"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/created"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/fail/prefail"
	failreturnbalance "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/fail/returnbalance"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/rejected/prerejected"
	rejectedreturnbalance "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/rejected/returnbalance"
	withdrawreviewnotify "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/review/notify"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/reviewing"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/successful/presuccessful"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/successful/spendbalance"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/transferring"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "withdraw"

var running sync.Map

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)

	transferring.Initialize(ctx, cancel, &running)
	created.Initialize(ctx, cancel, &running)
	approved.Initialize(ctx, cancel, &running)
	reviewing.Initialize(ctx, cancel, &running)
	prefail.Initialize(ctx, cancel, &running)
	failreturnbalance.Initialize(ctx, cancel, &running)
	rejectedreturnbalance.Initialize(ctx, cancel, &running)
	prerejected.Initialize(ctx, cancel, &running)
	presuccessful.Initialize(ctx, cancel, &running)
	spendbalance.Initialize(ctx, cancel, &running)
	withdrawreviewnotify.Initialize(ctx, cancel)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	withdrawreviewnotify.Finalize(ctx)
	spendbalance.Finalize(ctx)
	presuccessful.Finalize(ctx)
	prerejected.Finalize(ctx)
	rejectedreturnbalance.Finalize(ctx)
	failreturnbalance.Finalize(ctx)
	prefail.Finalize(ctx)
	reviewing.Finalize(ctx)
	approved.Finalize(ctx)
	created.Finalize(ctx)
	transferring.Finalize(ctx)
}
