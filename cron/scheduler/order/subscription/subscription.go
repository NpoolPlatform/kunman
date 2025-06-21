package subscription

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	"github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/created"
	paymentbookkeeping "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/bookkeeping"
	paymentreceived "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/received"
	paymentspend "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/spend"
	paymentwait "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/wait"
	"github.com/NpoolPlatform/kunman/framework/logger"
)

const subsystem = "ordersubscription"

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
	paymentwait.Initialize(ctx, cancel, &running)
	paymentspend.Initialize(ctx, cancel, &running)
	paymentbookkeeping.Initialize(ctx, cancel, &running)
	paymentreceived.Initialize(ctx, cancel, &running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	paymentreceived.Finalize(ctx)
	paymentbookkeeping.Finalize(ctx)
	paymentspend.Finalize(ctx)
	paymentwait.Finalize(ctx)
	created.Finalize(ctx)
}
