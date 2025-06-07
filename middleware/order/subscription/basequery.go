package subscriptionorder

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	orderbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/orderbase"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entordercoupon "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/ordercoupon"
	entorderlock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderlock"
	entorderstatebase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
	entpaymentbalancelock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalancelock"
	entpaymentbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbase"
	entsubscriptionorder "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/subscriptionorder"
	entsubscriptionorderstate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/subscriptionorderstate"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderBaseSelect
}

func (h *baseQueryHandler) selectOrderBase(stm *ent.OrderBaseQuery) *ent.OrderBaseSelect {
	return stm.Select(entorderbase.FieldID)
}

func (h *baseQueryHandler) queryOrderBase(cli *ent.Client) error {
	if h.OrderID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.OrderBase.Query().Where(entorderbase.DeletedAt(0))
	if h.OrderID != nil {
		stm.Where(entorderbase.EntID(*h.OrderID))
	}
	h.stmSelect = h.selectOrderBase(stm)
	return nil
}

func (h *baseQueryHandler) queryOrderBases(cli *ent.Client) (*ent.OrderBaseSelect, error) {
	stm, err := orderbasecrud.SetQueryConds(cli.OrderBase.Query(), h.OrderBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectOrderBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entorderbase.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldID),
			t.C(entorderbase.FieldID),
		)
	s.AppendSelect(
		t.C(entorderbase.FieldAppID),
		t.C(entorderbase.FieldUserID),
		t.C(entorderbase.FieldGoodID),
		t.C(entorderbase.FieldGoodType),
		t.C(entorderbase.FieldAppGoodID),
		t.C(entorderbase.FieldOrderType),
		t.C(entorderbase.FieldCreateMethod),
		t.C(entorderbase.FieldCreatedAt),
		t.C(entorderbase.FieldUpdatedAt),
	)
}

//nolint:gocyclo
func (h *baseQueryHandler) queryJoinSubscriptionOrder(s *sql.Selector) error {
	t := sql.Table(entsubscriptionorder.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entsubscriptionorder.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entsubscriptionorder.FieldDeletedAt), 0),
		)
	if h.SubscriptionOrderConds.ID != nil {
		id, ok := h.SubscriptionOrderConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(
			sql.EQ(t.C(entsubscriptionorder.FieldID), id),
		)
	}
	if h.SubscriptionOrderConds.IDs != nil {
		ids, ok := h.SubscriptionOrderConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(sql.In(t.C(entsubscriptionorder.FieldID), func() (_ids []interface{}) {
			for _, id := range ids {
				_ids = append(_ids, interface{}(id))
			}
			return _ids
		}()...))
	}
	if h.SubscriptionOrderConds.EntID != nil {
		id, ok := h.SubscriptionOrderConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(
			sql.EQ(t.C(entsubscriptionorder.FieldEntID), id),
		)
	}
	if h.SubscriptionOrderConds.EntIDs != nil {
		uids, ok := h.SubscriptionOrderConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(sql.In(t.C(entsubscriptionorder.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	if h.SubscriptionOrderConds.OrderID != nil {
		id, ok := h.SubscriptionOrderConds.OrderID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderid")
		}
		s.OnP(
			sql.EQ(t.C(entsubscriptionorder.FieldOrderID), id),
		)
	}
	if h.SubscriptionOrderConds.OrderIDs != nil {
		uids, ok := h.SubscriptionOrderConds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderids")
		}
		s.OnP(sql.In(t.C(entsubscriptionorder.FieldOrderID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	s.AppendSelect(
		t.C(entsubscriptionorder.FieldID),
		t.C(entsubscriptionorder.FieldEntID),
		t.C(entsubscriptionorder.FieldOrderID),
		t.C(entsubscriptionorder.FieldGoodValueUsd),
		t.C(entsubscriptionorder.FieldPaymentAmountUsd),
		t.C(entsubscriptionorder.FieldDiscountAmountUsd),
		t.C(entsubscriptionorder.FieldPromotionID),
		t.C(entsubscriptionorder.FieldDurationSeconds),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinOrderStateBase(s *sql.Selector) error {
	t := sql.Table(entorderstatebase.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entorderstatebase.FieldOrderID),
		)
	if h.OrderStateBaseConds.PaymentType != nil {
		_type, ok := h.OrderStateBaseConds.PaymentType.Val.(types.PaymentType)
		if !ok {
			return wlog.Errorf("invalid paymenttype")
		}
		s.OnP(
			sql.EQ(t.C(entorderstatebase.FieldPaymentType), _type.String()),
		)
	}
	if h.OrderStateBaseConds.PaymentTypes != nil {
		_types, ok := h.OrderStateBaseConds.PaymentTypes.Val.([]types.PaymentType)
		if !ok {
			return wlog.Errorf("invalid paymenttypes")
		}
		s.OnP(sql.In(t.C(entorderstatebase.FieldPaymentType), func() (__types []interface{}) {
			for _, _type := range _types {
				__types = append(__types, interface{}(_type.String()))
			}
			return __types
		}()...))
	}
	if h.OrderStateBaseConds.OrderState != nil {
		_state, ok := h.OrderStateBaseConds.OrderState.Val.(types.OrderState)
		if !ok {
			return wlog.Errorf("invalid orderstate")
		}
		s.OnP(
			sql.EQ(t.C(entorderstatebase.FieldOrderState), _state.String()),
		)
	}
	if h.OrderStateBaseConds.OrderStates != nil {
		states, ok := h.OrderStateBaseConds.OrderStates.Val.([]types.OrderState)
		if !ok {
			return wlog.Errorf("invalid orderstates")
		}
		s.OnP(sql.In(t.C(entorderstatebase.FieldOrderState), func() (_states []interface{}) {
			for _, _state := range states {
				_states = append(_states, interface{}(_state.String()))
			}
			return _states
		}()...))
	}
	s.AppendSelect(
		t.C(entorderstatebase.FieldPaymentType),
		t.C(entorderstatebase.FieldOrderState),
	)
	return nil
}

//nolint:gocyclo
func (h *baseQueryHandler) queryJoinSubscriptionOrderState(s *sql.Selector) error {
	t := sql.Table(entsubscriptionorderstate.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entsubscriptionorderstate.FieldOrderID),
		)
	if h.SubscriptionOrderStateConds.PaymentState != nil {
		_state, ok := h.SubscriptionOrderStateConds.PaymentState.Val.(types.PaymentState)
		if !ok {
			return wlog.Errorf("invalid paymentstate")
		}
		s.OnP(
			sql.EQ(t.C(entsubscriptionorderstate.FieldPaymentState), _state.String()),
		)
	}
	if h.SubscriptionOrderStateConds.PaymentStates != nil {
		states, ok := h.SubscriptionOrderStateConds.PaymentStates.Val.([]types.PaymentState)
		if !ok {
			return wlog.Errorf("invalid paymentstates")
		}
		s.OnP(sql.In(t.C(entsubscriptionorderstate.FieldPaymentState), func() (_states []interface{}) {
			for _, _state := range states {
				_states = append(_states, interface{}(_state.String()))
			}
			return
		}()...))
	}
	if h.SubscriptionOrderStateConds.UserSetCanceled != nil {
		b, ok := h.SubscriptionOrderStateConds.UserSetCanceled.Val.(bool)
		if !ok {
			return wlog.Errorf("invalid usersetcanceled")
		}
		s.OnP(
			sql.EQ(t.C(entsubscriptionorderstate.FieldUserSetCanceled), b),
		)
	}
	if h.SubscriptionOrderStateConds.AdminSetCanceled != nil {
		b, ok := h.SubscriptionOrderStateConds.AdminSetCanceled.Val.(bool)
		if !ok {
			return wlog.Errorf("invalid usersetcanceled")
		}
		s.OnP(
			sql.EQ(t.C(entsubscriptionorderstate.FieldAdminSetCanceled), b),
		)
	}
	if h.SubscriptionOrderStateConds.PaidAt != nil {
		b, ok := h.SubscriptionOrderStateConds.PaidAt.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid usersetcanceled")
		}
		switch h.SubscriptionOrderStateConds.PaidAt.Op {
		case cruder.EQ:
			s.OnP(
				sql.EQ(t.C(entsubscriptionorderstate.FieldPaidAt), b),
			)
		case cruder.LT:
			s.OnP(
				sql.LT(t.C(entsubscriptionorderstate.FieldPaidAt), b),
			)
		case cruder.LTE:
			s.OnP(
				sql.LTE(t.C(entsubscriptionorderstate.FieldPaidAt), b),
			)
		case cruder.GT:
			s.OnP(
				sql.GT(t.C(entsubscriptionorderstate.FieldPaidAt), b),
			)
		case cruder.GTE:
			s.OnP(
				sql.GTE(t.C(entsubscriptionorderstate.FieldPaidAt), b),
			)
		default:
			return wlog.Errorf("invalid paidat")
		}
	}
	s.AppendSelect(
		t.C(entsubscriptionorderstate.FieldPaymentID),
		t.C(entsubscriptionorderstate.FieldPaidAt),
		t.C(entsubscriptionorderstate.FieldUserSetPaid),
		t.C(entsubscriptionorderstate.FieldUserSetCanceled),
		t.C(entsubscriptionorderstate.FieldAdminSetCanceled),
		t.C(entsubscriptionorderstate.FieldPaymentState),
		t.C(entsubscriptionorderstate.FieldCancelState),
		t.C(entsubscriptionorderstate.FieldCanceledAt),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinPaymentBase(s *sql.Selector) {
	t1 := sql.Table(entpaymentbase.Table)
	t2 := sql.Table(entpaymentbalancelock.Table)
	t3 := sql.Table(entorderlock.Table)

	s.LeftJoin(t1).
		On(
			s.C(entorderbase.FieldEntID),
			t1.C(entpaymentbase.FieldOrderID),
		).
		OnP(
			sql.EQ(t1.C(entpaymentbase.FieldObseleteState), types.PaymentObseleteState_PaymentObseleteNone.String()),
		).
		LeftJoin(t2).
		On(
			t1.C(entpaymentbase.FieldEntID),
			t2.C(entpaymentbalancelock.FieldPaymentID),
		).
		LeftJoin(t3).
		On(
			t2.C(entpaymentbalancelock.FieldLedgerLockID),
			t3.C(entorderlock.FieldEntID),
		)
	s.AppendSelect(
		sql.As(t1.C(entpaymentbase.FieldEntID), "payment_id"),
		sql.As(t3.C(entorderlock.FieldEntID), "ledger_lock_id"),
	)
}

func (h *baseQueryHandler) queryJoinOrderCoupon(s *sql.Selector) error {
	t := sql.Table(entordercoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entordercoupon.FieldOrderID),
		).
		Distinct()
	if h.OrderCouponConds.CouponID != nil {
		uid, ok := h.OrderCouponConds.CouponID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid couponid")
		}
		s.OnP(
			sql.EQ(t.C(entordercoupon.FieldCouponID), uid),
		)
		s.Where(
			sql.EQ(t.C(entordercoupon.FieldCouponID), uid),
		)
	}
	if h.OrderCouponConds.CouponIDs != nil {
		uids, ok := h.OrderCouponConds.CouponIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid couponids")
		}
		s.OnP(sql.In(t.C(entordercoupon.FieldCouponID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
		s.Where(sql.In(t.C(entordercoupon.FieldCouponID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinSubscriptionOrder(s); err != nil {
			logger.Sugar().Errorw("queryJoinSubscriptionOrder", "Error", err)
		}
		if err := h.queryJoinOrderStateBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStateBase", "Error", err)
		}
		if err := h.queryJoinSubscriptionOrderState(s); err != nil {
			logger.Sugar().Errorw("queryJoinSubscriptionOrderState", "Error", err)
		}
		h.queryJoinPaymentBase(s)
		if err := h.queryJoinOrderCoupon(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderCoupon", "Error", err)
		}
	})
}
