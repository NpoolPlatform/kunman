package delegatedstaking

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
	appdelegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/app/delegatedstaking"
)

// TODO: check start mode with power rental start mode

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateDelegatedStaking(ctx context.Context) (*npool.AppDelegatedStaking, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDelegatedStaking(ctx); err != nil {
		return nil, err
	}

	dsHandler, err := appdelegatedstakingmw.NewHandler(
		ctx,
		appdelegatedstakingmw.WithID(h.ID, true),
		appdelegatedstakingmw.WithEntID(h.EntID, true),
		appdelegatedstakingmw.WithAppGoodID(h.AppGoodID, true),
		appdelegatedstakingmw.WithPurchasable(h.Purchasable, false),
		appdelegatedstakingmw.WithEnableProductPage(h.EnableProductPage, false),
		appdelegatedstakingmw.WithProductPage(h.ProductPage, false),
		appdelegatedstakingmw.WithOnline(h.Online, false),
		appdelegatedstakingmw.WithVisible(h.Visible, false),
		appdelegatedstakingmw.WithName(h.Name, false),
		appdelegatedstakingmw.WithDisplayIndex(h.DisplayIndex, false),
		appdelegatedstakingmw.WithBanner(h.Banner, false),
		appdelegatedstakingmw.WithServiceStartAt(h.ServiceStartAt, false),
		appdelegatedstakingmw.WithEnableSetCommission(h.EnableSetCommission, false),
		appdelegatedstakingmw.WithStartMode(h.StartMode, false),
	)
	if err != nil {
		return nil, err
	}

	if err := dsHandler.UpdateDelegatedStaking(ctx); err != nil {
		return nil, err
	}
	return h.GetDelegatedStaking(ctx)
}
