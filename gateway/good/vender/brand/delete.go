package brand

import (
	"context"

	brandmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/brand"
	brandmw "github.com/NpoolPlatform/kunman/middleware/good/vender/brand"
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

	brandHandler, err := brandmw.NewHandler(
		ctx,
		brandmw.WithID(h.ID, true),
		brandmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := brandHandler.DeleteBrand(ctx); err != nil {
		return nil, err
	}
	return info, nil
}
