package coupon

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	couponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event/coupon"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteEventCoupon(ctx context.Context, cli *ent.Client) error {
	if _, err := couponcrud.UpdateSet(
		cli.EventCoupon.UpdateOneID(*h.ID),
		&couponcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteEventCoupon(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetEventCoupon(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteEventCoupon(_ctx, cli)
	})
}
