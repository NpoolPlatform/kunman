package secret

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entappusersecret "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appusersecret"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	UserID       *uuid.UUID
	PasswordHash *string
	Salt         *string
	GoogleSecret *string
}

func CreateSet(c *ent.AppUserSecretCreate, req *Req) *ent.AppUserSecretCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.PasswordHash != nil {
		c.SetPasswordHash(*req.PasswordHash)
	}
	if req.Salt != nil {
		c.SetSalt(*req.Salt)
	}
	if req.GoogleSecret != nil {
		c.SetGoogleSecret(*req.GoogleSecret)
	}
	return c
}

func UpdateSet(u *ent.AppUserSecretUpdateOne, req *Req) *ent.AppUserSecretUpdateOne {
	if req.PasswordHash != nil {
		u.SetPasswordHash(*req.PasswordHash)
	}
	if req.Salt != nil {
		u.SetSalt(*req.Salt)
	}
	if req.GoogleSecret != nil {
		u.SetGoogleSecret(*req.GoogleSecret)
	}
	return u
}

type Conds struct {
	EntID  *cruder.Cond
	AppID  *cruder.Cond
	UserID *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.AppUserSecretQuery, conds *Conds) (*ent.AppUserSecretQuery, error) {
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
			q.Where(entappusersecret.EntID(id))
		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappusersecret.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}

	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entappusersecret.UserID(id))
		default:
			return nil, fmt.Errorf("invalid appusersecret field")
		}
	}
	return q, nil
}
