package ordercoupon

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercouponcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/coupon"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entordercoupon "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/ordercoupon"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderCouponSelect
}

func (h *baseQueryHandler) selectOrderCoupon(stm *ent.OrderCouponQuery) *ent.OrderCouponSelect {
	return stm.Select(entordercoupon.FieldID)
}

func (h *baseQueryHandler) queryOrderCoupon(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil && h.CouponID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.OrderCoupon.Query().Where(entordercoupon.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entordercoupon.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entordercoupon.EntID(*h.EntID))
	}
	if h.CouponID != nil {
		stm.Where(entordercoupon.CouponID(*h.CouponID))
	}
	h.stmSelect = h.selectOrderCoupon(stm)
	return nil
}

func (h *baseQueryHandler) queryOrderCoupons(cli *ent.Client) (*ent.OrderCouponSelect, error) {
	stm, err := ordercouponcrud.SetQueryConds(cli.OrderCoupon.Query(), h.OrderCouponConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectOrderCoupon(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entordercoupon.Table)
	s.AppendSelect(
		t.C(entordercoupon.FieldEntID),
		t.C(entordercoupon.FieldOrderID),
		t.C(entordercoupon.FieldCouponID),
		t.C(entordercoupon.FieldCreatedAt),
		t.C(entordercoupon.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoinOrder(s *sql.Selector) error {
	t := sql.Table(entorderbase.Table)
	s.Join(t).
		On(
			s.C(entordercoupon.FieldOrderID),
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
