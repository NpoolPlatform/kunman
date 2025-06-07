package displaycolor

import (
	"context"
	"time"

	appgooddisplaycolorcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/display/color"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteDisplayColor(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddisplaycolorcrud.UpdateSet(
		cli.AppGoodDisplayColor.UpdateOneID(*h.ID),
		&appgooddisplaycolorcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteDisplayColor(ctx context.Context) error {
	info, err := h.GetDisplayColor(ctx)
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
		return handler.deleteDisplayColor(_ctx, cli)
	})
}
