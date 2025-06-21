package coinfiat

import (
	"context"

	coinfiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat"
	coinfiatmw "github.com/NpoolPlatform/kunman/middleware/chain/coin/fiat"
)

func (h *Handler) CreateCoinFiat(ctx context.Context) (*coinfiatmwpb.CoinFiat, error) {
	handler, err := coinfiatmw.NewHandler(
		ctx,
		coinfiatmw.WithCoinTypeID(h.CoinTypeID, true),
		coinfiatmw.WithFiatID(h.FiatID, true),
		coinfiatmw.WithFeedType(h.FeedType, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.CreateCoinFiat(ctx)
}
