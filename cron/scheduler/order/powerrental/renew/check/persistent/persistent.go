package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/renew/check/types"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	powerrentaloutofgasmw "github.com/NpoolPlatform/kunman/middleware/order/powerrental/outofgas"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if _order.CreateOutOfGas {
		handler, err := powerrentaloutofgasmw.NewHandler(
			ctx,
			powerrentaloutofgasmw.WithOrderID(&_order.OrderID, true),
			powerrentaloutofgasmw.WithStartAt(&_order.FeeEndAt, true),
		)
		if err != nil {
			return err
		}

		if err := handler.CreateOutOfGas(ctx); err != nil {
			return err
		}
	}

	if _order.FinishOutOfGas {
		handler, err := powerrentaloutofgasmw.NewHandler(
			ctx,
			powerrentaloutofgasmw.WithEntID(&_order.OutOfGasEntID, true),
			powerrentaloutofgasmw.WithEndAt(&_order.OutOfGasFinishedAt, true),
		)
		if err != nil {
			return err
		}

		if err := handler.UpdateOutOfGas(ctx); err != nil {
			return err
		}
	}

	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithID(&_order.ID, true),
		powerrentalordermw.WithRenewState(&_order.NewRenewState, true),
		powerrentalordermw.WithRenewNotifyAt(&_order.NextRenewNotifyAt, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdatePowerRental(ctx)
}
