package comment

import (
	"context"
	"time"

	commentcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/comment"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteComment(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := commentcrud.UpdateSet(
		tx.Comment.UpdateOneID(*h.ID),
		&commentcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteComment(ctx context.Context) error {
	info, err := h.GetComment(ctx)
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
		return handler.deleteComment(ctx, tx)
	})
}
