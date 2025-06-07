package poster

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	topmostpostercrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost/poster"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updatePoster(ctx context.Context, cli *ent.Client) error {
	if _, err := topmostpostercrud.UpdateSet(
		cli.TopMostPoster.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdatePoster(ctx context.Context) error {
	info, err := h.GetPoster(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid poster")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updatePoster(_ctx, cli)
	})
}
