package poster

import (
	"context"

	postermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"
	postermw "github.com/NpoolPlatform/kunman/middleware/good/device/poster"
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

	posterHandler, err := postermw.NewHandler(
		ctx,
		postermw.WithID(h.ID, true),
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

	if err := posterHandler.UpdatePoster(ctx); err != nil {
		return nil, err
	}
	return h.GetPoster(ctx)
}
