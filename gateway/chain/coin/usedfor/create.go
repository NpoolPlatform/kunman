package coinusedfor

import (
	"context"

	coinusedformwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/usedfor"
	coinusedformw "github.com/NpoolPlatform/kunman/middleware/chain/coin/usedfor"
)

func (h *Handler) CreateCoinUsedFor(ctx context.Context) (*coinusedformwpb.CoinUsedFor, error) {
	handler, err := coinusedformw.NewHandler(
		ctx,
		coinusedformw.WithCoinTypeID(h.CoinTypeID, true),
		coinusedformw.WithUsedFor(h.UsedFor, true),
		coinusedformw.WithPriority(h.Priority, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.CreateCoinUsedFor(ctx)
}
