package ban

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entbanappuser "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/banappuser"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	Message   *string
	DeletedAt *uint32
}

func CreateSet(c *ent.BanAppUserCreate, req *Req) *ent.BanAppUserCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.BanAppUserUpdateOne, req *Req) *ent.BanAppUserUpdateOne {
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID  *cruder.Cond
	AppID  *cruder.Cond
	UserID *cruder.Cond
}

func SetQueryConds(q *ent.BanAppUserQuery, conds *Conds) (*ent.BanAppUserQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entbanappuser.EntID(id))
		default:
			return nil, fmt.Errorf("invalid banapp field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entbanappuser.AppID(id))
		default:
			return nil, fmt.Errorf("invalid banapp field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entbanappuser.UserID(id))
		default:
			return nil, fmt.Errorf("invalid banapp field")
		}
	}
	q.Where(entbanappuser.DeletedAt(0))
	return q, nil
}
