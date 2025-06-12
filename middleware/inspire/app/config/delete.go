package config

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	appconfigcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/app/config"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) DeleteAppConfig(ctx context.Context) error {
	info, err := h.GetAppConfig(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}
	h.ID = &info.ID

	now := uint32(time.Now().Unix())
	return db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := appconfigcrud.UpdateSet(
			tx.AppConfig.UpdateOneID(*h.ID),
			&appconfigcrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
