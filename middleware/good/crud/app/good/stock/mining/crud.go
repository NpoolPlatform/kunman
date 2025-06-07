package mining

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappmininggoodstock "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appmininggoodstock"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID             *uuid.UUID
	AppGoodStockID    *uuid.UUID
	MiningGoodStockID *uuid.UUID
	Reserved          *decimal.Decimal
	SpotQuantity      *decimal.Decimal
	Locked            *decimal.Decimal
	InService         *decimal.Decimal
	WaitStart         *decimal.Decimal
	Sold              *decimal.Decimal
	DeletedAt         *uint32
}

func CreateSet(c *ent.AppMiningGoodStockCreate, req *Req) *ent.AppMiningGoodStockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodStockID != nil {
		c.SetAppGoodStockID(*req.AppGoodStockID)
	}
	if req.MiningGoodStockID != nil {
		c.SetMiningGoodStockID(*req.MiningGoodStockID)
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

func UpdateSet(u *ent.AppMiningGoodStockUpdateOne, req *Req) *ent.AppMiningGoodStockUpdateOne {
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
	ID                 *cruder.Cond
	EntID              *cruder.Cond
	EntIDs             *cruder.Cond
	AppGoodStockID     *cruder.Cond
	AppGoodStockIDs    *cruder.Cond
	MiningGoodStockID  *cruder.Cond
	MiningGoodStockIDs *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.AppMiningGoodStockQuery, conds *Conds) (*ent.AppMiningGoodStockQuery, error) {
	q.Where(entappmininggoodstock.DeletedAt(0))
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
			q.Where(entappmininggoodstock.ID(id))
		default:
			return nil, wlog.Errorf("invalid appmininggoodstock field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappmininggoodstock.EntID(id))
		default:
			return nil, wlog.Errorf("invalid appmininggoodstock field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappmininggoodstock.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appmininggoodstock field")
		}
	}
	if conds.AppGoodStockID != nil {
		id, ok := conds.AppGoodStockID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodstockid")
		}
		switch conds.AppGoodStockID.Op {
		case cruder.EQ:
			q.Where(entappmininggoodstock.AppGoodStockID(id))
		default:
			return nil, wlog.Errorf("invalid appmininggoodstock field")
		}
	}
	if conds.AppGoodStockIDs != nil {
		ids, ok := conds.AppGoodStockIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodstockids")
		}
		switch conds.AppGoodStockIDs.Op {
		case cruder.IN:
			q.Where(entappmininggoodstock.AppGoodStockIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appmininggoodstock field")
		}
	}
	if conds.MiningGoodStockID != nil {
		id, ok := conds.MiningGoodStockID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid mininggoodstockid")
		}
		switch conds.MiningGoodStockID.Op {
		case cruder.EQ:
			q.Where(entappmininggoodstock.MiningGoodStockID(id))
		default:
			return nil, wlog.Errorf("invalid appmininggoodstock field")
		}
	}
	if conds.MiningGoodStockIDs != nil {
		ids, ok := conds.MiningGoodStockIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid mininggoodstockids")
		}
		switch conds.MiningGoodStockIDs.Op {
		case cruder.IN:
			q.Where(entappmininggoodstock.MiningGoodStockIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appmininggoodstock field")
		}
	}
	return q, nil
}
