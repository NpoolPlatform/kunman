package allocated

import (
	"time"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	inspiretypes "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcouponallocated "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/couponallocated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID            *uint32
	EntID         *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	CouponID      *uuid.UUID
	Used          *bool
	UsedByOrderID *uuid.UUID
	Denomination  *decimal.Decimal
	StartAt       *uint32
	CouponScope   *inspiretypes.CouponScope
	Cashable      *bool
	Extra         *string
	DeletedAt     *uint32
}

func CreateSet(c *ent.CouponAllocatedCreate, req *Req) *ent.CouponAllocatedCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.CouponID != nil {
		c.SetCouponID(*req.CouponID)
	}
	if req.Denomination != nil {
		c.SetDenomination(*req.Denomination)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.CouponScope != nil {
		c.SetCouponScope(req.CouponScope.String())
	}
	if req.Cashable != nil {
		c.SetCashable(*req.Cashable)
	}
	if req.Extra != nil {
		c.SetExtra(*req.Extra)
	}
	c.SetUsed(false)
	c.SetUsedAt(0)
	return c
}

func UpdateSet(u *ent.CouponAllocatedUpdateOne, req *Req) *ent.CouponAllocatedUpdateOne {
	if req.Used != nil && *req.Used && req.UsedByOrderID != nil {
		u.SetUsed(*req.Used)
		u.SetUsedAt(uint32(time.Now().Unix()))
		u.SetUsedByOrderID(*req.UsedByOrderID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID          *cruder.Cond
	EntIDs         *cruder.Cond
	AppID          *cruder.Cond
	UserID         *cruder.Cond
	CouponType     *cruder.Cond
	CouponID       *cruder.Cond
	CouponIDs      *cruder.Cond
	Used           *cruder.Cond
	UsedByOrderID  *cruder.Cond
	UsedByOrderIDs *cruder.Cond
	Extra          *cruder.Cond
}

func SetQueryConds(q *ent.CouponAllocatedQuery, conds *Conds) (*ent.CouponAllocatedQuery, error) { //nolint
	q.Where(entcouponallocated.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.EntID(id))
		default:
			return nil, wlog.Errorf("invalid allocated field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entcouponallocated.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid allocated ids")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.AppID(id))
		default:
			return nil, wlog.Errorf("invalid allocated field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.UserID(id))
		default:
			return nil, wlog.Errorf("invalid allocated field")
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid couponid")
		}
		switch conds.CouponID.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.CouponID(id))
		default:
			return nil, wlog.Errorf("invalid couponid field")
		}
	}
	if conds.CouponIDs != nil {
		ids, ok := conds.CouponIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid couponids")
		}
		switch conds.CouponIDs.Op {
		case cruder.IN:
			q.Where(entcouponallocated.CouponIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid couponids field")
		}
	}
	if conds.Used != nil {
		used, ok := conds.Used.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid used")
		}
		switch conds.Used.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.Used(used))
		default:
			return nil, wlog.Errorf("invalid allocated field")
		}
	}
	if conds.UsedByOrderID != nil {
		id, ok := conds.UsedByOrderID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid usedbyorderid")
		}
		switch conds.UsedByOrderID.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.UsedByOrderID(id))
		default:
			return nil, wlog.Errorf("invalid allocated field")
		}
	}
	if conds.UsedByOrderIDs != nil {
		ids, ok := conds.UsedByOrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid usedbyorderid")
		}
		switch conds.UsedByOrderIDs.Op {
		case cruder.IN:
			q.Where(entcouponallocated.UsedByOrderIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid allocated field")
		}
	}
	if conds.Extra != nil {
		id, ok := conds.Extra.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid extra")
		}
		switch conds.Extra.Op {
		case cruder.EQ:
			q.Where(entcouponallocated.ExtraContains(id))
		default:
			return nil, wlog.Errorf("invalid allocated field")
		}
	}
	return q, nil
}
