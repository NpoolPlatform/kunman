package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"
	delegatedstakingmw "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"
)

func (h *Handler) DeleteDelegatedStaking(ctx context.Context) (*npool.DelegatedStaking, error) {
	handler := &checkHandler{
		Handler: h,
	}
	if err := handler.checkDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetDelegatedStaking(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	dsHandler, err := delegatedstakingmw.NewHandler(
		ctx,
		delegatedstakingmw.WithID(h.ID, false),
		delegatedstakingmw.WithEntID(h.EntID, false),
		delegatedstakingmw.WithGoodID(h.GoodID, false),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := dsHandler.DeleteDelegatedStaking(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
