package currencyhistory

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	historycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/currency/history"
	entcoinbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	entcurrencyhis "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/currencyhistory"

	"entgo.io/ent/dialect/sql"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CurrencyHistorySelect
	stmCount  *ent.CurrencyHistorySelect
	infos     []*npool.Currency
	total     uint32
}

func (h *queryHandler) selectCurrencyHistory(stm *ent.CurrencyHistoryQuery) *ent.CurrencyHistorySelect {
	return stm.Select(entcurrencyhis.FieldID)
}

func (h *queryHandler) queryCurrencyHistories(cli *ent.Client) (*ent.CurrencyHistorySelect, error) {
	stm, err := historycrud.SetQueryConds(cli.CurrencyHistory.Query(), h.Conds)
	if err != nil {
		return nil, err
	}

	return h.selectCurrencyHistory(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcurrencyhis.Table)
	s.AppendSelect(
		sql.As(t.C(entcurrencyhis.FieldEntID), "ent_id"),
		sql.As(t.C(entcurrencyhis.FieldCoinTypeID), "coin_type_id"),
		sql.As(t.C(entcurrencyhis.FieldFeedType), "feed_type"),
		sql.As(t.C(entcurrencyhis.FieldMarketValueHigh), "market_value_high"),
		sql.As(t.C(entcurrencyhis.FieldMarketValueLow), "market_value_low"),
		sql.As(t.C(entcurrencyhis.FieldCreatedAt), "created_at"),
		sql.As(t.C(entcurrencyhis.FieldUpdatedAt), "updated_at"),
	)
}

func (h *queryHandler) queryJoinCoin(s *sql.Selector) error {
	t := sql.Table(entcoinbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entcurrencyhis.FieldCoinTypeID),
			t.C(entcoinbase.FieldEntID),
		)

	if h.Conds.CoinNames != nil {
		names, ok := h.Conds.CoinNames.Val.([]string)
		if !ok {
			return fmt.Errorf("invalid coinnames")
		}
		_names := []interface{}{}
		for _, _name := range names {
			_names = append(_names, _name)
		}
		s.Where(
			sql.In(t.C(entcoinbase.FieldName), _names...),
		)
	}

	s.AppendSelect(
		sql.As(t.C(entcoinbase.FieldName), "coin_name"),
		sql.As(t.C(entcoinbase.FieldLogo), "coin_logo"),
		sql.As(t.C(entcoinbase.FieldUnit), "coin_unit"),
		sql.As(t.C(entcoinbase.FieldEnv), "coin_env"),
	)
	return nil
}

func (h *queryHandler) queryJoin() (err error) {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err = h.queryJoinCoin(s); err != nil {
			return
		}
	})
	h.stmCount.Modify(func(s *sql.Selector) {
		if err = h.queryJoinCoin(s); err != nil {
			return
		}
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.FeedType = basetypes.CurrencyFeedType(basetypes.CurrencyFeedType_value[info.FeedTypeStr])
	}
}

func (h *Handler) GetCurrencies(ctx context.Context) ([]*npool.Currency, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_stm, err := handler.queryCurrencyHistories(cli)
		if err != nil {
			return err
		}
		handler.stmSelect = _stm

		_stm, err = handler.queryCurrencyHistories(cli)
		if err != nil {
			return err
		}
		handler.stmCount = _stm

		if err := handler.queryJoin(); err != nil {
			return err
		}

		_total, err := handler.stmCount.Count(ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Order(ent.Asc(entcurrencyhis.FieldCreatedAt)).
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
