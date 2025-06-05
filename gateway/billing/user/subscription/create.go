package subscription

import (
	"context"

	submwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/user/subscription"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/subscription"
	submwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/user/subscription"

	"github.com/google/uuid"
)

func (h *Handler) CreateSubscription(ctx context.Context) (*npool.UserSubscription, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.StartAt == nil {
		h.StartAt = func() *uint32 { u := uint32(0); return &u }()
	}
	if h.EndAt == nil {
		h.EndAt = func() *uint32 { u := uint32(0); return &u }()
	}

	if err := submwcli.CreateSubscription(ctx, &submwpb.SubscriptionReq{
		EntID:              h.EntID,
		AppID:              h.AppID,
		UserID:             h.UserID,
		PackageID:          h.PackageID,
		StartAt:            h.StartAt,
		EndAt:              h.EndAt,
		UsageState:         h.UsageState,
		SubscriptionCredit: h.SubscriptionCredit,
		AddonCredit:        h.AddonCredit,
	}); err != nil {
		return nil, err
	}
	return h.GetSubscription(ctx)
}
