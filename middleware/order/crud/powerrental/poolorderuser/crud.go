package poolorderuser

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entpoolorderuser "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/poolorderuser"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID           *uuid.UUID
	OrderID         *uuid.UUID
	PoolOrderUserID *uuid.UUID
	DeletedAt       *uint32
}

func CreateSet(c *ent.PoolOrderUserCreate, req *Req) *ent.PoolOrderUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.PoolOrderUserID != nil {
		c.SetPoolOrderUserID(*req.PoolOrderUserID)
	}
	return c
}

func UpdateSet(u *ent.PoolOrderUserUpdateOne, req *Req) *ent.PoolOrderUserUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID              *cruder.Cond
	IDs             *cruder.Cond
	EntID           *cruder.Cond
	EntIDs          *cruder.Cond
	OrderID         *cruder.Cond
	PoolOrderUserID *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.PoolOrderUserQuery, conds *Conds) (*ent.PoolOrderUserQuery, error) {
	q.Where(entpoolorderuser.DeletedAt(0))
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
			q.Where(entpoolorderuser.ID(id))
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
				q.Where(entpoolorderuser.IDIn(ids...))
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
			q.Where(entpoolorderuser.EntID(id))
		case cruder.NEQ:
			q.Where(entpoolorderuser.EntIDNEQ(id))
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
				q.Where(entpoolorderuser.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid powerrental field")
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
			q.Where(entpoolorderuser.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	if conds.PoolOrderUserID != nil {
		id, ok := conds.PoolOrderUserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid orderuserid")
		}
		switch conds.PoolOrderUserID.Op {
		case cruder.EQ:
			q.Where(entpoolorderuser.PoolOrderUserID(id))
		default:
			return nil, wlog.Errorf("invalid powerrental field")
		}
	}
	return q, nil
}
