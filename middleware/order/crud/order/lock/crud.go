package orderlock

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderlock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderlock"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	OrderID   *uuid.UUID
	UserID    *uuid.UUID
	LockType  *basetypes.OrderLockType
	DeletedAt *uint32
}

func CreateSet(c *ent.OrderLockCreate, req *Req) *ent.OrderLockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.LockType != nil {
		c.SetLockType(req.LockType.String())
	}
	return c
}

func UpdateSet(u *ent.OrderLockUpdateOne, req *Req) *ent.OrderLockUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	IDs      *cruder.Cond
	EntID    *cruder.Cond
	EntIDs   *cruder.Cond
	OrderID  *cruder.Cond
	OrderIDs *cruder.Cond
	UserID   *cruder.Cond
	UserIDs  *cruder.Cond
	LockType *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.OrderLockQuery, conds *Conds) (*ent.OrderLockQuery, error) {
	q.Where(entorderlock.DeletedAt(0))
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
			q.Where(entorderlock.ID(id))
		default:
			return nil, wlog.Errorf("invalid orderlock field")
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
				q.Where(entorderlock.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid orderlock field")
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
			q.Where(entorderlock.EntID(id))
		case cruder.NEQ:
			q.Where(entorderlock.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid orderlock field")
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
				q.Where(entorderlock.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid orderlock field")
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
			q.Where(entorderlock.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid orderlock field")
		}
	}
	if conds.OrderIDs != nil {
		ids, ok := conds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderids")
		}
		if len(ids) > 0 {
			switch conds.OrderIDs.Op {
			case cruder.IN:
				q.Where(entorderlock.OrderIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid orderlock field")
			}
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entorderlock.UserID(id))
		default:
			return nil, wlog.Errorf("invalid orderlock field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userids")
		}
		if len(ids) > 0 {
			switch conds.UserIDs.Op {
			case cruder.IN:
				q.Where(entorderlock.UserIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid orderlock field")
			}
		}
	}
	if conds.LockType != nil {
		_type, ok := conds.LockType.Val.(basetypes.OrderLockType)
		if !ok {
			return nil, wlog.Errorf("invalid locktype")
		}
		switch conds.LockType.Op {
		case cruder.EQ:
			q.Where(entorderlock.LockType(_type.String()))
		default:
			return nil, wlog.Errorf("invalid orderlock field")
		}
	}
	return q, nil
}
