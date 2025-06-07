package displaycolor

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgooddisplaycolorcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/display/color"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateDisplayColor(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddisplaycolorcrud.UpdateSet(
		cli.AppGoodDisplayColor.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateDisplayColor(ctx context.Context) error {
	info, err := h.GetDisplayColor(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid displaycolor")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updateDisplayColor(_ctx, cli)
	})
}
