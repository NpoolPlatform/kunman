package ordercoupon

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entordercoupon "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/ordercoupon"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	OrderID   *uuid.UUID
	CouponID  *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.OrderCouponCreate, req *Req) *ent.OrderCouponCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.CouponID != nil {
		c.SetCouponID(*req.CouponID)
	}
	return c
}

func UpdateSet(u *ent.OrderCouponUpdateOne, req *Req) *ent.OrderCouponUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	IDs       *cruder.Cond
	EntID     *cruder.Cond
	EntIDs    *cruder.Cond
	OrderID   *cruder.Cond
	OrderIDs  *cruder.Cond
	CouponID  *cruder.Cond
	CouponIDs *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.OrderCouponQuery, conds *Conds) (*ent.OrderCouponQuery, error) {
	q.Where(entordercoupon.DeletedAt(0))
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
			q.Where(entordercoupon.ID(id))
		default:
			return nil, wlog.Errorf("invalid ordercoupon field")
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
				q.Where(entordercoupon.IDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid ordercoupon field")
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
			q.Where(entordercoupon.EntID(id))
		case cruder.NEQ:
			q.Where(entordercoupon.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid ordercoupon field")
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
				q.Where(entordercoupon.EntIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid ordercoupon field")
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
			q.Where(entordercoupon.OrderID(id))
		default:
			return nil, wlog.Errorf("invalid ordercoupon field")
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
				q.Where(entordercoupon.OrderIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid ordercoupon field")
			}
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid couponid")
		}
		switch conds.CouponID.Op {
		case cruder.EQ:
			q.Where(entordercoupon.CouponID(id))
		default:
			return nil, wlog.Errorf("invalid ordercoupon field")
		}
	}
	if conds.CouponIDs != nil {
		ids, ok := conds.CouponIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid couponids")
		}
		if len(ids) > 0 {
			switch conds.CouponIDs.Op {
			case cruder.IN:
				q.Where(entordercoupon.CouponIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid ordercoupon field")
			}
		}
	}
	return q, nil
}
