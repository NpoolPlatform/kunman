//nolint:dupl
package commission

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcommission "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/commission"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID               *uint32
	EntID            *uuid.UUID
	AppID            *uuid.UUID
	UserID           *uuid.UUID
	GoodID           *uuid.UUID
	AppGoodID        *uuid.UUID
	AmountOrPercent  *decimal.Decimal
	EndAt            *uint32
	StartAt          *uint32
	SettleType       *types.SettleType
	SettleMode       *types.SettleMode
	SettleAmountType *types.SettleAmountType
	SettleInterval   *types.SettleInterval
	Threshold        *decimal.Decimal
	DeletedAt        *uint32
}

func CreateSet(c *ent.CommissionCreate, req *Req) *ent.CommissionCreate {
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
	if req.SettleType != nil {
		c.SetSettleType(req.SettleType.String())
	}
	if req.SettleAmountType != nil {
		c.SetSettleAmountType(req.SettleAmountType.String())
	}
	if req.SettleMode != nil {
		c.SetSettleMode(req.SettleMode.String())
	}
	if req.SettleInterval != nil {
		c.SetSettleInterval(req.SettleInterval.String())
	}
	if req.Threshold != nil {
		c.SetThreshold(*req.Threshold)
	}
	return c
}

func UpdateSet(u *ent.CommissionUpdateOne, req *Req) *ent.CommissionUpdateOne {
	if req.AmountOrPercent != nil {
		u = u.SetAmountOrPercent(*req.AmountOrPercent)
	}
	if req.StartAt != nil {
		u = u.SetStartAt(*req.StartAt)
	}
	if req.Threshold != nil {
		u.SetThreshold(*req.Threshold)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	GoodID     *cruder.Cond
	AppGoodID  *cruder.Cond
	SettleType *cruder.Cond
	EndAt      *cruder.Cond
	UserIDs    *cruder.Cond
	GoodIDs    *cruder.Cond
	StartAt    *cruder.Cond
}

func SetQueryConds(q *ent.CommissionQuery, conds *Conds) (*ent.CommissionQuery, error) { //nolint
	q.Where(entcommission.DeletedAt(0))
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
			q.Where(entcommission.ID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcommission.EntID(id))
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
			q.Where(entcommission.AppID(id))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entcommission.UserID(id))
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
			q.Where(entcommission.GoodID(id))
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
			q.Where(entcommission.AppGoodID(id))
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
			q.Where(entcommission.SettleType(settleType.String()))
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
			q.Where(entcommission.EndAtLT(at))
		case cruder.LTE:
			q.Where(entcommission.EndAtLTE(at))
		case cruder.GT:
			q.Where(entcommission.EndAtGT(at))
		case cruder.GTE:
			q.Where(entcommission.EndAtGTE(at))
		case cruder.EQ:
			q.Where(entcommission.EndAt(at))
		case cruder.NEQ:
			q.Where(entcommission.EndAtNEQ(at))
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
			q.Where(entcommission.StartAtLT(at))
		case cruder.LTE:
			q.Where(entcommission.StartAtLTE(at))
		case cruder.GT:
			q.Where(entcommission.StartAtGT(at))
		case cruder.GTE:
			q.Where(entcommission.StartAtGTE(at))
		case cruder.EQ:
			q.Where(entcommission.StartAt(at))
		case cruder.NEQ:
			q.Where(entcommission.StartAtNEQ(at))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entcommission.UserIDIn(ids...))
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
			q.Where(entcommission.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid commission field")
		}
	}
	return q, nil
}
