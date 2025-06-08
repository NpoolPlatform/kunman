package poster

import (
	"context"

	postermwcli "github.com/NpoolPlatform/good-middleware/pkg/client/device/poster"
	postermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"

	"github.com/google/uuid"
)

func (h *Handler) CreatePoster(ctx context.Context) (*postermwpb.Poster, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := postermwcli.CreatePoster(ctx, &postermwpb.PosterReq{
		EntID:        h.EntID,
		DeviceTypeID: h.DeviceTypeID,
		Poster:       h.Poster,
		Index:        h.Index,
	}); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
