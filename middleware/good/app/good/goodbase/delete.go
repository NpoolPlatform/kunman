package goodbase

import (
	"context"
	"time"

	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteGoodBase(ctx context.Context, cli *ent.Client) error {
	if _, err := appgoodbasecrud.UpdateSet(
		cli.AppGoodBase.UpdateOneID(*h.ID),
		&appgoodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteGoodBase(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	info, err := h.GetGoodBase(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = func() *uint32 { id := info.ID(); return &id }()
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteGoodBase(_ctx, cli)
	})
}
