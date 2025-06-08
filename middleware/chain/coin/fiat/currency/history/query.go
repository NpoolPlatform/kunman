package currencyhistory

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat/currency"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	historycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/fiat/currency/history"
	entcoinbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	entcurrencyhis "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinfiatcurrencyhistory"
	entfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiat"

	"entgo.io/ent/dialect/sql"
)

type queryHandler struct {
	*Handler
	stm   *ent.CoinFiatCurrencyHistorySelect
	infos []*npool.Currency
	total uint32
}

func (h *queryHandler) selectCurrencyHistory(stm *ent.CoinFiatCurrencyHistoryQuery) {
	h.stm = stm.Select(
		entcurrencyhis.FieldID,
		entcurrencyhis.FieldEntID,
		entcurrencyhis.FieldCoinTypeID,
		entcurrencyhis.FieldFeedType,
		entcurrencyhis.FieldMarketValueHigh,
		entcurrencyhis.FieldMarketValueLow,
		entcurrencyhis.FieldCreatedAt,
		entcurrencyhis.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCurrencyHistories(ctx context.Context, cli *ent.Client) error {
	stm, err := historycrud.SetQueryConds(cli.CoinFiatCurrencyHistory.Query(), h.Conds)
	if err != nil {
		return err
	}

	_total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(_total)
	h.selectCurrencyHistory(stm)
	return nil
}

func (h *queryHandler) queryJoinCoin(s *sql.Selector) {
	t1 := sql.Table(entcoinbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entcurrencyhis.FieldCoinTypeID),
			t1.C(entcoinbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entcoinbase.FieldName), "coin_name"),
			sql.As(t1.C(entcoinbase.FieldLogo), "coin_logo"),
			sql.As(t1.C(entcoinbase.FieldUnit), "coin_unit"),
			sql.As(t1.C(entcoinbase.FieldEnv), "coin_env"),
		)
}

func (h *queryHandler) queryJoinFiat(s *sql.Selector) {
	t := sql.Table(entfiat.Table)
	s.
		LeftJoin(t).
		On(
			s.C(entcurrencyhis.FieldFiatID),
			t.C(entfiat.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entfiat.FieldName), "fiat_name"),
			sql.As(t.C(entfiat.FieldLogo), "fiat_logo"),
			sql.As(t.C(entfiat.FieldUnit), "fiat_unit"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinCoin(s)
		h.queryJoinFiat(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
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
		if err := handler.queryCurrencyHistories(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Order(ent.Desc(entcurrencyhis.FieldCreatedAt)).
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
