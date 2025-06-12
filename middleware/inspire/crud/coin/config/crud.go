package config

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoinconfig "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coinconfig"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID         *uint32
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	CoinTypeID *uuid.UUID
	MaxValue   *decimal.Decimal
	Allocated  *decimal.Decimal
	DeletedAt  *uint32
}

func CreateSet(c *ent.CoinConfigCreate, req *Req) *ent.CoinConfigCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.MaxValue != nil {
		c.SetMaxValue(*req.MaxValue)
	}
	if req.Allocated != nil {
		c.SetAllocated(*req.Allocated)
	}
	return c
}

func UpdateSet(u *ent.CoinConfigUpdateOne, req *Req) *ent.CoinConfigUpdateOne {
	if req.MaxValue != nil {
		u.SetMaxValue(*req.MaxValue)
	}
	if req.Allocated != nil {
		u.SetAllocated(*req.Allocated)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppID      *cruder.Cond
	CoinTypeID *cruder.Cond
	ID         *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.CoinConfigQuery, conds *Conds) (*ent.CoinConfigQuery, error) {
	q.Where(entcoinconfig.DeletedAt(0))
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
			q.Where(entcoinconfig.EntID(id))
		default:
			return nil, wlog.Errorf("invalid coinconfig field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entcoinconfig.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid coinconfig field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcoinconfig.AppID(id))
		default:
			return nil, wlog.Errorf("invalid coinconfig field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcoinconfig.CoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid coinconfig field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entcoinconfig.ID(id))
		case cruder.NEQ:
			q.Where(entcoinconfig.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid coinconfig field")
		}
	}
	return q, nil
}
