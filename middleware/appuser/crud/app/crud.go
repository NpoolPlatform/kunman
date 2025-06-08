package app

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entapp "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/app"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID       *uuid.UUID
	CreatedBy   *uuid.UUID
	Name        *string
	Logo        *string
	Description *string
}

func CreateSet(c *ent.AppCreate, req *Req) *ent.AppCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.CreatedBy != nil {
		c.SetCreatedBy(*req.CreatedBy)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	return c
}

func UpdateSet(u *ent.AppUpdateOne, req *Req) *ent.AppUpdateOne {
	if req.EntID != nil {
		u.SetEntID(*req.EntID)
	}
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Logo != nil {
		u.SetLogo(*req.Logo)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	EntID     *cruder.Cond
	EntIDs    *cruder.Cond
	CreatedBy *cruder.Cond
	Name      *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AppQuery, conds *Conds) (*ent.AppQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entapp.ID(id))
		case cruder.NEQ:
			q.Where(entapp.IDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entapp.EntID(id))
		case cruder.NEQ:
			q.Where(entapp.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entapp.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.CreatedBy != nil {
		createdBy, ok := conds.CreatedBy.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid creator")
		}
		switch conds.CreatedBy.Op {
		case cruder.EQ:
			q.Where(entapp.CreatedBy(createdBy))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(entapp.Name(name))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	return q, nil
}
