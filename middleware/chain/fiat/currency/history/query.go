package currencyhistory

import (
	"context"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"

	historycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat/currency/history"
	entfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiat"
	entcurrencyhis "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiatcurrencyhistory"

	"entgo.io/ent/dialect/sql"
)

type queryHandler struct {
	*Handler
	stm   *ent.FiatCurrencyHistorySelect
	infos []*npool.Currency
	total uint32
}

func (h *queryHandler) selectCurrencyHistory(stm *ent.FiatCurrencyHistoryQuery) {
	h.stm = stm.Select(
		entcurrencyhis.FieldID,
		entcurrencyhis.FieldEntID,
		entcurrencyhis.FieldFiatID,
		entcurrencyhis.FieldFeedType,
		entcurrencyhis.FieldMarketValueHigh,
		entcurrencyhis.FieldMarketValueLow,
		entcurrencyhis.FieldCreatedAt,
		entcurrencyhis.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCurrencyHistories(ctx context.Context, cli *ent.Client) error {
	stm, err := historycrud.SetQueryConds(cli.FiatCurrencyHistory.Query(), h.Conds)
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
	t1 := sql.Table(entfiat.Table)
	s.LeftJoin(t1).
		On(
			s.C(entcurrencyhis.FieldFiatID),
			t1.C(entfiat.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entfiat.FieldName), "fiat_name"),
			sql.As(t1.C(entfiat.FieldLogo), "fiat_logo"),
			sql.As(t1.C(entfiat.FieldUnit), "fiat_unit"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinCoin(s)
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
		handler.stm.
			Order(ent.Asc(entcurrencyhis.FieldCreatedAt)).
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		handler.queryJoin()
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
