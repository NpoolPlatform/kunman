package location

import (
	"context"

	locationmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	locationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateLocation(ctx context.Context) (*locationmwpb.Location, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkLocation(ctx); err != nil {
		return nil, err
	}

	if err := locationmwcli.UpdateLocation(ctx, &locationmwpb.LocationReq{
		ID:       h.ID,
		EntID:    h.EntID,
		Country:  h.Country,
		Province: h.Province,
		City:     h.City,
		Address:  h.Address,
		BrandID:  h.BrandID,
	}); err != nil {
		return nil, err
	}
	return h.GetLocation(ctx)
}
