package exchange

import (
	"github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entcreditexchange "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/exchange"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/agi/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID                *uint32
	EntID             *uuid.UUID
	AppID             *uuid.UUID
	UsageType         *types.UsageType
	Credit            *uint32
	ExchangeThreshold *uint32
	Path              *string
	DeletedAt         *uint32
}

func CreateSet(c *ent.ExchangeCreate, req *Req) *ent.ExchangeCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UsageType != nil {
		c.SetUsageType(req.UsageType.String())
	}
	if req.Credit != nil {
		c.SetCredit(*req.Credit)
	}
	if req.ExchangeThreshold != nil {
		c.SetExchangeThreshold(*req.ExchangeThreshold)
	}
	if req.Path != nil {
		c.SetPath(*req.Path)
	}
	return c
}

func UpdateSet(u *ent.ExchangeUpdateOne, req *Req) *ent.ExchangeUpdateOne {
	if req.UsageType != nil {
		u.SetUsageType(req.UsageType.String())
	}
	if req.Credit != nil {
		u.SetCredit(*req.Credit)
	}
	if req.ExchangeThreshold != nil {
		u.SetExchangeThreshold(*req.ExchangeThreshold)
	}
	if req.Path != nil {
		u.SetPath(*req.Path)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	IDs       *cruder.Cond
	EntID     *cruder.Cond
	EntIDs    *cruder.Cond
	AppID     *cruder.Cond
	UsageType *cruder.Cond
	Path      *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.ExchangeQuery, conds *Conds) (*ent.ExchangeQuery, error) {
	q.Where(entcreditexchange.DeletedAt(0))
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
			q.Where(entcreditexchange.ID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entcreditexchange.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcreditexchange.EntID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entcreditexchange.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcreditexchange.AppID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.UsageType != nil {
		e, ok := conds.UsageType.Val.(types.UsageType)
		if !ok {
			return nil, wlog.Errorf("invalid usagetype")
		}
		switch conds.UsageType.Op {
		case cruder.EQ:
			q.Where(entcreditexchange.UsageType(e.String()))
		case cruder.NEQ:
			q.Where(entcreditexchange.UsageTypeNEQ(e.String()))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.Path != nil {
		s, ok := conds.Path.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid path")
		}
		switch conds.Path.Op {
		case cruder.EQ:
			q.Where(entcreditexchange.Path(s))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}

	return q, nil
}
