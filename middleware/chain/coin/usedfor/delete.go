package coinusedfor

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/usedfor"
	coinusedforcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/usedfor"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
)

func (h *Handler) DeleteCoinUsedFor(ctx context.Context) (*npool.CoinUsedFor, error) {
	info, err := h.GetCoinUsedFor(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := coinusedforcrud.UpdateSet(
			cli.CoinUsedFor.UpdateOneID(info.ID),
			&coinusedforcrud.Req{
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
