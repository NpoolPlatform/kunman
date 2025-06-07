package appsimulatepowerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

func (h *Handler) ExistSimulate(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySimulate(cli); err != nil {
			return err
		}
		handler.queryJoin()
		count, err := handler.stmSelect.Limit(1).Count(_ctx)
		exist = count > 0
		return err
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistSimulateConds(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.querySimulates(cli); err != nil {
			return err
		}
		handler.queryJoin()
		count, err := handler.stmSelect.Limit(1).Count(_ctx)
		exist = count > 0
		return err
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
