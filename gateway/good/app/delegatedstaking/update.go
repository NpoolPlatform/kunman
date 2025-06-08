package delegatedstaking

import (
	"context"

	appdelegatedstakingmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/delegatedstaking"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/delegatedstaking"
	appdelegatedstakingmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/delegatedstaking"
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
	if err := appdelegatedstakingmwcli.UpdateDelegatedStaking(ctx, &appdelegatedstakingmwpb.DelegatedStakingReq{
		ID:                  h.ID,
		EntID:               h.EntID,
		AppGoodID:           h.AppGoodID,
		Purchasable:         h.Purchasable,
		EnableProductPage:   h.EnableProductPage,
		ProductPage:         h.ProductPage,
		Online:              h.Online,
		Visible:             h.Visible,
		Name:                h.Name,
		DisplayIndex:        h.DisplayIndex,
		Banner:              h.Banner,
		ServiceStartAt:      h.ServiceStartAt,
		EnableSetCommission: h.EnableSetCommission,
		StartMode:           h.StartMode,
	}); err != nil {
		return nil, err
	}
	return h.GetDelegatedStaking(ctx)
}
