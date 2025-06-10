package poster

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/poster"
	appgoodpostermw "github.com/NpoolPlatform/kunman/middleware/good/app/good/poster"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdatePoster(ctx context.Context) (*npool.Poster, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPoster(ctx); err != nil {
		return nil, err
	}

	posterHandler, err := appgoodpostermw.NewHandler(
		ctx,
		appgoodpostermw.WithID(h.ID, true),
		appgoodpostermw.WithEntID(h.EntID, true),
		appgoodpostermw.WithPoster(h.Poster, true),
		appgoodpostermw.WithIndex(func() *uint8 {
			if h.Index == nil {
				return nil
			}
			u := uint8(*h.Index)
			return &u
		}(), true),
	)
	if err != nil {
		return nil, err
	}

	if err := posterHandler.UpdatePoster(ctx); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
