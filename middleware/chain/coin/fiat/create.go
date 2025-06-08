package coinfiat

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat"
	coinfiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/fiat"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) CreateCoinFiat(ctx context.Context) (*npool.CoinFiat, error) {
	// TODO: deduplicate

	h.Conds = &coinfiatcrud.Conds{
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		FiatID:     &cruder.Cond{Op: cruder.EQ, Val: *h.FiatID},
	}
	exist, err := h.ExistCoinConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coinfiat exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := coinfiatcrud.CreateSet(
			cli.CoinFiat.Create(),
			&coinfiatcrud.Req{
				CoinTypeID: h.CoinTypeID,
				FiatID:     h.FiatID,
				FeedType:   h.FeedType,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCoinFiat(ctx)
}
