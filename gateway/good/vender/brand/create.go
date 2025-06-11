package brand

import (
	"context"

	brandmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/brand"
	brandmw "github.com/NpoolPlatform/kunman/middleware/good/vender/brand"

	"github.com/google/uuid"
)

func (h *Handler) CreateBrand(ctx context.Context) (*brandmwpb.Brand, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := brandmw.NewHandler(
		ctx,
		brandmw.WithEntID(h.EntID, true),
		brandmw.WithName(h.Name, true),
		brandmw.WithLogo(h.Logo, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateBrand(ctx); err != nil {
		return nil, err
	}
	return h.GetBrand(ctx)
}
