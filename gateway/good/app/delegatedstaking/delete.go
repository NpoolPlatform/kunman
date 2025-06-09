package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
	appdelegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/app/delegatedstaking"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteDelegatedStaking(ctx context.Context) (*npool.AppDelegatedStaking, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetDelegatedStaking(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid delegatedstaking")
	}

	dsHandler, err := appdelegatedstakingmw.NewHandler(
		ctx,
		appdelegatedstakingmw.WithID(h.ID, false),
		appdelegatedstakingmw.WithEntID(h.EntID, false),
		appdelegatedstakingmw.WithAppGoodID(h.AppGoodID, false),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := dsHandler.DeleteDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
