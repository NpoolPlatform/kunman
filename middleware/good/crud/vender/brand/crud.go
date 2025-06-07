package brand

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entvendorbrand "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/vendorbrand"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	Name      *string
	Logo      *string
	DeletedAt *uint32
}

func CreateSet(c *ent.VendorBrandCreate, req *Req) *ent.VendorBrandCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	return c
}

func UpdateSet(u *ent.VendorBrandUpdateOne, req *Req) *ent.VendorBrandUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Logo != nil {
		u.SetLogo(*req.Logo)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID    *cruder.Cond
	EntID *cruder.Cond
	Name  *cruder.Cond
}

func SetQueryConds(q *ent.VendorBrandQuery, conds *Conds) (*ent.VendorBrandQuery, error) {
	q.Where(entvendorbrand.DeletedAt(0))
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
			q.Where(entvendorbrand.ID(id))
		case cruder.NEQ:
			q.Where(entvendorbrand.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid vendorbrand field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entvendorbrand.EntID(id))
		case cruder.NEQ:
			q.Where(entvendorbrand.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid vendorbrand field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(entvendorbrand.Name(name))
		default:
			return nil, wlog.Errorf("invalid vendorbrand field")
		}
	}
	return q, nil
}
