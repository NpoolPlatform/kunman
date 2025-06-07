package description

import (
	"context"
	"time"

	appgooddescriptioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/description"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteDescription(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddescriptioncrud.UpdateSet(
		cli.AppGoodDescription.UpdateOneID(*h.ID),
		&appgooddescriptioncrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteDescription(ctx context.Context) error {
	info, err := h.GetDescription(ctx)
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
		return handler.deleteDescription(_ctx, cli)
	})
}
