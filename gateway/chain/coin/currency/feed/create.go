package currencyfeed

import (
	"context"

	feedmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency/feed"
	feedmw "github.com/NpoolPlatform/kunman/middleware/chain/coin/currency/feed"
)

func (h *Handler) CreateFeed(ctx context.Context) (*feedmwpb.Feed, error) {
	handler, err := feedmw.NewHandler(
		ctx,
		feedmw.WithCoinTypeID(h.CoinTypeID, true),
		feedmw.WithFeedType(h.FeedType, true),
		feedmw.WithFeedCoinName(h.FeedCoinName, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.CreateFeed(ctx)
}
