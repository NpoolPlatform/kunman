package brand

import (
	"context"

	brandmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/brand"
	brandmw "github.com/NpoolPlatform/kunman/middleware/good/vender/brand"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateBrand(ctx context.Context) (*brandmwpb.Brand, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkBrand(ctx); err != nil {
		return nil, err
	}

	brandHandler, err := brandmw.NewHandler(
		ctx,
		brandmw.WithID(h.ID, true),
		brandmw.WithEntID(h.EntID, true),
		brandmw.WithName(h.Name, false),
		brandmw.WithLogo(h.Logo, false),
	)
	if err != nil {
		return nil, err
	}

	if err := brandHandler.UpdateBrand(ctx); err != nil {
		return nil, err
	}
	return h.GetBrand(ctx)
}
