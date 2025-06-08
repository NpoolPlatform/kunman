package currencyfeed

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency/feed"
	currencyfeedcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/currency/feed"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
)

func (h *Handler) UpdateFeed(ctx context.Context) (*npool.Feed, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := currencyfeedcrud.UpdateSet(
			cli.CurrencyFeed.UpdateOneID(*h.ID),
			&currencyfeedcrud.Req{
				FeedCoinName: h.FeedCoinName,
				Disabled:     h.Disabled,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetFeed(ctx)
}
