package apponeshot

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entapponeshot "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsubscriptiononeshot"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID     *uuid.UUID
	AppGoodID *uuid.UUID
	USDPrice  *decimal.Decimal
	DeletedAt *uint32
}

func CreateSet(c *ent.AppSubscriptionOneShotCreate, req *Req) *ent.AppSubscriptionOneShotCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.USDPrice != nil {
		c.SetUsdPrice(*req.USDPrice)
	}
	return c
}

func UpdateSet(u *ent.AppSubscriptionOneShotUpdateOne, req *Req) *ent.AppSubscriptionOneShotUpdateOne {
	if req.USDPrice != nil {
		u.SetUsdPrice(*req.USDPrice)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AppSubscriptionOneShotQuery, conds *Conds) (*ent.AppSubscriptionOneShotQuery, error) {
	q.Where(entapponeshot.DeletedAt(0))
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
			q.Where(entapponeshot.ID(id))
		default:
			return nil, wlog.Errorf("invalid apponeshot field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entapponeshot.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid apponeshot field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entapponeshot.EntID(id))
		default:
			return nil, wlog.Errorf("invalid apponeshot field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entapponeshot.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid apponeshot field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entapponeshot.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid apponeshot field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entapponeshot.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid apponeshot field")
		}
	}
	return q, nil
}
