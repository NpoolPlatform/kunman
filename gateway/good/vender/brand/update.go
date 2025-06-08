package brand

import (
	"context"

	brandmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
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

	if err := brandmwcli.UpdateBrand(ctx, &brandmwpb.BrandReq{
		ID:    h.ID,
		EntID: h.EntID,
		Name:  h.Name,
		Logo:  h.Logo,
	}); err != nil {
		return nil, err
	}
	return h.GetBrand(ctx)
}
