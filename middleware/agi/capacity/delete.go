package capacity

import (
	"context"
	"time"

	capacitycrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/capacity"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteCapacity(ctx context.Context, tx *ent.Tx) error {
	if _, err := capacitycrud.UpdateSet(
		tx.Capacity.UpdateOneID(*h.ID),
		&capacitycrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteCapacity(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetCapacity(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.deleteCapacity(_ctx, tx)
	})
}

func (h *Handler) DeleteCapacityWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetCapacity(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return handler.deleteCapacity(ctx, tx)
}
