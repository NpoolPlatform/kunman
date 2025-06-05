package record

import (
	"context"
	"time"

	recordcrud "github.com/NpoolPlatform/kunman/middleware/billing/crud/user/credit/record"
	"github.com/NpoolPlatform/kunman/middleware/billing/db"
	"github.com/NpoolPlatform/kunman/middleware/billing/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteRecord(ctx context.Context, cli *ent.Client) error {
	if _, err := recordcrud.UpdateSet(
		cli.UserCreditRecord.UpdateOneID(*h.ID),
		&recordcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteRecord(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetRecord(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteRecord(_ctx, cli)
	})
}
