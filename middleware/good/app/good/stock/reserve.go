package appstock

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"

	"github.com/shopspring/decimal"
)

type reserveHandler struct {
	*stockAppGoodQuery
}

func (h *reserveHandler) constructGoodStockSQL(table string, id uint32) string {
	sql := fmt.Sprintf(
		`update %v
		set
		  spot_quantity = spot_quantity - %v,
		  app_reserved = app_reserved + %v`,
		table,
		*h.Reserved,
		*h.Reserved,
	)
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  spot_quantity >= %v`,
		id,
		*h.Reserved,
	)
	sql += ` and
		in_service + wait_start + locked + app_reserved + spot_quantity = total`
	return sql
}

func (h *reserveHandler) constructAppGoodStockSQL(table string, id uint32) string {
	sql := fmt.Sprintf(
		`update %v
		set
		  spot_quantity = spot_quantity + %v,
		  reserved = reserved + %v`,
		table,
		*h.Reserved,
		*h.Reserved,
	)
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  spot_quantity <= reserved`,
		id,
	)
	return sql
}

func (h *reserveHandler) reserveStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructGoodStockSQL("stocks_v1", stock.stock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *reserveHandler) reserveMiningGoodStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok || stock.miningGoodStocks == nil {
		return wlog.Errorf("invalid stock")
	}
	appMiningGoodStock, ok := stock.appMiningGoodStocks[*h.EntID]
	if !ok {
		return wlog.Errorf("invalid appmininggoodstock")
	}
	miningGoodStock, ok := stock.miningGoodStocks[appMiningGoodStock.MiningGoodStockID]
	if !ok {
		return wlog.Errorf("invalid mininggoodstock")
	}
	sql := h.constructGoodStockSQL("mining_good_stocks", miningGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *reserveHandler) reserveAppStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructAppGoodStockSQL("app_stocks", stock.appGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *reserveHandler) reserveAppMiningGoodStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok || stock.appMiningGoodStocks == nil {
		return wlog.Errorf("invalid stock")
	}
	appMiningGoodStock, ok := stock.appMiningGoodStocks[*h.EntID]
	if !ok {
		return wlog.Errorf("invalid stock")
	}
	sql := h.constructAppGoodStockSQL("app_mining_good_stocks", appMiningGoodStock.ID)
	return h.execSQL(ctx, tx, sql)
}

func (h *Handler) ReserveStock(ctx context.Context) error {
	handler := &reserveHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
	}

	if err := handler.getStockAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}
	stock, ok := handler.stocks[*h.AppGoodID]
	if !ok || stock.powerRental == nil {
		return wlog.Errorf("invalid appgoodid")
	}
	if stock.powerRental.UnitLockDeposit.Equal(decimal.NewFromInt(0)) {
		return wlog.Errorf("permission denied")
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.reserveAppStock(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.reserveAppMiningGoodStock(ctx, tx); err != nil {
				return wlog.WrapError(err)
			}
		}
		if err := handler.reserveStock(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.reserveMiningGoodStock(ctx, tx); err != nil {
				return wlog.WrapError(err)
			}
		}
		return nil
	})
}
