package subscription

import (
	"context"

	subscriptionmwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/user/subscription"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/subscription"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteSubscription(ctx context.Context) (*npool.UserSubscription, error) {
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
