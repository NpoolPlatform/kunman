package orderstatebase

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderstatebase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID         *uuid.UUID
	OrderID       *uuid.UUID
	OrderState    *types.OrderState
	StartMode     *types.OrderStartMode
	StartAt       *uint32
	LastBenefitAt *uint32
	BenefitState  *types.BenefitState
	PaymentType   *types.PaymentType
	DeletedAt     *uint32
}

func CreateSet(c *ent.OrderStateBaseCreate, req *Req) *ent.OrderStateBaseCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.OrderState != nil {
		c.SetOrderState(req.OrderState.String())
	}
	if req.StartMode != nil {
		c.SetStartMode(req.StartMode.String())
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.LastBenefitAt != nil {
		c.SetLastBenefitAt(*req.LastBenefitAt)
	}
	if req.BenefitState != nil {
		c.SetBenefitState(req.BenefitState.String())
	}
	if req.PaymentType != nil {
		c.SetPaymentType(req.PaymentType.String())
	}
	return c
}

func UpdateSet(u *ent.OrderStateBaseUpdateOne, req *Req) *ent.OrderStateBaseUpdateOne {
	if req.OrderState != nil {
		u.SetOrderState(req.OrderState.String())
	}
	if req.StartMode != nil {
		u.SetStartMode(req.StartMode.String())
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.LastBenefitAt != nil {
		u.SetLastBenefitAt(*req.LastBenefitAt)
	}
	if req.BenefitState != nil {
		u.SetBenefitState(req.BenefitState.String())
	}
	if req.PaymentType != nil {
		u.SetPaymentType(req.PaymentType.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID            *cruder.Cond
	IDs           *cruder.Cond
	EntID         *cruder.Cond
	EntIDs        *cruder.Cond
	OrderID       *cruder.Cond
	OrderState    *cruder.Cond
	OrderStates   *cruder.Cond
	StartMode     *cruder.Cond
	LastBenefitAt *cruder.Cond
	StartAt       *cruder.Cond
	BenefitState  *cruder.Cond
	PaymentType   *cruder.Cond
	PaymentTypes  *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.OrderStateBaseQuery, conds *Conds) (*ent.OrderStateBaseQuery, error) {
	q.Where(entorderstatebase.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entorderstatebase.ID(id))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		if len(ids) > 0 {
			switch conds.IDs.Op {
			case cruder.IN:
				q.Where(entorderstatebase.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid order field")
			}
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entorderstatebase.EntID(id))
		case cruder.NEQ:
			q.Where(entorderstatebase.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		if len(ids) > 0 {
			switch conds.EntIDs.Op {
			case cruder.IN:
				q.Where(entorderstatebase.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid order field")
			}
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entorderstatebase.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.OrderState != nil {
		state, ok := conds.OrderState.Val.(types.OrderState)
		if !ok {
			return nil, wlog.Errorf("invalid orderstate")
		}
		switch conds.OrderState.Op {
		case cruder.EQ:
			q.Where(entorderstatebase.OrderState(state.String()))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.OrderStates != nil {
		states, ok := conds.OrderStates.Val.([]types.OrderState)
		if !ok {
			return nil, wlog.Errorf("invalid orderstate")
		}
		switch conds.OrderStates.Op {
		case cruder.IN:
			q.Where(entorderstatebase.OrderStateIn(func() (_states []string) {
				for _, state := range states {
					_states = append(_states, state.String())
				}
				return
			}()...))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.StartMode != nil {
		startmode, ok := conds.StartMode.Val.(types.OrderStartMode)
		if !ok {
			return nil, wlog.Errorf("invalid startmode")
		}
		switch conds.StartMode.Op {
		case cruder.EQ:
			q.Where(entorderstatebase.StartMode(startmode.String()))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.BenefitState != nil {
		state, ok := conds.BenefitState.Val.(types.BenefitState)
		if !ok {
			return nil, wlog.Errorf("invalid benefitstate")
		}
		switch conds.BenefitState.Op {
		case cruder.EQ:
			q.Where(entorderstatebase.BenefitState(state.String()))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.PaymentType != nil {
		paymenttype, ok := conds.PaymentType.Val.(types.PaymentType)
		if !ok {
			return nil, wlog.Errorf("invalid paymenttype")
		}
		switch conds.PaymentType.Op {
		case cruder.EQ:
			q.Where(entorderstatebase.PaymentType(paymenttype.String()))
		case cruder.NEQ:
			q.Where(entorderstatebase.PaymentTypeNEQ(paymenttype.String()))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	if conds.PaymentTypes != nil {
		paymenttypes, ok := conds.PaymentTypes.Val.([]types.PaymentType)
		if !ok {
			return nil, wlog.Errorf("invalid paymenttypes")
		}
		_types := []string{}
		for _, _type := range paymenttypes {
			_types = append(_types, _type.String())
		}
		switch conds.PaymentTypes.Op {
		case cruder.IN:
			q.Where(entorderstatebase.PaymentTypeIn(_types...))
		default:
			return nil, wlog.Errorf("invalid order field")
		}
	}
	return q, nil
}
