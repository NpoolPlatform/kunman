package goodmalfunction

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodmalfunction "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodmalfunction"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID             *uuid.UUID
	GoodID            *uuid.UUID
	Title             *string
	Message           *string
	StartAt           *uint32
	DurationSeconds   *uint32
	CompensateSeconds *uint32
	DeletedAt         *uint32
}

func CreateSet(c *ent.GoodMalfunctionCreate, req *Req) *ent.GoodMalfunctionCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.DurationSeconds != nil {
		c.SetDurationSeconds(*req.DurationSeconds)
	}
	if req.CompensateSeconds != nil {
		c.SetCompensateSeconds(*req.CompensateSeconds)
	}
	return c
}

func UpdateSet(u *ent.GoodMalfunctionUpdateOne, req *Req) *ent.GoodMalfunctionUpdateOne {
	if req.Title != nil {
		u.SetTitle(*req.Title)
	}
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.DurationSeconds != nil {
		u.SetDurationSeconds(*req.DurationSeconds)
	}
	if req.CompensateSeconds != nil {
		u.SetCompensateSeconds(*req.CompensateSeconds)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID     *cruder.Cond
	EntID  *cruder.Cond
	EntIDs *cruder.Cond
	GoodID *cruder.Cond
}

func SetQueryConds(q *ent.GoodMalfunctionQuery, conds *Conds) (*ent.GoodMalfunctionQuery, error) {
	q.Where(entgoodmalfunction.DeletedAt(0))
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
			q.Where(entgoodmalfunction.ID(id))
		default:
			return nil, wlog.Errorf("invalid goodmalfunction field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodmalfunction.EntID(id))
		default:
			return nil, wlog.Errorf("invalid goodmalfunction field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entgoodmalfunction.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid goodmalfunction field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodmalfunction.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid goodmalfunction field")
		}
	}
	return q, nil
}
