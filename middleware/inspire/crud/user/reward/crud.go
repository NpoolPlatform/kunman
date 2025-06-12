package reward

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entuserreward "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/userreward"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                   *uint32
	EntID                *uuid.UUID
	AppID                *uuid.UUID
	UserID               *uuid.UUID
	ActionCredits        *decimal.Decimal
	CouponAmount         *decimal.Decimal
	CouponCashableAmount *decimal.Decimal
	DeletedAt            *uint32
}

func CreateSet(c *ent.UserRewardCreate, req *Req) *ent.UserRewardCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.ActionCredits != nil {
		c.SetActionCredits(*req.ActionCredits)
	}
	if req.CouponAmount != nil {
		c.SetCouponAmount(*req.CouponAmount)
	}
	if req.CouponCashableAmount != nil {
		c.SetCouponCashableAmount(*req.CouponCashableAmount)
	}
	return c
}

func UpdateSet(u *ent.UserRewardUpdateOne, req *Req) *ent.UserRewardUpdateOne {
	if req.ActionCredits != nil {
		u.SetActionCredits(*req.ActionCredits)
	}
	if req.CouponAmount != nil {
		u.SetCouponAmount(*req.CouponAmount)
	}
	if req.CouponCashableAmount != nil {
		u.SetCouponCashableAmount(*req.CouponCashableAmount)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID  *cruder.Cond
	EntIDs *cruder.Cond
	AppID  *cruder.Cond
	UserID *cruder.Cond
	ID     *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.UserRewardQuery, conds *Conds) (*ent.UserRewardQuery, error) {
	q.Where(entuserreward.DeletedAt(0))
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
			q.Where(entuserreward.EntID(id))
		default:
			return nil, wlog.Errorf("invalid userreward field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entuserreward.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid userreward field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entuserreward.AppID(id))
		default:
			return nil, wlog.Errorf("invalid userreward field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entuserreward.UserID(id))
		default:
			return nil, wlog.Errorf("invalid userreward field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entuserreward.ID(id))
		case cruder.NEQ:
			q.Where(entuserreward.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid userreward field")
		}
	}
	return q, nil
}
