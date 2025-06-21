package currencyfeed

import (
	"context"

	feedmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency/feed"
	feedmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat/currency/feed"
)

func (h *Handler) CreateFeed(ctx context.Context) (*feedmwpb.Feed, error) {
	handler, err := feedmw.NewHandler(
		ctx,
		feedmw.WithFiatID(h.FiatID, true),
		feedmw.WithFeedType(h.FeedType, true),
		feedmw.WithFeedFiatName(h.FeedFiatName, true),
	)
	if err != nil {
		return nil, err
	}

	return handler.CreateFeed(ctx)
}
