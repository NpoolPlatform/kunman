package appsimulatepowerrental

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappsimulatepowerrental "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsimulatepowerrental"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type Req struct {
	EntID                *uuid.UUID
	AppGoodID            *uuid.UUID
	OrderUnits           *decimal.Decimal
	OrderDurationSeconds *uint32
	DeletedAt            *uint32
}

func CreateSet(c *ent.AppSimulatePowerRentalCreate, req *Req) *ent.AppSimulatePowerRentalCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.OrderUnits != nil {
		c.SetOrderUnits(*req.OrderUnits)
	}
	if req.OrderDurationSeconds != nil {
		c.SetOrderDurationSeconds(*req.OrderDurationSeconds)
	}
	return c
}

func UpdateSet(u *ent.AppSimulatePowerRentalUpdateOne, req *Req) *ent.AppSimulatePowerRentalUpdateOne {
	if req.OrderUnits != nil {
		u.SetOrderUnits(*req.OrderUnits)
	}
	if req.OrderDurationSeconds != nil {
		u.SetOrderDurationSeconds(*req.OrderDurationSeconds)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	EntID     *cruder.Cond
	AppGoodID *cruder.Cond
}

func SetQueryConds(q *ent.AppSimulatePowerRentalQuery, conds *Conds) (*ent.AppSimulatePowerRentalQuery, error) {
	q.Where(entappsimulatepowerrental.DeletedAt(0))
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
			q.Where(entappsimulatepowerrental.ID(id))
		default:
			return nil, wlog.Errorf("invalid appsimulatepowerrental field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappsimulatepowerrental.EntID(id))
		default:
			return nil, wlog.Errorf("invalid appsimulatepowerrental field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappsimulatepowerrental.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid appsimulatepowerrental field")
		}
	}
	return q, nil
}
