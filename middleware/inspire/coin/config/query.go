package config

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	devicecrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coin/config"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db"
	"github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoinconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coinconfig"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coin/config"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CoinConfigSelect
	stmCount  *ent.CoinConfigSelect
	infos     []*npool.CoinConfig
	total     uint32
}

func (h *queryHandler) selectCoinConfig(stm *ent.CoinConfigQuery) {
	h.stmSelect = stm.Select(entcoinconfig.FieldID)
}

func (h *queryHandler) queryCoinConfig(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.CoinConfig.Query().Where(entcoinconfig.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcoinconfig.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcoinconfig.EntID(*h.EntID))
	}
	h.selectCoinConfig(stm)
	return nil
}

func (h *queryHandler) queryCoinConfigs(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.CoinConfig.Query(), h.Conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.total = uint32(total)
	h.selectCoinConfig(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entcoinconfig.Table)
	s.LeftJoin(t1).
		On(
			s.C(entcoinconfig.FieldEntID),
			t1.C(entcoinconfig.FieldEntID),
		).
		AppendSelect(
			t1.C(entcoinconfig.FieldEntID),
			t1.C(entcoinconfig.FieldAppID),
			t1.C(entcoinconfig.FieldCoinTypeID),
			t1.C(entcoinconfig.FieldMaxValue),
			t1.C(entcoinconfig.FieldAllocated),
			t1.C(entcoinconfig.FieldCreatedAt),
			t1.C(entcoinconfig.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
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
		amount, err := decimal.NewFromString(info.MaxValue)
		if err != nil {
			info.MaxValue = decimal.NewFromInt(0).String()
		} else {
			info.MaxValue = amount.String()
		}
		amount, err = decimal.NewFromString(info.Allocated)
		if err != nil {
			info.Allocated = decimal.NewFromInt(0).String()
		} else {
			info.Allocated = amount.String()
		}
	}
}

func (h *Handler) GetCoinConfig(ctx context.Context) (*npool.CoinConfig, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinConfig(cli); err != nil {
			return wlog.WrapError(err)
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
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetCoinConfigs(ctx context.Context) ([]*npool.CoinConfig, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCoinConfigs(_ctx, cli); err != nil {
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
