package deposit

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entdeposit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/deposit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	EntID         *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	AccountID     *uuid.UUID
	CollectingTID *uuid.UUID
	Incoming      *decimal.Decimal
	Outcoming     *decimal.Decimal
	ScannableAt   *uint32
	DeletedAt     *uint32
}

func CreateSet(c *ent.DepositCreate, req *Req) *ent.DepositCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}

	c.SetIncoming(decimal.NewFromInt(0))
	c.SetOutcoming(decimal.NewFromInt(0))
	c.SetScannableAt(uint32(time.Now().Unix()))

	return c
}

func UpdateSet(u *ent.DepositUpdateOne, req *Req) *ent.DepositUpdateOne {
	if req.CollectingTID != nil {
		u.SetCollectingTid(*req.CollectingTID)
	}
	if req.Incoming != nil {
		u.SetIncoming(*req.Incoming)
	}
	if req.Outcoming != nil {
		u.SetOutcoming(*req.Outcoming)
	}
	if req.ScannableAt != nil {
		u.SetScannableAt(*req.ScannableAt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	accountcrud.Conds
	AppID       *cruder.Cond
	UserID      *cruder.Cond
	AccountID   *cruder.Cond
	ScannableAt *cruder.Cond
}

func SetQueryConds(q *ent.DepositQuery, conds *Conds) (*ent.DepositQuery, error) { //nolint
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid deposit entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entdeposit.EntID(id))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid deposit id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entdeposit.ID(id))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid deposit appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entdeposit.AppID(id))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid deposit userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entdeposit.UserID(id))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid deposit accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entdeposit.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	if conds.ScannableAt != nil {
		at, ok := conds.ScannableAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid deposit scannableat")
		}
		switch conds.ScannableAt.Op {
		case cruder.LT:
			q.Where(entdeposit.ScannableAtLT(at))
		case cruder.GT:
			q.Where(entdeposit.ScannableAtGT(at))
		default:
			return nil, fmt.Errorf("invalid deposit field")
		}
	}
	q.Where(entdeposit.DeletedAt(0))
	return q, nil
}
