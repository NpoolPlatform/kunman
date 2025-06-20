package history

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entauthhistory "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/authhistory"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID    *uuid.UUID
	AppID    *uuid.UUID
	UserID   *uuid.UUID
	Resource *string
	Method   *string
	Allowed  *bool
}

func CreateSet(c *ent.AuthHistoryCreate, req *Req) *ent.AuthHistoryCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Resource != nil {
		c.SetResource(*req.Resource)
	}
	if req.Method != nil {
		c.SetMethod(*req.Method)
	}
	if req.Allowed != nil {
		c.SetAllowed(*req.Allowed)
	}
	return c
}

func UpdateSet(u *ent.AuthHistoryUpdateOne, in *Req) *ent.AuthHistoryUpdateOne {
	return u
}

type Conds struct {
	EntID    *cruder.Cond
	AppID    *cruder.Cond
	UserID   *cruder.Cond
	Resource *cruder.Cond
	Method   *cruder.Cond
	Allowed  *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.AuthHistoryQuery, conds *Conds) (*ent.AuthHistoryQuery, error) {
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
			q.Where(entauthhistory.EntID(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entauthhistory.AppID(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entauthhistory.UserID(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Resource != nil {
		res, ok := conds.Resource.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid resource")
		}
		switch conds.Resource.Op {
		case cruder.EQ:
			q.Where(entauthhistory.Resource(res))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Method != nil {
		method, ok := conds.Method.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid method")
		}
		switch conds.Method.Op {
		case cruder.EQ:
			q.Where(entauthhistory.Method(method))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Allowed != nil {
		allowed, ok := conds.Allowed.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid allowed")
		}
		switch conds.Allowed.Op {
		case cruder.EQ:
			q.Where(entauthhistory.Allowed(allowed))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	q.Where(entauthhistory.DeletedAt(0))
	return q, nil
}
