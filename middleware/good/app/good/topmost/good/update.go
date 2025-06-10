package topmostgood

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	topmostgoodcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateTopMostGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostgoodcrud.UpdateSet(
		tx.TopMostGood.UpdateOneID(*h.ID),
		&topmostgoodcrud.Req{
			DisplayIndex: h.DisplayIndex,
			UnitPrice:    h.UnitPrice,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateTopMostGood(ctx context.Context) error {
	info, err := h.GetTopMostGood(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid topmostgood")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateTopMostGood(_ctx, tx)
	})
}
