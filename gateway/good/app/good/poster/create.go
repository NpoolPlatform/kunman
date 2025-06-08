package poster

import (
	"context"

	appgoodpostermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/poster"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/poster"
	appgoodpostermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/poster"

	"github.com/google/uuid"
)

func (h *Handler) CreatePoster(ctx context.Context) (*npool.Poster, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	if err := appgoodpostermwcli.CreatePoster(ctx, &appgoodpostermwpb.PosterReq{
		EntID:     h.EntID,
		AppGoodID: h.AppGoodID,
		Poster:    h.Poster,
		Index:     h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
