package topmostgood

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
	topmostgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkTopMostGood(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetTopMostGood(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid topmostgood")
	}

	goodHandler, err := topmostgoodmw.NewHandler(
		ctx,
		topmostgoodmw.WithID(h.ID, true),
		topmostgoodmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := goodHandler.DeleteTopMostGood(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
