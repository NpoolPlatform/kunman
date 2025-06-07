package quota

import (
	"context"
	"time"

	quotacrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/subscription/quota"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteQuota(ctx context.Context, tx *ent.Tx) error {
	if _, err := quotacrud.UpdateSet(
		tx.Quota.UpdateOneID(*h.ID),
		&quotacrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteQuota(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetQuota(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.deleteQuota(_ctx, tx)
	})
}

func (h *Handler) DeleteQuotaWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetQuota(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return handler.deleteQuota(ctx, tx)
}
