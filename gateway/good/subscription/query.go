package subscription

import (
	"context"

	subscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription"
	subscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/subscription"
)

func (h *Handler) GetSubscription(ctx context.Context) (*subscriptionmwpb.Subscription, error) {
	handler, err := subscriptionmw.NewHandler(
		ctx,
		subscriptionmw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetSubscription(ctx)
}

func (h *Handler) GetSubscriptions(ctx context.Context) ([]*subscriptionmwpb.Subscription, error) {
	handler, err := subscriptionmw.NewHandler(
		ctx,
		subscriptionmw.WithConds(&subscriptionmwpb.Conds{}),
		subscriptionmw.WithOffset(h.Offset),
		subscriptionmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetSubscriptions(ctx)
}
