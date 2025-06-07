package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
)

type existHandler struct {
	*baseQueryHandler
}

func (h *Handler) ExistCompensate(ctx context.Context) (exist bool, err error) {
	handler := &existHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCompensate(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		count, err := handler.stmSelect.Limit(1).Count(_ctx)
		exist = count > 0
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistCompensateConds(ctx context.Context) (exist bool, err error) {
	handler := &existHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCompensates(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		count, err := handler.stmSelect.Limit(1).Count(_ctx)
		exist = count > 0
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
