package subscription

import (
	"context"

	subscriptionmwpb "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/subscription"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteSubscription(ctx context.Context) (*subscriptionmwpb.Subscription, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSubscription(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetSubscription(ctx)
	if err != nil {
		return nil, err
	}
	if err := subscriptionmwcli.DeleteSubscription(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
