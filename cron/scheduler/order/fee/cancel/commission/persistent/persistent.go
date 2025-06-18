package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/commission/types"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgermw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentFeeOrder) error {
	state := ordertypes.OrderState_OrderStateCancelAchievement

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

func (p *handler) withDeductLockedCommission(ctx context.Context, order *types.PersistentFeeOrder) error {
	for _, revoke := range order.CommissionRevokes {
		handler, err := ledgermw.NewHandler(
			ctx,
			ledgermw.WithLockID(&revoke.LockID, true),
			ledgermw.WithIOSubType(ledgertypes.IOSubType_CommissionRevoke.Enum(), true),
			ledgermw.WithIOExtra(&revoke.IOExtra, true),
			ledgermw.WithStatementIDs(revoke.StatementIDs, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.SettleBalances(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentFeeOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if err := p.withUpdateOrderState(ctx, _order); err != nil {
		return err
	}
	if err := p.withDeductLockedCommission(ctx, _order); err != nil {
		return err
	}

	return nil
}
