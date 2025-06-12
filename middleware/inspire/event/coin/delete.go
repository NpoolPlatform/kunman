package coin

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	eventcoincrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event/coin"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteEventCoin(ctx context.Context, cli *ent.Client) error {
	if _, err := eventcoincrud.UpdateSet(
		cli.EventCoin.UpdateOneID(*h.ID),
		&eventcoincrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteEventCoin(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetEventCoin(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteEventCoin(_ctx, cli)
	})
}
