//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type inServiceHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *inServiceHandler) constructSQL(table string, lock *ent.AppStockLock, checkTotal bool, id uint32) string {
	sql := fmt.Sprintf(
		`update %v
		set
		  in_service = in_service + %v,
		  wait_start = wait_start - %v`,
		table,
		lock.Units,
		lock.Units,
	)
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  wait_start >= %v`,
		id,
		lock.Units,
	)
	if checkTotal {
		sql += ` and
		  in_service + wait_start + locked + app_reserved + spot_quantity = total`
	}
	return sql
}

func (h *inServiceHandler) inServiceStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructSQL("stocks_v1", lock, true, stock.stock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *inServiceHandler) inServiceMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.miningGoodStocks == nil {
		return wlog.Errorf("invalid stock")
	}
	appMiningGoodStock, ok := stock.appMiningGoodStocks[lock.AppStockID]
	if !ok {
		return wlog.Errorf("invalid appmininggoodstock")
	}
	miningGoodStock, ok := stock.miningGoodStocks[appMiningGoodStock.MiningGoodStockID]
	if !ok {
		return wlog.Errorf("invalid mininggoodstock")
	}
	sql := h.constructSQL("mining_good_stocks", lock, true, miningGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *inServiceHandler) inServiceAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructSQL("app_stocks", lock, false, stock.appGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *inServiceHandler) inServiceAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appMiningGoodStocks == nil {
		return wlog.Errorf("invalid stock")
	}
	appMiningGoodStock, ok := stock.appMiningGoodStocks[lock.AppStockID]
	if !ok {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructSQL("app_mining_good_stocks", lock, false, appMiningGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *Handler) InServiceStock(ctx context.Context) error {
	handler := &inServiceHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockInService.Enum(),
		},
	}

	if err := handler.lockOp.getLocks(ctx); err != nil {
		return wlog.WrapError(err)
	}
	h.Stocks = handler.lockOp.lock2Stocks()
	if err := handler.getStockAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockOp.locks {
			if err := handler.inServiceAppStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.inServiceAppMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.inServiceStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.inServiceMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
		}
		if err := handler.lockOp.updateLocks(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
