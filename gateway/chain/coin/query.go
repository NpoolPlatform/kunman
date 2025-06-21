package coin

import (
	"context"

	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
)

func (h *Handler) GetCoins(ctx context.Context) ([]*coinmwpb.Coin, uint32, error) {
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithConds(&coinmwpb.Conds{}),
		coinmw.WithOffset(h.Offset),
		coinmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetCoins(ctx)
}
