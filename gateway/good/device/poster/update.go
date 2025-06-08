package poster

import (
	"context"

	postermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/poster"
	postermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdatePoster(ctx context.Context) (*postermwpb.Poster, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPoster(ctx); err != nil {
		return nil, err
	}

	if err := postermwcli.UpdatePoster(ctx, &postermwpb.PosterReq{
		ID:     h.ID,
		EntID:  h.EntID,
		Poster: h.Poster,
		Index:  h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
