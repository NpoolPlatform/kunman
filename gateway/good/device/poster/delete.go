package poster

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	postermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"
	postermw "github.com/NpoolPlatform/kunman/middleware/good/device/poster"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeletePoster(ctx context.Context) (*postermwpb.Poster, error) {
	info, err := h.GetPoster(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("poster not exist")
	}
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPoster(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	posterHandler, err := postermw.NewHandler(
		ctx,
		postermw.WithID(h.ID, true),
		postermw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := posterHandler.DeletePoster(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
