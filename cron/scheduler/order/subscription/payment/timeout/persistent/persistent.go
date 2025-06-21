package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/timeout/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
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

	state := ordertypes.OrderState_OrderStatePreCancel

	handler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithID(&_order.ID, true),
		subscriptionordermw.WithOrderState(&state, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdateSubscriptionOrder(ctx)
}
