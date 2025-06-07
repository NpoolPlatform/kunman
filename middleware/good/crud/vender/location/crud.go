package location

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entvendorlocation "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/vendorlocation"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	Country   *string
	Province  *string
	City      *string
	Address   *string
	BrandID   *uuid.UUID
	DeletedAt *uint32
}

//nolint:dupl
func CreateSet(c *ent.VendorLocationCreate, req *Req) *ent.VendorLocationCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Country != nil {
		c.SetCountry(*req.Country)
	}
	if req.Province != nil {
		c.SetProvince(*req.Province)
	}
	if req.City != nil {
		c.SetCity(*req.City)
	}
	if req.Address != nil {
		c.SetAddress(*req.Address)
	}
	if req.BrandID != nil {
		c.SetBrandID(*req.BrandID)
	}
	return c
}

//nolint:dupl
func UpdateSet(u *ent.VendorLocationUpdateOne, req *Req) *ent.VendorLocationUpdateOne {
	if req.Country != nil {
		u.SetCountry(*req.Country)
	}
	if req.Province != nil {
		u.SetProvince(*req.Province)
	}
	if req.City != nil {
		u.SetCity(*req.City)
	}
	if req.Address != nil {
		u.SetAddress(*req.Address)
	}
	if req.BrandID != nil {
		u.SetBrandID(*req.BrandID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	EntID    *cruder.Cond
	Country  *cruder.Cond
	Province *cruder.Cond
	City     *cruder.Cond
	Address  *cruder.Cond
	BrandID  *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.VendorLocationQuery, conds *Conds) (*ent.VendorLocationQuery, error) {
	q.Where(entvendorlocation.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entvendorlocation.ID(id))
		case cruder.NEQ:
			q.Where(entvendorlocation.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid vendorlocation field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entvendorlocation.EntID(id))
		case cruder.NEQ:
			q.Where(entvendorlocation.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid vendorlocation field")
		}
	}
	if conds.Country != nil {
		country, ok := conds.Country.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid country")
		}
		switch conds.Country.Op {
		case cruder.EQ:
			q.Where(entvendorlocation.Country(country))
		default:
			return nil, wlog.Errorf("invalid vendorlocation field")
		}
	}
	if conds.Province != nil {
		province, ok := conds.Province.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid province")
		}
		switch conds.Province.Op {
		case cruder.EQ:
			q.Where(entvendorlocation.Province(province))
		default:
			return nil, wlog.Errorf("invalid vendorlocation field")
		}
	}
	if conds.City != nil {
		country, ok := conds.City.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid city")
		}
		switch conds.City.Op {
		case cruder.EQ:
			q.Where(entvendorlocation.City(country))
		default:
			return nil, wlog.Errorf("invalid vendorlocation field")
		}
	}
	if conds.Address != nil {
		country, ok := conds.Address.Val.(string)
		if !ok {
			return nil, wlog.Errorf("invalid address")
		}
		switch conds.Address.Op {
		case cruder.EQ:
			q.Where(entvendorlocation.Address(country))
		default:
			return nil, wlog.Errorf("invalid vendorlocation field")
		}
	}
	if conds.BrandID != nil {
		id, ok := conds.BrandID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid brandid")
		}
		switch conds.BrandID.Op {
		case cruder.EQ:
			q.Where(entvendorlocation.BrandID(id))
		default:
			return nil, wlog.Errorf("invalid vendorlocation field")
		}
	}
	return q, nil
}
