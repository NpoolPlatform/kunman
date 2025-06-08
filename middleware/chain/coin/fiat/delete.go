package coinfiat

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat"
	coinfiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/fiat"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
)

func (h *Handler) DeleteCoinFiat(ctx context.Context) (*npool.CoinFiat, error) {
	info, err := h.GetCoinFiat(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := coinfiatcrud.UpdateSet(
			cli.CoinFiat.UpdateOneID(*h.ID),
			&coinfiatcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
