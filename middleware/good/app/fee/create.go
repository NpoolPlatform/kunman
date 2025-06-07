package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appfeecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/fee"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
	"github.com/google/uuid"
)

type createHandler struct {
	*appFeeGoodQueryHandler
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

func (h *createHandler) createAppFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := appfeecrud.CreateSet(
		tx.AppFee.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateFee(ctx context.Context) error {
	handler := &createHandler{
		appFeeGoodQueryHandler: &appFeeGoodQueryHandler{
			Handler: h,
		},
	}
	if err := handler.requireFeeGood(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if h.UnitValue == nil {
		h.UnitValue = &handler.fee.UnitValue
	} else if h.UnitValue.LessThan(handler.fee.UnitValue) {
		return wlog.Errorf("invalid unitvalue")
	}
	if h.AppGoodID == nil {
		id := uuid.New()
		h.AppGoodID = &id
		h.AppGoodBaseReq.EntID = &id
	}

	if err := handler.formalizeMinOrderDurationSeconds(); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.createAppFee(_ctx, tx)
	})
}
