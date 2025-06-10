package poster

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/poster"
	topmostpostermw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good/poster"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeletePoster(ctx context.Context) (*npool.Poster, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPoster(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := h.GetPoster(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid poster")
	}

	posterHandler, err := topmostpostermw.NewHandler(
		ctx,
		topmostpostermw.WithID(h.ID, true),
		topmostpostermw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := posterHandler.DeletePoster(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	return info, nil
}
