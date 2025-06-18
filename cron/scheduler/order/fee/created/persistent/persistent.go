package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/created/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentFeeOrder)
	if !ok {
		return fmt.Errorf("invalid feeorder")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	handler, err := feeordermw.NewHandler(
		ctx,
		feeordermw.WithID(&_order.ID, true),
		feeordermw.WithOrderState(ordertypes.OrderState_OrderStateWaitPayment.Enum(), true),
	)
	if err != nil {
		return err
	}

	return handler.UpdateFeeOrder(ctx)
}
