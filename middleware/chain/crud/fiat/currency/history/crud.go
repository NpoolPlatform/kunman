package currencyhistory

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entfiatcurrencyhis "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiatcurrencyhistory"
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

func CreateSet(c *ent.FiatCurrencyHistoryCreate, req *Req) *ent.FiatCurrencyHistoryCreate {
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

func UpdateSet(u *ent.FiatCurrencyHistoryUpdateOne, req *Req) *ent.FiatCurrencyHistoryUpdateOne {
	if req.MarketValueHigh != nil {
		u = u.SetMarketValueHigh(*req.MarketValueHigh)
	}
	if req.MarketValueLow != nil {
		u = u.SetMarketValueLow(*req.MarketValueLow)
	}

	return u
}

type Conds struct {
	EntID   *cruder.Cond
	FiatID  *cruder.Cond
	FiatIDs *cruder.Cond
	StartAt *cruder.Cond
	EndAt   *cruder.Cond
}

// nolint:funlen,gocyclo
func SetQueryConds(q *ent.FiatCurrencyHistoryQuery, conds *Conds) (*ent.FiatCurrencyHistoryQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entfiatcurrencyhis.EntID(id))
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
			q.Where(entfiatcurrencyhis.FiatID(id))
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
			q.Where(entfiatcurrencyhis.FiatIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	if conds.StartAt != nil {
		at, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid startat")
		}
		switch conds.StartAt.Op {
		case cruder.EQ:
			q.Where(entfiatcurrencyhis.CreatedAtGTE(at))
		case cruder.GTE:
			q.Where(entfiatcurrencyhis.CreatedAtGTE(at))
		case cruder.LTE:
			q.Where(entfiatcurrencyhis.CreatedAtLTE(at))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	if conds.EndAt != nil {
		at, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid endat")
		}
		switch conds.EndAt.Op {
		case cruder.EQ:
			q.Where(entfiatcurrencyhis.CreatedAtLTE(at))
		case cruder.GTE:
			q.Where(entfiatcurrencyhis.CreatedAtGTE(at))
		case cruder.LTE:
			q.Where(entfiatcurrencyhis.CreatedAtLTE(at))
		default:
			return nil, fmt.Errorf("invalid fiatcurrency field")
		}
	}
	q.Where(entfiatcurrencyhis.DeletedAt(0))
	return q, nil
}
