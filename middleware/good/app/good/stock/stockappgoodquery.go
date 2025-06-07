package appstock

import (
	"context"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entappmininggoodstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appmininggoodstock"
	entappgoodstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appstock"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entmininggoodstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/mininggoodstock"
	entpowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/powerrental"
	entstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/stock"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
)

type stockAppGood struct {
	goodBase            *ent.GoodBase
	powerRental         *ent.PowerRental
	appGoodBase         *ent.AppGoodBase
	appGoodStock        *ent.AppStock
	appMiningGoodStocks map[uuid.UUID]*ent.AppMiningGoodStock
	stock               *ent.Stock
	miningGoodStocks    map[uuid.UUID]*ent.MiningGoodStock
}

type stockAppGoodQuery struct {
	*Handler
	stocks                   map[uuid.UUID]*stockAppGood
	appGoodIDs               []uuid.UUID
	appGoodStockEntIDs       []uuid.UUID
	appMiningGoodStockEntIDs []uuid.UUID
	miningGoodStockEntIDs    []uuid.UUID
	stockGoodIDs             []uuid.UUID
	appMiningGoodStocks      []*ent.AppMiningGoodStock
	miningGoodStocks         []*ent.MiningGoodStock
}

func (h *stockAppGoodQuery) _getAppGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	appGoodStocks, err := cli.
		AppStock.
		Query().
		Where(
			entappgoodstock.EntIDIn(h.appGoodStockEntIDs...),
			entappgoodstock.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, stock := range appGoodStocks {
		_stock, ok := h.stocks[stock.AppGoodID]
		if !ok {
			return wlog.Errorf("invalid stock")
		}
		_stock.appGoodStock = stock
	}
	return nil
}

func (h *stockAppGoodQuery) getAppMiningGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	h.appMiningGoodStocks, err = cli.
		AppMiningGoodStock.
		Query().
		Where(
			entappmininggoodstock.EntIDIn(h.appMiningGoodStockEntIDs...),
			entappmininggoodstock.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, stock := range h.appMiningGoodStocks {
		h.appGoodStockEntIDs = append(h.appGoodStockEntIDs, stock.AppGoodStockID)
		h.miningGoodStockEntIDs = append(h.miningGoodStockEntIDs, stock.MiningGoodStockID)
	}
	return nil
}

func (h *stockAppGoodQuery) formalizeMiningGoodStocks() {
	for _, stock := range h.stocks {
		stock.miningGoodStocks = map[uuid.UUID]*ent.MiningGoodStock{}
		for _, miningGoodStock := range h.miningGoodStocks {
			if miningGoodStock.GoodStockID == stock.stock.EntID {
				stock.miningGoodStocks[miningGoodStock.EntID] = miningGoodStock
			}
		}
	}
}

func (h *stockAppGoodQuery) formalizeAppMiningGoodStocks() {
	for _, stock := range h.stocks {
		stock.appMiningGoodStocks = map[uuid.UUID]*ent.AppMiningGoodStock{}
		for _, appMiningGoodStock := range h.appMiningGoodStocks {
			if stock.appGoodStock != nil && appMiningGoodStock.AppGoodStockID == stock.appGoodStock.EntID {
				stock.appMiningGoodStocks[appMiningGoodStock.EntID] = appMiningGoodStock
			}
		}
	}
}

func (h *stockAppGoodQuery) getAppGoodBases(ctx context.Context, cli *ent.Client) (err error) {
	appGoodBases, err := cli.
		AppGoodBase.
		Query().
		Where(
			entappgoodbase.EntIDIn(h.appGoodIDs...),
			entappgoodbase.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, appGoodBase := range appGoodBases {
		h.stocks[appGoodBase.EntID] = &stockAppGood{}
		h.stocks[appGoodBase.EntID].appGoodBase = appGoodBase
		h.stockGoodIDs = append(h.stockGoodIDs, appGoodBase.GoodID)
	}
	return nil
}

func (h *stockAppGoodQuery) getMiningGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	h.miningGoodStocks, err = cli.
		MiningGoodStock.
		Query().
		Where(
			entmininggoodstock.EntIDIn(h.miningGoodStockEntIDs...),
			entmininggoodstock.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *stockAppGoodQuery) getGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	stocks, err := cli.
		Stock.
		Query().
		Where(
			entstock.GoodIDIn(h.stockGoodIDs...),
			entstock.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, _stock := range h.stocks {
		for _, stock := range stocks {
			if _stock.goodBase != nil && _stock.goodBase.EntID == stock.GoodID {
				_stock.stock = stock
				break
			}
		}
	}
	return nil
}

func (h *stockAppGoodQuery) getGoodBases(ctx context.Context, cli *ent.Client) (err error) {
	goodIDs := []uuid.UUID{}
	for _, stock := range h.stocks {
		if stock.appGoodBase == nil {
			return wlog.Errorf("invalid appgoodbase")
		}
		goodIDs = append(goodIDs, stock.appGoodBase.GoodID)
	}

	goodBases, err := cli.
		GoodBase.
		Query().
		Where(
			entgoodbase.EntIDIn(goodIDs...),
			entgoodbase.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, stock := range h.stocks {
		for _, goodBase := range goodBases {
			if stock.appGoodBase != nil && stock.appGoodBase.GoodID == goodBase.EntID {
				stock.goodBase = goodBase
				break
			}
		}
	}
	return nil
}

func (h *stockAppGoodQuery) getPowerRentals(ctx context.Context, cli *ent.Client) (err error) {
	powerRentalGoodIDs := func() (uids []uuid.UUID) {
		for _, stock := range h.stocks {
			if stock.goodBase == nil {
				continue
			}
			switch stock.goodBase.GoodType {
			case types.GoodType_PowerRental.String():
			case types.GoodType_LegacyPowerRental.String():
			default:
				continue
			}
			uids = append(uids, stock.goodBase.EntID)
		}
		return
	}()

	powerRentals, err := cli.
		PowerRental.
		Query().
		Where(
			entpowerrental.GoodIDIn(powerRentalGoodIDs...),
			entpowerrental.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, stock := range h.stocks {
		for _, powerRental := range powerRentals {
			if stock.goodBase != nil && stock.goodBase.EntID == powerRental.GoodID {
				stock.powerRental = powerRental
				break
			}
		}
	}
	return nil
}

func (h *stockAppGoodQuery) getAppGoods(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getAppGoodBases(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.getGoodBases(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return h.getPowerRentals(_ctx, cli)
	})
}

func (h *stockAppGoodQuery) formalizeAppGoodIDs() error {
	if h.AppGoodID != nil {
		h.appGoodIDs = append(h.appGoodIDs, *h.AppGoodID)
	}
	for _, stock := range h.Stocks {
		h.appGoodIDs = append(h.appGoodIDs, *stock.AppGoodID)
	}
	if len(h.appGoodIDs) == 0 {
		return wlog.Errorf("invalid appgoodids")
	}
	return nil
}

func (h *stockAppGoodQuery) formalizeStockEntID(appGoodID, entID uuid.UUID) error {
	_stock, ok := h.stocks[appGoodID]
	if !ok {
		return wlog.Errorf("invalid stock")
	}
	if _stock.powerRental == nil || _stock.powerRental.StockMode != types.GoodStockMode_GoodStockByMiningPool.String() {
		h.appGoodStockEntIDs = append(h.appGoodStockEntIDs, entID)
		return nil
	}
	h.appMiningGoodStockEntIDs = append(h.appMiningGoodStockEntIDs, entID)
	return nil
}

func (h *stockAppGoodQuery) formalizeStockEntIDs() error {
	if h.AppGoodID != nil && h.EntID != nil {
		if err := h.formalizeStockEntID(*h.AppGoodID, *h.EntID); err != nil {
			return wlog.WrapError(err)
		}
	}
	for _, stock := range h.Stocks {
		if err := h.formalizeStockEntID(*stock.AppGoodID, *stock.EntID); err != nil {
			logger.Sugar().Errorw("formalizeStockEntID", "Stock", *stock, "Error", err)
		}
	}
	if len(h.appGoodStockEntIDs) == 0 && len(h.appMiningGoodStockEntIDs) == 0 {
		return wlog.Errorf("invalid stock")
	}
	return nil
}

func (h *stockAppGoodQuery) getAppGoodStocks(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if len(h.appMiningGoodStockEntIDs) > 0 {
			if err := h.getAppMiningGoodStocks(_ctx, cli); err != nil {
				return wlog.WrapError(err)
			}
		}
		if len(h.appGoodStockEntIDs) > 0 {
			if err := h._getAppGoodStocks(_ctx, cli); err != nil {
				return wlog.WrapError(err)
			}
		}
		if len(h.miningGoodStockEntIDs) > 0 {
			if err := h.getMiningGoodStocks(_ctx, cli); err != nil {
				return wlog.WrapError(err)
			}
		}
		if err := h.getGoodStocks(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		h.formalizeMiningGoodStocks()
		h.formalizeAppMiningGoodStocks()
		return nil
	})
}

func (h *stockAppGoodQuery) getStockAppGoods(ctx context.Context) error {
	h.stocks = map[uuid.UUID]*stockAppGood{}
	if err := h.formalizeAppGoodIDs(); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.getAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.formalizeStockEntIDs(); err != nil {
		return wlog.WrapError(err)
	}
	return h.getAppGoodStocks(ctx)
}

func (h *stockAppGoodQuery) getStockGoods(ctx context.Context) error {
	h.stocks = map[uuid.UUID]*stockAppGood{}
	if err := h.formalizeAppGoodIDs(); err != nil {
		return wlog.WrapError(err)
	}
	return h.getAppGoods(ctx)
}

func (h *stockAppGoodQuery) stockByMiningPool(appGoodID uuid.UUID) bool {
	stock, ok := h.stocks[appGoodID]
	if !ok {
		return false
	}
	return stock.powerRental != nil && stock.powerRental.StockMode == types.GoodStockMode_GoodStockByMiningPool.String()
}
