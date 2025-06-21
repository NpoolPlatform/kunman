package coinfiat

import (
	"context"
	"fmt"

	coinfiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat"
	coinfiatmw "github.com/NpoolPlatform/kunman/middleware/chain/coin/fiat"
)

func (h *Handler) DeleteCoinFiat(ctx context.Context) (*coinfiatmwpb.CoinFiat, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid coinfiatid")
	}

	handler, err := coinfiatmw.NewHandler(
		ctx,
		coinfiatmw.WithID(h.ID, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.DeleteCoinFiat(ctx)
}
