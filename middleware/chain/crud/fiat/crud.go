package currencytype

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entfiat "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiat"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	Name      *string
	Logo      *string
	Unit      *string
	DeletedAt *uint32
}

func CreateSet(c *ent.FiatCreate, req *Req) *ent.FiatCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	if req.Unit != nil {
		c.SetUnit(*req.Unit)
	}
	return c
}

func UpdateSet(u *ent.FiatUpdateOne, req *Req) *ent.FiatUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Logo != nil {
		u.SetLogo(*req.Logo)
	}
	if req.Unit != nil {
		u.SetUnit(*req.Unit)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID  *cruder.Cond
	EntIDs *cruder.Cond
	Name   *cruder.Cond
	Unit   *cruder.Cond
}

func SetQueryConds(q *ent.FiatQuery, conds *Conds) (*ent.FiatQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entfiat.EntID(id))
		default:
			return nil, fmt.Errorf("invalid fiat field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entfiat.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid fiat field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(entfiat.Name(name))
		default:
			return nil, fmt.Errorf("invalid fiat field")
		}
	}
	if conds.Unit != nil {
		unit, ok := conds.Unit.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Unit.Op {
		case cruder.EQ:
			q.Where(entfiat.Unit(unit))
		default:
			return nil, fmt.Errorf("invalid fiat field")
		}
	}
	q.Where(entfiat.DeletedAt(0))
	return q, nil
}
