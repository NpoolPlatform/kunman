package brand

import (
	"context"
	"time"

	brandcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/vender/brand"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteBrand(ctx context.Context, cli *ent.Client) error {
	if _, err := brandcrud.UpdateSet(
		cli.VendorBrand.UpdateOneID(*h.ID),
		&brandcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteBrand(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetBrand(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteBrand(_ctx, cli)
	})
}
