package oneshot

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	apponeshotcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/subscription/oneshot"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*appOneShotGoodQueryHandler
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

func (h *updateHandler) updateAppOneShot(ctx context.Context, tx *ent.Tx) error {
	if _, err := apponeshotcrud.UpdateSet(
		tx.AppSubscriptionOneShot.UpdateOneID(h.appOneShot.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateOneShot(ctx context.Context) error {
	handler := &updateHandler{
		appOneShotGoodQueryHandler: &appOneShotGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireAppOneShotGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if h.USDPrice != nil && h.USDPrice.LessThan(handler.oneShot.UsdPrice) {
		return wlog.Errorf("invalid usdprice")
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updateAppOneShot(_ctx, tx)
	})
}
