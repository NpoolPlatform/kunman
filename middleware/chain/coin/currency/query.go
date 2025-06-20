package currency

import (
	"context"
	"fmt"
	"time"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency"

	"github.com/NpoolPlatform/kunman/middleware/chain/db"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	coincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin"
	currencycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/currency"
	entcoinbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	entcoinextra "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinextra"
	entcurrency "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/currency"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CoinBaseSelect
	stmCount  *ent.CoinBaseSelect
	infos     []*npool.Currency
	total     uint32
}

func (h *queryHandler) selectCoinBase(stm *ent.CoinBaseQuery) *ent.CoinBaseSelect {
	return stm.Select(entcoinbase.FieldCreatedAt)
}

func (h *queryHandler) queryCoinBase(ctx context.Context, cli *ent.Client) error {
	_stm1, err := currencycrud.SetQueryConds(cli.Currency.Query(), &currencycrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *h.EntID},
	})
	if err != nil {
		return err
	}
	_info1, err := _stm1.Only(ctx)
	if err != nil {
		return err
	}

	_stm2, err := coincrud.SetQueryConds(cli.CoinBase.Query(), &coincrud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: _info1.CoinTypeID},
	})
	if err != nil {
		return err
	}

	h.stmSelect = h.selectCoinBase(_stm2)
	return nil
}

func (h *queryHandler) queryCoinBases(cli *ent.Client) (*ent.CoinBaseSelect, error) {
	stm, err := coincrud.SetQueryConds(cli.CoinBase.Query(), &coincrud.Conds{
		EntID:  h.Conds.CoinTypeID,
		EntIDs: h.Conds.CoinTypeIDs,
	})
	if err != nil {
		return nil, err
	}

	return h.selectCoinBase(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcoinbase.Table)
	s.AppendSelect(
		sql.As(t.C(entcoinbase.FieldEntID), "coin_type_id"),
		sql.As(t.C(entcoinbase.FieldName), "coin_name"),
		sql.As(t.C(entcoinbase.FieldLogo), "coin_logo"),
		sql.As(t.C(entcoinbase.FieldUnit), "coin_unit"),
		sql.As(t.C(entcoinbase.FieldEnv), "coin_env"),
	)
}

func (h *queryHandler) queryJoinCoinExtra(s *sql.Selector) {
	t := sql.Table(entcoinextra.Table)
	s.LeftJoin(t).
		On(
			s.C(entcoinbase.FieldEntID),
			t.C(entcoinextra.FieldCoinTypeID),
		).
		AppendSelect(
			sql.As(t.C(entcoinextra.FieldStableUsd), "stable_usd"),
		)
}

func (h *queryHandler) queryJoinCurrency(s *sql.Selector) error {
	t := sql.Table(entcurrency.Table)
	s.LeftJoin(t).
		On(
			s.C(entcoinbase.FieldEntID),
			t.C(entcurrency.FieldCoinTypeID),
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
			return fmt.Errorf("invalid currency field op")
		}
	}
	if h.Conds != nil && h.Conds.FeedType != nil {
		feedType, ok := h.Conds.FeedType.Val.(basetypes.CurrencyFeedType)
		if !ok {
			return fmt.Errorf("invalid feedtype")
		}
		switch h.Conds.FeedType.Op {
		case cruder.EQ:
			s.Where(
				sql.EQ(t.C(entcurrency.FieldFeedType), feedType.String()),
			)
		default:
			return fmt.Errorf("invalid currency field op")
		}
	}
	return nil
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinCoinExtra(s)
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
		if info.StableUSD {
			info.MarketValueHigh = decimal.NewFromInt(1).String()
			info.MarketValueLow = decimal.NewFromInt(1).String()
			info.CreatedAt = uint32(time.Now().Unix())
			info.UpdatedAt = uint32(time.Now().Unix())
			info.FeedType = basetypes.CurrencyFeedType_StableUSDHardCode
		}
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
		if err := handler.queryCoinBase(_ctx, cli); err != nil {
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
		handler.stmSelect, err = handler.queryCoinBases(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryCoinBases(cli)
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
		handler.stmSelect, err = handler.queryCoinBases(cli)
		if err != nil {
			return err
		}

		handler.queryJoin()

		handler.stmSelect.
			Offset(0).
			Limit(1)
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
