package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/payment/commission/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	state := ordertypes.OrderState_OrderStatePaymentUnlockAccount

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
