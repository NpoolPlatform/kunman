package appsubscription

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappsubscription "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsubscription"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID          *uuid.UUID
	AppGoodID      *uuid.UUID
	USDPrice       *decimal.Decimal
	ProductID      *string
	PlanID         *string
	TrialUnits     *uint32
	TrialUSDPrice  *decimal.Decimal
	PriceFiatID    *uuid.UUID
	FiatPrice      *decimal.Decimal
	TrialFiatPrice *decimal.Decimal
	DeletedAt      *uint32
}

func CreateSet(c *ent.AppSubscriptionCreate, req *Req) *ent.AppSubscriptionCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.USDPrice != nil {
		c.SetUsdPrice(*req.USDPrice)
	}
	if req.ProductID != nil {
		c.SetProductID(*req.ProductID)
	}
	if req.PlanID != nil {
		c.SetPlanID(*req.PlanID)
	}
	if req.TrialUnits != nil {
		c.SetTrialUnits(*req.TrialUnits)
	}
	if req.TrialUSDPrice != nil {
		c.SetTrialUsdPrice(*req.TrialUSDPrice)
	}
	if req.PriceFiatID != nil {
		c.SetPriceFiatID(*req.PriceFiatID)
	}
	if req.FiatPrice != nil {
		c.SetFiatPrice(*req.FiatPrice)
	}
	if req.TrialFiatPrice != nil {
		c.SetTrialFiatPrice(*req.TrialFiatPrice)
	}
	return c
}

func UpdateSet(u *ent.AppSubscriptionUpdateOne, req *Req) *ent.AppSubscriptionUpdateOne {
	if req.USDPrice != nil {
		u.SetUsdPrice(*req.USDPrice)
	}
	if req.PlanID != nil {
		u.SetPlanID(*req.PlanID)
	}
	if req.TrialUnits != nil {
		u.SetTrialUnits(*req.TrialUnits)
	}
	if req.TrialUSDPrice != nil {
		u.SetTrialUsdPrice(*req.TrialUSDPrice)
	}
	if req.PriceFiatID != nil {
		u.SetPriceFiatID(*req.PriceFiatID)
	}
	if req.FiatPrice != nil {
		u.SetFiatPrice(*req.FiatPrice)
	}
	if req.TrialFiatPrice != nil {
		u.SetTrialFiatPrice(*req.TrialFiatPrice)
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
func SetQueryConds(q *ent.AppSubscriptionQuery, conds *Conds) (*ent.AppSubscriptionQuery, error) {
	q.Where(entappsubscription.DeletedAt(0))
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
			q.Where(entappsubscription.ID(id))
		default:
			return nil, wlog.Errorf("invalid appsubscription field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappsubscription.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appsubscription field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappsubscription.EntID(id))
		default:
			return nil, wlog.Errorf("invalid appsubscription field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappsubscription.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appsubscription field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappsubscription.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid appsubscription field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entappsubscription.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appsubscription field")
		}
	}
	return q, nil
}
