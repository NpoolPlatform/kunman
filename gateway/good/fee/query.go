package fee

import (
	"context"

	feemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
	feemw "github.com/NpoolPlatform/kunman/middleware/good/fee"
)

func (h *Handler) GetFee(ctx context.Context) (*feemwpb.Fee, error) {
	handler, err := feemw.NewHandler(
		ctx,
		feemw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetFee(ctx)
}

func (h *Handler) GetFees(ctx context.Context) ([]*feemwpb.Fee, uint32, error) {
	handler, err := feemw.NewHandler(
		ctx,
		feemw.WithConds(&feemwpb.Conds{}),
		feemw.WithOffset(h.Offset),
		feemw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetFees(ctx)
}
