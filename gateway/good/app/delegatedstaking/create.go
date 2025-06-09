package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
	delegatedstakingmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/delegatedstaking"
	appdelegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/app/delegatedstaking"
	delegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"

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

	handler, err := delegatedstakingmw.NewHandler(
		ctx,
		delegatedstakingmw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.delegatedstaking, err = handler.GetDelegatedStaking(ctx)
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

	handler, err := appdelegatedstakingmw.NewHandler(
		ctx,
		appdelegatedstakingmw.WithEntID(h.EntID, true),
		appdelegatedstakingmw.WithAppID(h.AppID, true),
		appdelegatedstakingmw.WithAppGoodID(h.AppGoodID, true),
		appdelegatedstakingmw.WithPurchasable(h.Purchasable, true),
		appdelegatedstakingmw.WithEnableProductPage(h.EnableProductPage, true),
		appdelegatedstakingmw.WithProductPage(h.ProductPage, true),
		appdelegatedstakingmw.WithOnline(h.Online, true),
		appdelegatedstakingmw.WithVisible(h.Visible, true),
		appdelegatedstakingmw.WithName(h.Name, true),
		appdelegatedstakingmw.WithDisplayIndex(h.DisplayIndex, true),
		appdelegatedstakingmw.WithBanner(h.Banner, true),
		appdelegatedstakingmw.WithServiceStartAt(h.ServiceStartAt, true),
		appdelegatedstakingmw.WithEnableSetCommission(h.EnableSetCommission, true),
		appdelegatedstakingmw.WithStartMode(h.StartMode, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateDelegatedStaking(ctx); err != nil {
		return nil, err
	}
	return h.GetDelegatedStaking(ctx)
}
