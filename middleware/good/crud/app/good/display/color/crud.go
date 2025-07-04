package appgooddisplaycolor

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgooddisplaycolor "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgooddisplaycolor"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppGoodID *uuid.UUID
	Color     *string
	Index     *uint8
	DeletedAt *uint32
}

func CreateSet(c *ent.AppGoodDisplayColorCreate, req *Req) *ent.AppGoodDisplayColorCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.Color != nil {
		c.SetColor(*req.Color)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

func UpdateSet(u *ent.AppGoodDisplayColorUpdateOne, req *Req) *ent.AppGoodDisplayColorUpdateOne {
	if req.Color != nil {
		u.SetColor(*req.Color)
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
func SetQueryConds(q *ent.AppGoodDisplayColorQuery, conds *Conds) (*ent.AppGoodDisplayColorQuery, error) {
	q.Where(entappgooddisplaycolor.DeletedAt(0))
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
			q.Where(entappgooddisplaycolor.ID(id))
		default:
			return nil, wlog.Errorf("invalid appgooddisplaycolor field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappgooddisplaycolor.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgooddisplaycolor field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappgooddisplaycolor.EntID(id))
		default:
			return nil, wlog.Errorf("invalid appgooddisplaycolor field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappgooddisplaycolor.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgooddisplaycolor field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappgooddisplaycolor.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid appgooddisplaycolor field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entappgooddisplaycolor.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgooddisplaycolor field")
		}
	}
	return q, nil
}
