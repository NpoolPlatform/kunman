package comment

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	commentcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/comment"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateComment(ctx context.Context, tx *ent.Tx) error {
	if _, err := commentcrud.UpdateSet(
		tx.Comment.UpdateOneID(*h.ID),
		&commentcrud.Req{
			Content:    h.Content,
			Anonymous:  h.Anonymous,
			Hide:       h.Hide,
			HideReason: h.HideReason,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateComment(ctx context.Context) error {
	info, err := h.GetComment(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid comment")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateComment(ctx, tx)
	})
}
