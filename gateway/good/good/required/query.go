package required

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	requiredmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"
	requiredmw "github.com/NpoolPlatform/kunman/middleware/good/good/required"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) GetRequired(ctx context.Context) (*requiredmwpb.Required, error) {
	handler, err := requiredmw.NewHandler(
		ctx,
		requiredmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetRequired(ctx)
}

func (h *Handler) GetRequireds(ctx context.Context) ([]*requiredmwpb.Required, uint32, error) {
	conds := &requiredmwpb.Conds{}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	handler, err := requiredmw.NewHandler(
		ctx,
		requiredmw.WithConds(conds),
		requiredmw.WithOffset(h.Offset),
		requiredmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetRequireds(ctx)
}
