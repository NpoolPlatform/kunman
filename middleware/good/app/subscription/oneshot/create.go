package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	apponeshotcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/subscription/oneshot"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/google/uuid"
)

type createHandler struct {
	*appOneShotGoodQueryHandler
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

func (h *createHandler) createAppOneShot(ctx context.Context, tx *ent.Tx) error {
	if _, err := apponeshotcrud.CreateSet(
		tx.AppSubscriptionOneShot.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateOneShot(ctx context.Context) error {
	handler := &createHandler{
		appOneShotGoodQueryHandler: &appOneShotGoodQueryHandler{
			Handler: h,
		},
	}
	if err := handler.requireOneShotGood(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if h.USDPrice == nil {
		h.USDPrice = &handler.oneShot.UsdPrice
	} else if h.USDPrice.LessThan(handler.oneShot.UsdPrice) {
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
		return handler.createAppOneShot(_ctx, tx)
	})
}
