package coinbase

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcoinbase "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID          *uuid.UUID
	Name           *string
	Logo           *string
	Presale        *bool
	Unit           *string
	ENV            *string
	ReservedAmount *decimal.Decimal
	ForPay         *bool
	Disabled       *bool
	DeletedAt      *uint32
}

func CreateSet(c *ent.CoinBaseCreate, req *Req) *ent.CoinBaseCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	if req.Presale != nil {
		c.SetPresale(*req.Presale)
	}
	if req.Unit != nil {
		c.SetUnit(*req.Unit)
	}
	if req.ENV != nil {
		c.SetEnv(*req.ENV)
	}
	if req.ReservedAmount != nil {
		c.SetReservedAmount(*req.ReservedAmount)
	}
	if req.ForPay != nil {
		c.SetForPay(*req.ForPay)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	return c
}

func UpdateSet(u *ent.CoinBaseUpdateOne, req *Req) *ent.CoinBaseUpdateOne {
	if req.Logo != nil {
		u = u.SetLogo(*req.Logo)
	}
	if req.Presale != nil {
		u = u.SetPresale(*req.Presale)
	}
	if req.ReservedAmount != nil {
		u = u.SetReservedAmount(*req.ReservedAmount)
	}
	if req.ForPay != nil {
		u = u.SetForPay(*req.ForPay)
	}
	if req.Disabled != nil {
		u = u.SetDisabled(*req.Disabled)
	}
	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}

	return u
}

type Conds struct {
	EntID    *cruder.Cond
	EntIDs   *cruder.Cond
	Name     *cruder.Cond
	Unit     *cruder.Cond
	ENV      *cruder.Cond
	Presale  *cruder.Cond
	ForPay   *cruder.Cond
	Disabled *cruder.Cond
	Names    *cruder.Cond
}

// nolint:funlen,gocyclo
func SetQueryConds(q *ent.CoinBaseQuery, conds *Conds) (*ent.CoinBaseQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcoinbase.EntID(id))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entcoinbase.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(
				entcoinbase.Name(name),
			)
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.Unit != nil {
		unit, ok := conds.Unit.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid unit")
		}
		switch conds.Unit.Op {
		case cruder.EQ:
			q.Where(
				entcoinbase.Unit(unit),
			)
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.ENV != nil {
		env, ok := conds.ENV.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid env")
		}
		switch conds.ENV.Op {
		case cruder.EQ:
			q.Where(entcoinbase.Env(env))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.Presale != nil {
		presale, ok := conds.Presale.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid presale")
		}
		switch conds.Presale.Op {
		case cruder.EQ:
			q.Where(entcoinbase.Presale(presale))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.ForPay != nil {
		forPay, ok := conds.ForPay.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid forpay")
		}
		switch conds.ForPay.Op {
		case cruder.EQ:
			q.Where(entcoinbase.ForPay(forPay))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.Disabled != nil {
		disabled, ok := conds.Disabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid disabled")
		}
		switch conds.Disabled.Op {
		case cruder.EQ:
			q.Where(entcoinbase.Disabled(disabled))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	if conds.Names != nil {
		names, ok := conds.Names.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid names")
		}
		switch conds.Names.Op {
		case cruder.IN:
			q.Where(entcoinbase.NameIn(names...))
		default:
			return nil, fmt.Errorf("invalid coinbase field")
		}
	}
	q.Where(entcoinbase.DeletedAt(0))
	return q, nil
}
