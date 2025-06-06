package change

import (
	"context"
	"time"

	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/user/subscription/change"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteSubscriptionChange(ctx context.Context, cli *ent.Client) error {
	if _, err := subscriptioncrud.UpdateSet(
		cli.UserSubscriptionChange.UpdateOneID(*h.ID),
		&subscriptioncrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteSubscriptionChange(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetSubscriptionChange(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteSubscriptionChange(_ctx, cli)
	})
}
