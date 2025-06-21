package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/commission/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	state := ordertypes.OrderState_OrderStatePaymentUnlockAccount

	handler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithID(&order.ID, true),
		subscriptionordermw.WithOrderState(&state, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdateSubscriptionOrder(ctx)
}

func (p *handler) withCreateCommission(ctx context.Context, order *types.PersistentOrder) error {
	handler, err := ledgerstatementmw.NewHandler(
		ctx,
		ledgerstatementmw.WithReqs(order.LedgerStatements, true),
	)
	if err != nil {
		return err
	}

	_, err = handler.CreateStatements(ctx)
	return err
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
	if err := p.withCreateCommission(ctx, _order); err != nil {
		return err
	}

	return nil
}
