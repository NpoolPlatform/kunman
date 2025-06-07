package manufacturer

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entdevicemanufacturer "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/devicemanufacturer"
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

func CreateSet(c *ent.DeviceManufacturerCreate, req *Req) *ent.DeviceManufacturerCreate {
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

func UpdateSet(u *ent.DeviceManufacturerUpdateOne, req *Req) *ent.DeviceManufacturerUpdateOne {
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

func SetQueryConds(q *ent.DeviceManufacturerQuery, conds *Conds) (*ent.DeviceManufacturerQuery, error) {
	q.Where(entdevicemanufacturer.DeletedAt(0))
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
			q.Where(entdevicemanufacturer.ID(id))
		case cruder.NEQ:
			q.Where(entdevicemanufacturer.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid devicemanufacturer field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entdevicemanufacturer.EntID(id))
		case cruder.NEQ:
			q.Where(entdevicemanufacturer.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid devicemanufacturer field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(entdevicemanufacturer.Name(name))
		default:
			return nil, wlog.Errorf("invalid devicemanufacturer field")
		}
	}
	return q, nil
}
