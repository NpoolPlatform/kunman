package label

import (
	"context"
	"time"

	appgoodlabelcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/label"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteLabel(ctx context.Context, cli *ent.Client) error {
	if _, err := appgoodlabelcrud.UpdateSet(
		cli.AppGoodLabel.UpdateOneID(*h.ID),
		&appgoodlabelcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteLabel(ctx context.Context) error {
	info, err := h.GetLabel(ctx)
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
		return handler.deleteLabel(_ctx, cli)
	})
}
