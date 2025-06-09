package goodbenefit

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entgoodbenefit "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/goodbenefit"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	EntID         *uuid.UUID
	GoodID        *uuid.UUID
	AccountID     *uuid.UUID
	TransactionID *uuid.UUID
	Backup        *bool
	DeletedAt     *uint32
}

func CreateSet(c *ent.GoodBenefitCreate, req *Req) *ent.GoodBenefitCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.Backup != nil {
		c.SetBackup(*req.Backup)
	}
	return c
}

func UpdateSet(u *ent.GoodBenefitUpdateOne, req *Req) *ent.GoodBenefitUpdateOne {
	if req.Backup != nil {
		u.SetBackup(*req.Backup)
	}
	if req.TransactionID != nil {
		u.SetTransactionID(*req.TransactionID)
	}
	if req.AccountID != nil {
		u.SetAccountID(*req.AccountID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	accountcrud.Conds
	GoodID    *cruder.Cond
	AccountID *cruder.Cond
	Backup    *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.GoodBenefitQuery, conds *Conds) (*ent.GoodBenefitQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodbenefit entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.EntID(id))
		default:
			return nil, fmt.Errorf("invalid goodbenefit field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid goodbenefit id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.ID(id))
		default:
			return nil, fmt.Errorf("invalid goodbenefit field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodbenefit goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid goodbenefit field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodbenefit accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid goodbenefit field")
		}
	}
	if conds.Backup != nil {
		backup, ok := conds.Backup.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid goodbenefit backup")
		}
		switch conds.Backup.Op {
		case cruder.EQ:
			q.Where(entgoodbenefit.Backup(backup))
		default:
			return nil, fmt.Errorf("invalid goodbenefit field")
		}
	}
	q.Where(entgoodbenefit.DeletedAt(0))
	return q, nil
}
