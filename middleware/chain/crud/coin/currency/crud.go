package currency

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcurrency "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/currency"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID           *uuid.UUID
	CoinTypeID      *uuid.UUID
	FeedType        *basetypes.CurrencyFeedType
	MarketValueHigh *decimal.Decimal
	MarketValueLow  *decimal.Decimal
}

func CreateSet(c *ent.CurrencyCreate, req *Req) *ent.CurrencyCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
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

func UpdateSet(u *ent.CurrencyUpdateOne, req *Req) *ent.CurrencyUpdateOne {
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
	FeedType    *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
}

func SetQueryConds(q *ent.CurrencyQuery, conds *Conds) (*ent.CurrencyQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcurrency.EntID(id))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.FeedType != nil {
		feedType, ok := conds.FeedType.Val.(basetypes.CurrencyFeedType)
		if !ok {
			return nil, fmt.Errorf("invalid feedtype")
		}
		switch conds.FeedType.Op {
		case cruder.EQ:
			q.Where(entcurrency.FeedType(feedType.String()))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcurrency.CoinTypeID(id))
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
			q.Where(entcurrency.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid currency field")
		}
	}
	q.Where(entcurrency.DeletedAt(0))
	return q, nil
}
