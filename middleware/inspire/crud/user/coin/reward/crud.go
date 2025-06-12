package reward

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entusercoinreward "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/usercoinreward"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppID       *uuid.UUID
	UserID      *uuid.UUID
	CoinTypeID  *uuid.UUID
	CoinRewards *decimal.Decimal
	DeletedAt   *uint32
}

func CreateSet(c *ent.UserCoinRewardCreate, req *Req) *ent.UserCoinRewardCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.CoinRewards != nil {
		c.SetCoinRewards(*req.CoinRewards)
	}
	return c
}

func UpdateSet(u *ent.UserCoinRewardUpdateOne, req *Req) *ent.UserCoinRewardUpdateOne {
	if req.CoinRewards != nil {
		u.SetCoinRewards(*req.CoinRewards)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	UserIDs    *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	ID         *cruder.Cond
	CoinTypeID *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.UserCoinRewardQuery, conds *Conds) (*ent.UserCoinRewardQuery, error) {
	q.Where(entusercoinreward.DeletedAt(0))
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
			q.Where(entusercoinreward.EntID(id))
		default:
			return nil, wlog.Errorf("invalid usercoinreward field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entusercoinreward.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid usercoinreward field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entusercoinreward.UserIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid usercoinreward field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entusercoinreward.AppID(id))
		default:
			return nil, wlog.Errorf("invalid usercoinreward field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entusercoinreward.UserID(id))
		default:
			return nil, wlog.Errorf("invalid usercoinreward field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entusercoinreward.ID(id))
		case cruder.NEQ:
			q.Where(entusercoinreward.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid usercoinreward field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entusercoinreward.CoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid usercoinreward field")
		}
	}
	return q, nil
}
