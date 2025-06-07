package appstock

import (
	"context"

	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appmininggoodstockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/stock/mining"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entappmininggoodstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appmininggoodstock"
	entappstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appstock"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entpowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/powerrental"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/stock"
	appmininggoodstockmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/stock/mining"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect           *ent.AppStockSelect
	stmCount            *ent.AppStockSelect
	infos               []*npool.Stock
	appMiningGoodStocks []*appmininggoodstockmwpb.StockInfo
	total               uint32
}

func (h *queryHandler) selectStock(stm *ent.AppStockQuery) *ent.AppStockSelect {
	return stm.Select(entappstock.FieldID)
}

func (h *queryHandler) queryStock(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.AppStock.Query().Where(entappstock.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappstock.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappstock.EntID(*h.EntID))
	}
	h.stmSelect = h.selectStock(stm)
	return nil
}

func (h *queryHandler) queryStocks(cli *ent.Client) (*ent.AppStockSelect, error) {
	return nil, wlog.Errorf("NOT IMPLEMENTED")
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappstock.Table)
	s.Join(t).
		On(
			s.C(entappstock.FieldID),
			t.C(entappstock.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappstock.FieldEntID), "ent_id"),
			sql.As(t.C(entappstock.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entappstock.FieldReserved), "reserved"),
			sql.As(t.C(entappstock.FieldSpotQuantity), "spot_quantity"),
			sql.As(t.C(entappstock.FieldLocked), "locked"),
			sql.As(t.C(entappstock.FieldWaitStart), "wait_start"),
			sql.As(t.C(entappstock.FieldInService), "in_service"),
			sql.As(t.C(entappstock.FieldSold), "sold"),
			sql.As(t.C(entappstock.FieldCreatedAt), "created_at"),
			sql.As(t.C(entappstock.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	t3 := sql.Table(entpowerrental.Table)

	s.Join(t1).
		On(
			s.C(entappstock.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		Join(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		LeftJoin(t3).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t3.C(entpowerrental.FieldGoodID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldAppID),
			t1.C(entappgoodbase.FieldGoodID),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			sql.As(t2.C(entgoodbase.FieldName), "good_name"),
			t3.C(entpowerrental.FieldStockMode),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinAppGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinAppGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) getAppMiningGoodStocks(ctx context.Context, cli *ent.Client) error {
	appGoodStockIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			if info.StockModeStr != types.GoodStockMode_GoodStockByMiningPool.String() {
				continue
			}
			uids = append(uids, uuid.MustParse(info.EntID))
		}
		return
	}()

	stm, err := appmininggoodstockcrud.SetQueryConds(
		cli.AppMiningGoodStock.Query(),
		&appmininggoodstockcrud.Conds{
			AppGoodStockIDs: &cruder.Cond{Op: cruder.IN, Val: appGoodStockIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entappmininggoodstock.FieldID,
		entappmininggoodstock.FieldEntID,
		entappmininggoodstock.FieldAppGoodStockID,
		entappmininggoodstock.FieldMiningGoodStockID,
		entappmininggoodstock.FieldReserved,
		entappmininggoodstock.FieldSpotQuantity,
		entappmininggoodstock.FieldLocked,
		entappmininggoodstock.FieldWaitStart,
		entappmininggoodstock.FieldInService,
		entappmininggoodstock.FieldSold,
	).Scan(ctx, &h.appMiningGoodStocks)
}

func (h *queryHandler) formalize() {
	appMiningGoodStocks := map[string][]*appmininggoodstockmwpb.StockInfo{}

	for _, stock := range h.appMiningGoodStocks {
		stock.Reserved = func() string { amount, _ := decimal.NewFromString(stock.Reserved); return amount.String() }()
		stock.SpotQuantity = func() string { amount, _ := decimal.NewFromString(stock.SpotQuantity); return amount.String() }()
		stock.Locked = func() string { amount, _ := decimal.NewFromString(stock.Locked); return amount.String() }()
		stock.WaitStart = func() string { amount, _ := decimal.NewFromString(stock.WaitStart); return amount.String() }()
		stock.InService = func() string { amount, _ := decimal.NewFromString(stock.InService); return amount.String() }()
		stock.Sold = func() string { amount, _ := decimal.NewFromString(stock.Sold); return amount.String() }()
		appMiningGoodStocks[stock.AppGoodStockID] = append(appMiningGoodStocks[stock.AppGoodStockID], stock)
	}
	for _, info := range h.infos {
		info.Reserved = func() string { amount, _ := decimal.NewFromString(info.Reserved); return amount.String() }()
		info.SpotQuantity = func() string { amount, _ := decimal.NewFromString(info.SpotQuantity); return amount.String() }()
		info.Locked = func() string { amount, _ := decimal.NewFromString(info.Locked); return amount.String() }()
		info.InService = func() string { amount, _ := decimal.NewFromString(info.InService); return amount.String() }()
		info.WaitStart = func() string { amount, _ := decimal.NewFromString(info.WaitStart); return amount.String() }()
		info.Sold = func() string { amount, _ := decimal.NewFromString(info.Sold); return amount.String() }()
		info.StockMode = types.GoodStockMode(types.GoodStockMode_value[info.StockModeStr])
		info.AppMiningGoodStocks = appMiningGoodStocks[info.EntID]
	}
}

func (h *Handler) GetStock(ctx context.Context) (*npool.Stock, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryStock(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.getAppMiningGoodStocks(_ctx, cli)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetStocks(ctx context.Context) ([]*npool.Stock, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryStocks(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryStocks(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(total)

		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.getAppMiningGoodStocks(_ctx, cli)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
