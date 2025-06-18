package persistent

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/achievement/types"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	achievementmw "github.com/NpoolPlatform/kunman/middleware/inspire/achievement"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentFeeOrder)
	if !ok {
		return wlog.Errorf("invalid feeorder")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	achievementHandler, err := achievementmw.NewHandler(
		ctx,
		achievementmw.WithOrderID(&_order.OrderID, true),
	)
	if err != nil {
		return err
	}

	if err := achievementHandler.ExpropriateAchievement(ctx); err != nil {
		return wlog.WrapError(err)
	}

	feeHandler, err := feeordermw.NewHandler(
		ctx,
		feeordermw.WithID(&_order.ID, true),
		feeordermw.WithOrderState(ordertypes.OrderState_OrderStateReturnCanceledBalance.Enum(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return wlog.WrapError(feeHandler.UpdateFeeOrder(ctx))
}
