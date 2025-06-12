package allocated

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated"
	entcoinallocated "github.com/NpoolPlatform/kunman/middleware/inspire/db/ent/generated/coinallocated"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID           *uint32
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	CoinConfigID *uuid.UUID
	CoinTypeID   *uuid.UUID
	UserID       *uuid.UUID
	Value        *decimal.Decimal
	Extra        *string
	DeletedAt    *uint32
}

func CreateSet(c *ent.CoinAllocatedCreate, req *Req) *ent.CoinAllocatedCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CoinConfigID != nil {
		c.SetCoinConfigID(*req.CoinConfigID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.Value != nil {
		c.SetValue(*req.Value)
	}
	if req.Extra != nil {
		c.SetExtra(*req.Extra)
	}
	return c
}

func UpdateSet(u *ent.CoinAllocatedUpdateOne, req *Req) *ent.CoinAllocatedUpdateOne {
	if req.Value != nil {
		u.SetValue(*req.Value)
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
	UserID       *cruder.Cond
	CoinConfigID *cruder.Cond
	CoinTypeID   *cruder.Cond
	ID           *cruder.Cond
	Extra        *cruder.Cond
}

//nolint:funlen
func SetQueryConds(q *ent.CoinAllocatedQuery, conds *Conds) (*ent.CoinAllocatedQuery, error) {
	q.Where(entcoinallocated.DeletedAt(0))
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
			q.Where(entcoinallocated.EntID(id))
		default:
			return nil, wlog.Errorf("invalid coinallocated field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entcoinallocated.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid coinallocated field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcoinallocated.AppID(id))
		default:
			return nil, wlog.Errorf("invalid coinallocated field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entcoinallocated.UserID(id))
		default:
			return nil, wlog.Errorf("invalid coinallocated field")
		}
	}
	if conds.CoinConfigID != nil {
		id, ok := conds.CoinConfigID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid coinconfigid")
		}
		switch conds.CoinConfigID.Op {
		case cruder.EQ:
			q.Where(entcoinallocated.CoinConfigID(id))
		default:
			return nil, wlog.Errorf("invalid coinallocated field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcoinallocated.CoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid coinallocated field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entcoinallocated.ID(id))
		case cruder.NEQ:
			q.Where(entcoinallocated.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid coinallocated field")
		}
	}
	if conds.Extra != nil {
		id, ok := conds.Extra.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid extra")
		}
		switch conds.Extra.Op {
		case cruder.EQ:
			q.Where(entcoinallocated.ExtraContains(id))
		default:
			return nil, wlog.Errorf("invalid coinallocated field")
		}
	}
	return q, nil
}
