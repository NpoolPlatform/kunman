package allocated

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	allocatedcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coin/allocated"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) ExistCoinAllocatedConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := allocatedcrud.SetQueryConds(
			cli.CoinAllocated.Query(),
			h.Conds,
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}

	return exist, nil
}
