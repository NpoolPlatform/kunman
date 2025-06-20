package paymentbalancelock

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entpaymentbalancelock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalancelock"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	PaymentID    *uuid.UUID
	LedgerLockID *uuid.UUID
	DeletedAt    *uint32
}

func CreateSet(c *ent.PaymentBalanceLockCreate, req *Req) *ent.PaymentBalanceLockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.PaymentID != nil {
		c.SetPaymentID(*req.PaymentID)
	}
	if req.LedgerLockID != nil {
		c.SetLedgerLockID(*req.LedgerLockID)
	}
	return c
}

func UpdateSet(u *ent.PaymentBalanceLockUpdateOne, req *Req) *ent.PaymentBalanceLockUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID           *cruder.Cond
	IDs          *cruder.Cond
	EntID        *cruder.Cond
	EntIDs       *cruder.Cond
	PaymentID    *cruder.Cond
	LedgerLockID *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.PaymentBalanceLockQuery, conds *Conds) (*ent.PaymentBalanceLockQuery, error) {
	q.Where(entpaymentbalancelock.DeletedAt(0))
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
			q.Where(entpaymentbalancelock.ID(id))
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
				q.Where(entpaymentbalancelock.IDIn(ids...))
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
			q.Where(entpaymentbalancelock.EntID(id))
		case cruder.NEQ:
			q.Where(entpaymentbalancelock.EntIDNEQ(id))
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
				q.Where(entpaymentbalancelock.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid payment field")
			}
		}
	}
	if conds.PaymentID != nil {
		id, ok := conds.PaymentID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderid")
		}
		switch conds.PaymentID.Op {
		case cruder.EQ:
			q.Where(entpaymentbalancelock.PaymentID(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	if conds.LedgerLockID != nil {
		id, ok := conds.LedgerLockID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid ledgerlockid")
		}
		switch conds.LedgerLockID.Op {
		case cruder.EQ:
			q.Where(entpaymentbalancelock.LedgerLockID(id))
		default:
			return nil, wlog.Errorf("invalid payment field")
		}
	}
	return q, nil
}
