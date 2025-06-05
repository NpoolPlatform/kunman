package subscription

import (
	"context"

	submwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/subscription"
	submwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/subscription"

	"github.com/google/uuid"
)

func (h *Handler) CreateSubscription(ctx context.Context) (*submwpb.Subscription, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := submwcli.CreateSubscription(ctx, &submwpb.SubscriptionReq{
		EntID:       h.EntID,
		AppID:       h.AppID,
		PackageName: h.PackageName,
		UsdPrice:    h.UsdPrice,
		Description: h.Description,
		SortOrder:   h.SortOrder,
		PackageType: h.PackageType,
		Credit:      h.Credit,
		ResetType:   h.ResetType,
		QPSLimit:    h.QPSLimit,
	}); err != nil {
		return nil, err
	}
	return h.GetSubscription(ctx)
}
