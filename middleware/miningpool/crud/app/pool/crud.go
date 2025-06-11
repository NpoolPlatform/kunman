package apppool

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated"
	entapppool "github.com/NpoolPlatform/kunman/middleware/miningpool/db/ent/generated/apppool"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	PoolID    *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.AppPoolCreate, req *Req) *ent.AppPoolCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.PoolID != nil {
		c.SetPoolID(*req.PoolID)
	}
	return c
}

func UpdateSet(u *ent.AppPoolUpdateOne, req *Req) (*ent.AppPoolUpdateOne, error) {
	if req.AppID != nil {
		u = u.SetAppID(*req.AppID)
	}
	if req.PoolID != nil {
		u = u.SetPoolID(*req.PoolID)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u, nil
}

type Conds struct {
	ID      *cruder.Cond
	EntID   *cruder.Cond
	AppID   *cruder.Cond
	PoolID  *cruder.Cond
	EntIDs  *cruder.Cond
	PoolIDs *cruder.Cond
}

func SetQueryConds(q *ent.AppPoolQuery, conds *Conds) (*ent.AppPoolQuery, error) { //nolint
	q.Where(entapppool.DeletedAt(0))
	if conds == nil {
		return nil, wlog.Errorf("have no any conds")
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entapppool.ID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entapppool.EntID(id))
		default:
			return nil, wlog.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entapppool.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid entids field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entapppool.AppID(id))
		default:
			return nil, wlog.Errorf("invalid appid field")
		}
	}
	if conds.PoolID != nil {
		id, ok := conds.PoolID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid poolid")
		}
		switch conds.PoolID.Op {
		case cruder.EQ:
			q.Where(entapppool.PoolID(id))
		default:
			return nil, wlog.Errorf("invalid poolid field")
		}
	}
	if conds.PoolIDs != nil {
		ids, ok := conds.PoolIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid poolids")
		}
		if len(ids) > 0 {
			switch conds.PoolIDs.Op {
			case cruder.IN:
				q.Where(entapppool.PoolIDIn(ids...))
			default:
				return nil, wlog.Errorf("invalid compensate field")
			}
		}
	}
	return q, nil
}
