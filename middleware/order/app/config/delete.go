package appconfig

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appconfigcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/app/config"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteAppConfig(ctx context.Context, tx *ent.Tx) error {
	if _, err := appconfigcrud.UpdateSet(
		tx.AppConfig.UpdateOneID(*h.ID),
		&appconfigcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteAppConfig(ctx context.Context) error {
	info, err := h.GetAppConfig(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.deleteAppConfig(ctx, tx)
	})
}
