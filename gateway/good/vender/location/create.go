package location

import (
	"context"

	locationmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/location"
	locationmw "github.com/NpoolPlatform/kunman/middleware/good/vender/location"

	"github.com/google/uuid"
)

func (h *Handler) CreateLocation(ctx context.Context) (*locationmwpb.Location, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := locationmw.NewHandler(
		ctx,
		locationmw.WithEntID(h.EntID, true),
		locationmw.WithCountry(h.Country, true),
		locationmw.WithProvince(h.Province, true),
		locationmw.WithCity(h.City, true),
		locationmw.WithAddress(h.Address, true),
		locationmw.WithBrandID(h.BrandID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateLocation(ctx); err != nil {
		return nil, err
	}
	return h.GetLocation(ctx)
}
