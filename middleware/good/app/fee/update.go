package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appfeecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/fee"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type updateHandler struct {
	*appFeeGoodQueryHandler
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

func (h *updateHandler) updateAppFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := appfeecrud.UpdateSet(
		tx.AppFee.UpdateOneID(h.appFee.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateFee(ctx context.Context) error {
	handler := &updateHandler{
		appFeeGoodQueryHandler: &appFeeGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireAppFeeGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if h.UnitValue != nil && h.UnitValue.LessThan(handler.fee.UnitValue) {
		return wlog.Errorf("invalid unitvalue")
	}

	if err := handler.formalizeMinOrderDurationSeconds(); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updateAppFee(_ctx, tx)
	})
}
