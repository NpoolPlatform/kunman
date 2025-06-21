package coinusedfor

import (
	"context"
	"fmt"

	coinusedformwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/usedfor"
	coinusedformw "github.com/NpoolPlatform/kunman/middleware/chain/coin/usedfor"
)

func (h *Handler) DeleteCoinUsedFor(ctx context.Context) (*coinusedformwpb.CoinUsedFor, error) {
	if h.ID == nil || h.EntID == nil {
		return nil, fmt.Errorf("invalid coinusedforid")
	}

	handler, err := coinusedformw.NewHandler(
		ctx,
		coinusedformw.WithID(h.ID, true),
		coinusedformw.WithEntID(h.EntID, true),
	)
	if err != nil {
		return nil, err
	}

	info, err := handler.GetCoinUsedFor(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid coinusedfor")
	}
	if info.ID != *h.ID {
		return nil, fmt.Errorf("invalid coinusedforid")
	}
	return handler.DeleteCoinUsedFor(ctx)
}
