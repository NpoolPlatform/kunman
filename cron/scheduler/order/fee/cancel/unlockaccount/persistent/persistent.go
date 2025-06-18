package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/unlockaccount/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	paymentaccmw "github.com/NpoolPlatform/kunman/middleware/account/payment"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	state := ordertypes.OrderState_OrderStateCanceled

	handler, err := feeordermw.NewHandler(
		ctx,
		feeordermw.WithID(&order.ID, true),
		feeordermw.WithOrderState(&state, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdateFeeOrder(ctx)
}

func (p *handler) withUnlockPaymentAccount(ctx context.Context, order *types.PersistentOrder) error {
	// TODO: use UpdateAccounts in future
	for _, paymentAccountID := range order.PaymentAccountIDs {
		locked := false
		handler, err := paymentaccmw.NewHandler(
			ctx,
			paymentaccmw.WithID(&paymentAccountID, true),
			paymentaccmw.WithLocked(&locked, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.UpdateAccount(ctx); err != nil {
			return err
		}
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
	if err := p.withUnlockPaymentAccount(ctx, _order); err != nil {
		return err
	}

	return nil
}
