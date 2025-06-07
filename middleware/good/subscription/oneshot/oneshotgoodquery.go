package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	entgoodbase "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/goodbase"
	entoneshot "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated/subscriptiononeshot"
)

type oneShotGoodQueryHandler struct {
	*Handler
	oneShot  *ent.SubscriptionOneShot
	goodBase *ent.GoodBase
}

func (h *oneShotGoodQueryHandler) _getOneShotGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.SubscriptionOneShot.Query()
		if h.ID != nil {
			stm.Where(entoneshot.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entoneshot.EntID(*h.EntID))
		}
		if h.GoodID != nil {
			stm.Where(entoneshot.GoodID(*h.GoodID))
		}
		if h.oneShot, err = stm.Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return wlog.WrapError(err)
		}
		if h.goodBase, err = cli.
			GoodBase.
			Query().
			Where(
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

func (h *oneShotGoodQueryHandler) getOneShotGood(ctx context.Context) error {
	return h._getOneShotGood(ctx, false)
}

func (h *oneShotGoodQueryHandler) requireOneShotGood(ctx context.Context) error {
	return h._getOneShotGood(ctx, true)
}
