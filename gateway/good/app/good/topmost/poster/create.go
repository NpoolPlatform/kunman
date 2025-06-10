package poster

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/poster"
	topmostpostermw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/poster"

	"github.com/google/uuid"
)

func (h *Handler) CreatePoster(ctx context.Context) (*npool.Poster, error) {
	if err := h.CheckTopMost(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := topmostpostermw.NewHandler(
		ctx,
		topmostpostermw.WithEntID(h.EntID, true),
		topmostpostermw.WithTopMostID(h.TopMostID, true),
		topmostpostermw.WithPoster(h.Poster, true),
		topmostpostermw.WithIndex(func() *uint8 {
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
