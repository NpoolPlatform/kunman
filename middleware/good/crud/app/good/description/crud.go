package appgooddescription

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgooddescription "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgooddescription"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID       *uuid.UUID
	AppGoodID   *uuid.UUID
	Description *string
	Index       *uint8
	DeletedAt   *uint32
}

func CreateSet(c *ent.AppGoodDescriptionCreate, req *Req) *ent.AppGoodDescriptionCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

func UpdateSet(u *ent.AppGoodDescriptionUpdateOne, req *Req) *ent.AppGoodDescriptionUpdateOne {
	if req.Description != nil {
		u.SetDescription(*req.Description)
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
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AppGoodDescriptionQuery, conds *Conds) (*ent.AppGoodDescriptionQuery, error) {
	q.Where(entappgooddescription.DeletedAt(0))
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
			q.Where(entappgooddescription.ID(id))
		default:
			return nil, wlog.Errorf("invalid appgooddescription field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappgooddescription.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgooddescription field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappgooddescription.EntID(id))
		default:
			return nil, wlog.Errorf("invalid appgooddescription field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappgooddescription.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgooddescription field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappgooddescription.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid appgooddescription field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entappgooddescription.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgooddescription field")
		}
	}
	return q, nil
}
