package scheduler

import (
	"context"

	"github.com/NpoolPlatform/kunman/cron/scheduler/base/retry"
	"github.com/NpoolPlatform/kunman/cron/scheduler/benefit"
	"github.com/NpoolPlatform/kunman/cron/scheduler/couponwithdraw"
	"github.com/NpoolPlatform/kunman/cron/scheduler/deposit"
	"github.com/NpoolPlatform/kunman/cron/scheduler/gasfeeder"
	"github.com/NpoolPlatform/kunman/cron/scheduler/good"
	"github.com/NpoolPlatform/kunman/cron/scheduler/limitation"
	"github.com/NpoolPlatform/kunman/cron/scheduler/notif/announcement"
	notifbenefit "github.com/NpoolPlatform/kunman/cron/scheduler/notif/benefit"
	"github.com/NpoolPlatform/kunman/cron/scheduler/notif/notification"
	"github.com/NpoolPlatform/kunman/cron/scheduler/order"
	paymentcollector "github.com/NpoolPlatform/kunman/cron/scheduler/payment/collector"
	paymentobselete "github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete"
	"github.com/NpoolPlatform/kunman/cron/scheduler/txqueue"
	"github.com/NpoolPlatform/kunman/cron/scheduler/withdraw"
)

func Finalize(ctx context.Context) {
	notifbenefit.Finalize(ctx)
	benefit.Finalize(ctx)
	deposit.Finalize(ctx)
	withdraw.Finalize(ctx)
	couponwithdraw.Finalize(ctx)
	limitation.Finalize(ctx)
	txqueue.Finalize(ctx)
	announcement.Finalize(ctx)
	notification.Finalize(ctx)
	gasfeeder.Finalize(ctx)
	order.Finalize(ctx)
	paymentcollector.Finalize(ctx)
	paymentobselete.Finalize(ctx)
	good.Finalize(ctx)
}

func Initialize(ctx context.Context, cancel context.CancelFunc) {
	paymentobselete.Initialize(ctx, cancel)
	paymentcollector.Initialize(ctx, cancel)
	order.Initialize(ctx, cancel)
	gasfeeder.Initialize(ctx, cancel)
	announcement.Initialize(ctx, cancel)
	notification.Initialize(ctx, cancel)
	txqueue.Initialize(ctx, cancel)
	limitation.Initialize(ctx, cancel)
	withdraw.Initialize(ctx, cancel)
	couponwithdraw.Initialize(ctx, cancel)
	deposit.Initialize(ctx, cancel)
	benefit.Initialize(ctx, cancel)
	notifbenefit.Initialize(ctx, cancel)
	retry.Initialize(ctx)
	good.Initialize(ctx, cancel)
}

type initializer struct {
	init  func(context.Context, context.CancelFunc)
	final func(context.Context)
}

var subsystems = map[string]initializer{
	"paymentobselete":  {paymentobselete.Initialize, paymentobselete.Finalize},
	"paymentcollector": {paymentcollector.Initialize, paymentcollector.Finalize},
	"order":            {order.Initialize, order.Finalize},
	"gasfeeder":        {gasfeeder.Initialize, gasfeeder.Finalize},
	"announcement":     {announcement.Initialize, announcement.Finalize},
	"notification":     {notification.Initialize, notification.Finalize},
	"txqueue":          {txqueue.Initialize, txqueue.Finalize},
	"limitation":       {limitation.Initialize, limitation.Finalize},
	"withdraw":         {withdraw.Initialize, withdraw.Finalize},
	"couponwithdraw":   {couponwithdraw.Initialize, couponwithdraw.Finalize},
	"deposit":          {deposit.Initialize, deposit.Finalize},
	"benefit":          {benefit.Initialize, benefit.Finalize},
	"notifbenefit":     {notifbenefit.Initialize, notifbenefit.Finalize},
	"good":             {good.Initialize, good.Finalize},
}

func FinalizeSubsystem(ctx context.Context, system string) {
	_finalizer, ok := subsystems[system]
	if !ok {
		return
	}
	_finalizer.final(ctx)
}

func InitializeSubsystem(system string) {
	_initializer, ok := subsystems[system]
	if !ok {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	_initializer.init(ctx, cancel)
}
