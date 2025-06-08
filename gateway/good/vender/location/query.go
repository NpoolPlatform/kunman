package location

import (
	"context"

	locationmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	locationmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/location"
)

func (h *Handler) GetLocation(ctx context.Context) (*locationmwpb.Location, error) {
	return locationmwcli.GetLocation(ctx, *h.EntID)
}

func (h *Handler) GetLocations(ctx context.Context) ([]*locationmwpb.Location, uint32, error) {
	conds := &locationmwpb.Conds{}
	if h.BrandID != nil {
		conds.BrandID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.BrandID}
	}
	return locationmwcli.GetLocations(ctx, conds, h.Offset, h.Limit)
}
