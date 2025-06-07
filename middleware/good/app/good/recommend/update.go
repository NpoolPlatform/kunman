package recommend

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	recommendcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/recommend"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateRecommend(ctx context.Context, tx *ent.Tx) error {
	if _, err := recommendcrud.UpdateSet(
		tx.Recommend.UpdateOneID(*h.ID),
		&recommendcrud.Req{
			Message:        h.Message,
			RecommendIndex: h.RecommendIndex,
			Hide:           h.Hide,
			HideReason:     h.HideReason,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateRecommend(ctx context.Context) error {
	info, err := h.GetRecommend(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid recommend")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateRecommend(ctx, tx)
	})
}
