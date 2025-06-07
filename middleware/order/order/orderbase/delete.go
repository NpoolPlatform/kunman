package orderbase

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	orderbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/orderbase"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteOrderBase(ctx context.Context, cli *ent.Client) error {
	if _, err := orderbasecrud.UpdateSet(
		cli.OrderBase.UpdateOneID(*h.ID),
		&orderbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteOrderBase(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	info, err := h.GetOrderBase(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = func() *uint32 { id := info.ID(); return &id }()
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteOrderBase(_ctx, cli)
	})
}
