package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appdelegatedstakingmwcli "github.com/NpoolPlatform/kunman/middleware/good/app/delegatedstaking"
	delegatedstakingmwcli "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
	appdelegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/delegatedstaking"
	delegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/delegatedstaking"

	"github.com/google/uuid"
)

type CreateHander struct {
	*Handler
	delegatedstaking *delegatedstakingmwpb.DelegatedStaking
}

func (h *CreateHander) getDelegatedStaking(ctx context.Context) (err error) {
	if h.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}

	h.delegatedstaking, err = delegatedstakingmwcli.GetDelegatedStaking(ctx, *h.GoodID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

// TODO: check start mode with power rental start mode
func (h *Handler) CreateDelegatedStaking(ctx context.Context) (*npool.AppDelegatedStaking, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.AppGoodID == nil {
		h.AppGoodID = func() *string { s := uuid.NewString(); return &s }()
	}

	createH := &CreateHander{Handler: h}

	if err := createH.getDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := appdelegatedstakingmwcli.CreateDelegatedStaking(ctx, &appdelegatedstakingmwpb.DelegatedStakingReq{
		EntID:               h.EntID,
		AppID:               h.AppID,
		GoodID:              h.GoodID,
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
