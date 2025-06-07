package outofgas

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entoutofgas "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/outofgas"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	OrderID   *uuid.UUID
	StartAt   *uint32
	EndAt     *uint32
	DeletedAt *uint32
}

func CreateSet(c *ent.OutOfGasCreate, req *Req) *ent.OutOfGasCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	return c
}

func UpdateSet(u *ent.OutOfGasUpdateOne, req *Req) *ent.OutOfGasUpdateOne {
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID    *cruder.Cond
	EntIDs   *cruder.Cond
	ID       *cruder.Cond
	IDs      *cruder.Cond
	OrderID  *cruder.Cond
	OrderIDs *cruder.Cond
	StartAt  *cruder.Cond
	EndAt    *cruder.Cond
	StartEnd *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.OutOfGasQuery, conds *Conds) (*ent.OutOfGasQuery, error) {
	q.Where(entoutofgas.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entoutofgas.EntID(id))
		case cruder.NEQ:
			q.Where(entoutofgas.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid outofgas field")
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
				q.Where(entoutofgas.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid outofgas field")
			}
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entoutofgas.ID(id))
		case cruder.NEQ:
			q.Where(entoutofgas.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid outofgas field")
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
				q.Where(entoutofgas.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid outofgas field")
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
			q.Where(entoutofgas.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid outofgas field")
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
				q.Where(entoutofgas.OrderIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid outofgas field")
			}
		}
	}
	if conds.StartAt != nil {
		start, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid startat")
		}
		switch conds.StartAt.Op {
		case cruder.LT:
			q.Where(entoutofgas.StartAtLT(start))
		case cruder.LTE:
			q.Where(entoutofgas.StartAtLTE(start))
		case cruder.GT:
			q.Where(entoutofgas.StartAtGT(start))
		case cruder.GTE:
			q.Where(entoutofgas.StartAtGTE(start))
		case cruder.EQ:
			q.Where(entoutofgas.StartAt(start))
		default:
			return nil, wlog.Errorf("invalid outofgas field")
		}
	}
	if conds.EndAt != nil {
		end, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid endat")
		}
		switch conds.EndAt.Op {
		case cruder.LT:
			q.Where(entoutofgas.EndAtLT(end))
		case cruder.LTE:
			q.Where(entoutofgas.EndAtLTE(end))
		case cruder.GT:
			q.Where(entoutofgas.EndAtGT(end))
		case cruder.GTE:
			q.Where(entoutofgas.EndAtGTE(end))
		case cruder.EQ:
			q.Where(entoutofgas.EndAtEQ(end))
		default:
			return nil, wlog.Errorf("invalid outofgas field")
		}
	}
	if conds.StartEnd != nil {
		ats, ok := conds.StartEnd.Val.([]uint32)
		if !ok || len(ats) != 2 {
			return nil, wlog.Errorf("invalid startend")
		}
		switch conds.StartEnd.Op {
		case cruder.OVERLAP:
			q.Where(
				entoutofgas.Or(
					entoutofgas.And(
						entoutofgas.StartAtLTE(ats[0]),
						entoutofgas.EndAtGTE(ats[0]),
					),
					entoutofgas.And(
						entoutofgas.StartAtLTE(ats[1]),
						entoutofgas.EndAtGTE(ats[1]),
					),
					entoutofgas.And(
						entoutofgas.StartAtGTE(ats[0]),
						entoutofgas.EndAtLTE(ats[1]),
					),
				),
			)
		default:
			return nil, wlog.Errorf("invalid outofgas field")
		}
	}
	return q, nil
}
