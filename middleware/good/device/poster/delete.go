package poster

import (
	"context"
	"time"

	devicepostercrud "github.com/NpoolPlatform/kunman/middleware/good/crud/device/poster"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deletePoster(ctx context.Context, cli *ent.Client) error {
	if _, err := devicepostercrud.UpdateSet(
		cli.DevicePoster.UpdateOneID(*h.ID),
		&devicepostercrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeletePoster(ctx context.Context) error {
	info, err := h.GetPoster(ctx)
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
		return handler.deletePoster(_ctx, cli)
	})
}
