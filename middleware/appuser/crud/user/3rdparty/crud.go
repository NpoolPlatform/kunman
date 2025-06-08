package thirdparty

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entappuserthirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuserthirdparty"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID              *uuid.UUID
	AppID              *uuid.UUID
	UserID             *uuid.UUID
	ThirdPartyID       *uuid.UUID
	ThirdPartyUserID   *string
	ThirdPartyUsername *string
	ThirdPartyAvatar   *string
}

func CreateSet(c *ent.AppUserThirdPartyCreate, req *Req) *ent.AppUserThirdPartyCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.ThirdPartyUserID != nil {
		c.SetThirdPartyUserID(*req.ThirdPartyUserID)
	}
	if req.ThirdPartyID != nil {
		c.SetThirdPartyID(*req.ThirdPartyID)
	}
	if req.ThirdPartyUsername != nil {
		c.SetThirdPartyUsername(*req.ThirdPartyUsername)
	}
	if req.ThirdPartyAvatar != nil {
		c.SetThirdPartyAvatar(*req.ThirdPartyAvatar)
	}
	return c
}

func UpdateSet(u *ent.AppUserThirdPartyUpdateOne, req *Req) *ent.AppUserThirdPartyUpdateOne {
	if req.UserID != nil {
		u.SetUserID(*req.UserID)
	}
	if req.ThirdPartyUsername != nil {
		u.SetThirdPartyUsername(*req.ThirdPartyUsername)
	}
	if req.ThirdPartyAvatar != nil {
		u.SetThirdPartyAvatar(*req.ThirdPartyAvatar)
	}
	return u
}

type Conds struct {
	EntID            *cruder.Cond
	AppID            *cruder.Cond
	UserID           *cruder.Cond
	ThirdPartyUserID *cruder.Cond
	ThirdPartyID     *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.AppUserThirdPartyQuery, conds *Conds) (*ent.AppUserThirdPartyQuery, error) {
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
			q.Where(entappuserthirdparty.EntID(id))

		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappuserthirdparty.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entappuserthirdparty.UserID(id))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.ThirdPartyUserID != nil {
		id, ok := conds.ThirdPartyUserID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid 3rd userid")
		}
		switch conds.ThirdPartyUserID.Op {
		case cruder.EQ:
			q.Where(entappuserthirdparty.ThirdPartyUserID(id))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.ThirdPartyID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid 3rd id")
		}
		switch conds.ThirdPartyID.Op {
		case cruder.EQ:
			q.Where(entappuserthirdparty.ThirdPartyID(id))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	return q, nil
}
