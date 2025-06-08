package delegatedstaking

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appdelegatedstakingmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/delegatedstaking"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/delegatedstaking"
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
	if err := appdelegatedstakingmwcli.DeleteDelegatedStaking(ctx, h.ID, h.EntID, h.AppGoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
