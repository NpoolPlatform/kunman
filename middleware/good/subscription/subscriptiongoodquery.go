package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entsubscription "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/subscription"
)

type subscriptionGoodQueryHandler struct {
	*Handler
	subscription *ent.Subscription
	goodBase     *ent.GoodBase
}

func (h *subscriptionGoodQueryHandler) _getSubscriptionGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.Subscription.Query()
		if h.ID != nil {
			stm.Where(entsubscription.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entsubscription.EntID(*h.EntID))
		}
		if h.GoodID != nil {
			stm.Where(entsubscription.GoodID(*h.GoodID))
		}
		if h.subscription, err = stm.Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		if h.goodBase, err = cli.
			GoodBase.
			Query().
			Where(
				entgoodbase.EntID(h.subscription.GoodID),
				entgoodbase.DeletedAt(0),
			).Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *subscriptionGoodQueryHandler) getSubscriptionGood(ctx context.Context) error {
	return h._getSubscriptionGood(ctx, false)
}

func (h *subscriptionGoodQueryHandler) requireSubscriptionGood(ctx context.Context) error {
	return h._getSubscriptionGood(ctx, true)
}
