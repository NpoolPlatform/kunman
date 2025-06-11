package location

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	locationmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/location"
	locationmw "github.com/NpoolPlatform/kunman/middleware/good/vender/location"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) GetLocation(ctx context.Context) (*locationmwpb.Location, error) {
	handler, err := locationmw.NewHandler(
		ctx,
		locationmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetLocation(ctx)
}

func (h *Handler) GetLocations(ctx context.Context) ([]*locationmwpb.Location, uint32, error) {
	conds := &locationmwpb.Conds{}
	if h.BrandID != nil {
		conds.BrandID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.BrandID}
	}
	handler, err := locationmw.NewHandler(
		ctx,
		locationmw.WithConds(conds),
		locationmw.WithOffset(h.Offset),
		locationmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetLocations(ctx)
}
