package oneshot

import (
	"context"

	oneshotmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription/oneshot"
	oneshotmw "github.com/NpoolPlatform/kunman/middleware/good/subscription/oneshot"
)

func (h *Handler) GetOneShot(ctx context.Context) (*oneshotmwpb.OneShot, error) {
	handler, err := oneshotmw.NewHandler(
		ctx,
		oneshotmw.WithGoodID(h.GoodID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetOneShot(ctx)
}

func (h *Handler) GetOneShots(ctx context.Context) ([]*oneshotmwpb.OneShot, error) {
	handler, err := oneshotmw.NewHandler(
		ctx,
		oneshotmw.WithConds(&oneshotmwpb.Conds{}),
		oneshotmw.WithOffset(h.Offset),
		oneshotmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, err
	}

	return handler.GetOneShots(ctx)
}
