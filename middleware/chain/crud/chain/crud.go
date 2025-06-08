package chainbase

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entchainbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/chainbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
)

type Req struct {
	EntID      *uuid.UUID
	Name       *string
	Logo       *string
	NativeUnit *string
	AtomicUnit *string
	UnitExp    *uint32
	ENV        *string
	ChainID    *string
	Nickname   *string
	GasType    *basetypes.GasType
	DeletedAt  *uint32
}

func CreateSet(c *ent.ChainBaseCreate, req *Req) *ent.ChainBaseCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	if req.NativeUnit != nil {
		c.SetNativeUnit(*req.NativeUnit)
	}
	if req.AtomicUnit != nil {
		c.SetAtomicUnit(*req.AtomicUnit)
	}
	if req.UnitExp != nil {
		c.SetUnitExp(*req.UnitExp)
	}
	if req.ENV != nil {
		c.SetEnv(*req.ENV)
	}
	if req.ChainID != nil {
		c.SetChainID(*req.ChainID)
	}
	if req.Nickname != nil {
		c.SetNickname(*req.Nickname)
	}
	if req.GasType != nil {
		c.SetGasType(req.GasType.String())
	}
	return c
}

func UpdateSet(u *ent.ChainBaseUpdateOne, req *Req) *ent.ChainBaseUpdateOne {
	if req.Logo != nil {
		u = u.SetLogo(*req.Logo)
	}
	if req.NativeUnit != nil {
		u = u.SetNativeUnit(*req.NativeUnit)
	}
	if req.AtomicUnit != nil {
		u = u.SetAtomicUnit(*req.AtomicUnit)
	}
	if req.UnitExp != nil {
		u = u.SetUnitExp(*req.UnitExp)
	}
	if req.ENV != nil {
		u = u.SetEnv(*req.ENV)
	}
	if req.ChainID != nil {
		u = u.SetChainID(*req.ChainID)
	}
	if req.Nickname != nil {
		u = u.SetNickname(*req.Nickname)
	}
	if req.GasType != nil {
		u = u.SetGasType(req.GasType.String())
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID      *cruder.Cond
	Name       *cruder.Cond
	ENV        *cruder.Cond
	NativeUnit *cruder.Cond
	ChainID    *cruder.Cond
	Nickname   *cruder.Cond
	ChainType  *cruder.Cond
	EntIDs     *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.ChainBaseQuery, conds *Conds) (*ent.ChainBaseQuery, error) {
	if conds.EntID != nil {
		name, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(
				entchainbase.EntID(name),
			)
		default:
			return nil, fmt.Errorf("invalid chainbase field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(
				entchainbase.Name(name),
			)
		default:
			return nil, fmt.Errorf("invalid chainbase field")
		}
	}
	if conds.ENV != nil {
		env, ok := conds.ENV.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid env")
		}
		switch conds.ENV.Op {
		case cruder.EQ:
			q.Where(entchainbase.Env(env))
		default:
			return nil, fmt.Errorf("invalid chainbase field")
		}
	}
	if conds.NativeUnit != nil {
		unit, ok := conds.NativeUnit.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid nativeunit")
		}
		switch conds.NativeUnit.Op {
		case cruder.EQ:
			q.Where(entchainbase.NativeUnit(unit))
		default:
			return nil, fmt.Errorf("invalid chainbase field")
		}
	}
	if conds.ChainID != nil {
		chainID, ok := conds.ChainID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid chainID")
		}
		switch conds.ChainID.Op {
		case cruder.EQ:
			q.Where(entchainbase.ChainID(chainID))
		default:
			return nil, fmt.Errorf("invalid chainbase field")
		}
	}
	if conds.Nickname != nil {
		nickname, ok := conds.Nickname.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid nickname")
		}
		switch conds.Nickname.Op {
		case cruder.EQ:
			q.Where(entchainbase.Nickname(nickname))
		default:
			return nil, fmt.Errorf("invalid chainbase field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entchainbase.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid chainbase field")
		}
	}
	q.Where(entchainbase.DeletedAt(0))
	return q, nil
}
