//nolint:dupl
package config

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entappcommissionconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/appcommissionconfig"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID              *uint32
	EntID           *uuid.UUID
	AppID           *uuid.UUID
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

func CreateSet(c *ent.AppCommissionConfigCreate, req *Req) *ent.AppCommissionConfigCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.ThresholdAmount != nil {
		c.SetThresholdAmount(*req.ThresholdAmount)
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
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	if req.Level != nil {
		c.SetLevel(*req.Level)
	}
	return c
}

func UpdateSet(u *ent.AppCommissionConfigUpdateOne, req *Req) *ent.AppCommissionConfigUpdateOne {
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
	SettleType      *cruder.Cond
	StartAt         *cruder.Cond
	EndAt           *cruder.Cond
	EntIDs          *cruder.Cond
	ThresholdAmount *cruder.Cond
	Invites         *cruder.Cond
	Disabled        *cruder.Cond
	Level           *cruder.Cond
	ID              *cruder.Cond
}

func SetQueryConds(q *ent.AppCommissionConfigQuery, conds *Conds) (*ent.AppCommissionConfigQuery, error) { //nolint
	q.Where(entappcommissionconfig.DeletedAt(0))
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
			q.Where(entappcommissionconfig.EntID(id))
		case cruder.NEQ:
			q.Where(entappcommissionconfig.EntIDNEQ(id))
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
			q.Where(entappcommissionconfig.AppID(id))
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
			q.Where(entappcommissionconfig.SettleType(settleType.String()))
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
			q.Where(entappcommissionconfig.EndAtLT(at))
		case cruder.LTE:
			q.Where(entappcommissionconfig.EndAtLTE(at))
		case cruder.GT:
			q.Where(entappcommissionconfig.EndAtGT(at))
		case cruder.GTE:
			q.Where(entappcommissionconfig.EndAtGTE(at))
		case cruder.EQ:
			q.Where(entappcommissionconfig.EndAt(at))
		case cruder.NEQ:
			q.Where(entappcommissionconfig.EndAtNEQ(at))
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
			q.Where(entappcommissionconfig.StartAtLT(at))
		case cruder.LTE:
			q.Where(entappcommissionconfig.StartAtLTE(at))
		case cruder.GT:
			q.Where(entappcommissionconfig.StartAtGT(at))
		case cruder.GTE:
			q.Where(entappcommissionconfig.StartAtGTE(at))
		case cruder.EQ:
			q.Where(entappcommissionconfig.StartAt(at))
		case cruder.NEQ:
			q.Where(entappcommissionconfig.StartAtNEQ(at))
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
			q.Where(entappcommissionconfig.EntIDIn(ids...))
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
			q.Where(entappcommissionconfig.ThresholdAmount(id))
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
			q.Where(entappcommissionconfig.Invites(id))
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
			q.Where(entappcommissionconfig.Disabled(value))
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
			q.Where(entappcommissionconfig.Level(id))
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
			q.Where(entappcommissionconfig.ID(id))
		case cruder.NEQ:
			q.Where(entappcommissionconfig.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	return q, nil
}
