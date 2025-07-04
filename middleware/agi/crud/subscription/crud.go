package subscription

import (
	"github.com/NpoolPlatform/kunman/framework/wlog"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated"
	entsubscription "github.com/NpoolPlatform/kunman/middleware/agi/db/ent/generated/subscription"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID                 *uint32
	EntID              *uuid.UUID
	AppID              *uuid.UUID
	UserID             *uuid.UUID
	AppGoodID          *uuid.UUID
	NextExtendAt       *uint32
	PermanentQuota     *uint32
	ConsumedQuota      *uint32
	PayWithCoinBalance *bool
	SubscriptionID     *string
	FiatPaymentChannel *ordertypes.FiatPaymentChannel
	LastPaymentAt      *uint32
	LastUpdatedEventID *string
	ActivatedAt        *uint32
	ActivatedEventID   *string
	DeletedAt          *uint32
}

func CreateSet(c *ent.SubscriptionCreate, req *Req) *ent.SubscriptionCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.NextExtendAt != nil {
		c.SetNextExtendAt(*req.NextExtendAt)
	}
	if req.PermanentQuota != nil {
		c.SetPermanentQuota(*req.PermanentQuota)
	}
	if req.ConsumedQuota != nil {
		c.SetConsumedQuota(*req.ConsumedQuota)
	}
	if req.PayWithCoinBalance != nil {
		c.SetPayWithCoinBalance(*req.PayWithCoinBalance)
	}
	if req.SubscriptionID != nil {
		c.SetSubscriptionID(*req.SubscriptionID)
	}
	if req.FiatPaymentChannel != nil {
		c.SetFiatPaymentChannel(req.FiatPaymentChannel.String())
	}
	if req.LastPaymentAt != nil {
		c.SetLastPaymentAt(*req.LastPaymentAt)
	}
	if req.LastUpdatedEventID != nil {
		c.SetLastUpdatedEventID(*req.LastUpdatedEventID)
	}
	if req.ActivatedAt != nil {
		c.SetActivatedAt(*req.ActivatedAt)
	}
	if req.ActivatedEventID != nil {
		c.SetActivatedEventID(*req.ActivatedEventID)
	}
	return c
}

func UpdateSet(u *ent.SubscriptionUpdateOne, req *Req) *ent.SubscriptionUpdateOne {
	if req.NextExtendAt != nil {
		u.SetNextExtendAt(*req.NextExtendAt)
	}
	if req.PermanentQuota != nil {
		u.SetPermanentQuota(*req.PermanentQuota)
	}
	if req.ConsumedQuota != nil {
		u.SetConsumedQuota(*req.ConsumedQuota)
	}
	if req.PayWithCoinBalance != nil {
		u.SetPayWithCoinBalance(*req.PayWithCoinBalance)
	}
	if req.SubscriptionID != nil {
		u.SetSubscriptionID(*req.SubscriptionID)
	}
	if req.FiatPaymentChannel != nil {
		u.SetFiatPaymentChannel(req.FiatPaymentChannel.String())
	}
	if req.LastPaymentAt != nil {
		u.SetLastPaymentAt(*req.LastPaymentAt)
	}
	if req.LastUpdatedEventID != nil {
		u.SetLastUpdatedEventID(*req.LastUpdatedEventID)
	}
	if req.ActivatedAt != nil {
		u.SetActivatedAt(*req.ActivatedAt)
	}
	if req.ActivatedEventID != nil {
		u.SetActivatedEventID(*req.ActivatedEventID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppID      *cruder.Cond
	AppIDs     *cruder.Cond
	UserID     *cruder.Cond
	UserIDs    *cruder.Cond
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.SubscriptionQuery, conds *Conds) (*ent.SubscriptionQuery, error) {
	q.Where(entsubscription.DeletedAt(0))
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
			q.Where(entsubscription.ID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entsubscription.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entsubscription.EntID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entsubscription.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entsubscription.AppID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entsubscription.AppIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entsubscription.UserID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.UserIDs != nil {
		ids, ok := conds.UserIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid userids")
		}
		switch conds.UserIDs.Op {
		case cruder.IN:
			q.Where(entsubscription.UserIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entsubscription.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entsubscription.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid subscription field")
		}
	}

	return q, nil
}
