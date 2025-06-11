package poster

import (
	"context"

	postermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"
	postermw "github.com/NpoolPlatform/kunman/middleware/good/device/poster"

	"github.com/google/uuid"
)

func (h *Handler) CreatePoster(ctx context.Context) (*postermwpb.Poster, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := postermw.NewHandler(
		ctx,
		postermw.WithEntID(h.EntID, true),
		postermw.WithDeviceTypeID(h.DeviceTypeID, true),
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
