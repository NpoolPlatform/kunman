package gooduser

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodusercrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/gooduser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	gooduserent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/gooduser"
)

func (h *Handler) ExistGoodUser(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			GoodUser.
			Query().
			Where(
				gooduserent.EntID(*h.EntID),
				gooduserent.DeletedAt(0),
			).
			Exist(_ctx)
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

func (h *Handler) ExistGoodUserConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := goodusercrud.SetQueryConds(cli.GoodUser.Query(), h.Conds)
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
