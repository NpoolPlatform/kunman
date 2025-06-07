package goodrewardhistory

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodrewardhistory "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodrewardhistory"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID         *uuid.UUID
	GoodID        *uuid.UUID
	CoinTypeID    *uuid.UUID
	RewardDate    *uint32
	TID           *uuid.UUID
	Amount        *decimal.Decimal
	UnitAmount    *decimal.Decimal
	UnitNetAmount *decimal.Decimal
}

func CreateSet(c *ent.GoodRewardHistoryCreate, req *Req) *ent.GoodRewardHistoryCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.RewardDate != nil {
		c.SetRewardDate(*req.RewardDate)
	}
	if req.TID != nil {
		c.SetTid(*req.TID)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.UnitAmount != nil {
		c.SetUnitAmount(*req.UnitAmount)
	}
	if req.UnitNetAmount != nil {
		c.SetUnitNetAmount(*req.UnitNetAmount)
	}
	return c
}

func UpdateSet(u *ent.GoodRewardHistoryUpdateOne, req *Req) *ent.GoodRewardHistoryUpdateOne {
	return u
}

type Conds struct {
	ID          *cruder.Cond
	EntID       *cruder.Cond
	GoodID      *cruder.Cond
	GoodIDs     *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	RewardDate  *cruder.Cond
	StartAt     *cruder.Cond
	EndAt       *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.GoodRewardHistoryQuery, conds *Conds) (*ent.GoodRewardHistoryQuery, error) {
	q.Where(entgoodrewardhistory.DeletedAt(0))
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
			q.Where(entgoodrewardhistory.ID(id))
		default:
			return nil, wlog.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodrewardhistory.EntID(id))
		default:
			return nil, wlog.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodrewardhistory.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entgoodrewardhistory.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entgoodrewardhistory.CoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entgoodrewardhistory.CoinTypeIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.RewardDate != nil {
		date, ok := conds.RewardDate.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid rewarddate")
		}
		switch conds.RewardDate.Op {
		case cruder.EQ:
			q.Where(entgoodrewardhistory.RewardDate(date))
		case cruder.LTE:
			q.Where(entgoodrewardhistory.RewardDateLTE(date))
		case cruder.GTE:
			q.Where(entgoodrewardhistory.RewardDateGTE(date))
		default:
			return nil, wlog.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.StartAt != nil {
		date, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid startat")
		}
		q.Where(entgoodrewardhistory.RewardDateGTE(date))
	}
	if conds.EndAt != nil {
		date, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid endat")
		}
		q.Where(entgoodrewardhistory.RewardDateLTE(date))
	}
	return q, nil
}
