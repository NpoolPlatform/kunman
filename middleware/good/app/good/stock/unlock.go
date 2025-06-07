//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/shopspring/decimal"
)

type unlockHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *unlockHandler) constructGoodStockSQL(table string, lock *ent.AppStockLock, id uint32) (string, error) {
	platformLocked := lock.Units.Sub(lock.AppSpotUnits)
	if platformLocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return "", wlog.Errorf("invalid appspotunits")
	}
	sql := fmt.Sprintf(
		`update %v
		set
		  spot_quantity = spot_quantity + %v,
		  locked = locked - %v,
		  app_reserved = app_reserved + %v`,
		table,
		platformLocked,
		lock.Units,
		lock.AppSpotUnits,
	)
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  locked >= %v`,
		id,
		lock.Units,
	)
	sql += ` and
		in_service + wait_start + locked + app_reserved + spot_quantity = total`
	return sql, nil
}

func (h *unlockHandler) constructAppGoodStockSQL(table string, lock *ent.AppStockLock, id uint32) string {
	sql := fmt.Sprintf(
		`update %v
		set
		  spot_quantity = spot_quantity + %v,
		  locked = locked - %v`,
		table,
		lock.AppSpotUnits,
		lock.Units,
	)
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  locked >= %v
		and
		  spot_quantity <= reserved`,
		id,
		lock.Units,
	)
	return sql
}

func (h *unlockHandler) unlockStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructGoodStockSQL("stocks_v1", lock, stock.stock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *unlockHandler) unlockMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
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
	sql, err := h.constructGoodStockSQL("mining_good_stocks", lock, miningGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *unlockHandler) unlockAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructAppGoodStockSQL("app_stocks", lock, stock.appGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *unlockHandler) unlockAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appMiningGoodStocks == nil {
		return wlog.Errorf("invalid stock")
	}
	appMiningGoodStock, ok := stock.appMiningGoodStocks[lock.AppStockID]
	if !ok {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructAppGoodStockSQL("app_mining_good_stocks", lock, appMiningGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

//nolint:gocyclo
func (h *Handler) UnlockStock(ctx context.Context) error {
	handler := &unlockHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockRollback.Enum(),
		},
	}

	if err := handler.lockOp.getLocks(ctx); err != nil {
		if ent.IsNotFound(err) && h.Rollback != nil && *h.Rollback {
			return nil
		}
		return wlog.WrapError(err)
	}
	if h.Rollback == nil || !*h.Rollback {
		handler.lockOp.state = types.AppStockLockState_AppStockCanceled.Enum()
	}
	h.Stocks = handler.lockOp.lock2Stocks()
	if err := handler.getStockAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockOp.locks {
			if err := handler.unlockAppStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.unlockAppMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.unlockStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.unlockMiningGoodStock(ctx, lock, tx); err != nil {
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
