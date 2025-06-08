package poster

import (
	"context"

	postermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good/poster"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/poster"
	postermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good/poster"

	"github.com/google/uuid"
)

func (h *Handler) CreatePoster(ctx context.Context) (*npool.Poster, error) {
	if err := h.CheckTopMostGood(ctx); err != nil {
		return nil, err
	}
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := postermwcli.CreatePoster(ctx, &postermwpb.PosterReq{
		EntID:         h.EntID,
		TopMostGoodID: h.TopMostGoodID,
		Poster:        h.Poster,
		Index:         h.Index,
	}); err != nil {
		return nil, err
	}

	return h.GetPoster(ctx)
}
