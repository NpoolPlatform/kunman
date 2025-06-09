package apppool

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	apppoolcrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/app/pool"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	entapppool "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/apppool"
)

func (h *Handler) ExistPool(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		count, err := cli.
			AppPool.
			Query().
			Where(
				entapppool.EntID(*h.EntID),
				entapppool.DeletedAt(0),
			).
			Limit(1).
			Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}

		exist = count > 0

		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistPoolConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := apppoolcrud.SetQueryConds(cli.AppPool.Query(), h.Conds)
		if err != nil {
			return wlog.WrapError(err)
		}
		count, err := stm.Limit(1).Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}

		exist = count > 0

		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
