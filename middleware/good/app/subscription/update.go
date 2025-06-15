package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	appsubscriptioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/subscription"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*appSubscriptionGoodQueryHandler
}

func (h *updateHandler) updateAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodbasecrud.UpdateSet(
		tx.AppGoodBase.UpdateOneID(h.appGoodBase.ID),
		h.AppGoodBaseReq,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) updateAppSubscription(ctx context.Context, tx *ent.Tx) error {
	if _, err := appsubscriptioncrud.UpdateSet(
		tx.AppSubscription.UpdateOneID(h.appSubscription.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateSubscription(ctx context.Context) error {
	handler := &updateHandler{
		appSubscriptionGoodQueryHandler: &appSubscriptionGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireAppSubscriptionGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if h.USDPrice != nil && h.USDPrice.LessThan(handler.subscription.UsdPrice) {
		return wlog.Errorf("invalid usdprice")
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updateAppSubscription(_ctx, tx)
	})
}
