package currencyfeed

import (
	"context"

	feedmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency/feed"
	feedmw "github.com/NpoolPlatform/kunman/middleware/chain/coin/currency/feed"
)

func (h *Handler) UpdateFeed(ctx context.Context) (*feedmwpb.Feed, error) {
	handler, err := feedmw.NewHandler(
		ctx,
		feedmw.WithID(h.ID, true),
		feedmw.WithFeedCoinName(h.FeedCoinName, true),
		feedmw.WithDisabled(h.Disabled, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.UpdateFeed(ctx)
}
