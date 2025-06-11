package location

import (
	"context"

	locationmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/location"
	locationmw "github.com/NpoolPlatform/kunman/middleware/good/vender/location"
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

	locationHandler, err := locationmw.NewHandler(
		ctx,
		locationmw.WithID(h.ID, true),
		locationmw.WithEntID(h.EntID, true),
		locationmw.WithCountry(h.Country, false),
		locationmw.WithProvince(h.Province, false),
		locationmw.WithCity(h.City, false),
		locationmw.WithAddress(h.Address, false),
		locationmw.WithBrandID(h.BrandID, false),
	)
	if err != nil {
		return nil, err
	}

	if err := locationHandler.UpdateLocation(ctx); err != nil {
		return nil, err
	}
	return h.GetLocation(ctx)
}
