package persistent

import (
	"context"
	"fmt"

	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/wait/types"
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

	req := &feeordermwpb.FeeOrderReq{
		ID:           &_order.ID,
		OrderState:   &_order.NewOrderState,
		PaymentState: _order.NewPaymentState,
	}
	return feeordermwcli.UpdateFeeOrder(ctx, req)
}
