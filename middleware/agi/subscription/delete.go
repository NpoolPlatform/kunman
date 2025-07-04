package subscription

import (
	"context"
	"time"

	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/subscription"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteSubscription(ctx context.Context, tx *ent.Tx) error {
	if _, err := subscriptioncrud.UpdateSet(
		tx.Subscription.UpdateOneID(*h.ID),
		&subscriptioncrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteSubscription(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetSubscription(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.deleteSubscription(_ctx, tx)
	})
}

func (h *Handler) DeleteSubscriptionWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetSubscription(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return handler.deleteSubscription(ctx, tx)
}
