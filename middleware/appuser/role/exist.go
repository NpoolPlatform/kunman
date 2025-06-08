package role

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	rolecrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/role"
	entapprole "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/approle"
)

func (h *Handler) ExistRole(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			AppRole.
			Query().
			Where(
				entapprole.EntID(*h.EntID),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistRoleConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := rolecrud.SetQueryConds(cli.AppRole.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
