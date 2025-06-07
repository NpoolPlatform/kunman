//nolint:dupl
package quota

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
)

func (h *Handler) ExistQuota(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryQuota(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistQuotaConds(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err = handler.queryQuotas(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
