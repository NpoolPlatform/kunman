package contract

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated"
	entcontract "github.com/NpoolPlatform/kunman/middleware/account/db/ent/generated/contract"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/account/v1"
	accountcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/account"

	"github.com/google/uuid"
)

type Req struct {
	EntID                *uuid.UUID
	GoodID               *uuid.UUID
	DelegatedStakingID   *uuid.UUID
	AccountID            *uuid.UUID
	Backup               *bool
	ContractOperatorType *basetypes.ContractOperatorType
	DeletedAt            *uint32
}

func CreateSet(c *ent.ContractCreate, req *Req) *ent.ContractCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.DelegatedStakingID != nil {
		c.SetDelegatedStakingID(*req.DelegatedStakingID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.Backup != nil {
		c.SetBackup(*req.Backup)
	}
	if req.ContractOperatorType != nil {
		c.SetContractOperatorType(req.ContractOperatorType.String())
	}
	return c
}

func UpdateSet(u *ent.ContractUpdateOne, req *Req) *ent.ContractUpdateOne {
	if req.Backup != nil {
		u.SetBackup(*req.Backup)
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
	GoodID               *cruder.Cond
	AccountID            *cruder.Cond
	Backup               *cruder.Cond
	ContractOperatorType *cruder.Cond
	DelegatedStakingID   *cruder.Cond
	DelegatedStakingIDs  *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.ContractQuery, conds *Conds) (*ent.ContractQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid contract entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcontract.EntID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid contract id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entcontract.ID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid contract goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entcontract.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid contract accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entcontract.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.Backup != nil {
		backup, ok := conds.Backup.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid contract backup")
		}
		switch conds.Backup.Op {
		case cruder.EQ:
			q.Where(entcontract.Backup(backup))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.ContractOperatorType != nil {
		contractOperatorType, ok := conds.ContractOperatorType.Val.(basetypes.ContractOperatorType)
		if !ok {
			return nil, fmt.Errorf("invalid account ContractOperatorType")
		}
		switch conds.ContractOperatorType.Op {
		case cruder.EQ:
			q.Where(entcontract.ContractOperatorType(contractOperatorType.String()))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	if conds.DelegatedStakingID != nil {
		id, ok := conds.DelegatedStakingID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid contract DelegatedStakingID")
		}
		switch conds.DelegatedStakingID.Op {
		case cruder.EQ:
			q.Where(entcontract.DelegatedStakingID(id))
		default:
			return nil, fmt.Errorf("invalid contract field")
		}
	}
	if conds.DelegatedStakingIDs != nil {
		ids, ok := conds.DelegatedStakingIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid account DelegatedStakingIDs")
		}
		switch conds.DelegatedStakingIDs.Op {
		case cruder.IN:
			q.Where(entcontract.DelegatedStakingIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid account field")
		}
	}
	q.Where(entcontract.DeletedAt(0))
	return q, nil
}
