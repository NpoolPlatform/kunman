package subscription

import (
	"context"

	submwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/user/subscription"
	subchangemwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/user/subscription/change"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/subscription"
	submwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/user/subscription"
	subchangemwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/user/subscription/change"
)

type updateHandler struct {
	*checkHandler
}

func (h *updateHandler) subscriptionChange(ctx context.Context) error {
	if h.PackageID == nil {
		return nil
	}
	sub, err := submwcli.GetSubscription(ctx, *h.EntID)
	if err != nil {
		return err
	}
	if sub.PackageID == *h.PackageID {
		return nil
	}
	if err := subchangemwcli.CreateSubscriptionChange(ctx, &subchangemwpb.SubscriptionChangeReq{
		AppID:              h.AppID,
		UserID:             h.UserID,
		UserSubscriptionID: h.EntID,
		OldPackageID:       &sub.PackageID,
		NewPackageID:       h.PackageID,
	}); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateSubscription(ctx context.Context) (*npool.UserSubscription, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSubscription(ctx); err != nil {
		return nil, err
	}
	if err := handler.subscriptionChange(ctx); err != nil {
		return nil, err
	}
	if err := submwcli.UpdateSubscription(ctx, &submwpb.SubscriptionReq{
		ID:                 h.ID,
		EntID:              h.EntID,
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
