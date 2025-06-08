package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	delegatedstakingmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/delegatedstaking"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/delegatedstaking"
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
	if err := delegatedstakingmwcli.DeleteDelegatedStaking(ctx, h.ID, h.EntID, h.GoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
