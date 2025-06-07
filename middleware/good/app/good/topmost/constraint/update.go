package constraint

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constraintcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost/constraint"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateConstraint(ctx context.Context, tx *ent.Tx) error {
	if _, err := constraintcrud.UpdateSet(
		tx.TopMostConstraint.UpdateOneID(*h.ID),
		&constraintcrud.Req{
			TargetValue: h.TargetValue,
			Index:       h.Index,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateConstraint(ctx context.Context) error {
	info, err := h.GetConstraint(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid constraint")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateConstraint(_ctx, tx)
	})
}
