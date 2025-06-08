package poster

import (
	"context"

	topmostpostermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/poster"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/poster"
	topmostpostermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/poster"

	"github.com/google/uuid"
)

func (h *Handler) CreatePoster(ctx context.Context) (*npool.Poster, error) {
	if err := h.CheckTopMost(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := topmostpostermwcli.CreatePoster(ctx, &topmostpostermwpb.PosterReq{
		EntID:     h.EntID,
		TopMostID: h.TopMostID,
		Poster:    h.Poster,
		Index:     h.Index,
	}); err != nil {
		return nil, err
	}

	return h.GetPoster(ctx)
}
