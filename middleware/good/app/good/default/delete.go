package appdefaultgood

import (
	"context"
	"time"

	appdefaultgoodcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/default"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteDefault(ctx context.Context, cli *ent.Client) error {
	if _, err := appdefaultgoodcrud.UpdateSet(
		cli.AppDefaultGood.UpdateOneID(*h.ID),
		&appdefaultgoodcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteDefault(ctx context.Context) error {
	info, err := h.GetDefault(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	h.ID = &info.ID

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteDefault(_ctx, cli)
	})
}
