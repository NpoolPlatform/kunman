package subscription

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entsubscription "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/subscription"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID               *uuid.UUID
	GoodID              *uuid.UUID
	DurationDisplayType *types.GoodDurationType
	DurationUnits       *uint32
	DurationQuota       *uint32
	DailyBonusQuota     *uint32
	USDPrice            *decimal.Decimal
	DeletedAt           *uint32
}

func CreateSet(c *ent.SubscriptionCreate, req *Req) *ent.SubscriptionCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.DurationDisplayType != nil {
		c.SetDurationDisplayType(req.DurationDisplayType.String())
	}
	if req.DurationUnits != nil {
		c.SetDurationUnits(*req.DurationUnits)
	}
	if req.DurationQuota != nil {
		c.SetDurationQuota(*req.DurationQuota)
	}
	if req.DailyBonusQuota != nil {
		c.SetDailyBonusQuota(*req.DailyBonusQuota)
	}
	if req.USDPrice != nil {
		c.SetUsdPrice(*req.USDPrice)
	}
	return c
}

func UpdateSet(u *ent.SubscriptionUpdateOne, req *Req) *ent.SubscriptionUpdateOne {
	if req.DurationDisplayType != nil {
		u.SetDurationDisplayType(req.DurationDisplayType.String())
	}
	if req.DurationUnits != nil {
		u.SetDurationUnits(*req.DurationUnits)
	}
	if req.DurationQuota != nil {
		u.SetDurationQuota(*req.DurationQuota)
	}
	if req.DailyBonusQuota != nil {
		u.SetDailyBonusQuota(*req.DailyBonusQuota)
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
func SetQueryConds(q *ent.SubscriptionQuery, conds *Conds) (*ent.SubscriptionQuery, error) {
	q.Where(entsubscription.DeletedAt(0))
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
			q.Where(entsubscription.ID(id))
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
			q.Where(entsubscription.IDIn(ids...))
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
			q.Where(entsubscription.EntID(id))
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
			q.Where(entsubscription.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entsubscription.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entsubscription.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	return q, nil
}
