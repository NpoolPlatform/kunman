package persistent

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/spend/types"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgermw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"

	"github.com/google/uuid"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) updateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	handler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithID(&order.ID, true),
		subscriptionordermw.WithOrderState(ordertypes.OrderState_OrderStateTransferGoodStockLocked.Enum(), true),
	)
	if err != nil {
		return err
	}

	return handler.UpdateSubscriptionOrder(ctx)
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	return p.updateOrderState(ctx, order)
}

func (p *handler) withSpendLockedBalance(ctx context.Context, order *types.PersistentOrder) error {
	handler, err := ledgermw.NewHandler(
		ctx,
		ledgermw.WithLockID(&order.LedgerLockID, true),
		ledgermw.WithStatementIDs(func() (statementIDs []string) {
			for range order.PaymentBalances {
				statementIDs = append(statementIDs, uuid.NewString())
			}
			return
		}(), true),
		ledgermw.WithIOExtra(&order.BalanceOutcomingExtra, true),
		ledgermw.WithIOSubType(ledgertypes.IOSubType_Payment.Enum(), true),
	)
	if err != nil {
		return err
	}

	_, err = handler.SettleBalances(ctx)
	return err
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return wlog.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if len(_order.PaymentBalances) == 0 {
		return wlog.WrapError(p.updateOrderState(ctx, _order))
	}

	if err := p.withUpdateOrderState(ctx, _order); err != nil {
		return err
	}
	if err := p.withSpendLockedBalance(ctx, _order); err != nil {
		return err
	}

	return nil
}
