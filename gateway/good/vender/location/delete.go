package location

import (
	"context"

	locationmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	locationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
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
	if err := locationmwcli.DeleteLocation(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
