package topmost

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"
	topmostmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteTopMost(ctx context.Context) (*npool.TopMost, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkTopMost(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetTopMost(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid topmost")
	}

	topMostHandler, err := topmostmw.NewHandler(
		ctx,
		topmostmw.WithID(h.ID, true),
		topmostmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := topMostHandler.DeleteTopMost(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
