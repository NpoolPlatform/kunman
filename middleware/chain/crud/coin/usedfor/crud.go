package coinusedfor

import (
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/chain/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcoinusedfor "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinusedfor"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	CoinTypeID *uuid.UUID
	UsedFor    *types.CoinUsedFor
	Priority   *uint32
	DeletedAt  *uint32
}

func CreateSet(c *ent.CoinUsedForCreate, req *Req) *ent.CoinUsedForCreate {
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.UsedFor != nil {
		c.SetUsedFor(req.UsedFor.String())
	}
	if req.Priority != nil {
		c.SetPriority(*req.Priority)
	}
	return c
}

func UpdateSet(u *ent.CoinUsedForUpdateOne, req *Req) *ent.CoinUsedForUpdateOne {
	if req.Priority != nil {
		u.SetPriority(*req.Priority)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	UsedFor     *cruder.Cond
	UsedFors    *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.CoinUsedForQuery, conds *Conds) (*ent.CoinUsedForQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcoinusedfor.EntID(id))
		default:
			return nil, fmt.Errorf("invalid usedfor field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcoinusedfor.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid usedfor field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entcoinusedfor.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid usedfor field")
		}
	}
	if conds.UsedFor != nil {
		usedFor, ok := conds.UsedFor.Val.(types.CoinUsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid usedfor")
		}
		switch conds.UsedFor.Op {
		case cruder.EQ:
			q.Where(entcoinusedfor.UsedFor(usedFor.String()))
		default:
			return nil, fmt.Errorf("invalid usedfor field")
		}
	}
	if conds.UsedFors != nil {
		usedFors, ok := conds.UsedFors.Val.([]types.CoinUsedFor)
		if !ok {
			return nil, fmt.Errorf("invalid usedfors")
		}
		_usedFors := []string{}
		for _, usedFor := range usedFors {
			_usedFors = append(_usedFors, usedFor.String())
		}
		switch conds.UsedFors.Op {
		case cruder.IN:
			q.Where(entcoinusedfor.UsedForIn(_usedFors...))
		default:
			return nil, fmt.Errorf("invalid usedfor field")
		}
	}
	q.Where(entcoinusedfor.DeletedAt(0))
	return q, nil
}
