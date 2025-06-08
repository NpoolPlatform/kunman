package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/subscription"
	subscriptionmw "github.com/NpoolPlatform/kunman/middleware/agi/subscription"
)

func (h *Handler) GetSubscription(ctx context.Context) (*npool.Subscription, error) {
	if h.EntID == nil && h.AppID == nil && h.UserID == nil {
		return nil, wlog.Errorf("invalid id")
	}
	if h.EntID == nil {
		if h.AppID == nil || h.UserID == nil {
			return nil, wlog.Errorf("invalid id")
		}
	}

	handler, err := subscriptionmw.NewHandler(
		ctx,
		subscriptionmw.WithEntID(h.EntID, false),
		subscriptionmw.WithAppID(h.AppID, false),
		subscriptionmw.WithUserID(h.UserID, false),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetSubscription(ctx)
}

func (h *Handler) GetSubscriptions(ctx context.Context) ([]*npool.Subscription, error) {
	handler, err := subscriptionmw.NewHandler(
		ctx,
		subscriptionmw.WithAppID(h.AppID, false),
		subscriptionmw.WithAppGoodID(h.AppGoodID, false),
		subscriptionmw.WithOffset(h.Offset),
		subscriptionmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetSubscriptions(ctx)
}
