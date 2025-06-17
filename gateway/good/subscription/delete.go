package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	subscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription"
	subscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/subscription"
)

func (h *Handler) DeleteSubscription(ctx context.Context) (*subscriptionmwpb.Subscription, error) {
	info, err := h.GetSubscription(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	handler, err := subscriptionmw.NewHandler(
		ctx,
		subscriptionmw.WithID(h.ID, true),
		subscriptionmw.WithEntID(h.EntID, true),
		subscriptionmw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.DeleteSubscription(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return info, nil
}
