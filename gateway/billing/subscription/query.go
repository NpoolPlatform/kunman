package subscription

import (
	"context"

	submwpb "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/subscription"
)

func (h *Handler) GetSubscription(ctx context.Context) (*submwpb.Subscription, error) {
	return submwcli.GetSubscription(ctx, *h.EntID)
}

func (h *Handler) GetSubscriptions(ctx context.Context) ([]*submwpb.Subscription, error) {
	return submwcli.GetSubscriptions(ctx, &submwpb.Conds{}, h.Offset, h.Limit)
}

func (h *Handler) GetSubscriptionsCount(ctx context.Context) (uint32, error) {
	return submwcli.GetSubscriptionsCount(ctx, &submwpb.Conds{})
}
