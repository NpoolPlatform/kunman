package poster

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/poster"
	appgoodpostermw "github.com/NpoolPlatform/kunman/middleware/good/app/good/poster"

	"github.com/google/uuid"
)

func (h *Handler) CreatePoster(ctx context.Context) (*npool.Poster, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := appgoodpostermw.NewHandler(
		ctx,
		appgoodpostermw.WithEntID(h.EntID, true),
		appgoodpostermw.WithAppGoodID(h.AppGoodID, true),
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

	if err := handler.CreatePoster(ctx); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
