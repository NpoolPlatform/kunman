package fee

import (
	"context"
	"sync"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/cron/scheduler/config"
	cancelachievement "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/achievement"
	cancelbookkeeping "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/bookkeeping"
	cancelcheck "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/check"
	cancelcommission "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/commission"
	"github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/precancel"
	cancelrestorestock "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/restorestock"
	"github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/returnbalance"
	cancelunlockaccount "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/unlockaccount"
	"github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/created"
	paidcheck "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/paid/check"
	paidstock "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/paid/stock"
	paymentachievement "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/achievement"
	"github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/bookkeeping"
	paymentcommission "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/commission"
	"github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/received"
	"github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/spend"
	paymentstock "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/stock"
	"github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/timeout"
	paymentunlockaccount "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/unlockaccount"
	paymentwait "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/wait"
)

const subsystem = "orderfee"

var running sync.Map

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	logger.Sugar().Infow(
		"Initialize",
		"Subsystem", subsystem,
	)

	paymentunlockaccount.Initialize(ctx, cancel, &running)
	paymentachievement.Initialize(ctx, cancel, &running)
	bookkeeping.Initialize(ctx, cancel, &running)
	paymentwait.Initialize(ctx, cancel, &running)
	paymentcommission.Initialize(ctx, cancel, &running)
	cancelcommission.Initialize(ctx, cancel, &running)
	received.Initialize(ctx, cancel, &running)
	spend.Initialize(ctx, cancel, &running)
	paidstock.Initialize(ctx, cancel, &running)
	paymentstock.Initialize(ctx, cancel, &running)
	timeout.Initialize(ctx, cancel, &running)
	paidcheck.Initialize(ctx, cancel, &running)
	precancel.Initialize(ctx, cancel, &running)
	cancelbookkeeping.Initialize(ctx, cancel, &running)
	cancelunlockaccount.Initialize(ctx, cancel, &running)
	cancelcheck.Initialize(ctx, cancel, &running)
	cancelachievement.Initialize(ctx, cancel, &running)
	cancelrestorestock.Initialize(ctx, cancel, &running)
	returnbalance.Initialize(ctx, cancel, &running)
	created.Initialize(ctx, cancel, &running)
}

func Finalize(ctx context.Context) {
	if b := config.SupportSubsystem(subsystem); !b {
		return
	}
	created.Finalize(ctx)
	returnbalance.Finalize(ctx)
	cancelrestorestock.Finalize(ctx)
	cancelachievement.Finalize(ctx)
	cancelcheck.Finalize(ctx)
	cancelunlockaccount.Finalize(ctx)
	cancelbookkeeping.Finalize(ctx)
	precancel.Finalize(ctx)
	paidcheck.Finalize(ctx)
	timeout.Finalize(ctx)
	paymentstock.Finalize(ctx)
	paidstock.Finalize(ctx)
	spend.Finalize(ctx)
	received.Finalize(ctx)
	cancelcommission.Finalize(ctx)
	paymentcommission.Finalize(ctx)
	paymentwait.Finalize(ctx)
	bookkeeping.Finalize(ctx)
	paymentachievement.Finalize(ctx)
	paymentunlockaccount.Finalize(ctx)
}
