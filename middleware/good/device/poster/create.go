package poster

import (
	"context"

	devicepostercrud "github.com/NpoolPlatform/kunman/middleware/good/crud/device/poster"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createPoster(ctx context.Context, cli *ent.Client) error {
	if _, err := devicepostercrud.CreateSet(
		cli.DevicePoster.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreatePoster(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.createPoster(_ctx, cli)
	})
}
