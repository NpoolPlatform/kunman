package description

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgooddescriptioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/description"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateDescription(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddescriptioncrud.UpdateSet(
		cli.AppGoodDescription.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateDescription(ctx context.Context) error {
	info, err := h.GetDescription(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid description")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updateDescription(_ctx, cli)
	})
}
