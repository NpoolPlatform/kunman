package oneshot

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entoneshot "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/subscriptiononeshot"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID     *uuid.UUID
	GoodID    *uuid.UUID
	GoodType  *types.GoodType
	Name      *string
	Quota     *uint32
	USDPrice  *decimal.Decimal
	DeletedAt *uint32
}

func CreateSet(c *ent.SubscriptionOneShotCreate, req *Req) *ent.SubscriptionOneShotCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.GoodType != nil {
		c.SetGoodType(req.GoodType.String())
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Quota != nil {
		c.SetQuota(*req.Quota)
	}
	if req.USDPrice != nil {
		c.SetUsdPrice(*req.USDPrice)
	}
	return c
}

func UpdateSet(u *ent.SubscriptionOneShotUpdateOne, req *Req) *ent.SubscriptionOneShotUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Quota != nil {
		u.SetQuota(*req.Quota)
	}
	if req.USDPrice != nil {
		u.SetUsdPrice(*req.USDPrice)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	IDs     *cruder.Cond
	EntID   *cruder.Cond
	EntIDs  *cruder.Cond
	GoodID  *cruder.Cond
	GoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.SubscriptionOneShotQuery, conds *Conds) (*ent.SubscriptionOneShotQuery, error) {
	q.Where(entoneshot.DeletedAt(0))
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
			q.Where(entoneshot.ID(id))
		default:
			return nil, wlog.Errorf("invalid oneshot field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entoneshot.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid oneshot field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entoneshot.EntID(id))
		default:
			return nil, wlog.Errorf("invalid oneshot field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entoneshot.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid oneshot field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entoneshot.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid oneshot field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entoneshot.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid oneshot field")
		}
	}
	return q, nil
}
