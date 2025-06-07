package stock

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	stockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/stock"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) _deleteStock(ctx context.Context, cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid stockid")
	}

	stm := cli.Stock.Query()
	if h.ID != nil {
		stm.Where(entstock.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entstock.EntID(*h.EntID))
	}
	if h.GoodID != nil {
		stm.Where(entstock.GoodID(*h.GoodID))
	}
	info, err := stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return wlog.WrapError(err)
	}

	if _, err := stockcrud.UpdateSet(
		info.Update(),
		&stockcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

// Only for test. Stock should always be deleted with good
func (h *Handler) DeleteStock(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler._deleteStock(ctx, cli)
	})
}
