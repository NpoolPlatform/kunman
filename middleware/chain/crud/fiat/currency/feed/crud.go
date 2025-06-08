package currencyfeed

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcurrencyfeed "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/fiatcurrencyfeed"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	FiatID       *uuid.UUID
	FeedType     *basetypes.CurrencyFeedType
	FeedFiatName *string
	Disabled     *bool
}

func CreateSet(c *ent.FiatCurrencyFeedCreate, req *Req) *ent.FiatCurrencyFeedCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.FiatID != nil {
		c.SetFiatID(*req.FiatID)
	}
	if req.FeedType != nil {
		c.SetFeedType(req.FeedType.String())
	}
	if req.FeedFiatName != nil {
		c.SetFeedFiatName(*req.FeedFiatName)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	return c
}

func UpdateSet(u *ent.FiatCurrencyFeedUpdateOne, req *Req) *ent.FiatCurrencyFeedUpdateOne {
	if req.FeedFiatName != nil {
		u.SetFeedFiatName(*req.FeedFiatName)
	}
	if req.Disabled != nil {
		u.SetDisabled(*req.Disabled)
	}
	return u
}

type Conds struct {
	EntID    *cruder.Cond
	FiatID   *cruder.Cond
	FiatIDs  *cruder.Cond
	Disabled *cruder.Cond
	FeedType *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.FiatCurrencyFeedQuery, conds *Conds) (*ent.FiatCurrencyFeedQuery, error) {
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.EntID(id))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.FiatID != nil {
		id, ok := conds.FiatID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatid")
		}
		switch conds.FiatID.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.FiatID(id))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.FiatIDs != nil {
		ids, ok := conds.FiatIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fiatids")
		}
		switch conds.FiatIDs.Op {
		case cruder.IN:
			q.Where(entcurrencyfeed.FiatIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.Disabled != nil {
		disabled, ok := conds.Disabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid disabled")
		}
		switch conds.Disabled.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.Disabled(disabled))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.FeedType != nil {
		feedType, ok := conds.FeedType.Val.(basetypes.CurrencyFeedType)
		if !ok {
			return nil, fmt.Errorf("invalid feedtype")
		}
		switch conds.FeedType.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.FeedType(feedType.String()))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	q.Where(entcurrencyfeed.DeletedAt(0))
	return q, nil
}
