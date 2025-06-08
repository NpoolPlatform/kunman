package poster

import (
	"context"

	topmostpostermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost/good/poster"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good/poster"
	topmostpostermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good/poster"
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
