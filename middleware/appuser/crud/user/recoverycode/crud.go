package user

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entrecoverycode "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/recoverycode"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	Used      *bool
	DeletedAt *uint32
}

func CreateSet(c *ent.RecoveryCodeCreate, req *Req) *ent.RecoveryCodeCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Used != nil {
		c.SetUsed(*req.Used)
	}
	return c
}

func UpdateSet(u *ent.RecoveryCodeUpdateOne, req *Req) *ent.RecoveryCodeUpdateOne {
	if req.Used != nil {
		u.SetUsed(*req.Used)
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
	UserID *cruder.Cond
	AppID  *cruder.Cond
	Code   *cruder.Cond
	Used   *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.RecoveryCodeQuery, conds *Conds) (*ent.RecoveryCodeQuery, error) {
	q.Where(entrecoverycode.DeletedAt(0))
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
			q.Where(entrecoverycode.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.UserID(id))
		default:
			return nil, fmt.Errorf("invalid userid field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.Code != nil {
		code, ok := conds.Code.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid code")
		}
		switch conds.Code.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.Code(code))
		default:
			return nil, fmt.Errorf("invalid code field")
		}
	}
	if conds.Used != nil {
		used, ok := conds.Used.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid used")
		}
		switch conds.Used.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.Used(used))
		default:
			return nil, fmt.Errorf("invalid used field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entrecoverycode.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	return q, nil
}
