package poster

import (
	"context"

	appgoodpostermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/poster"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/poster"
	appgoodpostermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/poster"
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

	if err := appgoodpostermwcli.UpdatePoster(ctx, &appgoodpostermwpb.PosterReq{
		ID:     h.ID,
		EntID:  h.EntID,
		Poster: h.Poster,
		Index:  h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
