package displayname

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgooddisplaynamecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/display/name"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateDisplayName(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddisplaynamecrud.UpdateSet(
		cli.AppGoodDisplayName.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateDisplayName(ctx context.Context) error {
	info, err := h.GetDisplayName(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid displayname")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updateDisplayName(_ctx, cli)
	})
}
