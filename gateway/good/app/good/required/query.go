package required

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	requiredappgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) GetRequired(ctx context.Context) (*requiredappgoodmwpb.Required, error) {
	handler, err := requiredappgoodmw.NewHandler(
		ctx,
		requiredappgoodmw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetRequired(ctx)
}

func (h *Handler) GetRequireds(ctx context.Context) ([]*requiredappgoodmwpb.Required, uint32, error) {
	conds := &requiredappgoodmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}

	handler, err := requiredappgoodmw.NewHandler(
		ctx,
		requiredappgoodmw.WithConds(conds),
		requiredappgoodmw.WithOffset(h.Offset),
		requiredappgoodmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetRequireds(ctx)
}
