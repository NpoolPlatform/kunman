package scope

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcouponscope "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/couponscope"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	GoodID      *uuid.UUID
	CouponID    *uuid.UUID
	CouponScope *types.CouponScope
	DeletedAt   *uint32
}

func CreateSet(c *ent.CouponScopeCreate, req *Req) *ent.CouponScopeCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.CouponID != nil {
		c.SetCouponID(*req.CouponID)
	}
	if req.CouponScope != nil {
		c.SetCouponScope(req.CouponScope.String())
	}
	return c
}

func UpdateSet(u *ent.CouponScopeUpdateOne, req *Req) *ent.CouponScopeUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	GoodID      *cruder.Cond
	CouponID    *cruder.Cond
	CouponIDs   *cruder.Cond
	CouponScope *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.CouponScopeQuery, conds *Conds) (*ent.CouponScopeQuery, error) {
	q.Where(entcouponscope.DeletedAt(0))
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
			q.Where(entcouponscope.EntID(id))
		default:
			return nil, wlog.Errorf("invalid entid field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entcouponscope.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid goodid field")
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid couponid")
		}
		switch conds.CouponID.Op {
		case cruder.EQ:
			q.Where(entcouponscope.CouponID(id))
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
			q.Where(entcouponscope.CouponScope(scope.String()))
		case cruder.NEQ:
			q.Where(entcouponscope.CouponScopeNEQ(scope.String()))
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
			q.Where(entcouponscope.CouponIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid couponids field")
		}
	}
	return q, nil
}
