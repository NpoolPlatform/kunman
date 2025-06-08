package poster

import (
	"context"

	topmostpostermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good/poster"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost/good/poster"
	topmostpostermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/poster"
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

	if err := topmostpostermwcli.UpdatePoster(ctx, &topmostpostermwpb.PosterReq{
		ID:     h.ID,
		EntID:  h.EntID,
		Index:  h.Index,
		Poster: h.Poster,
	}); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
