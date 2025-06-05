package subscription

import (
	"context"

	submwpb "github.com/NpoolPlatform/kunman/message/billing/middleware/v1/subscription"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateSubscription(ctx context.Context) (*submwpb.Subscription, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSubscription(ctx); err != nil {
		return nil, err
	}
	if err := submwcli.UpdateSubscription(ctx, &submwpb.SubscriptionReq{
		ID:          h.ID,
		EntID:       h.EntID,
		PackageName: h.PackageName,
		UsdPrice:    h.UsdPrice,
		Description: h.Description,
		SortOrder:   h.SortOrder,
		Credit:      h.Credit,
		ResetType:   h.ResetType,
		QPSLimit:    h.QPSLimit,
	}); err != nil {
		return nil, err
	}
	return h.GetSubscription(ctx)
}
