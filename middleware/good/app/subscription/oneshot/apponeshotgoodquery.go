package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entappgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appgoodbase"
	entapponeshot "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/appsubscriptiononeshot"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entoneshot "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/subscriptiononeshot"
)

type appOneShotGoodQueryHandler struct {
	*Handler
	oneShot     *ent.SubscriptionOneShot
	goodBase    *ent.GoodBase
	appOneShot  *ent.AppSubscriptionOneShot
	appGoodBase *ent.AppGoodBase
}

func (h *appOneShotGoodQueryHandler) _getAppOneShotGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppSubscriptionOneShot.Query()
		if h.ID != nil {
			stm.Where(entapponeshot.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entapponeshot.EntID(*h.EntID))
		}
		if h.AppGoodID != nil {
			stm.Where(entapponeshot.AppGoodID(*h.AppGoodID))
		}
		if h.appOneShot, err = stm.Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		if h.appGoodBase, err = cli.
			AppGoodBase.
			Query().
			Where(
				entappgoodbase.EntID(h.appOneShot.AppGoodID),
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

func (h *appOneShotGoodQueryHandler) getAppOneShotGood(ctx context.Context) error {
	if err := h._getAppOneShotGood(ctx, false); err != nil {
		return wlog.WrapError(err)
	}
	if h.appGoodBase == nil {
		return nil
	}
	h.AppGoodBaseReq.GoodID = &h.appGoodBase.GoodID
	return h._getOneShotGood(ctx, false)
}

func (h *appOneShotGoodQueryHandler) requireAppOneShotGood(ctx context.Context) error {
	if err := h._getAppOneShotGood(ctx, true); err != nil {
		return wlog.WrapError(err)
	}
	h.AppGoodBaseReq.GoodID = &h.appGoodBase.GoodID
	return h._getOneShotGood(ctx, true)
}

func (h *appOneShotGoodQueryHandler) _getOneShotGood(ctx context.Context, must bool) (err error) {
	if h.AppGoodBaseReq.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if h.oneShot, err = cli.SubscriptionOneShot.Query().Where(
			entoneshot.GoodID(*h.AppGoodBaseReq.GoodID),
			entoneshot.DeletedAt(0),
		).Only(ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		if h.goodBase, err = cli.GoodBase.Query().Where(
			entgoodbase.EntID(h.oneShot.GoodID),
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

func (h *appOneShotGoodQueryHandler) getOneShotGood(ctx context.Context) error { //nolint
	return h._getOneShotGood(ctx, false)
}

func (h *appOneShotGoodQueryHandler) requireOneShotGood(ctx context.Context) error {
	return h._getOneShotGood(ctx, true)
}
