package coin

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	devicecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/event/coin"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoinconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coinconfig"
	enteventcoin "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/eventcoin"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event/coin"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.EventCoinSelect
	stmCount  *ent.EventCoinSelect
	infos     []*npool.EventCoin
	total     uint32
}

func (h *queryHandler) selectEventCoin(stm *ent.EventCoinQuery) {
	h.stmSelect = stm.Select(enteventcoin.FieldID)
}

func (h *queryHandler) queryEventCoin(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.EventCoin.Query().Where(enteventcoin.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enteventcoin.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enteventcoin.EntID(*h.EntID))
	}
	h.selectEventCoin(stm)
	return nil
}

func (h *queryHandler) queryEventCoins(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.EventCoin.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)
	h.selectEventCoin(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(enteventcoin.Table)
	s.LeftJoin(t1).
		On(
			s.C(enteventcoin.FieldEntID),
			t1.C(enteventcoin.FieldEntID),
		).
		AppendSelect(
			t1.C(enteventcoin.FieldEntID),
			t1.C(enteventcoin.FieldAppID),
			t1.C(enteventcoin.FieldEventID),
			t1.C(enteventcoin.FieldCoinConfigID),
			t1.C(enteventcoin.FieldCoinValue),
			t1.C(enteventcoin.FieldCoinPerUsd),
			t1.C(enteventcoin.FieldCreatedAt),
			t1.C(enteventcoin.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinCoinConfig(s *sql.Selector) {
	t := sql.Table(entcoinconfig.Table)
	s.LeftJoin(t).
		On(
			s.C(enteventcoin.FieldCoinConfigID),
			t.C(entcoinconfig.FieldEntID),
		)

	s.AppendSelect(
		sql.As(t.C(entcoinconfig.FieldCoinTypeID), "coin_type_id"),
	)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinCoinConfig(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmSelect.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.CoinValue)
		if err != nil {
			info.CoinValue = decimal.NewFromInt(0).String()
		} else {
			info.CoinValue = amount.String()
		}
		amount, err = decimal.NewFromString(info.CoinPerUSD)
		if err != nil {
			info.CoinPerUSD = decimal.NewFromInt(0).String()
		} else {
			info.CoinPerUSD = amount.String()
		}
	}
}

func (h *Handler) GetEventCoin(ctx context.Context) (*npool.EventCoin, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEventCoin(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
		return handler.scan(_ctx)
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

func (h *Handler) GetEventCoins(ctx context.Context) ([]*npool.EventCoin, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryEventCoins(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos, handler.total, nil
}
