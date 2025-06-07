package outofgas

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	outofgascrud "github.com/NpoolPlatform/kunman/middleware/order/crud/outofgas"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entoutofgas "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/outofgas"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OutOfGasSelect
}

func (h *baseQueryHandler) selectOutOfGas(stm *ent.OutOfGasQuery) *ent.OutOfGasSelect {
	return stm.Select(entoutofgas.FieldID)
}

func (h *baseQueryHandler) queryOutOfGas(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.OutOfGas.Query().Where(entoutofgas.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entoutofgas.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entoutofgas.EntID(*h.EntID))
	}
	h.stmSelect = h.selectOutOfGas(stm)
	return nil
}

func (h *baseQueryHandler) queryOutOfGases(cli *ent.Client) (*ent.OutOfGasSelect, error) {
	stm, err := outofgascrud.SetQueryConds(cli.OutOfGas.Query(), h.OutOfGasConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectOutOfGas(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entoutofgas.Table)
	s.AppendSelect(
		t.C(entoutofgas.FieldEntID),
		t.C(entoutofgas.FieldOrderID),
		t.C(entoutofgas.FieldStartAt),
		t.C(entoutofgas.FieldEndAt),
		t.C(entoutofgas.FieldCreatedAt),
		t.C(entoutofgas.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoinOrder(s *sql.Selector) error {
	t := sql.Table(entorderbase.Table)
	s.Join(t).
		On(
			s.C(entoutofgas.FieldOrderID),
			t.C(entorderbase.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entorderbase.FieldDeletedAt), 0),
		)
	if h.OrderBaseConds.AppID != nil {
		id, ok := h.OrderBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldAppID), id),
		)
	}
	if h.OrderBaseConds.UserID != nil {
		id, ok := h.OrderBaseConds.UserID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid userid")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldUserID), id),
		)
	}
	if h.OrderBaseConds.GoodID != nil {
		id, ok := h.OrderBaseConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldGoodID), id),
		)
	}
	if h.OrderBaseConds.AppGoodID != nil {
		id, ok := h.OrderBaseConds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldAppGoodID), id),
		)
	}
	s.AppendSelect(
		t.C(entorderbase.FieldAppID),
		t.C(entorderbase.FieldUserID),
		t.C(entorderbase.FieldGoodID),
		t.C(entorderbase.FieldGoodType),
		t.C(entorderbase.FieldAppGoodID),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinOrder(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrder", "Error", err)
		}
	})
}
