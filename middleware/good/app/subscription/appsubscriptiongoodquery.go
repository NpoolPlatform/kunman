package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entappsubscription "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsubscription"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entsubscription "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/subscription"
)

type appSubscriptionGoodQueryHandler struct {
	*Handler
	subscription    *ent.Subscription
	goodBase        *ent.GoodBase
	appSubscription *ent.AppSubscription
	appGoodBase     *ent.AppGoodBase
}

func (h *appSubscriptionGoodQueryHandler) _getAppSubscriptionGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppSubscription.Query()
		if h.ID != nil {
			stm.Where(entappsubscription.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entappsubscription.EntID(*h.EntID))
		}
		if h.AppGoodID != nil {
			stm.Where(entappsubscription.AppGoodID(*h.AppGoodID))
		}
		if h.appSubscription, err = stm.Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		if h.appGoodBase, err = cli.
			AppGoodBase.
			Query().
			Where(
				entappgoodbase.EntID(h.appSubscription.AppGoodID),
				entappgoodbase.DeletedAt(0),
			).Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *appSubscriptionGoodQueryHandler) getAppSubscriptionGood(ctx context.Context) error {
	if err := h._getAppSubscriptionGood(ctx, false); err != nil {
		return wlog.WrapError(err)
	}
	if h.appGoodBase == nil {
		return nil
	}
	h.AppGoodBaseReq.GoodID = &h.appGoodBase.GoodID
	return h._getSubscriptionGood(ctx, false)
}

func (h *appSubscriptionGoodQueryHandler) requireAppSubscriptionGood(ctx context.Context) error {
	if err := h._getAppSubscriptionGood(ctx, true); err != nil {
		return wlog.WrapError(err)
	}
	h.AppGoodBaseReq.GoodID = &h.appGoodBase.GoodID
	return h._getSubscriptionGood(ctx, true)
}

func (h *appSubscriptionGoodQueryHandler) _getSubscriptionGood(ctx context.Context, must bool) (err error) {
	if h.AppGoodBaseReq.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if h.subscription, err = cli.Subscription.Query().Where(
			entsubscription.GoodID(*h.AppGoodBaseReq.GoodID),
			entsubscription.DeletedAt(0),
		).Only(ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		if h.goodBase, err = cli.GoodBase.Query().Where(
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

func (h *appSubscriptionGoodQueryHandler) getSubscriptionGood(ctx context.Context) error { //nolint
	return h._getSubscriptionGood(ctx, false)
}

func (h *appSubscriptionGoodQueryHandler) requireSubscriptionGood(ctx context.Context) error {
	return h._getSubscriptionGood(ctx, true)
}
