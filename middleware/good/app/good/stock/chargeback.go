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

type chargeBackHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *chargeBackHandler) constructSQL(table string, lock *ent.AppStockLock, returnSpotQuantity, checkTotal bool, id uint32) (string, error) {
	sql := fmt.Sprintf(
		`update %v
		set
		  sold = sold - %v`,
		table,
		lock.Units,
	)
	if returnSpotQuantity {
		sql += fmt.Sprintf(
			`, spot_quantity = spot_quantity + %v`,
			lock.Units,
		)
	}
	switch lock.LockState {
	case types.AppStockLockState_AppStockInService.String():
		sql += fmt.Sprintf(
			`, in_service = in_service - %v
			where
			  in_service >= %v `,
			lock.Units,
			lock.Units,
		)
	case types.AppStockLockState_AppStockWaitStart.String():
		sql += fmt.Sprintf(
			`, wait_start = wait_start - %v
			where
			  wait_start >= %v `,
			lock.Units,
			lock.Units,
		)
	default:
		return "", wlog.Errorf("invalid lockstate")
	}
	sql += fmt.Sprintf(
		`and
		  id = %v
		and
		  deleted_at = 0
		and
		  sold >= %v`,
		id,
		lock.Units,
	)
	if checkTotal {
		sql += ` and
		  in_service + wait_start + locked + app_reserved + spot_quantity = total`
	}
	return sql, nil
}

func (h *chargeBackHandler) chargeBackStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructSQL("stocks_v1", lock, true, true, stock.stock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *chargeBackHandler) chargeBackMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
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
	sql, err := h.constructSQL("mining_good_stocks", lock, true, true, miningGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *chargeBackHandler) chargeBackAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructSQL("app_stocks", lock, false, false, stock.appGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *chargeBackHandler) chargeBackAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appMiningGoodStocks == nil {
		return wlog.Errorf("invalid stock")
	}
	appMiningGoodStock, ok := stock.appMiningGoodStocks[lock.AppStockID]
	if !ok {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructSQL("app_mining_good_stocks", lock, false, false, appMiningGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *Handler) ChargeBackStock(ctx context.Context) error {
	handler := &chargeBackHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockChargeBack.Enum(),
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
			if err := handler.chargeBackAppStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.chargeBackAppMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.chargeBackStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.chargeBackMiningGoodStock(ctx, lock, tx); err != nil {
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
