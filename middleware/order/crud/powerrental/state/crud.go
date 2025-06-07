package powerrentalstate

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entpowerrentalstate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/powerrentalstate"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID             *uuid.UUID
	OrderID           *uuid.UUID
	CancelState       *types.OrderState
	CanceledAt        *uint32
	PaymentID         *uuid.UUID
	PaidAt            *uint32
	UserSetPaid       *bool
	UserSetCanceled   *bool
	AdminSetCanceled  *bool
	PaymentState      *types.PaymentState
	OutOfGasSeconds   *uint32
	CompensateSeconds *uint32
	RenewState        *types.OrderRenewState
	RenewNotifyAt     *uint32
	DeletedAt         *uint32
}

func CreateSet(c *ent.PowerRentalStateCreate, req *Req) *ent.PowerRentalStateCreate {
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
	return c
}

func UpdateSet(u *ent.PowerRentalStateUpdateOne, req *Req) *ent.PowerRentalStateUpdateOne {
	if req.CancelState != nil {
		u.SetCancelState(req.CancelState.String())
	}
	if req.CanceledAt != nil {
		u.SetCanceledAt(*req.CanceledAt)
	}
	if req.PaymentID != nil {
		u.SetPaymentID(*req.PaymentID)
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
	if req.OutOfGasSeconds != nil {
		u.SetOutofgasSeconds(*req.OutOfGasSeconds)
	}
	if req.CompensateSeconds != nil {
		u.SetCompensateSeconds(*req.CompensateSeconds)
	}
	if req.RenewState != nil {
		u.SetRenewState(req.RenewState.String())
	}
	if req.RenewNotifyAt != nil {
		u.SetRenewNotifyAt(*req.RenewNotifyAt)
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
	RenewState       *cruder.Cond
	RenewNotifyAt    *cruder.Cond
	UserSetCanceled  *cruder.Cond
	AdminSetCanceled *cruder.Cond
	PaidAt           *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.PowerRentalStateQuery, conds *Conds) (*ent.PowerRentalStateQuery, error) {
	q.Where(entpowerrentalstate.DeletedAt(0))
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
			q.Where(entpowerrentalstate.ID(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
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
				q.Where(entpowerrentalstate.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid powerrental field")
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
			q.Where(entpowerrentalstate.EntID(id))
		case cruder.NEQ:
			q.Where(entpowerrentalstate.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
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
				q.Where(entpowerrentalstate.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid powerrental field")
			}
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid powerrentalid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entpowerrentalstate.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.PaymentState != nil {
		state, ok := conds.PaymentState.Val.(types.PaymentState)
		if !ok {
			return nil, wlog.Errorf("invalid paymentstate")
		}
		switch conds.PaymentState.Op {
		case cruder.EQ:
			q.Where(entpowerrentalstate.PaymentState(state.String()))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.PaymentStates != nil {
		states, ok := conds.PaymentStates.Val.([]types.PaymentState)
		if !ok {
			return nil, wlog.Errorf("invalid paymentstate")
		}
		switch conds.PaymentState.Op {
		case cruder.IN:
			q.Where(entpowerrentalstate.PaymentStateIn(func() (_states []string) {
				for _, state := range states {
					_states = append(_states, state.String())
				}
				return
			}()...))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.RenewState != nil {
		state, ok := conds.RenewState.Val.(types.OrderRenewState)
		if !ok {
			return nil, wlog.Errorf("invalid renewstate")
		}
		switch conds.RenewState.Op {
		case cruder.EQ:
			q.Where(entpowerrentalstate.RenewState(state.String()))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.RenewNotifyAt != nil {
		at, ok := conds.RenewNotifyAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid renewnotifyat")
		}
		switch conds.RenewNotifyAt.Op {
		case cruder.EQ:
			q.Where(entpowerrentalstate.RenewNotifyAt(at))
		case cruder.LT:
			q.Where(entpowerrentalstate.RenewNotifyAtLT(at))
		case cruder.LTE:
			q.Where(entpowerrentalstate.RenewNotifyAtLTE(at))
		case cruder.GT:
			q.Where(entpowerrentalstate.RenewNotifyAtGT(at))
		case cruder.GTE:
			q.Where(entpowerrentalstate.RenewNotifyAtGTE(at))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	return q, nil
}
