package currencyfeed

import (
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated"
	entcurrencyfeed "github.com/NpoolPlatform/kunman/middleware/chain/db/ent/generated/currencyfeed"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID        *uuid.UUID
	CoinTypeID   *uuid.UUID
	FeedType     *basetypes.CurrencyFeedType
	FeedCoinName *string
	Disabled     *bool
}

func CreateSet(c *ent.CurrencyFeedCreate, req *Req) *ent.CurrencyFeedCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.FeedType != nil {
		c.SetFeedType(req.FeedType.String())
	}
	if req.FeedCoinName != nil {
		c.SetFeedCoinName(*req.FeedCoinName)
	}
	if req.Disabled != nil {
		c.SetDisabled(*req.Disabled)
	}
	return c
}

func UpdateSet(u *ent.CurrencyFeedUpdateOne, req *Req) *ent.CurrencyFeedUpdateOne {
	if req.FeedCoinName != nil {
		u.SetFeedCoinName(*req.FeedCoinName)
	}
	if req.Disabled != nil {
		u.SetDisabled(*req.Disabled)
	}
	return u
}

type Conds struct {
	EntID       *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	Disabled    *cruder.Cond
	FeedType    *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.CurrencyFeedQuery, conds *Conds) (*ent.CurrencyFeedQuery, error) {
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
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entcurrencyfeed.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid currencyfeed field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entcurrencyfeed.CoinTypeIDIn(ids...))
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
