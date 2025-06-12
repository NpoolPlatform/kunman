package goodachievement

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entgoodachievement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/goodachievement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID              *uuid.UUID
	AppID              *uuid.UUID
	UserID             *uuid.UUID
	GoodID             *uuid.UUID
	AppGoodID          *uuid.UUID
	TotalAmountUSD     *decimal.Decimal
	SelfAmountUSD      *decimal.Decimal
	TotalUnits         *decimal.Decimal
	SelfUnits          *decimal.Decimal
	TotalCommissionUSD *decimal.Decimal
	SelfCommissionUSD  *decimal.Decimal
	DeletedAt          *uint32
}

func CreateSet(c *ent.GoodAchievementCreate, req *Req) *ent.GoodAchievementCreate {
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
	if req.TotalAmountUSD != nil {
		c.SetTotalAmountUsd(*req.TotalAmountUSD)
	}
	if req.SelfAmountUSD != nil {
		c.SetSelfAmountUsd(*req.SelfAmountUSD)
	}
	if req.TotalUnits != nil {
		c.SetTotalUnits(*req.TotalUnits)
	}
	if req.SelfUnits != nil {
		c.SetSelfUnits(*req.SelfUnits)
	}
	if req.TotalCommissionUSD != nil {
		c.SetTotalCommissionUsd(*req.TotalCommissionUSD)
	}
	if req.SelfCommissionUSD != nil {
		c.SetSelfCommissionUsd(*req.SelfCommissionUSD)
	}

	return c
}

func UpdateSet(u *ent.GoodAchievementUpdateOne, req *Req) *ent.GoodAchievementUpdateOne {
	if req.TotalAmountUSD != nil {
		u = u.SetTotalAmountUsd(*req.TotalAmountUSD)
	}
	if req.SelfAmountUSD != nil {
		u = u.SetSelfAmountUsd(*req.SelfAmountUSD)
	}
	if req.TotalUnits != nil {
		u = u.SetTotalUnits(*req.TotalUnits)
	}
	if req.SelfUnits != nil {
		u = u.SetSelfUnits(*req.SelfUnits)
	}
	if req.TotalCommissionUSD != nil {
		u = u.SetTotalCommissionUsd(*req.TotalCommissionUSD)
	}
	if req.SelfCommissionUSD != nil {
		u = u.SetSelfCommissionUsd(*req.SelfCommissionUSD)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID     *cruder.Cond
	AppID     *cruder.Cond
	UserID    *cruder.Cond
	GoodID    *cruder.Cond
	AppGoodID *cruder.Cond
	UserIDs   *cruder.Cond
}

func SetQueryConds(q *ent.GoodAchievementQuery, conds *Conds) (*ent.GoodAchievementQuery, error) { //nolint
	q.Where(entgoodachievement.DeletedAt(0))
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
			q.Where(entgoodachievement.EntID(id))
		default:
			return nil, wlog.Errorf("invalid goodachievement field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entgoodachievement.AppID(id))
		default:
			return nil, wlog.Errorf("invalid goodachievement field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entgoodachievement.UserID(id))
		default:
			return nil, wlog.Errorf("invalid goodachievement field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entgoodachievement.UserIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid goodachievement field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodachievement.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid goodachievement field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entgoodachievement.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid goodachievement field")
		}
	}
	return q, nil
}
