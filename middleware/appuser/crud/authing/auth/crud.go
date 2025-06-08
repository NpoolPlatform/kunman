package auth

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entauth "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/auth"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	RoleID    *uuid.UUID
	UserID    *uuid.UUID
	Resource  *string
	Method    *string
	DeletedAt *uint32
}

func CreateSet(c *ent.AuthCreate, req *Req) *ent.AuthCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.RoleID != nil {
		c.SetRoleID(*req.RoleID)
	}
	if req.UserID != nil && req.RoleID == nil {
		c.SetUserID(*req.UserID)
	}
	if req.Resource != nil {
		c.SetResource(*req.Resource)
	}
	if req.Method != nil {
		c.SetMethod(*req.Method)
	}
	return c
}

func UpdateSet(u *ent.AuthUpdateOne, req *Req) *ent.AuthUpdateOne {
	if req.Resource != nil {
		u.SetResource(*req.Resource)
	}
	if req.Method != nil {
		u.SetMethod(*req.Method)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	EntID    *cruder.Cond
	EntIDs   *cruder.Cond
	AppID    *cruder.Cond
	RoleID   *cruder.Cond
	UserID   *cruder.Cond
	Resource *cruder.Cond
	Method   *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AuthQuery, conds *Conds) (*ent.AuthQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entauth.ID(id))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entauth.EntID(id))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entauth.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entauth.AppID(id))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.RoleID != nil {
		id, ok := conds.RoleID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid roleid")
		}
		switch conds.RoleID.Op {
		case cruder.EQ:
			q.Where(entauth.RoleID(id))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entauth.UserID(id))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.Resource != nil {
		res, ok := conds.Resource.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid resource")
		}
		switch conds.Resource.Op {
		case cruder.EQ:
			q.Where(entauth.Resource(res))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.Method != nil {
		method, ok := conds.Method.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid method")
		}
		switch conds.Method.Op {
		case cruder.EQ:
			q.Where(entauth.Method(method))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	q.Where(entauth.DeletedAt(0))
	return q, nil
}
