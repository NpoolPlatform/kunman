package appdefaultgood

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

func (h *Handler) ExistDefault(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDefault(cli); err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(ctx)
		return err
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistDefaultConds(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryDefaults(cli); err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(ctx)
		return err
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
