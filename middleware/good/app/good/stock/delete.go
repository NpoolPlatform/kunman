package appstock

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodstockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/stock"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appstock"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) _deleteStock(ctx context.Context, cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid stockid")
	}

	stm := cli.AppStock.Query()
	if h.ID != nil {
		stm.Where(entappgoodstock.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgoodstock.EntID(*h.EntID))
	}
	if h.AppGoodID != nil {
		stm.Where(entappgoodstock.AppGoodID(*h.AppGoodID))
	}
	info, err := stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return wlog.WrapError(err)
	}

	if _, err := appgoodstockcrud.UpdateSet(
		info.Update(),
		&appgoodstockcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

// Only for test. Stock should always be deleted with good
func (h *Handler) deleteStock(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler._deleteStock(ctx, cli)
	})
}
