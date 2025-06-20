//nolint:dupl
package country

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated"
	entcountry "github.com/NpoolPlatform/kunman/middleware/g11n/db/ent/generated/country"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	Country   *string
	Flag      *string
	Code      *string
	Short     *string
	DeletedAt *uint32
}

func CreateSet(c *ent.CountryCreate, req *Req) *ent.CountryCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Country != nil {
		c.SetCountry(*req.Country)
	}
	if req.Flag != nil {
		c.SetFlag(*req.Flag)
	}
	if req.Code != nil {
		c.SetCode(*req.Code)
	}
	if req.Short != nil {
		c.SetShort(*req.Short)
	}
	return c
}

func UpdateSet(u *ent.CountryUpdateOne, req *Req) *ent.CountryUpdateOne {
	if req.Country != nil {
		u.SetCountry(*req.Country)
	}
	if req.Flag != nil {
		u.SetFlag(*req.Flag)
	}
	if req.Code != nil {
		u.SetCode(*req.Code)
	}
	if req.Short != nil {
		u.SetShort(*req.Short)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	EntID     *cruder.Cond
	EntIDs    *cruder.Cond
	Country   *cruder.Cond
	Code      *cruder.Cond
	Short     *cruder.Cond
	Countries *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.CountryQuery, conds *Conds) (*ent.CountryQuery, error) {
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
				entcountry.ID(id),
				entcountry.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entcountry.IDNEQ(id),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid1")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(
				entcountry.EntID(id),
				entcountry.DeletedAt(0),
			)
		case cruder.NEQ:
			q.Where(
				entcountry.EntIDNEQ(id),
				entcountry.DeletedAt(0),
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
				entcountry.EntIDIn(ids...),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid entids field")
		}
	}
	if conds.Country != nil {
		country, ok := conds.Country.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid country")
		}
		switch conds.Country.Op {
		case cruder.EQ:
			q.Where(
				entcountry.Country(country),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid country field")
		}
	}
	if conds.Code != nil {
		code, ok := conds.Code.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid code")
		}
		switch conds.Code.Op {
		case cruder.EQ:
			q.Where(
				entcountry.Code(code),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid code field")
		}
	}
	if conds.Short != nil {
		short, ok := conds.Short.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid short")
		}
		switch conds.Short.Op {
		case cruder.EQ:
			q.Where(
				entcountry.Short(short),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid short field")
		}
	}
	if conds.Countries != nil {
		countries, ok := conds.Countries.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid countries")
		}
		switch conds.Countries.Op {
		case cruder.IN:
			q.Where(
				entcountry.CountryIn(countries...),
				entcountry.DeletedAt(0),
			)
		default:
			return nil, fmt.Errorf("invalid countries field")
		}
	}
	return q, nil
}
