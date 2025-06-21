package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/created/types"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
)

type orderHandler struct {
	*subscriptionordermwpb.SubscriptionOrder
	persistent chan interface{}
}

func (h *orderHandler) final(ctx context.Context) {
	persistentSubscriptionOrder := &types.PersistentSubscriptionOrder{
		SubscriptionOrder: h.SubscriptionOrder,
	}
	asyncfeed.AsyncFeed(ctx, persistentSubscriptionOrder, h.persistent)
}

func (h *orderHandler) exec(ctx context.Context) error {
	h.final(ctx)
	return nil
}
