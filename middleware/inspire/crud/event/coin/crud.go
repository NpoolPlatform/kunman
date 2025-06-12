package coin

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	enteventcoin "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/eventcoin"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID           *uint32
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	EventID      *uuid.UUID
	CoinConfigID *uuid.UUID
	CoinValue    *decimal.Decimal
	CoinPerUSD   *decimal.Decimal
	DeletedAt    *uint32
}

func CreateSet(c *ent.EventCoinCreate, req *Req) *ent.EventCoinCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.EventID != nil {
		c.SetEventID(*req.EventID)
	}
	if req.CoinConfigID != nil {
		c.SetCoinConfigID(*req.CoinConfigID)
	}
	if req.CoinValue != nil {
		c.SetCoinValue(*req.CoinValue)
	}
	if req.CoinPerUSD != nil {
		c.SetCoinPerUsd(*req.CoinPerUSD)
	}
	return c
}

func UpdateSet(u *ent.EventCoinUpdateOne, req *Req) *ent.EventCoinUpdateOne {
	if req.CoinValue != nil {
		u.SetCoinValue(*req.CoinValue)
	}
	if req.CoinPerUSD != nil {
		u.SetCoinPerUsd(*req.CoinPerUSD)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID        *cruder.Cond
	EntIDs       *cruder.Cond
	AppID        *cruder.Cond
	EventID      *cruder.Cond
	CoinConfigID *cruder.Cond
	ID           *cruder.Cond
	EventIDs     *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.EventCoinQuery, conds *Conds) (*ent.EventCoinQuery, error) {
	q.Where(enteventcoin.DeletedAt(0))
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
			q.Where(enteventcoin.EntID(id))
		default:
			return nil, wlog.Errorf("invalid eventcoin field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enteventcoin.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid eventcoin field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(enteventcoin.AppID(id))
		default:
			return nil, wlog.Errorf("invalid eventcoin field")
		}
	}
	if conds.CoinConfigID != nil {
		id, ok := conds.CoinConfigID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid coinconfigid")
		}
		switch conds.CoinConfigID.Op {
		case cruder.EQ:
			q.Where(enteventcoin.CoinConfigID(id))
		default:
			return nil, wlog.Errorf("invalid eventcoin field")
		}
	}
	if conds.EventID != nil {
		id, ok := conds.EventID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid eventid")
		}
		switch conds.EventID.Op {
		case cruder.EQ:
			q.Where(enteventcoin.EventID(id))
		default:
			return nil, wlog.Errorf("invalid eventcoin field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(enteventcoin.ID(id))
		case cruder.NEQ:
			q.Where(enteventcoin.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid eventcoin field")
		}
	}
	if conds.EventIDs != nil {
		ids, ok := conds.EventIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid eventids")
		}
		switch conds.EventIDs.Op {
		case cruder.IN:
			q.Where(enteventcoin.EventIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid eventcoin field")
		}
	}
	return q, nil
}
