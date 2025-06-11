package brand

import (
	"context"

	brandmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/brand"
	brandmw "github.com/NpoolPlatform/kunman/middleware/good/vender/brand"
)

func (h *Handler) GetBrand(ctx context.Context) (*brandmwpb.Brand, error) {
	handler, err := brandmw.NewHandler(
		ctx,
		brandmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetBrand(ctx)
}

func (h *Handler) GetBrands(ctx context.Context) ([]*brandmwpb.Brand, uint32, error) {
	handler, err := brandmw.NewHandler(
		ctx,
		brandmw.WithConds(&brandmwpb.Conds{}),
		brandmw.WithOffset(h.Offset),
		brandmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetBrands(ctx)
}
