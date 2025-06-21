package currency

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	fiatcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat"
	currencycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat/currency"
	entfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiat"
	entcurrency "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiatcurrency"

	"entgo.io/ent/dialect/sql"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.FiatSelect
	stmCount  *ent.FiatSelect
	infos     []*npool.Currency
	total     uint32
}

func (h *queryHandler) selectFiat(stm *ent.FiatQuery) *ent.FiatSelect {
	return stm.Select(entfiat.FieldCreatedAt)
}

func (h *queryHandler) queryFiat(ctx context.Context, cli *ent.Client) error {
	_stm1, err := currencycrud.SetQueryConds(cli.FiatCurrency.Query(), &currencycrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *h.EntID},
	})
	if err != nil {
		return err
	}
	_info1, err := _stm1.Only(ctx)
	if err != nil {
		return err
	}

	_stm2, err := fiatcrud.SetQueryConds(cli.Fiat.Query(), &fiatcrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: _info1.FiatID},
	})
	if err != nil {
		return err
	}

	h.stmSelect = h.selectFiat(_stm2)
	return nil
}

func (h *queryHandler) queryFiats(cli *ent.Client) (*ent.FiatSelect, error) {
	stm, err := fiatcrud.SetQueryConds(cli.Fiat.Query(), &fiatcrud.Conds{
		EntID:  h.Conds.FiatID,
		EntIDs: h.Conds.FiatIDs,
		Name:   h.Conds.FiatName,
	})
	if err != nil {
		return nil, err
	}

	return h.selectFiat(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entfiat.Table)
	s.AppendSelect(
		sql.As(t.C(entfiat.FieldEntID), "fiat_id"),
		sql.As(t.C(entfiat.FieldName), "fiat_name"),
		sql.As(t.C(entfiat.FieldLogo), "fiat_logo"),
		sql.As(t.C(entfiat.FieldUnit), "fiat_unit"),
	)
}

func (h *queryHandler) queryJoinCurrency(s *sql.Selector) error {
	t := sql.Table(entcurrency.Table)
	s.LeftJoin(t).
		On(
			s.C(entfiat.FieldEntID),
			t.C(entcurrency.FieldFiatID),
		).
		OnP(
			sql.EQ(t.C(entcurrency.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entcurrency.FieldID), "id"),
			sql.As(t.C(entcurrency.FieldEntID), "ent_id"),
			sql.As(t.C(entcurrency.FieldFeedType), "feed_type"),
			sql.As(t.C(entcurrency.FieldMarketValueHigh), "market_value_high"),
			sql.As(t.C(entcurrency.FieldMarketValueLow), "market_value_low"),
			sql.As(t.C(entcurrency.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcurrency.FieldUpdatedAt), "updated_at"),
		)

	if h.Conds != nil && h.Conds.EntID != nil {
		id, ok := h.Conds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid entid")
		}
		switch h.Conds.EntID.Op {
		case cruder.EQ:
			s.Where(
				sql.EQ(t.C(entcurrency.FieldEntID), id),
			)
		default:
			return fmt.Errorf("invalid fiat currency field op")
		}
	}
	if h.Conds != nil && h.Conds.FiatID != nil {
		id, ok := h.Conds.FiatID.Val.(string)
		if !ok {
			return fmt.Errorf("invalid fiatid")
		}
		switch h.Conds.FiatID.Op {
		case cruder.EQ:
			s.Where(
				sql.EQ(t.C(entcurrency.FieldFiatID), id),
			)
		default:
			return fmt.Errorf("invalid fiat currency field op")
		}
	}
	if h.Conds != nil && h.Conds.FiatIDs != nil {
		ids, ok := h.Conds.FiatIDs.Val.([]string)
		if !ok {
			return fmt.Errorf("invalid fiatids")
		}
		var valueInterfaces []interface{}
		for _, value := range ids {
			valueInterfaces = append(valueInterfaces, value)
		}
		switch h.Conds.FiatIDs.Op {
		case cruder.IN:
			s.Where(
				sql.In(t.C(entcurrency.FieldFiatID), valueInterfaces...),
			)
		default:
			return fmt.Errorf("invalid fiat currency field op")
		}
	}

	return nil
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinCurrency(s); err != nil {
			return
		}
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		if err := h.queryJoinCurrency(s); err != nil {
			return
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.FeedType = basetypes.CurrencyFeedType(basetypes.CurrencyFeedType_value[info.FeedTypeStr])
		if _, err := decimal.NewFromString(info.MarketValueHigh); err != nil {
			info.MarketValueHigh = decimal.NewFromInt(0).String()
		}
		if _, err := decimal.NewFromString(info.MarketValueLow); err != nil {
			info.MarketValueLow = decimal.NewFromInt(0).String()
		}
		if _, err := uuid.Parse(info.EntID); err != nil {
			info.EntID = uuid.Nil.String()
		}
	}
}

func (h *Handler) GetCurrency(ctx context.Context) (*npool.Currency, error) {
	if h.EntID == nil {
		return nil, fmt.Errorf("invalid entid")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryFiat(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		const singleRowLimit = 1
		handler.stmSelect.
			Offset(0).
			Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetCurrencies(ctx context.Context) ([]*npool.Currency, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryFiats(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryFiats(cli)
		if err != nil {
			return err
		}

		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stmSelect.
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

func (h *Handler) GetCurrencyOnly(ctx context.Context) (*npool.Currency, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryFiats(cli)
		if err != nil {
			return err
		}

		handler.queryJoin()

		handler.stmSelect.
			Offset(0).
			Limit(2)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("invalid currency")
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	handler.formalize()
	return handler.infos[0], nil
}
