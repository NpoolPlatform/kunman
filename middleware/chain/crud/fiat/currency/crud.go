package currency

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entfiatcurrency "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiatcurrency"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID           *uuid.UUID
	FiatID          *uuid.UUID
	FeedType        *basetypes.CurrencyFeedType
	MarketValueHigh *decimal.Decimal
	MarketValueLow  *decimal.Decimal
}

func CreateSet(c *ent.FiatCurrencyCreate, req *Req) *ent.FiatCurrencyCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
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

func UpdateSet(u *ent.FiatCurrencyUpdateOne, req *Req) *ent.FiatCurrencyUpdateOne {
	if req.MarketValueHigh != nil {
		u = u.SetMarketValueHigh(*req.MarketValueHigh)
	}
	if req.MarketValueLow != nil {
		u = u.SetMarketValueLow(*req.MarketValueLow)
	}

	return u
}

type Conds struct {
	EntID    *cruder.Cond
	FiatID   *cruder.Cond
	FiatIDs  *cruder.Cond
	FiatName *cruder.Cond
}

func SetQueryConds(q *ent.FiatCurrencyQuery, conds *Conds) (*ent.FiatCurrencyQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entfiatcurrency.EntID(id))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	if conds.FiatID != nil {
		id, ok := conds.FiatID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatid")
		}
		switch conds.FiatID.Op {
		case cruder.EQ:
			q.Where(entfiatcurrency.FiatID(id))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	if conds.FiatIDs != nil {
		ids, ok := conds.FiatIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatids")
		}
		switch conds.FiatIDs.Op {
		case cruder.IN:
			q.Where(entfiatcurrency.FiatIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	q.Where(entfiatcurrency.DeletedAt(0))
	return q, nil
}
