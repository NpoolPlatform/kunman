package coupon

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	enteventcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/eventcoupon"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	EventID   *uuid.UUID
	CouponID  *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.EventCouponCreate, req *Req) *ent.EventCouponCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.EventID != nil {
		c.SetEventID(*req.EventID)
	}
	if req.CouponID != nil {
		c.SetCouponID(*req.CouponID)
	}
	return c
}

func UpdateSet(u *ent.EventCouponUpdateOne, req *Req) *ent.EventCouponUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID    *cruder.Cond
	EntIDs   *cruder.Cond
	AppID    *cruder.Cond
	EventID  *cruder.Cond
	CouponID *cruder.Cond
	ID       *cruder.Cond
	EventIDs *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.EventCouponQuery, conds *Conds) (*ent.EventCouponQuery, error) {
	q.Where(enteventcoupon.DeletedAt(0))
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
			q.Where(enteventcoupon.EntID(id))
		default:
			return nil, wlog.Errorf("invalid eventcoupon field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enteventcoupon.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid eventcoupon field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(enteventcoupon.AppID(id))
		default:
			return nil, wlog.Errorf("invalid eventcoupon field")
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid couponid")
		}
		switch conds.CouponID.Op {
		case cruder.EQ:
			q.Where(enteventcoupon.CouponID(id))
		default:
			return nil, wlog.Errorf("invalid eventcoupon field")
		}
	}
	if conds.EventID != nil {
		id, ok := conds.EventID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid eventid")
		}
		switch conds.EventID.Op {
		case cruder.EQ:
			q.Where(enteventcoupon.EventID(id))
		default:
			return nil, wlog.Errorf("invalid eventcoupon field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(enteventcoupon.ID(id))
		case cruder.NEQ:
			q.Where(enteventcoupon.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid eventcoupon field")
		}
	}
	if conds.EventIDs != nil {
		ids, ok := conds.EventIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid eventids")
		}
		switch conds.EventIDs.Op {
		case cruder.IN:
			q.Where(enteventcoupon.EventIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid eventcoupon field")
		}
	}
	return q, nil
}
