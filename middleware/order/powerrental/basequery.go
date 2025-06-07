//nolint:dupl
package powerrental

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	orderbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/orderbase"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entordercoupon "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/ordercoupon"
	entorderlock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderlock"
	entorderstatebase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
	entpaymentbalancelock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalancelock"
	entpaymentbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbase"
	entpoolorderuser "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/poolorderuser"
	entpowerrental "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/powerrental"
	entpowerrentalstate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/powerrentalstate"
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
		t.C(entorderbase.FieldSimulate),
		t.C(entorderbase.FieldCreatedAt),
		t.C(entorderbase.FieldUpdatedAt),
	)
}

//nolint:gocyclo,funlen
func (h *baseQueryHandler) queryJoinPowerRental(s *sql.Selector) error {
	t := sql.Table(entpowerrental.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entpowerrental.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entpowerrental.FieldDeletedAt), 0),
		)
	if h.PowerRentalConds.ID != nil {
		id, ok := h.PowerRentalConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrental.FieldID), id),
		)
	}
	if h.PowerRentalConds.IDs != nil {
		ids, ok := h.PowerRentalConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(sql.In(t.C(entpowerrental.FieldID), func() (_ids []interface{}) {
			for _, id := range ids {
				_ids = append(_ids, interface{}(id))
			}
			return _ids
		}()...))
	}
	if h.PowerRentalConds.EntID != nil {
		id, ok := h.PowerRentalConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrental.FieldEntID), id),
		)
	}
	if h.PowerRentalConds.EntIDs != nil {
		uids, ok := h.PowerRentalConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(sql.In(t.C(entpowerrental.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	if h.PowerRentalConds.OrderID != nil {
		id, ok := h.PowerRentalConds.OrderID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderid")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrental.FieldOrderID), id),
		)
	}
	if h.PowerRentalConds.OrderIDs != nil {
		uids, ok := h.PowerRentalConds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderids")
		}
		s.OnP(sql.In(t.C(entpowerrental.FieldOrderID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	if h.PowerRentalConds.GoodStockMode != nil {
		_type, ok := h.PowerRentalConds.GoodStockMode.Val.(goodtypes.GoodStockMode)
		if !ok {
			return wlog.Errorf("invalid goodstockmode")
		}
		switch h.PowerRentalConds.GoodStockMode.Op {
		case cruder.EQ:
			s.OnP(
				sql.EQ(t.C(entpowerrental.FieldGoodStockMode), _type.String()),
			)
		case cruder.NEQ:
			s.OnP(
				sql.NEQ(t.C(entpowerrental.FieldGoodStockMode), _type.String()),
			)
		}
	}
	s.AppendSelect(
		t.C(entpowerrental.FieldID),
		t.C(entpowerrental.FieldEntID),
		t.C(entpowerrental.FieldOrderID),
		t.C(entpowerrental.FieldAppGoodStockID),
		t.C(entpowerrental.FieldUnits),
		t.C(entpowerrental.FieldGoodValueUsd),
		t.C(entpowerrental.FieldPaymentAmountUsd),
		t.C(entpowerrental.FieldDiscountAmountUsd),
		t.C(entpowerrental.FieldPromotionID),
		t.C(entpowerrental.FieldInvestmentType),
		t.C(entpowerrental.FieldGoodStockMode),
		t.C(entpowerrental.FieldDurationSeconds),
	)
	return nil
}

//nolint:gocyclo,funlen
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
		switch h.OrderStateBaseConds.PaymentType.Op {
		case cruder.EQ:
			s.OnP(
				sql.EQ(t.C(entorderstatebase.FieldPaymentType), _type.String()),
			)
		case cruder.NEQ:
			s.OnP(
				sql.NEQ(t.C(entorderstatebase.FieldPaymentType), _type.String()),
			)
		}
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
		switch h.OrderStateBaseConds.OrderState.Op {
		case cruder.EQ:
			s.OnP(
				sql.EQ(t.C(entorderstatebase.FieldOrderState), _state.String()),
			)
		case cruder.NEQ:
			s.OnP(
				sql.NEQ(t.C(entorderstatebase.FieldOrderState), _state.String()),
			)
		}
	}
	if h.OrderStateBaseConds.OrderStates != nil {
		states, ok := h.OrderStateBaseConds.OrderStates.Val.([]types.OrderState)
		if !ok {
			return wlog.Errorf("invalid orderstate")
		}
		switch h.OrderStateBaseConds.OrderStates.Op {
		case cruder.IN:
			s.OnP(sql.In(t.C(entorderstatebase.FieldOrderState), func() (_states []interface{}) {
				for _, _state := range states {
					_states = append(_states, interface{}(_state.String()))
				}
				return _states
			}()...))
		case cruder.NIN:
			s.OnP(sql.NotIn(t.C(entorderstatebase.FieldOrderState), func() (_states []interface{}) {
				for _, _state := range states {
					_states = append(_states, interface{}(_state.String()))
				}
				return _states
			}()...))
		default:
			return wlog.Errorf("invalid orderstates")
		}
	}
	if h.OrderStateBaseConds.StartMode != nil {
		mode, ok := h.OrderStateBaseConds.StartMode.Val.(types.OrderStartMode)
		if !ok {
			return wlog.Errorf("invalid startmode")
		}
		s.OnP(
			sql.EQ(t.C(entorderstatebase.FieldStartMode), mode.String()),
		)
	}
	if h.OrderStateBaseConds.LastBenefitAt != nil {
		at, ok := h.OrderStateBaseConds.LastBenefitAt.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid lastbenefitat")
		}
		s.OnP(
			sql.EQ(t.C(entorderstatebase.FieldLastBenefitAt), at),
		)
	}
	if h.OrderStateBaseConds.StartAt != nil {
		at, ok := h.OrderStateBaseConds.StartAt.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid lastbenefitat")
		}
		switch h.OrderStateBaseConds.StartAt.Op {
		case cruder.LT:
			s.OnP(
				sql.LT(t.C(entorderstatebase.FieldStartAt), at),
			)
		case cruder.LTE:
			s.OnP(
				sql.LTE(t.C(entorderstatebase.FieldStartAt), at),
			)
		case cruder.GT:
			s.OnP(
				sql.GT(t.C(entorderstatebase.FieldStartAt), at),
			)
		case cruder.GTE:
			s.OnP(
				sql.GTE(t.C(entorderstatebase.FieldStartAt), at),
			)
		case cruder.EQ:
			s.OnP(
				sql.EQ(t.C(entorderstatebase.FieldStartAt), at),
			)
		}
	}
	if h.OrderStateBaseConds.BenefitState != nil {
		_state, ok := h.OrderStateBaseConds.BenefitState.Val.(types.BenefitState)
		if !ok {
			return wlog.Errorf("invalid benefitstate")
		}
		s.OnP(
			sql.EQ(t.C(entorderstatebase.FieldBenefitState), _state.String()),
		)
	}
	s.AppendSelect(
		t.C(entorderstatebase.FieldPaymentType),
		t.C(entorderstatebase.FieldOrderState),
		t.C(entorderstatebase.FieldStartMode),
		t.C(entorderstatebase.FieldStartAt),
		t.C(entorderstatebase.FieldLastBenefitAt),
		t.C(entorderstatebase.FieldBenefitState),
	)
	return nil
}

//nolint:funlen,gocyclo
func (h *baseQueryHandler) queryJoinPowerRentalState(s *sql.Selector) error {
	t := sql.Table(entpowerrentalstate.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entpowerrentalstate.FieldOrderID),
		)
	if h.PowerRentalStateConds.PaymentState != nil {
		_state, ok := h.PowerRentalStateConds.PaymentState.Val.(types.PaymentState)
		if !ok {
			return wlog.Errorf("invalid paymentstate")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrentalstate.FieldPaymentState), _state.String()),
		)
	}
	if h.PowerRentalStateConds.PaymentStates != nil {
		states, ok := h.PowerRentalStateConds.PaymentStates.Val.([]types.PaymentState)
		if !ok {
			return wlog.Errorf("invalid paymentstates")
		}
		s.OnP(sql.In(t.C(entpowerrentalstate.FieldPaymentState), func() (_states []interface{}) {
			for _, _state := range states {
				_states = append(_states, interface{}(_state.String()))
			}
			return
		}()...))
	}
	if h.PowerRentalStateConds.UserSetCanceled != nil {
		b, ok := h.PowerRentalStateConds.UserSetCanceled.Val.(bool)
		if !ok {
			return wlog.Errorf("invalid usersetcanceled")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrentalstate.FieldUserSetCanceled), b),
		)
	}
	if h.PowerRentalStateConds.AdminSetCanceled != nil {
		b, ok := h.PowerRentalStateConds.AdminSetCanceled.Val.(bool)
		if !ok {
			return wlog.Errorf("invalid adminsetcanceled")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrentalstate.FieldAdminSetCanceled), b),
		)
	}
	if h.PowerRentalStateConds.PaidAt != nil {
		b, ok := h.PowerRentalStateConds.PaidAt.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid paidat")
		}
		switch h.PowerRentalStateConds.PaidAt.Op {
		case cruder.EQ:
			s.OnP(
				sql.EQ(t.C(entpowerrentalstate.FieldPaidAt), b),
			)
		case cruder.LT:
			s.OnP(
				sql.LT(t.C(entpowerrentalstate.FieldPaidAt), b),
			)
		case cruder.LTE:
			s.OnP(
				sql.LTE(t.C(entpowerrentalstate.FieldPaidAt), b),
			)
		case cruder.GT:
			s.OnP(
				sql.GT(t.C(entpowerrentalstate.FieldPaidAt), b),
			)
		case cruder.GTE:
			s.OnP(
				sql.GTE(t.C(entpowerrentalstate.FieldPaidAt), b),
			)
		default:
			return wlog.Errorf("invalid paidat")
		}
	}
	if h.PowerRentalStateConds.RenewState != nil {
		state, ok := h.PowerRentalStateConds.RenewState.Val.(types.OrderRenewState)
		if !ok {
			return wlog.Errorf("invalid renewstate")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrentalstate.FieldRenewState), state.String()),
		)
	}
	if h.PowerRentalStateConds.RenewNotifyAt != nil {
		b, ok := h.PowerRentalStateConds.RenewNotifyAt.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid renewnotifyat")
		}
		switch h.PowerRentalStateConds.RenewNotifyAt.Op {
		case cruder.EQ:
			s.OnP(
				sql.EQ(t.C(entpowerrentalstate.FieldRenewNotifyAt), b),
			)
		case cruder.LT:
			s.OnP(
				sql.LT(t.C(entpowerrentalstate.FieldRenewNotifyAt), b),
			)
		case cruder.LTE:
			s.OnP(
				sql.LTE(t.C(entpowerrentalstate.FieldRenewNotifyAt), b),
			)
		case cruder.GT:
			s.OnP(
				sql.GT(t.C(entpowerrentalstate.FieldRenewNotifyAt), b),
			)
		case cruder.GTE:
			s.OnP(
				sql.GTE(t.C(entpowerrentalstate.FieldRenewNotifyAt), b),
			)
		default:
			return wlog.Errorf("invalid renewnotifyat")
		}
	}
	s.AppendSelect(
		t.C(entpowerrentalstate.FieldPaymentID),
		t.C(entpowerrentalstate.FieldPaidAt),
		t.C(entpowerrentalstate.FieldUserSetPaid),
		t.C(entpowerrentalstate.FieldUserSetCanceled),
		t.C(entpowerrentalstate.FieldAdminSetCanceled),
		t.C(entpowerrentalstate.FieldPaymentState),
		t.C(entpowerrentalstate.FieldOutofgasSeconds),
		t.C(entpowerrentalstate.FieldCompensateSeconds),
		t.C(entpowerrentalstate.FieldCancelState),
		t.C(entpowerrentalstate.FieldCanceledAt),
		t.C(entpowerrentalstate.FieldRenewState),
		t.C(entpowerrentalstate.FieldRenewNotifyAt),
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

func (h *baseQueryHandler) queryJoinStockLock(s *sql.Selector) {
	t := sql.Table(entorderlock.Table)
	s.LeftJoin(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entorderlock.FieldOrderID),
		).
		OnP(
			sql.And(
				sql.EQ(t.C(entorderlock.FieldLockType), types.OrderLockType_LockStock.String()),
				sql.EQ(t.C(entorderlock.FieldDeletedAt), 0),
			),
		)
	s.AppendSelect(
		sql.As(t.C(entorderlock.FieldEntID), "app_good_stock_lock_id"),
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
		id, ok := h.OrderCouponConds.CouponID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid couponid")
		}
		s.OnP(
			sql.EQ(t.C(entordercoupon.FieldCouponID), id),
		)
		s.Where(
			sql.EQ(t.C(entordercoupon.FieldCouponID), id),
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

func (h *baseQueryHandler) queryJoinPoolOrderUser(s *sql.Selector) {
	t := sql.Table(entpoolorderuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entpoolorderuser.FieldOrderID),
		).
		OnP(
			sql.And(
				sql.EQ(t.C(entpoolorderuser.FieldDeletedAt), 0),
			),
		)
	s.AppendSelect(
		sql.As(t.C(entpoolorderuser.FieldPoolOrderUserID), "mining_pool_order_user_id"),
	)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinPowerRental(s); err != nil {
			logger.Sugar().Errorw("queryJoinPowerRental", "Error", err)
		}
		if err := h.queryJoinOrderStateBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStateBase", "Error", err)
		}
		if err := h.queryJoinPowerRentalState(s); err != nil {
			logger.Sugar().Errorw("queryJoinPowerRentalState", "Error", err)
		}
		h.queryJoinPaymentBase(s)
		h.queryJoinStockLock(s)
		h.queryJoinPoolOrderUser(s)
		if err := h.queryJoinOrderCoupon(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderCoupon", "Error", err)
		}
	})
}
