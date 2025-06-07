package displaycolor

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

func (h *Handler) ExistDisplayColor(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDisplayColor(cli); err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return err
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistDisplayColorConds(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryDisplayColors(cli); err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return err
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
