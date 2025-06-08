package poster

import (
	"context"

	postermwcli "github.com/NpoolPlatform/kunman/middleware/good/device/poster"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	postermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"
)

func (h *Handler) GetPoster(ctx context.Context) (*postermwpb.Poster, error) {
	return postermwcli.GetPoster(ctx, *h.EntID)
}

func (h *Handler) GetPosters(ctx context.Context) ([]*postermwpb.Poster, uint32, error) {
	conds := &postermwpb.Conds{}
	if h.DeviceTypeID != nil {
		conds.DeviceTypeID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.DeviceTypeID}
	}
	return postermwcli.GetPosters(ctx, conds, h.Offset, h.Limit)
}
