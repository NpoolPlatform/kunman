package poster

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/poster"
	postermw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good/poster"

	"github.com/google/uuid"
)

func (h *Handler) CreatePoster(ctx context.Context) (*npool.Poster, error) {
	if err := h.CheckTopMostGood(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := postermw.NewHandler(
		ctx,
		postermw.WithEntID(h.EntID, true),
		postermw.WithTopMostGoodID(h.TopMostGoodID, true),
		postermw.WithPoster(h.Poster, true),
		postermw.WithIndex(func() *uint8 {
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
