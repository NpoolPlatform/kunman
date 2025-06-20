package appstock

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *lockHandler) constructGoodStockSQL(table string, stock *LockStock, id uint32) (string, error) {
	platformLocked := *stock.Locked
	if stock.AppSpotLocked != nil {
		platformLocked = platformLocked.Sub(*stock.AppSpotLocked)
	}
	if platformLocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return "", wlog.Errorf("invalid appspotlocked")
	}
	sql := fmt.Sprintf(
		`update %v
		set
		  spot_quantity = spot_quantity - %v,
		  locked = locked + %v`,
		table,
		platformLocked,
		*stock.Locked,
	)
	if stock.AppSpotLocked != nil {
		sql += fmt.Sprintf(
			`, app_reserved = app_reserved - %v`,
			*stock.AppSpotLocked,
		)
	}
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  spot_quantity >= %v`,
		id,
		platformLocked,
	)
	if stock.AppSpotLocked != nil {
		sql += fmt.Sprintf(
			` and
			  app_reserved >= %v`,
			*stock.AppSpotLocked,
		)
	}
	sql += ` and
		in_service + wait_start + locked + app_reserved + spot_quantity = total`
	return sql, nil
}

func (h *lockHandler) constructAppGoodStockSQL(table string, stock *LockStock, id uint32) string {
	sql := fmt.Sprintf(
		`update %v
		set
		  locked = locked + %v`,
		table,
		*stock.Locked,
	)
	if stock.AppSpotLocked != nil {
		sql += fmt.Sprintf(
			`, spot_quantity = spot_quantity - %v`,
			*stock.AppSpotLocked,
		)
	}
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0`,
		id,
	)
	if stock.AppSpotLocked != nil {
		sql += fmt.Sprintf(
			` and
			  spot_quantity >= %v
			and
			  spot_quantity - %v <= reserved`,
			*stock.AppSpotLocked,
			*stock.AppSpotLocked,
		)
	}
	return sql
}

func (h *lockHandler) lockStock(ctx context.Context, stock *LockStock, tx *ent.Tx) (err error) {
	_stock, ok := h.stocks[*stock.AppGoodID]
	if !ok || _stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructGoodStockSQL("stocks_v1", stock, _stock.stock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *lockHandler) lockMiningGoodStock(ctx context.Context, stock *LockStock, tx *ent.Tx) (err error) {
	_stock, ok := h.stocks[*stock.AppGoodID]
	if !ok || _stock.miningGoodStocks == nil {
		return wlog.Errorf("invalid mininggoodstock")
	}
	appMiningGoodStock, ok := _stock.appMiningGoodStocks[*stock.EntID]
	if !ok {
		return wlog.Errorf("invalid appmininggoodstock")
	}
	miningGoodStock, ok := _stock.miningGoodStocks[appMiningGoodStock.MiningGoodStockID]
	if !ok {
		return wlog.Errorf("invalid mininggoodstock")
	}
	sql, err := h.constructGoodStockSQL("mining_good_stocks", stock, miningGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *lockHandler) lockAppStock(ctx context.Context, stock *LockStock, tx *ent.Tx) (err error) {
	_stock, ok := h.stocks[*stock.AppGoodID]
	if !ok || _stock.appGoodStock == nil {
		return wlog.Errorf("invalid appstock")
	}
	sql := h.constructAppGoodStockSQL("app_stocks", stock, _stock.appGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *lockHandler) lockAppMiningGoodStock(ctx context.Context, stock *LockStock, tx *ent.Tx) (err error) {
	_stock, ok := h.stocks[*stock.AppGoodID]
	if !ok || _stock.appMiningGoodStocks == nil {
		return wlog.Errorf("invalid appmininggoodstock")
	}
	appMiningGoodStock, ok := _stock.appMiningGoodStocks[*stock.EntID]
	if !ok {
		return wlog.Errorf("invalid appmininggoodstock")
	}
	sql := h.constructAppGoodStockSQL("app_mining_good_stocks", stock, appMiningGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *Handler) LockStock(ctx context.Context) error {
	handler := &lockHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
		},
	}

	if err := handler.getStockAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		stock := &LockStock{
			EntID:         h.EntID,
			AppGoodID:     h.AppGoodID,
			Locked:        h.Locked,
			AppSpotLocked: h.AppSpotLocked,
		}
		if err := handler.lockAppStock(ctx, stock, tx); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.lockAppMiningGoodStock(ctx, stock, tx); err != nil {
				return wlog.WrapError(err)
			}
		}
		if err := handler.lockStock(ctx, stock, tx); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.lockMiningGoodStock(ctx, stock, tx); err != nil {
				return wlog.WrapError(err)
			}
		}
		if err := handler.lockOp.createLocks(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *Handler) LockStocks(ctx context.Context) error {
	handler := &lockHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
		},
	}

	if err := handler.getStockAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, stock := range h.Stocks {
			if err := handler.lockAppStock(ctx, stock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(*stock.AppGoodID) {
				if err := handler.lockAppMiningGoodStock(ctx, stock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.lockStock(ctx, stock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(*stock.AppGoodID) {
				if err := handler.lockMiningGoodStock(ctx, stock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
		}
		if err := handler.lockOp.createLocks(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
