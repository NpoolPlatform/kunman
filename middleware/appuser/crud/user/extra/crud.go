package extra

import (
	"context"
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated"
	entappuserextra "github.com/NpoolPlatform/kunman/middleware/appuser/db/ent/generated/appuserextra"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID         *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	FirstName     *string
	LastName      *string
	Organization  *string
	IDNumber      *string
	PostalCode    *string
	Age           *uint32
	Birthday      *uint32
	Avatar        *string
	Username      *string
	Gender        *string
	AddressFields []string
}

func CreateSet(c *ent.AppUserExtraCreate, req *Req) *ent.AppUserExtraCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.FirstName != nil {
		c.SetFirstName(*req.FirstName)
	}
	if req.LastName != nil {
		c.SetLastName(*req.LastName)
	}
	if req.Organization != nil {
		c.SetOrganization(*req.Organization)
	}
	if req.IDNumber != nil {
		c.SetIDNumber(*req.IDNumber)
	}
	if req.PostalCode != nil {
		c.SetPostalCode(*req.PostalCode)
	}
	if req.Age != nil {
		c.SetAge(*req.Age)
	}
	if req.Birthday != nil {
		c.SetBirthday(*req.Birthday)
	}
	if req.Avatar != nil {
		c.SetAvatar(*req.Avatar)
	}
	if req.Username != nil {
		c.SetUsername(*req.Username)
	}
	if req.Gender != nil {
		c.SetGender(*req.Gender)
	}
	if req.AddressFields != nil {
		c.SetAddressFields(req.AddressFields)
	}
	return c
}

func UpdateSet(ctx context.Context, u *ent.AppUserExtraUpdateOne, req *Req) (*ent.AppUserExtraUpdateOne, error) {
	if req.Username != nil {
		u.SetUsername(*req.Username)
	}
	if req.FirstName != nil {
		u.SetFirstName(*req.FirstName)
	}
	if req.LastName != nil {
		u.SetLastName(*req.LastName)
	}
	if req.AddressFields != nil {
		u.SetAddressFields(req.AddressFields)
	}
	if req.Gender != nil {
		u.SetGender(*req.Gender)
	}
	if req.PostalCode != nil {
		u.SetPostalCode(*req.PostalCode)
	}
	if req.IDNumber != nil {
		u.SetIDNumber(*req.IDNumber)
	}
	if req.Organization != nil {
		u.SetOrganization(*req.Organization)
	}
	if req.Age != nil {
		u.SetAge(*req.Age)
	}
	if req.Birthday != nil {
		u.SetBirthday(*req.Birthday)
	}
	if req.Avatar != nil {
		u.SetAvatar(*req.Avatar)
	}
	if req.LastName != nil {
		u.SetLastName(*req.LastName)
	}

	return u, nil
}

type Conds struct {
	EntID    *cruder.Cond
	AppID    *cruder.Cond
	UserID   *cruder.Cond
	IDNumber *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.AppUserExtraQuery, conds *Conds) (*ent.AppUserExtraQuery, error) {
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
			q.Where(entappuserextra.EntID(id))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappuserextra.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entappuserextra.UserID(id))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.IDNumber != nil {
		idNumber, ok := conds.IDNumber.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid id number")
		}
		switch conds.IDNumber.Op {
		case cruder.EQ:
			q.Where(entappuserextra.IDNumber(idNumber))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	return q, nil
}
