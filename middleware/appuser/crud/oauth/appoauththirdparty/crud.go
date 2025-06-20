package auth

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entappoauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appoauththirdparty"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	ThirdPartyID *uuid.UUID
	ClientID     *string
	ClientSecret *string
	CallbackURL  *string
	Salt         *string
	DeletedAt    *uint32
}

func CreateSet(c *ent.AppOAuthThirdPartyCreate, req *Req) *ent.AppOAuthThirdPartyCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.ThirdPartyID != nil {
		c.SetThirdPartyID(*req.ThirdPartyID)
	}
	if req.ClientID != nil {
		c.SetClientID(*req.ClientID)
	}
	if req.ClientSecret != nil {
		c.SetClientSecret(*req.ClientSecret)
	}
	if req.CallbackURL != nil {
		c.SetCallbackURL(*req.CallbackURL)
	}
	if req.Salt != nil {
		c.SetSalt(*req.Salt)
	}
	return c
}

func UpdateSet(u *ent.AppOAuthThirdPartyUpdateOne, req *Req) *ent.AppOAuthThirdPartyUpdateOne {
	if req.ClientID != nil {
		u.SetClientID(*req.ClientID)
	}
	if req.ClientSecret != nil {
		u.SetClientSecret(*req.ClientSecret)
	}
	if req.CallbackURL != nil {
		u.SetCallbackURL(*req.CallbackURL)
	}
	if req.Salt != nil {
		u.SetSalt(*req.Salt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID            *cruder.Cond
	EntID         *cruder.Cond
	EntIDs        *cruder.Cond
	ClientName    *cruder.Cond
	ClientTag     *cruder.Cond
	DecryptSecret *cruder.Cond
	AppID         *cruder.Cond
	ClientID      *cruder.Cond
	ClientSecret  *cruder.Cond
	ThirdPartyID  *cruder.Cond
	ThirdPartyIDs *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.AppOAuthThirdPartyQuery, conds *Conds) (*ent.AppOAuthThirdPartyQuery, error) {
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
			q.Where(entappoauththirdparty.ID(id))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappoauththirdparty.EntID(id))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappoauththirdparty.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappoauththirdparty.AppID(id))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.ThirdPartyID != nil {
		id, ok := conds.ThirdPartyID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid thirdpartyid")
		}
		switch conds.ThirdPartyID.Op {
		case cruder.EQ:
			q.Where(entappoauththirdparty.ThirdPartyID(id))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.ThirdPartyIDs != nil {
		ids, ok := conds.ThirdPartyIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid thirdpartyids")
		}
		switch conds.ThirdPartyIDs.Op {
		case cruder.IN:
			q.Where(entappoauththirdparty.ThirdPartyIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.ClientID != nil {
		res, ok := conds.ClientID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid clientid")
		}
		switch conds.ClientID.Op {
		case cruder.EQ:
			q.Where(entappoauththirdparty.ClientID(res))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.ClientSecret != nil {
		res, ok := conds.ClientSecret.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid clientsecret")
		}
		switch conds.ClientSecret.Op {
		case cruder.EQ:
			q.Where(entappoauththirdparty.ClientSecret(res))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	q.Where(entappoauththirdparty.DeletedAt(0))
	return q, nil
}
