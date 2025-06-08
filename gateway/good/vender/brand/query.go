package brand

import (
	"context"

	brandmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (h *Handler) GetBrand(ctx context.Context) (*brandmwpb.Brand, error) {
	return brandmwcli.GetBrand(ctx, *h.EntID)
}

func (h *Handler) GetBrands(ctx context.Context) ([]*brandmwpb.Brand, uint32, error) {
	return brandmwcli.GetBrands(ctx, &brandmwpb.Conds{}, h.Offset, h.Limit)
}
