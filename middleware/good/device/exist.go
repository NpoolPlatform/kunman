package device

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

func (h *Handler) ExistDeviceType(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDeviceType(cli); err != nil {
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

func (h *Handler) ExistDeviceTypeConds(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryDeviceTypes(cli); err != nil {
			return wlog.WrapError(err)
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
