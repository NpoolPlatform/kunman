package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/timeout/types"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
)

type orderHandler struct {
	*subscriptionordermwpb.SubscriptionOrder
	persistent chan interface{}
}

func (h *orderHandler) final(ctx context.Context) {
	persistentOrder := &types.PersistentOrder{
		SubscriptionOrder: h.SubscriptionOrder,
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
}

func (h *orderHandler) exec(ctx context.Context) error {
	h.final(ctx)
	return nil
}
