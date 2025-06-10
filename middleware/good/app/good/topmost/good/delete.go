package topmostgood

import (
	"context"
	"time"

	topmostgoodcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteTopMostGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostgoodcrud.UpdateSet(
		tx.TopMostGood.UpdateOneID(*h.ID),
		&topmostgoodcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteTopMostGood(ctx context.Context) error {
	info, err := h.GetTopMostGood(ctx)
	if err != nil {
		return err
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
		return handler.deleteTopMostGood(_ctx, tx)
	})
}
