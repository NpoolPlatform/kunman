package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	appsubscriptioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/subscription"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/google/uuid"
)

type createHandler struct {
	*appSubscriptionGoodQueryHandler
}

func (h *createHandler) createAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodbasecrud.CreateSet(
		tx.AppGoodBase.Create(),
		h.AppGoodBaseReq,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *createHandler) createAppSubscription(ctx context.Context, tx *ent.Tx) error {
	if _, err := appsubscriptioncrud.CreateSet(
		tx.AppSubscription.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateSubscription(ctx context.Context) error {
	handler := &createHandler{
		appSubscriptionGoodQueryHandler: &appSubscriptionGoodQueryHandler{
			Handler: h,
		},
	}
	if err := handler.requireSubscriptionGood(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if h.USDPrice == nil {
		h.USDPrice = &handler.subscription.UsdPrice
	} else if h.USDPrice.LessThan(handler.subscription.UsdPrice) {
		return wlog.Errorf("invalid usdprice")
	}
	if h.AppGoodID == nil {
		id := uuid.New()
		h.AppGoodID = &id
		h.AppGoodBaseReq.EntID = &id
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.createAppSubscription(_ctx, tx)
	})
}
