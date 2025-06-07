package location

import (
	"context"
	"time"

	locationcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/vender/location"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteLocation(ctx context.Context, cli *ent.Client) error {
	if _, err := locationcrud.UpdateSet(
		cli.VendorLocation.UpdateOneID(*h.ID),
		&locationcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteLocation(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	info, err := h.GetLocation(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}
	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteLocation(_ctx, cli)
	})
}
