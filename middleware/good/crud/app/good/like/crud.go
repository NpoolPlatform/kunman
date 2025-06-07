package like

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entlike "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/like"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	UserID    *uuid.UUID
	GoodID    *uuid.UUID
	AppGoodID *uuid.UUID
	Like      *bool
	DeletedAt *uint32
}

func CreateSet(c *ent.LikeCreate, req *Req) *ent.LikeCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.Like != nil {
		c.SetLike(*req.Like)
	}
	return c
}

func UpdateSet(u *ent.LikeUpdateOne, req *Req) *ent.LikeUpdateOne {
	if req.Like != nil {
		u.SetLike(*req.Like)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	UserID     *cruder.Cond
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.LikeQuery, conds *Conds) (*ent.LikeQuery, error) {
	q.Where(entlike.DeletedAt(0))
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
			q.Where(entlike.ID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entlike.EntID(id))
		default:
			return nil, wlog.Errorf("invalid id field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entlike.UserID(id))
		default:
			return nil, wlog.Errorf("invalid userid field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entlike.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid appgoodid field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entlike.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid appgoodids field")
		}
	}
	return q, nil
}
