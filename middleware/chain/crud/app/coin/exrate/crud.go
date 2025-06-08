package exrate

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entexrate "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/exchangerate"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID            *uint32
	EntID         *uuid.UUID
	AppID         *uuid.UUID
	CoinTypeID    *uuid.UUID
	MarketValue   *decimal.Decimal
	SettlePercent *uint32
	SettleTips    []string
	Setter        *uuid.UUID
	DeletedAt     *uint32
}

func CreateSet(c *ent.ExchangeRateCreate, req *Req) *ent.ExchangeRateCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.MarketValue != nil {
		c.SetMarketValue(*req.MarketValue)
		settleValue := *req.MarketValue
		if req.SettlePercent != nil {
			settleValue = settleValue.Mul(decimal.NewFromInt(int64(*req.SettlePercent)))
			settleValue = settleValue.Div(decimal.NewFromInt(100)) //nolint
		}
		c.SetSettleValue(settleValue)
	}

	if req.SettlePercent != nil {
		c.SetSettlePercent(*req.SettlePercent)
	}
	if len(req.SettleTips) > 0 {
		c.SetSettleTips(req.SettleTips)
	}
	if req.Setter != nil {
		c.SetSetter(*req.Setter)
	}
	return c
}

func UpdateSet(u *ent.ExchangeRateUpdateOne, req *Req) *ent.ExchangeRateUpdateOne {
	settlePercent, _ := u.Mutation().SettlePercent()
	marketValue, _ := u.Mutation().MarketValue()

	if req.MarketValue != nil {
		u = u.SetMarketValue(*req.MarketValue)
		marketValue = *req.MarketValue
	}
	if req.SettlePercent != nil {
		u = u.SetSettlePercent(*req.SettlePercent)
		settlePercent = *req.SettlePercent
	}

	if len(req.SettleTips) > 0 {
		u = u.SetSettleTips(req.SettleTips)
	}

	settleValue, _ := u.Mutation().SettleValue() //nolint
	settleValue = marketValue.Mul(decimal.NewFromInt(int64(settlePercent)))
	settleValue = settleValue.Div(decimal.NewFromInt(100)) //nolint
	u = u.SetSettleValue(settleValue)

	if req.DeletedAt != nil {
		u = u.SetDeletedAt(*req.DeletedAt)
	}

	return u
}

type Conds struct {
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	CoinTypeID *cruder.Cond
}

func SetQueryConds(q *ent.ExchangeRateQuery, conds *Conds) (*ent.ExchangeRateQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entexrate.EntID(id))
		default:
			return nil, fmt.Errorf("invalid exrate field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entexrate.AppID(id))
		default:
			return nil, fmt.Errorf("invalid exrate field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entexrate.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid exrate field")
		}
	}
	q.Where(entexrate.DeletedAt(0))
	return q, nil
}
