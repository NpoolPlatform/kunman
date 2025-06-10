package poster

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/poster"
	appgoodpostermw "github.com/NpoolPlatform/kunman/middleware/good/app/good/poster"
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

	posterHandler, err := appgoodpostermw.NewHandler(
		ctx,
		appgoodpostermw.WithID(h.ID, true),
		appgoodpostermw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := posterHandler.DeletePoster(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
