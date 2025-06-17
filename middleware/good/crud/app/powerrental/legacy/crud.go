package applegacypowerrental

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entapplegacypowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/applegacypowerrental"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID             *uuid.UUID
	AppGoodID         *uuid.UUID
	TechniqueFeeRatio *decimal.Decimal
	DeletedAt         *uint32
}

func CreateSet(c *ent.AppLegacyPowerRentalCreate, req *Req) *ent.AppLegacyPowerRentalCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.TechniqueFeeRatio != nil {
		c.SetTechniqueFeeRatio(*req.TechniqueFeeRatio)
	}
	return c
}

func UpdateSet(u *ent.AppLegacyPowerRentalUpdateOne, req *Req) *ent.AppLegacyPowerRentalUpdateOne {
	if req.TechniqueFeeRatio != nil {
		u.SetTechniqueFeeRatio(*req.TechniqueFeeRatio)
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
func SetQueryConds(q *ent.AppLegacyPowerRentalQuery, conds *Conds) (*ent.AppLegacyPowerRentalQuery, error) {
	q.Where(entapplegacypowerrental.DeletedAt(0))
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
			q.Where(entapplegacypowerrental.ID(id))
		default:
			return nil, wlog.Errorf("invalid applegacypowerrental field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entapplegacypowerrental.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid applegacypowerrental field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entapplegacypowerrental.EntID(id))
		default:
			return nil, wlog.Errorf("invalid applegacypowerrental field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entapplegacypowerrental.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid applegacypowerrental field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entapplegacypowerrental.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid applegacypowerrental field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entapplegacypowerrental.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid applegacypowerrental field")
		}
	}
	return q, nil
}
