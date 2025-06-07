package subscription

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/subscription"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*subscriptionGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.goodBase == nil {
		return wlog.Errorf("invalid goodbase")
	}
	if _, err := goodbasecrud.UpdateSet(
		tx.GoodBase.UpdateOneID(h.goodBase.ID),
		&goodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteSubscription(ctx context.Context, tx *ent.Tx) error {
	if h.subscription == nil {
		return wlog.Errorf("invalid subscription")
	}
	if _, err := subscriptioncrud.UpdateSet(
		tx.Subscription.UpdateOneID(h.subscription.ID),
		&subscriptioncrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteSubscription(ctx context.Context) error {
	handler := &deleteHandler{
		subscriptionGoodQueryHandler: &subscriptionGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getSubscriptionGood(ctx); err != nil {
		return err
	}
	if handler.subscription == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.deleteSubscription(_ctx, tx)
	})
}
