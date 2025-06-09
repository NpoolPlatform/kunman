package rootuser

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	rootusercrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/rootuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	rootuserent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/rootuser"
)

func (h *Handler) ExistRootUser(ctx context.Context) (bool, error) {
	if h.EntID == nil {
		return false, wlog.Errorf("invalid entid")
	}
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			RootUser.
			Query().
			Where(
				rootuserent.EntID(*h.EntID),
				rootuserent.DeletedAt(0),
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

func (h *Handler) ExistRootUserConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := rootusercrud.SetQueryConds(cli.RootUser.Query(), h.Conds)
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
