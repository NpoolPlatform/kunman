package appstocklock

import (
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID             *uuid.UUID
	AppStockID        *uuid.UUID
	AppGoodID         *uuid.UUID
	Units             *decimal.Decimal
	AppSpotUnits      *decimal.Decimal
	AppStockLockState *types.AppStockLockState
	ExLockID          *uuid.UUID
	DeletedAt         *uint32
}

func CreateSet(c *ent.AppStockLockCreate, req *Req) *ent.AppStockLockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppStockID != nil {
		c.SetAppStockID(*req.AppStockID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.Units != nil {
		c.SetUnits(*req.Units)
	}
	if req.AppSpotUnits != nil {
		c.SetAppSpotUnits(*req.AppSpotUnits)
	}
	if req.AppStockLockState != nil {
		c.SetLockState(req.AppStockLockState.String())
	}
	if req.ExLockID != nil {
		c.SetExLockID(*req.ExLockID)
	}
	return c
}

func UpdateSet(u *ent.AppStockLockUpdateOne, req *Req) *ent.AppStockLockUpdateOne {
	if req.AppStockLockState != nil {
		u.SetLockState(req.AppStockLockState.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}
