package achievement

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entachievement "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/achievement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID           *uuid.UUID
	AppID           *uuid.UUID
	UserID          *uuid.UUID
	GoodID          *uuid.UUID
	AppGoodID       *uuid.UUID
	CoinTypeID      *uuid.UUID
	TotalAmount     *decimal.Decimal
	SelfAmount      *decimal.Decimal
	TotalUnits      *decimal.Decimal
	SelfUnits       *decimal.Decimal
	TotalCommission *decimal.Decimal
	SelfCommission  *decimal.Decimal
}

func CreateSet(c *ent.AchievementCreate, req *Req) *ent.AchievementCreate {
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
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.TotalAmount != nil {
		c.SetTotalAmount(*req.TotalAmount)
	}
	if req.SelfAmount != nil {
		c.SetSelfAmount(*req.SelfAmount)
	}
	if req.TotalUnits != nil {
		c.SetTotalUnitsV1(*req.TotalUnits)
	}
	if req.SelfUnits != nil {
		c.SetSelfUnitsV1(*req.SelfUnits)
	}
	if req.TotalCommission != nil {
		c.SetTotalCommission(*req.TotalCommission)
	}
	if req.SelfCommission != nil {
		c.SetSelfCommission(*req.SelfCommission)
	}

	return c
}

func UpdateSet(u *ent.AchievementUpdateOne, req *Req) *ent.AchievementUpdateOne {
	if req.TotalAmount != nil {
		u = u.SetTotalAmount(*req.TotalAmount)
	}
	if req.SelfAmount != nil {
		u = u.SetSelfAmount(*req.SelfAmount)
	}
	if req.TotalUnits != nil {
		u = u.SetTotalUnitsV1(*req.TotalUnits)
	}
	if req.SelfUnits != nil {
		u = u.SetSelfUnitsV1(*req.SelfUnits)
	}
	if req.TotalCommission != nil {
		u = u.SetTotalCommission(*req.TotalCommission)
	}
	if req.SelfCommission != nil {
		u = u.SetSelfCommission(*req.SelfCommission)
	}
	return u
}

type Conds struct {
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	GoodID     *cruder.Cond
	AppGoodID  *cruder.Cond
	CoinTypeID *cruder.Cond
	UserIDs    *cruder.Cond
}

func SetQueryConds(q *ent.AchievementQuery, conds *Conds) (*ent.AchievementQuery, error) { //nolint
	q.Where(entachievement.DeletedAt(0))
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
			q.Where(entachievement.EntID(id))
		default:
			return nil, wlog.Errorf("invalid general field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entachievement.AppID(id))
		default:
			return nil, wlog.Errorf("invalid general field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entachievement.UserID(id))
		default:
			return nil, wlog.Errorf("invalid general field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entachievement.UserIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid general field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entachievement.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid general field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entachievement.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid general field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entachievement.CoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid general field")
		}
	}
	return q, nil
}
