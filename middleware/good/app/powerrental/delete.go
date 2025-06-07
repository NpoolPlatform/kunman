package powerrental

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	apppowerrentalcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/powerrental"
	"github.com/NpoolPlatform/kunman/middleware/good/db"
	"github.com/NpoolPlatform/kunman/middleware/good/db/ent/generated"
)

type deleteHandler struct {
	*powerRentalAppGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodbasecrud.UpdateSet(
		tx.AppGoodBase.UpdateOneID(h._ent.appGoodBase.ID),
		&appgoodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *deleteHandler) deletePowerRental(ctx context.Context, tx *ent.Tx) error {
	if _, err := apppowerrentalcrud.UpdateSet(
		tx.AppPowerRental.UpdateOneID(h._ent.appPowerRental.ID),
		&apppowerrentalcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeletePowerRental(ctx context.Context) error {
	handler := &deleteHandler{
		powerRentalAppGoodQueryHandler: &powerRentalAppGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getAppPowerRentalAppGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if handler._ent.appPowerRental == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.deletePowerRental(_ctx, tx)
	})
}
