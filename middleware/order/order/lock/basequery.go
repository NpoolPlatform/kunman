package orderlock

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	orderlockcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/lock"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entorderlock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderlock"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderLockSelect
}

func (h *baseQueryHandler) selectOrderLock(stm *ent.OrderLockQuery) *ent.OrderLockSelect {
	return stm.Select(entorderlock.FieldID)
}

func (h *baseQueryHandler) queryOrderLock(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.OrderLock.Query().Where(entorderlock.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entorderlock.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entorderlock.EntID(*h.EntID))
	}
	h.stmSelect = h.selectOrderLock(stm)
	return nil
}

func (h *baseQueryHandler) queryOrderLocks(cli *ent.Client) (*ent.OrderLockSelect, error) {
	stm, err := orderlockcrud.SetQueryConds(cli.OrderLock.Query(), h.OrderLockConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectOrderLock(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entorderlock.Table)
	s.Join(t1).On(
		s.C(entorderlock.FieldID),
		t1.C(entorderlock.FieldID),
	).AppendSelect(
		t1.C(entorderlock.FieldEntID),
		t1.C(entorderlock.FieldUserID),
		t1.C(entorderlock.FieldOrderID),
		t1.C(entorderlock.FieldLockType),
		t1.C(entorderlock.FieldCreatedAt),
		t1.C(entorderlock.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoinOrderBase(s *sql.Selector) error {
	t1 := sql.Table(entorderbase.Table)
	s.Join(t1).On(
		s.C(entorderlock.FieldOrderID),
		t1.C(entorderbase.FieldEntID),
	)
	if h.OrderBaseConds.AppID != nil {
		id, ok := h.OrderBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t1.C(entorderbase.FieldAppID), id),
		)
	}
	s.AppendSelect(
		t1.C(entorderbase.FieldAppID),
		sql.As(t1.C(entorderbase.FieldUserID), "order_user_id"),
		t1.C(entorderbase.FieldGoodType),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinOrderBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderBase", "Error", err)
		}
	})
}
