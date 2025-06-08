package brand

import (
	"context"

	brandmwcli "github.com/NpoolPlatform/kunman/middleware/good/vender/brand"
	brandmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/brand"

	"github.com/google/uuid"
)

func (h *Handler) CreateBrand(ctx context.Context) (*brandmwpb.Brand, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := brandmwcli.CreateBrand(ctx, &brandmwpb.BrandReq{
		EntID: h.EntID,
		Name:  h.Name,
		Logo:  h.Logo,
	}); err != nil {
		return nil, err
	}
	return h.GetBrand(ctx)
}
