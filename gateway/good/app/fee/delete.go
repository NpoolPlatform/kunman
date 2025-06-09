package appfee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteAppFee(ctx context.Context) (*npool.AppFee, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkAppFee(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetAppFee(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid fee")
	}

	feeHandler, err := appfeemw.NewHandler(
		ctx,
		appfeemw.WithID(h.ID, false),
		appfeemw.WithEntID(h.EntID, false),
		appfeemw.WithAppGoodID(h.AppGoodID, false),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := feeHandler.DeleteFee(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
