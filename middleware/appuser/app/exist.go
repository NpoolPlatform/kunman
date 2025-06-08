package app

import (
	"context"

	"github.com/NpoolPlatform/kunman/middleware/appuser/db"
	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"

	appcrud "github.com/NpoolPlatform/kunman/middleware/appuser/crud/app"
	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
)

func (h *Handler) ExistApp(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			App.
			Query().
			Where(
				entapp.EntID(*h.EntID),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistAppConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appcrud.SetQueryConds(cli.App.Query(), h.Conds)
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
