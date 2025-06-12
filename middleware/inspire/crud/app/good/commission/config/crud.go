//nolint:dupl
package config

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entappgoodcommissionconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appgoodcommissionconfig"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID              *uint32
	EntID           *uuid.UUID
	AppID           *uuid.UUID
	GoodID          *uuid.UUID
	AppGoodID       *uuid.UUID
	ThresholdAmount *decimal.Decimal
	AmountOrPercent *decimal.Decimal
	EndAt           *uint32
	StartAt         *uint32
	Invites         *uint32
	SettleType      *types.SettleType
	Disabled        *bool
	Level           *uint32
	DeletedAt       *uint32
}

func CreateSet(c *ent.AppGoodCommissionConfigCreate, req *Req) *ent.AppGoodCommissionConfigCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.AmountOrPercent != nil {
		c.SetAmountOrPercent(*req.AmountOrPercent)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	c.SetEndAt(0)
	if req.Invites != nil {
		c.SetInvites(*req.Invites)
	}
	if req.SettleType != nil {
		c.SetSettleType(req.SettleType.String())
	}
	if req.ThresholdAmount != nil {
		c.SetThresholdAmount(*req.ThresholdAmount)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	if req.Level != nil {
		c.SetLevel(*req.Level)
	}
	return c
}

func UpdateSet(u *ent.AppGoodCommissionConfigUpdateOne, req *Req) *ent.AppGoodCommissionConfigUpdateOne {
	if req.AmountOrPercent != nil {
		u.SetAmountOrPercent(*req.AmountOrPercent)
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.ThresholdAmount != nil {
		u.SetThresholdAmount(*req.ThresholdAmount)
	}
	if req.Invites != nil {
		u.SetInvites(*req.Invites)
	}
	if req.Disabled != nil {
		u.SetDisabled(*req.Disabled)
	}
	if req.Level != nil {
		u.SetLevel(*req.Level)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID           *cruder.Cond
	AppID           *cruder.Cond
	UserID          *cruder.Cond
	GoodID          *cruder.Cond
	AppGoodID       *cruder.Cond
	SettleType      *cruder.Cond
	StartAt         *cruder.Cond
	EndAt           *cruder.Cond
	EntIDs          *cruder.Cond
	GoodIDs         *cruder.Cond
	AppGoodIDs      *cruder.Cond
	ThresholdAmount *cruder.Cond
	Invites         *cruder.Cond
	Disabled        *cruder.Cond
	Level           *cruder.Cond
	ID              *cruder.Cond
}

func SetQueryConds(q *ent.AppGoodCommissionConfigQuery, conds *Conds) (*ent.AppGoodCommissionConfigQuery, error) { //nolint
	q.Where(entappgoodcommissionconfig.DeletedAt(0))
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
			q.Where(entappgoodcommissionconfig.EntID(id))
		case cruder.NEQ:
			q.Where(entappgoodcommissionconfig.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.AppID(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.SettleType != nil {
		settleType, ok := conds.SettleType.Val.(types.SettleType)
		if !ok {
			return nil, wlog.Errorf("invalid settletype")
		}
		switch conds.SettleType.Op {
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.SettleType(settleType.String()))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.EndAt != nil {
		at, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid endat")
		}
		switch conds.EndAt.Op {
		case cruder.LT:
			q.Where(entappgoodcommissionconfig.EndAtLT(at))
		case cruder.LTE:
			q.Where(entappgoodcommissionconfig.EndAtLTE(at))
		case cruder.GT:
			q.Where(entappgoodcommissionconfig.EndAtGT(at))
		case cruder.GTE:
			q.Where(entappgoodcommissionconfig.EndAtGTE(at))
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.EndAt(at))
		case cruder.NEQ:
			q.Where(entappgoodcommissionconfig.EndAtNEQ(at))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.StartAt != nil {
		at, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid startat")
		}
		switch conds.StartAt.Op {
		case cruder.LT:
			q.Where(entappgoodcommissionconfig.StartAtLT(at))
		case cruder.LTE:
			q.Where(entappgoodcommissionconfig.StartAtLTE(at))
		case cruder.GT:
			q.Where(entappgoodcommissionconfig.StartAtGT(at))
		case cruder.GTE:
			q.Where(entappgoodcommissionconfig.StartAtGTE(at))
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.StartAt(at))
		case cruder.NEQ:
			q.Where(entappgoodcommissionconfig.StartAtNEQ(at))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappgoodcommissionconfig.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entappgoodcommissionconfig.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entappgoodcommissionconfig.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.ThresholdAmount != nil {
		id, ok := conds.ThresholdAmount.Val.(decimal.Decimal)
		if !ok {
			return nil, wlog.Errorf("invalid thresholdamount")
		}
		switch conds.ThresholdAmount.Op {
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.ThresholdAmount(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.Invites != nil {
		id, ok := conds.Invites.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid invites")
		}
		switch conds.Invites.Op {
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.Invites(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.Disabled != nil {
		value, ok := conds.Disabled.Val.(bool)
		if !ok {
			return nil, wlog.Errorf("invalid disabled")
		}
		switch conds.Disabled.Op {
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.Disabled(value))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.Level != nil {
		id, ok := conds.Level.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid level")
		}
		switch conds.Level.Op {
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.Level(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entappgoodcommissionconfig.ID(id))
		case cruder.NEQ:
			q.Where(entappgoodcommissionconfig.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	return q, nil
}
