package lock

import (
	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID              *uint32
	EntID           *uuid.UUID
	LedgerID        *uuid.UUID
	StatementID     *uuid.UUID
	Amount          *decimal.Decimal
	LedgerLockState *types.LedgerLockState
	ExLockID        *uuid.UUID
	DeletedAt       *uint32
}

func CreateSet(c *ent.LedgerLockCreate, in *Req) *ent.LedgerLockCreate {
	if in.ID != nil {
		c.SetID(*in.ID)
	}
	if in.EntID != nil {
		c.SetEntID(*in.EntID)
	}
	if in.LedgerID != nil {
		c.SetLedgerID(*in.LedgerID)
	}
	if in.StatementID != nil {
		c.SetStatementID(*in.StatementID)
	}
	if in.Amount != nil {
		c.SetAmount(*in.Amount)
	}
	if in.LedgerLockState != nil {
		c.SetLockState(in.LedgerLockState.String())
	}
	if in.ExLockID != nil {
		c.SetExLockID(*in.ExLockID)
	}
	return c
}

func UpdateSet(u *ent.LedgerLockUpdateOne, req *Req) *ent.LedgerLockUpdateOne {
	if req.LedgerLockState != nil {
		u.SetLockState(req.LedgerLockState.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}
