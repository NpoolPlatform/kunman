package appstock

import (
	"context"

	appgoodstockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/stock"
	appmininggoodstockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/stock/mining"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type createHandler struct {
	*stockAppGoodQuery
}

func (h *createHandler) _createStock(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodstockcrud.CreateSet(
		tx.AppStock.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppGoodMiningStocks(ctx context.Context, tx *ent.Tx) error {
	for _, req := range h.AppMiningGoodStockReqs {
		if _, err := appmininggoodstockcrud.CreateSet(
			tx.AppMiningGoodStock.Create(),
			&appmininggoodstockcrud.Req{
				EntID:             req.EntID,
				AppGoodStockID:    h.EntID,
				MiningGoodStockID: req.MiningGoodStockID,
				Reserved:          req.Reserved,
				SpotQuantity:      req.Reserved,
			},
		).Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Only for test. Stock should always be created with good
func (h *Handler) createStock(ctx context.Context) error {
	handler := &createHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
	}
	if err := handler.getStockGoods(ctx); err != nil {
		return err
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.createAppGoodMiningStocks(_ctx, tx); err != nil {
				return err
			}
		}
		return handler._createStock(ctx, tx)
	})
}
