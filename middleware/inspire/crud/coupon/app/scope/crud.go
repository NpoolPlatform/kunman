package scope

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entappgoodscope "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appgoodscope"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	AppGoodID   *uuid.UUID
	CouponID    *uuid.UUID
	CouponScope *types.CouponScope
	DeletedAt   *uint32
	GoodID      *uuid.UUID
}

func CreateSet(c *ent.AppGoodScopeCreate, req *Req) *ent.AppGoodScopeCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.CouponID != nil {
		c.SetCouponID(*req.CouponID)
	}
	if req.CouponScope != nil {
		c.SetCouponScope(req.CouponScope.String())
	}
	return c
}

func UpdateSet(u *ent.AppGoodScopeUpdateOne, req *Req) *ent.AppGoodScopeUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	AppGoodID   *cruder.Cond
	CouponID    *cruder.Cond
	CouponIDs   *cruder.Cond
	CouponScope *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.AppGoodScopeQuery, conds *Conds) (*ent.AppGoodScopeQuery, error) {
	q.Where(entappgoodscope.DeletedAt(0))
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
			q.Where(entappgoodscope.EntID(id))
		default:
			return nil, wlog.Errorf("invalid entid field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappgoodscope.AppID(id))
		default:
			return nil, wlog.Errorf("invalid appid field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappgoodscope.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid appgoodid field")
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid couponid")
		}
		switch conds.CouponID.Op {
		case cruder.EQ:
			q.Where(entappgoodscope.CouponID(id))
		default:
			return nil, wlog.Errorf("invalid couponid field")
		}
	}
	if conds.CouponScope != nil {
		scope, ok := conds.CouponScope.Val.(types.CouponScope)
		if !ok {
			return nil, wlog.Errorf("invalid couponscope")
		}
		switch conds.CouponScope.Op {
		case cruder.EQ:
			q.Where(entappgoodscope.CouponScope(scope.String()))
		case cruder.NEQ:
			q.Where(entappgoodscope.CouponScopeNEQ(scope.String()))
		default:
			return nil, wlog.Errorf("invalid couponscope field")
		}
	}
	if conds.CouponIDs != nil {
		ids, ok := conds.CouponIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid couponids")
		}
		switch conds.CouponIDs.Op {
		case cruder.IN:
			q.Where(entappgoodscope.CouponIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid couponids field")
		}
	}
	return q, nil
}
