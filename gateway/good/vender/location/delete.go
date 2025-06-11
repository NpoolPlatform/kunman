package location

import (
	"context"

	locationmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/location"
	locationmw "github.com/NpoolPlatform/kunman/middleware/good/vender/location"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteLocation(ctx context.Context) (*locationmwpb.Location, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkLocation(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetLocation(ctx)
	if err != nil {
		return nil, err
	}

	locationHandler, err := locationmw.NewHandler(
		ctx,
		locationmw.WithID(h.ID, true),
		locationmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := locationHandler.DeleteLocation(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
