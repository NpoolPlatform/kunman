package coin

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcashcontrol "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/cashcontrol"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	CouponID    *uuid.UUID
	ControlType *types.ControlType
	Value       *decimal.Decimal
	DeletedAt   *uint32
}

func CreateSet(c *ent.CashControlCreate, req *Req) *ent.CashControlCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CouponID != nil {
		c.SetCouponID(*req.CouponID)
	}
	if req.ControlType != nil {
		c.SetControlType(req.ControlType.String())
	}
	c.SetValue(decimal.RequireFromString("0"))
	if req.Value != nil {
		c.SetValue(*req.Value)
	}
	return c
}

func UpdateSet(u *ent.CashControlUpdateOne, req *Req) *ent.CashControlUpdateOne {
	if req.Value != nil {
		u.SetValue(*req.Value)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	CouponID    *cruder.Cond
	CouponIDs   *cruder.Cond
	ControlType *cruder.Cond
}

func SetQueryConds(q *ent.CashControlQuery, conds *Conds) (*ent.CashControlQuery, error) { //nolint
	q.Where(entcashcontrol.DeletedAt(0))
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
			q.Where(entcashcontrol.EntID(id))
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
			q.Where(entcashcontrol.AppID(id))
		default:
			return nil, wlog.Errorf("invalid appid field")
		}
	}
	if conds.CouponID != nil {
		id, ok := conds.CouponID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CouponID.Op {
		case cruder.EQ:
			q.Where(entcashcontrol.CouponID(id))
		default:
			return nil, wlog.Errorf("invalid cointypeid field")
		}
	}
	if conds.CouponIDs != nil {
		ids, ok := conds.CouponIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid couponids")
		}
		switch conds.CouponIDs.Op {
		case cruder.IN:
			q.Where(entcashcontrol.CouponIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid couponids field")
		}
	}
	if conds.ControlType != nil {
		controlType, ok := conds.ControlType.Val.(types.ControlType)
		if !ok {
			return nil, wlog.Errorf("invalid control type")
		}
		switch conds.ControlType.Op {
		case cruder.EQ:
			q.Where(entcashcontrol.ControlType(controlType.String()))
		case cruder.NEQ:
			q.Where(entcashcontrol.ControlTypeNEQ(controlType.String()))
		default:
			return nil, wlog.Errorf("invalid control type field")
		}
	}
	return q, nil
}
