//nolint:dupl
package deviceinfo

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entdeviceinfo "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/deviceinfo"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID               *uint32
	EntID            *uuid.UUID
	Type             *string
	ManufacturerID   *uuid.UUID
	PowerConsumption *uint32
	ShipmentAt       *uint32
	DeletedAt        *uint32
}

func CreateSet(c *ent.DeviceInfoCreate, req *Req) *ent.DeviceInfoCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Type != nil {
		c.SetType(*req.Type)
	}
	if req.ManufacturerID != nil {
		c.SetManufacturerID(*req.ManufacturerID)
	}
	if req.PowerConsumption != nil {
		c.SetPowerConsumption(*req.PowerConsumption)
	}
	if req.ShipmentAt != nil {
		c.SetShipmentAt(*req.ShipmentAt)
	}
	return c
}

func UpdateSet(u *ent.DeviceInfoUpdateOne, req *Req) *ent.DeviceInfoUpdateOne {
	if req.Type != nil {
		u.SetType(*req.Type)
	}
	if req.ManufacturerID != nil {
		u.SetManufacturerID(*req.ManufacturerID)
	}
	if req.PowerConsumption != nil {
		u.SetPowerConsumption(*req.PowerConsumption)
	}
	if req.ShipmentAt != nil {
		u.SetShipmentAt(*req.ShipmentAt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	EntID          *cruder.Cond
	Type           *cruder.Cond
	ManufacturerID *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.DeviceInfoQuery, conds *Conds) (*ent.DeviceInfoQuery, error) {
	q.Where(entdeviceinfo.DeletedAt(0))
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
			q.Where(entdeviceinfo.ID(id))
		case cruder.NEQ:
			q.Where(entdeviceinfo.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid deviceinfo field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entdeviceinfo.EntID(id))
		case cruder.NEQ:
			q.Where(entdeviceinfo.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid deviceinfo field")
		}
	}
	if conds.Type != nil {
		_type, ok := conds.Type.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid type")
		}
		switch conds.Type.Op {
		case cruder.EQ:
			q.Where(entdeviceinfo.Type(_type))
		default:
			return nil, wlog.Errorf("invalid deviceinfo field")
		}
	}
	if conds.ManufacturerID != nil {
		manufacturer, ok := conds.ManufacturerID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid manufacturerid")
		}
		switch conds.ManufacturerID.Op {
		case cruder.EQ:
			q.Where(entdeviceinfo.ManufacturerID(manufacturer))
		default:
			return nil, wlog.Errorf("invalid deviceinfo field")
		}
	}
	return q, nil
}
