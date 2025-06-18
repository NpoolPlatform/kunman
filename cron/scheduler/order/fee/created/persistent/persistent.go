package persistent

import (
	"context"
	"fmt"

	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/created/types"
	feeordermwcli "github.com/NpoolPlatform/kunman/middleware/order/fee"
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

	return feeordermwcli.UpdateFeeOrder(ctx, &feeordermwpb.FeeOrderReq{
		ID:         &_order.ID,
		OrderState: ordertypes.OrderState_OrderStateWaitPayment.Enum(),
	})
}
