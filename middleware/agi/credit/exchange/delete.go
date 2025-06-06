package exchange

import (
	"context"
	"time"

	exchangecrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/credit/exchange"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteExchange(ctx context.Context, cli *ent.Client) error {
	if _, err := exchangecrud.UpdateSet(
		cli.Exchange.UpdateOneID(*h.ID),
		&exchangecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteExchange(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetExchange(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteExchange(_ctx, cli)
	})
}
