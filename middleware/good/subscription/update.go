package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/subscription"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/good/goodbase"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type updateHandler struct {
	*subscriptionGoodQueryHandler
	sqlGoodBase string
}

func (h *updateHandler) constructGoodBaseSQL(ctx context.Context) (err error) {
	handler, _ := goodbase1.NewHandler(
		ctx,
		goodbase1.WithID(&h.goodBase.ID, true),
	)
	handler.Req = *h.GoodBaseReq
	h.sqlGoodBase, err = handler.ConstructUpdateSQL()
	if err != nil {
		if err == cruder.ErrUpdateNothing {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) updateGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.sqlGoodBase == "" {
		return nil
	}
	rc, err := tx.ExecContext(ctx, h.sqlGoodBase)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail update subscription: %v", err)
	}
	return nil
}

func (h *updateHandler) updateSubscription(ctx context.Context, tx *ent.Tx) error {
	if _, err := subscriptioncrud.UpdateSet(
		tx.Subscription.UpdateOneID(*h.ID),
		&subscriptioncrud.Req{
			DurationDisplayType: h.DurationDisplayType,
			DurationUnits:       h.DurationUnits,
			DurationQuota:       h.DurationQuota,
			DailyBonusQuota:     h.DailyBonusQuota,
			USDPrice:            h.USDPrice,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateSubscription(ctx context.Context) error {
	handler := &updateHandler{
		subscriptionGoodQueryHandler: &subscriptionGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireSubscriptionGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	h.ID = &handler.subscription.ID
	if err := handler.constructGoodBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updateSubscription(_ctx, tx)
	})
}
