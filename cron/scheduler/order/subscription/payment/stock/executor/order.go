package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/stock/types"
	logger "github.com/NpoolPlatform/kunman/framework/logger"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
)

type orderHandler struct {
	*subscriptionordermwpb.SubscriptionOrder
	persistent chan interface{}
	done       chan interface{}
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"SubscriptionOrder", h.SubscriptionOrder,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		SubscriptionOrder: h.SubscriptionOrder,
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.done)
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	return nil
}
