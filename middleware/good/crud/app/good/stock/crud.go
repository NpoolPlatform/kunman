package appstock

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appstock"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID        *uuid.UUID
	AppGoodID    *uuid.UUID
	Reserved     *decimal.Decimal
	SpotQuantity *decimal.Decimal
	Locked       *decimal.Decimal
	InService    *decimal.Decimal
	WaitStart    *decimal.Decimal
	Sold         *decimal.Decimal
	DeletedAt    *uint32
}

func CreateSet(c *ent.AppStockCreate, req *Req) *ent.AppStockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.Reserved != nil {
		c.SetReserved(*req.Reserved)
	}
	if req.SpotQuantity != nil {
		c.SetSpotQuantity(*req.SpotQuantity)
	}
	if req.Locked != nil {
		c.SetLocked(*req.Locked)
	}
	if req.InService != nil {
		c.SetInService(*req.InService)
	}
	if req.WaitStart != nil {
		c.SetWaitStart(*req.WaitStart)
	}
	if req.Sold != nil {
		c.SetSold(*req.Sold)
	}
	return c
}

func UpdateSet(u *ent.AppStockUpdateOne, req *Req) *ent.AppStockUpdateOne {
	if req.Reserved != nil {
		u.SetReserved(*req.Reserved)
	}
	if req.SpotQuantity != nil {
		u.SetSpotQuantity(*req.SpotQuantity)
	}
	if req.Locked != nil {
		u.SetLocked(*req.Locked)
	}
	if req.InService != nil {
		u.SetInService(*req.InService)
	}
	if req.WaitStart != nil {
		u.SetWaitStart(*req.WaitStart)
	}
	if req.Sold != nil {
		u.SetSold(*req.Sold)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.AppStockQuery, conds *Conds) (*ent.AppStockQuery, error) {
	q.Where(entappstock.DeletedAt(0))
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
			q.Where(entappstock.ID(id))
		default:
			return nil, wlog.Errorf("invalid appstock field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappstock.EntID(id))
		default:
			return nil, wlog.Errorf("invalid appstock field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappstock.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appstock field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappstock.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid appstock field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entappstock.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appstock field")
		}
	}
	return q, nil
}
