package paymentbase

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entpaymentbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID         *uuid.UUID
	OrderID       *uuid.UUID
	ObseleteState *types.PaymentObseleteState
	DeletedAt     *uint32
}

func CreateSet(c *ent.PaymentBaseCreate, req *Req) *ent.PaymentBaseCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	return c
}

func UpdateSet(u *ent.PaymentBaseUpdateOne, req *Req) *ent.PaymentBaseUpdateOne {
	if req.ObseleteState != nil {
		u.SetObseleteState(req.ObseleteState.String())
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
	ObseleteState *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.PaymentBaseQuery, conds *Conds) (*ent.PaymentBaseQuery, error) {
	q.Where(entpaymentbase.DeletedAt(0))
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
			q.Where(entpaymentbase.ID(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
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
				q.Where(entpaymentbase.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid payment field")
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
			q.Where(entpaymentbase.EntID(id))
		case cruder.NEQ:
			q.Where(entpaymentbase.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
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
				q.Where(entpaymentbase.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid payment field")
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
			q.Where(entpaymentbase.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	if conds.ObseleteState != nil {
		_state, ok := conds.ObseleteState.Val.(types.PaymentObseleteState)
		if !ok {
			return nil, wlog.Errorf("invalid obseletestate")
		}
		switch conds.ObseleteState.Op {
		case cruder.EQ:
			q.Where(entpaymentbase.ObseleteState(_state.String()))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	return q, nil
}
