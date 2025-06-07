package topmost

import (
	"context"
	"time"

	topmostcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteTopMost(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := topmostcrud.UpdateSet(
		tx.TopMost.UpdateOneID(*h.ID),
		&topmostcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteTopMost(ctx context.Context) error {
	info, err := h.GetTopMost(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.deleteTopMost(_ctx, tx)
	})
}
