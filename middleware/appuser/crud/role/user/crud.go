//nolint:nolintlint,dupl
package approleuser

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entapproleuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/approleuser"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	RoleID    *uuid.UUID
	UserID    *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.AppRoleUserCreate, req *Req) *ent.AppRoleUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.RoleID != nil {
		c.SetRoleID(*req.RoleID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	return c
}

func UpdateSet(u *ent.AppRoleUserUpdateOne, req *Req) *ent.AppRoleUserUpdateOne {
	if req.RoleID != nil {
		u.SetRoleID(*req.RoleID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	EntID   *cruder.Cond
	AppID   *cruder.Cond
	RoleID  *cruder.Cond
	UserID  *cruder.Cond
	AppIDs  *cruder.Cond
	RoleIDs *cruder.Cond
	Genesis *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.AppRoleUserQuery, conds *Conds) (*ent.AppRoleUserQuery, error) {
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
			q.Where(entapproleuser.ID(id))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entapproleuser.EntID(id))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entapproleuser.AppID(id))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entapproleuser.AppIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entapproleuser.UserID(id))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.RoleID != nil {
		id, ok := conds.RoleID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.RoleID.Op {
		case cruder.EQ:
			q.Where(entapproleuser.RoleID(id))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.RoleIDs != nil {
		ids, ok := conds.RoleIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.RoleIDs.Op {
		case cruder.IN:
			q.Where(entapproleuser.RoleIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	return q, nil
}
