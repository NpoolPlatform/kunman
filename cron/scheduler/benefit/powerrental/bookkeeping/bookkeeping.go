package bookkeeping

import (
	"context"
	"sync"

	goodbookkeeping "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/good"
	simulatebookkeeping "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/simulate"
	userbookkeeping "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/user"
	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "benefitpowerrentalbookkeeping"

func Initialize(ctx context.Context, cancel context.CancelFunc, running *sync.Map) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)

	goodbookkeeping.Initialize(ctx, cancel, running)
	userbookkeeping.Initialize(ctx, cancel, running)
	simulatebookkeeping.Initialize(ctx, cancel, running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	userbookkeeping.Finalize(ctx)
	goodbookkeeping.Finalize(ctx)
	simulatebookkeeping.Finalize(ctx)
}
