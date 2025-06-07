package order

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/kunman/framework/logger"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	orderbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/orderbase"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entorderstatebase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderBaseSelect
}

func (h *baseQueryHandler) selectOrderBase(stm *ent.OrderBaseQuery) *ent.OrderBaseSelect {
	return stm.Select(entorderbase.FieldID)
}

func (h *baseQueryHandler) queryOrderBase(cli *ent.Client) error {
	if h.EntID == nil && h.ID == nil {
		return wlog.Errorf("invalid entid")
	}
	stm := cli.OrderBase.
		Query().
		Where(
			entorderbase.DeletedAt(0),
		)
	if h.ID != nil {
		stm.Where(entorderbase.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entorderbase.EntID(*h.EntID))
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
		).
		AppendSelect(
			t.C(entorderbase.FieldEntID),
			t.C(entorderbase.FieldAppID),
			t.C(entorderbase.FieldUserID),
			t.C(entorderbase.FieldGoodID),
			t.C(entorderbase.FieldGoodType),
			t.C(entorderbase.FieldAppGoodID),
			t.C(entorderbase.FieldParentOrderID),
			t.C(entorderbase.FieldOrderType),
			t.C(entorderbase.FieldCreateMethod),
			t.C(entorderbase.FieldSimulate),
			t.C(entorderbase.FieldCreatedAt),
			t.C(entorderbase.FieldUpdatedAt),
		)
}

//nolint:gocyclo
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

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinOrderStateBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStateBase", "Error", err)
		}
	})
}
