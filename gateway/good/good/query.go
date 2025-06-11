package good

import (
	"context"

	goodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good"
	goodmw "github.com/NpoolPlatform/kunman/middleware/good/good"
)

func (h *Handler) GetGoods(ctx context.Context) ([]*goodmwpb.Good, uint32, error) {
	handler, err := goodmw.NewHandler(
		ctx,
		goodmw.WithConds(&goodmwpb.Conds{}),
		goodmw.WithOffset(h.Offset),
		goodmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetGoods(ctx)
}
