package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/cancel/restorestock/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	appgoodstockmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/stock"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	state := ordertypes.OrderState_OrderStateDeductLockedCommission

	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithID(&order.ID, true),
		powerrentalordermw.WithOrderState(&state, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdatePowerRental(ctx)
}

func (p *handler) withUpdateStock(ctx context.Context, order *types.PersistentOrder) error {
	handler, err := appgoodstockmw.NewHandler(
		ctx,
		appgoodstockmw.WithLockID(&order.AppGoodStockLockID, true),
	)
	if err != nil {
		return err
	}

	switch order.CancelState {
	case ordertypes.OrderState_OrderStatePaymentTimeout:
		fallthrough //nolint
	case ordertypes.OrderState_OrderStateWaitPayment:
		return handler.UnlockStock(ctx)
	case ordertypes.OrderState_OrderStatePaid:
		fallthrough //nolint
	case ordertypes.OrderState_OrderStateInService:
		return handler.ChargeBackStock(ctx)
	}

	return nil
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if err := p.withUpdateOrderState(ctx, _order); err != nil {
		return err
	}
	if err := p.withUpdateStock(ctx, _order); err != nil {
		return err
	}

	return nil
}
