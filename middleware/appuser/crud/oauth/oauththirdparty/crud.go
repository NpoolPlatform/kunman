//nolint:nolintlint,dupl
package auth

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entoauththirdparty "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/oauththirdparty"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID          *uuid.UUID
	ClientName     *basetypes.SignMethod
	ClientTag      *string
	ClientLogoURL  *string
	ClientOAuthURL *string
	ResponseType   *string
	Scope          *string
	DeletedAt      *uint32
}

func CreateSet(c *ent.OAuthThirdPartyCreate, req *Req) *ent.OAuthThirdPartyCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.ClientName != nil {
		c.SetClientName(req.ClientName.String())
	}
	if req.ClientTag != nil {
		c.SetClientTag(*req.ClientTag)
	}
	if req.ClientLogoURL != nil {
		c.SetClientLogoURL(*req.ClientLogoURL)
	}
	if req.ClientOAuthURL != nil {
		c.SetClientOauthURL(*req.ClientOAuthURL)
	}
	if req.ResponseType != nil {
		c.SetResponseType(*req.ResponseType)
	}
	if req.Scope != nil {
		c.SetScope(*req.Scope)
	}
	return c
}

func UpdateSet(u *ent.OAuthThirdPartyUpdateOne, req *Req) *ent.OAuthThirdPartyUpdateOne {
	if req.ClientName != nil {
		u.SetClientName(req.ClientName.String())
	}
	if req.ClientTag != nil {
		u.SetClientTag(*req.ClientTag)
	}
	if req.ClientLogoURL != nil {
		u.SetClientLogoURL(*req.ClientLogoURL)
	}
	if req.ClientOAuthURL != nil {
		u.SetClientOauthURL(*req.ClientOAuthURL)
	}
	if req.ResponseType != nil {
		u.SetResponseType(*req.ResponseType)
	}
	if req.Scope != nil {
		u.SetScope(*req.Scope)
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
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.OAuthThirdPartyQuery, conds *Conds) (*ent.OAuthThirdPartyQuery, error) {
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
			q.Where(entoauththirdparty.ID(id))
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
			q.Where(entoauththirdparty.EntID(id))
		case cruder.NEQ:
			q.Where(entoauththirdparty.EntIDNEQ(id))
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
			q.Where(entoauththirdparty.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.ClientName != nil {
		clientName, ok := conds.ClientName.Val.(basetypes.SignMethod)
		if !ok {
			return nil, fmt.Errorf("invalid clientname")
		}
		switch conds.ClientName.Op {
		case cruder.EQ:
			q.Where(entoauththirdparty.ClientName(clientName.String()))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	if conds.ClientTag != nil {
		res, ok := conds.ClientTag.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid clienttag")
		}
		switch conds.ClientTag.Op {
		case cruder.EQ:
			q.Where(entoauththirdparty.ClientTag(res))
		default:
			return nil, fmt.Errorf("invalid auth field")
		}
	}
	q.Where(entoauththirdparty.DeletedAt(0))
	return q, nil
}
