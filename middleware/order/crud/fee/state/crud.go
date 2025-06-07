package feeorderstate

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entfeeorderstate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/feeorderstate"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID            *uuid.UUID
	OrderID          *uuid.UUID
	PaymentID        *uuid.UUID
	PaidAt           *uint32
	UserSetPaid      *bool
	UserSetCanceled  *bool
	AdminSetCanceled *bool
	PaymentState     *types.PaymentState
	CancelState      *types.OrderState
	CanceledAt       *uint32
	DeletedAt        *uint32
}

func CreateSet(c *ent.FeeOrderStateCreate, req *Req) *ent.FeeOrderStateCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.PaymentID != nil {
		c.SetPaymentID(*req.PaymentID)
	}
	if req.UserSetPaid != nil {
		c.SetUserSetPaid(*req.UserSetPaid)
	}
	if req.UserSetCanceled != nil {
		c.SetUserSetCanceled(*req.UserSetCanceled)
	}
	if req.AdminSetCanceled != nil {
		c.SetAdminSetCanceled(*req.AdminSetCanceled)
	}
	if req.PaymentState != nil {
		c.SetPaymentState(req.PaymentState.String())
	}
	return c
}

func UpdateSet(u *ent.FeeOrderStateUpdateOne, req *Req) *ent.FeeOrderStateUpdateOne {
	if req.OrderID != nil {
		u.SetOrderID(*req.OrderID)
	}
	if req.PaidAt != nil {
		u.SetPaidAt(*req.PaidAt)
	}
	if req.UserSetPaid != nil {
		u.SetUserSetPaid(*req.UserSetPaid)
	}
	if req.UserSetCanceled != nil {
		u.SetUserSetCanceled(*req.UserSetCanceled)
	}
	if req.AdminSetCanceled != nil {
		u.SetAdminSetCanceled(*req.AdminSetCanceled)
	}
	if req.PaymentState != nil {
		u.SetPaymentState(req.PaymentState.String())
	}
	if req.CancelState != nil {
		u.SetCancelState(req.CancelState.String())
	}
	if req.CanceledAt != nil {
		u.SetCanceledAt(*req.CanceledAt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID               *cruder.Cond
	IDs              *cruder.Cond
	EntID            *cruder.Cond
	EntIDs           *cruder.Cond
	OrderID          *cruder.Cond
	PaymentState     *cruder.Cond
	PaymentStates    *cruder.Cond
	UserSetCanceled  *cruder.Cond
	AdminSetCanceled *cruder.Cond
	PaidAt           *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.FeeOrderStateQuery, conds *Conds) (*ent.FeeOrderStateQuery, error) {
	q.Where(entfeeorderstate.DeletedAt(0))
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
			q.Where(entfeeorderstate.ID(id))
		default:
			return nil, wlog.Errorf("invalid feeorder field")
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
				q.Where(entfeeorderstate.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid feeorder field")
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
			q.Where(entfeeorderstate.EntID(id))
		case cruder.NEQ:
			q.Where(entfeeorderstate.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid feeorder field")
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
				q.Where(entfeeorderstate.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid feeorder field")
			}
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid feeorderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entfeeorderstate.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid feeorder field")
		}
	}
	if conds.PaymentState != nil {
		state, ok := conds.PaymentState.Val.(types.PaymentState)
		if !ok {
			return nil, wlog.Errorf("invalid paymentstate")
		}
		switch conds.PaymentState.Op {
		case cruder.EQ:
			q.Where(entfeeorderstate.PaymentState(state.String()))
		default:
			return nil, wlog.Errorf("invalid feeorder field")
		}
	}
	if conds.PaymentStates != nil {
		states, ok := conds.PaymentStates.Val.([]types.PaymentState)
		if !ok {
			return nil, wlog.Errorf("invalid paymentstate")
		}
		switch conds.PaymentState.Op {
		case cruder.IN:
			q.Where(entfeeorderstate.PaymentStateIn(func() (_states []string) {
				for _, state := range states {
					_states = append(_states, state.String())
				}
				return
			}()...))
		default:
			return nil, wlog.Errorf("invalid feeorder field")
		}
	}
	if conds.UserSetCanceled != nil {
		b, ok := conds.UserSetCanceled.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid usersetcanceled")
		}
		switch conds.UserSetCanceled.Op {
		case cruder.EQ:
			q.Where(entfeeorderstate.UserSetCanceled(b))
		default:
			return nil, wlog.Errorf("invalid feeorder field")
		}
	}
	if conds.AdminSetCanceled != nil {
		b, ok := conds.AdminSetCanceled.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid adminsetcanceled")
		}
		switch conds.AdminSetCanceled.Op {
		case cruder.EQ:
			q.Where(entfeeorderstate.AdminSetCanceled(b))
		default:
			return nil, wlog.Errorf("invalid feeorder field")
		}
	}
	return q, nil
}
