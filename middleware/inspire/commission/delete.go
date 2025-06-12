package commission

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	commissioncrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/commission"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

func (h *Handler) DeleteCommission(ctx context.Context) error {
	info, err := h.GetCommission(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	now := uint32(time.Now().Unix())
	return db.WithClient(ctx, func(_ctx context.Context, tx *ent.Client) error {
		if _, err := commissioncrud.UpdateSet(
			tx.Commission.UpdateOneID(*h.ID),
			&commissioncrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
