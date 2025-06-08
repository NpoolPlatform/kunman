package coinusedfor

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/usedfor"
	coinusedforcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/usedfor"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

func (h *Handler) CreateCoinUsedFor(ctx context.Context) (*npool.CoinUsedFor, error) {
	// TODO: deduplicate

	h.Conds = &coinusedforcrud.Conds{
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
		UsedFor:    &cruder.Cond{Op: cruder.EQ, Val: *h.UsedFor},
	}
	exist, err := h.ExistCoinUsedForConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("coinusedfor exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := coinusedforcrud.CreateSet(
			cli.CoinUsedFor.Create(),
			&coinusedforcrud.Req{
				CoinTypeID: h.CoinTypeID,
				UsedFor:    h.UsedFor,
				Priority:   h.Priority,
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

	return h.GetCoinUsedFor(ctx)
}
