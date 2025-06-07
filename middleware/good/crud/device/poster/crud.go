package deviceposter

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entdeviceposter "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/deviceposter"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	DeviceTypeID *uuid.UUID
	Poster       *string
	Index        *uint8
	DeletedAt    *uint32
}

func CreateSet(c *ent.DevicePosterCreate, req *Req) *ent.DevicePosterCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.DeviceTypeID != nil {
		c.SetDeviceTypeID(*req.DeviceTypeID)
	}
	if req.Poster != nil {
		c.SetPoster(*req.Poster)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

func UpdateSet(u *ent.DevicePosterUpdateOne, req *Req) *ent.DevicePosterUpdateOne {
	if req.Poster != nil {
		u.SetPoster(*req.Poster)
	}
	if req.Index != nil {
		u.SetIndex(*req.Index)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID            *cruder.Cond
	IDs           *cruder.Cond
	EntID         *cruder.Cond
	EntIDs        *cruder.Cond
	DeviceTypeID  *cruder.Cond
	DeviceTypeIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.DevicePosterQuery, conds *Conds) (*ent.DevicePosterQuery, error) {
	q.Where(entdeviceposter.DeletedAt(0))
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
			q.Where(entdeviceposter.ID(id))
		default:
			return nil, wlog.Errorf("invalid deviceposter field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entdeviceposter.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid deviceposter field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entdeviceposter.EntID(id))
		default:
			return nil, wlog.Errorf("invalid deviceposter field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entdeviceposter.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid deviceposter field")
		}
	}
	if conds.DeviceTypeID != nil {
		id, ok := conds.DeviceTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid devicetypeid")
		}
		switch conds.DeviceTypeID.Op {
		case cruder.EQ:
			q.Where(entdeviceposter.DeviceTypeID(id))
		default:
			return nil, wlog.Errorf("invalid deviceposter field")
		}
	}
	if conds.DeviceTypeIDs != nil {
		ids, ok := conds.DeviceTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid devicetypeids")
		}
		switch conds.DeviceTypeIDs.Op {
		case cruder.IN:
			q.Where(entdeviceposter.DeviceTypeIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid deviceposter field")
		}
	}
	return q, nil
}
