package coinextra

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcoinextra "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinextra"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID      *uuid.UUID
	CoinTypeID *uuid.UUID
	HomePage   *string
	Specs      *string
	StableUSD  *bool
	DeletedAt  *uint32
}

func CreateSet(c *ent.CoinExtraCreate, req *Req) *ent.CoinExtraCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.HomePage != nil {
		c.SetHomePage(*req.HomePage)
	}
	if req.Specs != nil {
		c.SetSpecs(*req.Specs)
	}
	if req.StableUSD != nil {
		c.SetStableUsd(*req.StableUSD)
	}
	return c
}

func UpdateSet(u *ent.CoinExtraUpdateOne, req *Req) *ent.CoinExtraUpdateOne {
	if req.HomePage != nil {
		u = u.SetHomePage(*req.HomePage)
	}
	if req.Specs != nil {
		u = u.SetSpecs(*req.Specs)
	}
	if req.StableUSD != nil {
		u = u.SetStableUsd(*req.StableUSD)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID      *cruder.Cond
	CoinTypeID *cruder.Cond
	StableUSD  *cruder.Cond
}

func SetQueryConds(q *ent.CoinExtraQuery, conds *Conds) (*ent.CoinExtraQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(
				entcoinextra.EntID(id),
			)
		default:
			return nil, fmt.Errorf("invalid coinextra field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(
				entcoinextra.CoinTypeID(id),
			)
		default:
			return nil, fmt.Errorf("invalid coinextra field")
		}
	}
	if conds.StableUSD != nil {
		stable, ok := conds.StableUSD.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid stableusd")
		}
		switch conds.StableUSD.Op {
		case cruder.EQ:
			q.Where(
				entcoinextra.StableUsd(stable),
			)
		default:
			return nil, fmt.Errorf("invalid coinextra field")
		}
	}
	return q, nil
}
