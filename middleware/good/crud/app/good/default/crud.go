package appdefaultgood

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappdefaultgood "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appdefaultgood"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uint32
	EntID      *uuid.UUID
	AppGoodID  *uuid.UUID
	CoinTypeID *uuid.UUID
	DeletedAt  *uint32
}

func CreateSet(c *ent.AppDefaultGoodCreate, req *Req) *ent.AppDefaultGoodCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	return c
}

func UpdateSet(u *ent.AppDefaultGoodUpdateOne, req *Req) *ent.AppDefaultGoodUpdateOne {
	if req.AppGoodID != nil {
		u.SetAppGoodID(*req.AppGoodID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	EntID       *cruder.Cond
	AppGoodID   *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.AppDefaultGoodQuery, conds *Conds) (*ent.AppDefaultGoodQuery, error) {
	q.Where(entappdefaultgood.DeletedAt(0))
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
			q.Where(entappdefaultgood.ID(id))
		default:
			return nil, wlog.Errorf("invalid appdefaultgood field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappdefaultgood.EntID(id))
		default:
			return nil, wlog.Errorf("invalid appdefaultgood field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappdefaultgood.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid appdefaultgood field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entappdefaultgood.CoinTypeID(id))
		default:
			return nil, wlog.Errorf("invalid appdefaultgood field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entappdefaultgood.CoinTypeIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appdefaultgood field")
		}
	}
	return q, nil
}
