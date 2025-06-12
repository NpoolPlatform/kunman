package config

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	commissionconfigcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/app/commission/config"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) DeleteCommissionConfig(ctx context.Context) error {
	info, err := h.GetCommissionConfig(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}
	h.ID = &info.ID

	now := uint32(time.Now().Unix())
	return db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissionconfigcrud.UpdateSet(
			tx.AppCommissionConfig.UpdateOneID(*h.ID),
			&commissionconfigcrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
