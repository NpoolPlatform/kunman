package currencyfeed

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency/feed"
	currencyfeedcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat/currency/feed"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

func (h *Handler) CreateFeed(ctx context.Context) (*npool.Feed, error) {
	if h.FeedFiatName == nil {
		return nil, fmt.Errorf("invalid feedfiatname")
	}

	// TODO: deduplicate

	h.Conds = &currencyfeedcrud.Conds{
		FiatID:   &cruder.Cond{Op: cruder.EQ, Val: *h.FiatID},
		FeedType: &cruder.Cond{Op: cruder.EQ, Val: *h.FeedType},
	}
	exist, err := h.ExistFeedConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("fiatfeed exist")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := currencyfeedcrud.CreateSet(
			cli.FiatCurrencyFeed.Create(),
			&currencyfeedcrud.Req{
				EntID:        h.EntID,
				FiatID:       h.FiatID,
				FeedType:     h.FeedType,
				FeedFiatName: h.FeedFiatName,
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
