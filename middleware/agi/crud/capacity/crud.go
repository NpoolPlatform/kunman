package capacity

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/agi/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entcapacity "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/capacity"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	AppGoodID   *uuid.UUID
	CapacityKey *types.CapacityKey
	Value       *string
	Description *string
	DeletedAt   *uint32
}

func CreateSet(c *ent.CapacityCreate, req *Req) *ent.CapacityCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.CapacityKey != nil {
		c.SetCapacityKey(req.CapacityKey.String())
	}
	if req.Value != nil {
		c.SetCapacityValue(*req.Value)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	return c
}

func UpdateSet(u *ent.CapacityUpdateOne, req *Req) *ent.CapacityUpdateOne {
	if req.Value != nil {
		u.SetCapacityValue(*req.Value)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.CapacityQuery, conds *Conds) (*ent.CapacityQuery, error) {
	q.Where(entcapacity.DeletedAt(0))
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
			q.Where(entcapacity.ID(id))
		default:
			return nil, wlog.Errorf("invalid capacity field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entcapacity.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid capacity field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcapacity.EntID(id))
		default:
			return nil, wlog.Errorf("invalid capacity field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entcapacity.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid capacity field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entcapacity.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid capacity field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entcapacity.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid capacity field")
		}
	}

	return q, nil
}
