package goodcoinreward

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodcoinreward "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodcoinreward"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID                 *uuid.UUID
	GoodID                *uuid.UUID
	CoinTypeID            *uuid.UUID
	RewardTID             *uuid.UUID
	NextRewardStartAmount *decimal.Decimal
	LastRewardAmount      *decimal.Decimal
	LastUnitRewardAmount  *decimal.Decimal
	TotalRewardAmount     *decimal.Decimal
	DeletedAt             *uint32
}

func CreateSet(c *ent.GoodCoinRewardCreate, req *Req) *ent.GoodCoinRewardCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	return c
}

func UpdateSet(u *ent.GoodCoinRewardUpdateOne, req *Req) *ent.GoodCoinRewardUpdateOne {
	if req.RewardTID != nil {
		u.SetRewardTid(*req.RewardTID)
	}
	if req.NextRewardStartAmount != nil {
		u.SetNextRewardStartAmount(*req.NextRewardStartAmount)
	}
	if req.LastRewardAmount != nil {
		u.SetLastRewardAmount(*req.LastRewardAmount)
	}
	if req.LastUnitRewardAmount != nil {
		u.SetLastUnitRewardAmount(*req.LastUnitRewardAmount)
	}
	if req.TotalRewardAmount != nil {
		u.SetTotalRewardAmount(*req.TotalRewardAmount)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	GoodID      *cruder.Cond
	GoodIDs     *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
}

func SetQueryConds(q *ent.GoodCoinRewardQuery, conds *Conds) (*ent.GoodCoinRewardQuery, error) { //nolint:gocyclo
	q.Where(entgoodcoinreward.DeletedAt(0))
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
			q.Where(entgoodcoinreward.EntID(id))
		default:
			return nil, wlog.Errorf("invalid goodcoinreward field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodcoinreward.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid goodcoinreward field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entgoodcoinreward.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid goodcoinreward field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entgoodcoinreward.CoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid goodcoinreward field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entgoodcoinreward.CoinTypeIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid goodcoinreward field")
		}
	}
	return q, nil
}
