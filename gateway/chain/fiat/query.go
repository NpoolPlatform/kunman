package fiat

import (
	"context"

	fiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	fiatmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
)

func (h *Handler) GetFiats(ctx context.Context) ([]*fiatmwpb.Fiat, uint32, error) {
	handler, err := fiatmw.NewHandler(
		ctx,
		fiatmw.WithConds(&fiatmwpb.Conds{}),
		fiatmw.WithOffset(h.Offset),
		fiatmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetFiats(ctx)
}
