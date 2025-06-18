package applang

import (
	"context"

	applangcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/applang"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
)

func (h *Handler) ExistAppLangCondsWithClient(ctx context.Context, cli *ent.Client) (exist bool, err error) {
	stm, err := applangcrud.SetQueryConds(cli.AppLang.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	if exist, err = stm.Exist(ctx); err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistAppLangConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = h.ExistAppLangCondsWithClient(_ctx, cli)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
