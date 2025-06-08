package brand

import (
	"context"

	brandmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteBrand(ctx context.Context) (*brandmwpb.Brand, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkBrand(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetBrand(ctx)
	if err != nil {
		return nil, err
	}
	if err := brandmwcli.DeleteBrand(ctx, h.ID, h.EntID); err != nil {
		return nil, err
	}
	return info, nil
}
