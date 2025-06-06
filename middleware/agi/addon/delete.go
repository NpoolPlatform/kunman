package addon

import (
	"context"
	"time"

	addoncrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/addon"
	"github.com/NpoolPlatform/kunman/middleware/agi/db"
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteAddon(ctx context.Context, cli *ent.Client) error {
	if _, err := addoncrud.UpdateSet(
		cli.Addon.UpdateOneID(*h.ID),
		&addoncrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteAddon(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetAddon(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteAddon(_ctx, cli)
	})
}
