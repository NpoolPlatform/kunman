package mining

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entmininggoodstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/mininggoodstock"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID             *uint32
	EntID          *uuid.UUID
	GoodStockID    *uuid.UUID
	PoolRootUserID *uuid.UUID
	PoolGoodUserID *uuid.UUID
	Total          *decimal.Decimal
	SpotQuantity   *decimal.Decimal
	Locked         *decimal.Decimal
	InService      *decimal.Decimal
	WaitStart      *decimal.Decimal
	Sold           *decimal.Decimal
	AppReserved    *decimal.Decimal
	State          *types.MiningGoodStockState
	DeletedAt      *uint32
}

func CreateSet(c *ent.MiningGoodStockCreate, req *Req) *ent.MiningGoodStockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodStockID != nil {
		c.SetGoodStockID(*req.GoodStockID)
	}
	if req.PoolRootUserID != nil {
		c.SetPoolRootUserID(*req.PoolRootUserID)
	}
	if req.PoolGoodUserID != nil {
		c.SetPoolGoodUserID(*req.PoolGoodUserID)
	}
	if req.Total != nil {
		c.SetTotal(*req.Total)
	}
	if req.SpotQuantity != nil {
		c.SetSpotQuantity(*req.SpotQuantity)
	}
	c.SetLocked(decimal.NewFromInt(0))
	c.SetInService(decimal.NewFromInt(0))
	c.SetWaitStart(decimal.NewFromInt(0))
	c.SetSold(decimal.NewFromInt(0))
	c.SetAppReserved(decimal.NewFromInt(0))
	if req.State != nil {
		c.SetState(req.State.String())
	}
	return c
}

func UpdateSet(u *ent.MiningGoodStockUpdateOne, req *Req) *ent.MiningGoodStockUpdateOne {
	if req.Total != nil {
		u.SetTotal(*req.Total)
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
	if req.AppReserved != nil {
		u.SetAppReserved(*req.AppReserved)
	}
	if req.PoolGoodUserID != nil {
		u.SetPoolGoodUserID(*req.PoolGoodUserID)
	}
	if req.State != nil {
		u.SetState(req.State.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID        *cruder.Cond
	GoodStockID  *cruder.Cond
	GoodStockIDs *cruder.Cond
}

func SetQueryConds(q *ent.MiningGoodStockQuery, conds *Conds) (*ent.MiningGoodStockQuery, error) {
	q.Where(entmininggoodstock.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entmininggoodstock.EntID(id))
		default:
			return nil, wlog.Errorf("invalid mininggoodstock field")
		}
	}
	if conds.GoodStockID != nil {
		id, ok := conds.GoodStockID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodstockid")
		}
		switch conds.GoodStockID.Op {
		case cruder.EQ:
			q.Where(entmininggoodstock.GoodStockID(id))
		default:
			return nil, wlog.Errorf("invalid mininggoodstock field")
		}
	}
	if conds.GoodStockIDs != nil {
		ids, ok := conds.GoodStockIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodstockids")
		}
		switch conds.GoodStockIDs.Op {
		case cruder.IN:
			q.Where(entmininggoodstock.GoodStockIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid mininggoodstock field")
		}
	}
	return q, nil
}
