package coupon

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoupon "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coupon"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                  *uint32
	EntID               *uuid.UUID
	CouponType          *types.CouponType
	AppID               *uuid.UUID
	Denomination        *decimal.Decimal
	Circulation         *decimal.Decimal
	IssuedBy            *uuid.UUID
	StartAt             *uint32
	EndAt               *uint32
	DurationDays        *uint32
	Message             *string
	Name                *string
	CouponConstraint    *types.CouponConstraint
	CouponScope         *types.CouponScope
	Threshold           *decimal.Decimal
	Allocated           *decimal.Decimal
	Random              *bool
	CashableProbability *decimal.Decimal
	DeletedAt           *uint32
}

func CreateSet(c *ent.CouponCreate, req *Req) *ent.CouponCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.CouponType != nil {
		c.SetCouponType(req.CouponType.String())
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.Denomination != nil {
		c.SetDenomination(*req.Denomination)
	}
	if req.Circulation != nil {
		c.SetCirculation(*req.Circulation)
	}
	if req.IssuedBy != nil {
		c.SetIssuedBy(*req.IssuedBy)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	if req.DurationDays != nil {
		c.SetDurationDays(*req.DurationDays)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Threshold != nil {
		c.SetThreshold(*req.Threshold)
	}
	if req.CouponConstraint != nil {
		c.SetCouponConstraint(req.CouponConstraint.String())
	}
	if req.CouponScope != nil {
		c.SetCouponScope(req.CouponScope.String())
	}
	if req.Allocated != nil {
		c.SetAllocated(*req.Allocated)
	}
	if req.Random != nil {
		c.SetRandom(*req.Random)
	}
	if req.CashableProbability != nil {
		c.SetCashableProbability(*req.CashableProbability)
	}
	return c
}

func UpdateSet(u *ent.CouponUpdateOne, req *Req) *ent.CouponUpdateOne {
	if req.Denomination != nil {
		u.SetDenomination(*req.Denomination)
	}
	if req.Circulation != nil {
		u.SetCirculation(*req.Circulation)
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.DurationDays != nil {
		u.SetDurationDays(*req.DurationDays)
	}
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Random != nil {
		u.SetRandom(*req.Random)
	}
	if req.Threshold != nil {
		u.SetThreshold(*req.Threshold)
	}
	if req.Allocated != nil {
		u.SetAllocated(*req.Allocated)
	}
	if req.CouponScope != nil {
		u.SetCouponScope(req.CouponScope.String())
	}
	if req.CouponConstraint != nil {
		u.SetCouponConstraint(req.CouponConstraint.String())
	}
	if req.CashableProbability != nil {
		u.SetCashableProbability(*req.CashableProbability)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	CouponType *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	GoodID     *cruder.Cond
	AppGoodID  *cruder.Cond
}

func SetQueryConds(q *ent.CouponQuery, conds *Conds) (*ent.CouponQuery, error) {
	q.Where(entcoupon.DeletedAt(0))
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
			q.Where(entcoupon.EntID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entcoupon.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid entids field")
		}
	}
	if conds.CouponType != nil {
		_type, ok := conds.CouponType.Val.(types.CouponType)
		if !ok {
			return nil, wlog.Errorf("invalid coupontype")
		}
		switch conds.CouponType.Op {
		case cruder.EQ:
			q.Where(entcoupon.CouponType(_type.String()))
		default:
			return nil, wlog.Errorf("invalid coupontype field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcoupon.AppID(id))
		default:
			return nil, wlog.Errorf("invalid appid field")
		}
	}
	return q, nil
}
