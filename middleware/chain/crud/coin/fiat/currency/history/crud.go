package currencyhistory

import (
	"fmt"

	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcurrencyhis "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/coinfiatcurrencyhistory"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	CoinTypeID      *uuid.UUID
	FiatID          *uuid.UUID
	FeedType        *basetypes.CurrencyFeedType
	MarketValueHigh *decimal.Decimal
	MarketValueLow  *decimal.Decimal
}

func CreateSet(c *ent.CoinFiatCurrencyHistoryCreate, req *Req) *ent.CoinFiatCurrencyHistoryCreate {
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FiatID != nil {
		c.SetFiatID(*req.FiatID)
	}
	if req.FeedType != nil {
		c.SetFeedType(req.FeedType.String())
	}
	if req.MarketValueHigh != nil {
		c.SetMarketValueHigh(*req.MarketValueHigh)
	}
	if req.MarketValueLow != nil {
		c.SetMarketValueLow(*req.MarketValueLow)
	}
	return c
}

func UpdateSet(u *ent.CoinFiatCurrencyHistoryUpdateOne, req *Req) *ent.CoinFiatCurrencyHistoryUpdateOne {
	if req.MarketValueHigh != nil {
		u = u.SetMarketValueHigh(*req.MarketValueHigh)
	}
	if req.MarketValueLow != nil {
		u = u.SetMarketValueLow(*req.MarketValueLow)
	}

	return u
}

type Conds struct {
	EntID       *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	StartAt     *cruder.Cond
	EndAt       *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.CoinFiatCurrencyHistoryQuery, conds *Conds) (*ent.CoinFiatCurrencyHistoryQuery, error) {
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcurrencyhis.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entcurrencyhis.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.StartAt != nil {
		at, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid startat")
		}
		switch conds.StartAt.Op {
		case cruder.EQ:
			q.Where(entcurrencyhis.CreatedAtGTE(at))
		case cruder.LTE:
			q.Where(entcurrencyhis.CreatedAtLTE(at))
		case cruder.GTE:
			q.Where(entcurrencyhis.CreatedAtGTE(at))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.EndAt != nil {
		at, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid endat")
		}
		switch conds.EndAt.Op {
		case cruder.EQ:
			q.Where(entcurrencyhis.CreatedAtLTE(at))
		case cruder.GTE:
			q.Where(entcurrencyhis.CreatedAtGTE(at))
		case cruder.LTE:
			q.Where(entcurrencyhis.CreatedAtLTE(at))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	q.Where(entcurrencyhis.DeletedAt(0))
	return q, nil
}
