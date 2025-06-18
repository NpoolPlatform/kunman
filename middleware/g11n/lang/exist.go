package lang

import (
	"context"
	"fmt"

	langcrud "github.com/NpoolPlatform/kunman/middleware/g11n/crud/lang"

	"github.com/NpoolPlatform/kunman/middleware/g11n/db"
	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
	entlang "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/lang"
)

func (h *Handler) ExistLang(ctx context.Context) (bool, error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Lang.
			Query().
			Where(
				entlang.EntID(*h.EntID),
				entlang.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistLangCondsWithClient(ctx context.Context, cli *ent.Client) (exist bool, err error) {
	stm, err := langcrud.SetQueryConds(cli.Lang.Query(), h.Conds)
	if err != nil {
		return false, err
	}
	if exist, err = stm.Exist(ctx); err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistLangConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = h.ExistLangCondsWithClient(_ctx, cli)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
