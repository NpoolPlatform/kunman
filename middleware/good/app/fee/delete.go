package fee

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appfeecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/fee"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	ent "github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*appFeeGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.appGoodBase == nil {
		return wlog.Errorf("invalid appgoodbase")
	}
	if _, err := appgoodbasecrud.UpdateSet(
		tx.AppGoodBase.UpdateOneID(h.appGoodBase.ID),
		&appgoodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := appfeecrud.UpdateSet(
		tx.AppFee.UpdateOneID(h.appFee.ID),
		&appfeecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteFee(ctx context.Context) error {
	handler := &deleteHandler{
		appFeeGoodQueryHandler: &appFeeGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getAppFeeGood(ctx); err != nil {
		return err
	}
	if handler.appFee == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.deleteAppFee(_ctx, tx)
	})
}
