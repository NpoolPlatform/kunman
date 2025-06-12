package allocated

import (
	"context"
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	allocatedcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coin/allocated"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteCoinAllocated(ctx context.Context, cli *ent.Client) error {
	if _, err := allocatedcrud.UpdateSet(
		cli.CoinAllocated.UpdateOneID(*h.ID),
		&allocatedcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteCoinAllocated(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetCoinAllocated(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteCoinAllocated(_ctx, cli)
	})
}
