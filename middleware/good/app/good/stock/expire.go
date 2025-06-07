//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
)

type expireHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *expireHandler) constructSQL(table string, lock *ent.AppStockLock, returnSpotQuantity, checkTotal bool, id uint32) string {
	sql := fmt.Sprintf(
		`update %v
		set
		  in_service = in_service - %v`,
		table,
		lock.Units,
	)
	if returnSpotQuantity {
		sql += fmt.Sprintf(
			`, spot_quantity = spot_quantity + %v`,
			lock.Units,
		)
	}
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  in_service >= %v`,
		id,
		lock.Units,
	)
	if checkTotal {
		sql += ` and
		  in_service + wait_start + locked + app_reserved + spot_quantity = total`
	}
	return sql
}

func (h *expireHandler) expireStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructSQL("stocks_v1", lock, true, true, stock.stock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *expireHandler) expireMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
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
	sql := h.constructSQL("mining_good_stocks", lock, true, true, miningGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *expireHandler) expireAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructSQL("app_stocks", lock, false, false, stock.appGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *expireHandler) expireAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appMiningGoodStocks == nil {
		return wlog.Errorf("invalid stock")
	}
	appMiningGoodStock, ok := stock.appMiningGoodStocks[lock.AppStockID]
	if !ok {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructSQL("app_mining_good_stocks", lock, false, false, appMiningGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *Handler) ExpireStock(ctx context.Context) error {
	handler := &expireHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockExpired.Enum(),
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
			if err := handler.expireAppStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.expireAppMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.expireStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.expireMiningGoodStock(ctx, lock, tx); err != nil {
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
