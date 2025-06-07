package good

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entdelegatedstaking "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/delegatedstaking"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID              *uuid.UUID
	GoodID             *uuid.UUID
	ContractCodeURL    *string
	ContractCodeBranch *string
	ContractState      *types.ContractState
	DeletedAt          *uint32
}

func CreateSet(c *ent.DelegatedStakingCreate, req *Req) *ent.DelegatedStakingCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.ContractCodeURL != nil {
		c.SetContractCodeURL(*req.ContractCodeURL)
	}
	if req.ContractCodeBranch != nil {
		c.SetContractCodeBranch(*req.ContractCodeBranch)
	}
	if req.ContractState != nil {
		c.SetContractState(req.ContractState.String())
	}
	return c
}

func UpdateSet(u *ent.DelegatedStakingUpdateOne, req *Req) *ent.DelegatedStakingUpdateOne {
	if req.ContractCodeURL != nil {
		u.SetContractCodeURL(*req.ContractCodeURL)
	}
	if req.ContractCodeBranch != nil {
		u.SetContractCodeBranch(*req.ContractCodeBranch)
	}
	if req.ContractState != nil {
		u.SetContractState(req.ContractState.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	IDs            *cruder.Cond
	EntID          *cruder.Cond
	EntIDs         *cruder.Cond
	GoodID         *cruder.Cond
	GoodIDs        *cruder.Cond
	ContractState  *cruder.Cond
	ContractStates *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.DelegatedStakingQuery, conds *Conds) (*ent.DelegatedStakingQuery, error) {
	q.Where(entdelegatedstaking.DeletedAt(0))
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
			q.Where(entdelegatedstaking.ID(id))
		case cruder.NEQ:
			q.Where(entdelegatedstaking.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entdelegatedstaking.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entdelegatedstaking.EntID(id))
		case cruder.NEQ:
			q.Where(entdelegatedstaking.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entdelegatedstaking.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entdelegatedstaking.GoodID(id))
		case cruder.NEQ:
			q.Where(entdelegatedstaking.GoodIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entdelegatedstaking.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.ContractState != nil {
		_mode, ok := conds.ContractState.Val.(types.ContractState)
		if !ok {
			return nil, wlog.Errorf("invalid contractstate")
		}
		switch conds.ContractState.Op {
		case cruder.EQ:
			q.Where(entdelegatedstaking.ContractState(_mode.String()))
		case cruder.NEQ:
			q.Where(entdelegatedstaking.ContractStateNEQ(_mode.String()))
		default:
			return nil, wlog.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.ContractStates != nil {
		_types, ok := conds.ContractStates.Val.([]types.ContractState)
		if !ok {
			return nil, wlog.Errorf("invalid contractstates")
		}
		es := []string{}
		for _, _type := range _types {
			es = append(es, _type.String())
		}
		switch conds.ContractStates.Op {
		case cruder.IN:
			q.Where(entdelegatedstaking.ContractStateIn(es...))
		default:
			return nil, wlog.Errorf("invalid delegatedstaking field")
		}
	}
	return q, nil
}
