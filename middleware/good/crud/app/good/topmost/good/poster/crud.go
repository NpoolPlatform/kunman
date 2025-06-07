package topmostgoodposter

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"

	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	enttopmostgoodposter "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/topmostgoodposter"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID         *uuid.UUID
	TopMostGoodID *uuid.UUID
	Poster        *string
	Index         *uint8
	DeletedAt     *uint32
}

func CreateSet(c *ent.TopMostGoodPosterCreate, req *Req) *ent.TopMostGoodPosterCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.TopMostGoodID != nil {
		c.SetTopMostGoodID(*req.TopMostGoodID)
	}
	if req.Poster != nil {
		c.SetPoster(*req.Poster)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

func UpdateSet(u *ent.TopMostGoodPosterUpdateOne, req *Req) *ent.TopMostGoodPosterUpdateOne {
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
	ID             *cruder.Cond
	IDs            *cruder.Cond
	EntID          *cruder.Cond
	EntIDs         *cruder.Cond
	TopMostGoodID  *cruder.Cond
	TopMostGoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.TopMostGoodPosterQuery, conds *Conds) (*ent.TopMostGoodPosterQuery, error) {
	q.Where(enttopmostgoodposter.DeletedAt(0))
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
			q.Where(enttopmostgoodposter.ID(id))
		default:
			return nil, wlog.Errorf("invalid topmostgoodposter field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(enttopmostgoodposter.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid topmostgoodposter field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttopmostgoodposter.EntID(id))
		default:
			return nil, wlog.Errorf("invalid topmostgoodposter field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttopmostgoodposter.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid topmostgoodposter field")
		}
	}
	if conds.TopMostGoodID != nil {
		id, ok := conds.TopMostGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid topmostid")
		}
		switch conds.TopMostGoodID.Op {
		case cruder.EQ:
			q.Where(enttopmostgoodposter.TopMostGoodID(id))
		default:
			return nil, wlog.Errorf("invalid topmostgoodposter field")
		}
	}
	if conds.TopMostGoodIDs != nil {
		ids, ok := conds.TopMostGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid topmostids")
		}
		switch conds.TopMostGoodIDs.Op {
		case cruder.IN:
			q.Where(enttopmostgoodposter.TopMostGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid topmostgoodposter field")
		}
	}
	return q, nil
}
