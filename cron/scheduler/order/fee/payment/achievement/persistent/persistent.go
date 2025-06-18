package persistent

import (
	"context"
	"fmt"

	orderstatementmwcli "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/statement/order"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/achievement/types"
	feeordermwcli "github.com/NpoolPlatform/kunman/middleware/order/fee"
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

	if len(_order.OrderStatements) > 0 {
		if err := orderstatementmwcli.CreateStatements(ctx, _order.OrderStatements); err != nil {
			return err
		}
	}
	return feeordermwcli.UpdateFeeOrder(ctx, &feeordermwpb.FeeOrderReq{
		ID:         &_order.ID,
		OrderState: ordertypes.OrderState_OrderStateAddCommission.Enum(),
	})
}
