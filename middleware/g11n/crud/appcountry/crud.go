//nolint:nolintlint,dupl
package appcountry

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
	entappcountry "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/appcountry"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	CountryID *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.AppCountryCreate, req *Req) *ent.AppCountryCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CountryID != nil {
		c.SetCountryID(*req.CountryID)
	}
	return c
}

func UpdateSet(u *ent.AppCountryUpdateOne, req *Req) *ent.AppCountryUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppID      *cruder.Cond
	CountryID  *cruder.Cond
	AppIDs     *cruder.Cond
	CountryIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppCountryQuery, conds *Conds) (*ent.AppCountryQuery, error) {
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
			q.Where(
				entappcountry.ID(id),
				entappcountry.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entappcountry.IDNEQ(id),
				entappcountry.DeletedAt(0),
			)
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
			q.Where(
				entappcountry.EntID(id),
				entappcountry.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entappcountry.EntIDNEQ(id),
				entappcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(
				entappcountry.EntIDIn(ids...),
				entappcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(
				entappcountry.AppID(id),
				entappcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(
				entappcountry.AppIDIn(ids...),
				entappcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid appids field")
		}
	}
	if conds.CountryID != nil {
		id, ok := conds.CountryID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langid")
		}
		switch conds.CountryID.Op {
		case cruder.EQ:
			q.Where(
				entappcountry.CountryID(id),
				entappcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid langid field")
		}
	}
	if conds.CountryIDs != nil {
		ids, ok := conds.CountryIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid langids")
		}
		switch conds.CountryIDs.Op {
		case cruder.IN:
			q.Where(
				entappcountry.CountryIDIn(ids...),
				entappcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid langids field")
		}
	}
	return q, nil
}
