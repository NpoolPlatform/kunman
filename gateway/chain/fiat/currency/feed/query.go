package currencyfeed

import (
	"context"

	feedmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency/feed"
	feedmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat/currency/feed"
)

func (h *Handler) GetFeeds(ctx context.Context) ([]*feedmwpb.Feed, uint32, error) {
	handler, err := feedmw.NewHandler(
		ctx,
		feedmw.WithConds(&feedmwpb.Conds{}),
		feedmw.WithOffset(h.Offset),
		feedmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	return handler.GetFeeds(ctx)
}
