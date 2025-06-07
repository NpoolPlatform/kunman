package displaycolor

import (
	"context"

	appgooddisplaycolorcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/display/color"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createDisplayColor(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddisplaycolorcrud.CreateSet(
		cli.AppGoodDisplayColor.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateDisplayColor(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.createDisplayColor(_ctx, cli)
	})
}
