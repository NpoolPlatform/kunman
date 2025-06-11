package poster

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	postermwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/device/poster"
	postermw "github.com/NpoolPlatform/kunman/middleware/good/device/poster"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) GetPoster(ctx context.Context) (*postermwpb.Poster, error) {
	handler, err := postermw.NewHandler(
		ctx,
		postermw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetPoster(ctx)
}

func (h *Handler) GetPosters(ctx context.Context) ([]*postermwpb.Poster, uint32, error) {
	conds := &postermwpb.Conds{}
	if h.DeviceTypeID != nil {
		conds.DeviceTypeID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.DeviceTypeID}
	}

	handler, err := postermw.NewHandler(
		ctx,
		postermw.WithConds(conds),
		postermw.WithOffset(h.Offset),
		postermw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetPosters(ctx)
}
