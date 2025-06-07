package powerrental

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	orderbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/orderbase"
	orderstatebasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/statebase"
	powerrentalcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/powerrental"
	poolorderusercrud "github.com/NpoolPlatform/kunman/middleware/order/crud/powerrental/poolorderuser"
	powerrentalstatecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/powerrental/state"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
)

type deleteHandler struct {
	*powerRentalQueryHandler
	now uint32
}

func (h *deleteHandler) deleteOrderBase(ctx context.Context, tx *ent.Tx) error {
	_, err := orderbasecrud.UpdateSet(
		tx.OrderBase.UpdateOneID(h._ent.OrderBaseID()),
		&orderbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) deleteOrderStateBase(ctx context.Context, tx *ent.Tx) error {
	_, err := orderstatebasecrud.UpdateSet(
		tx.OrderStateBase.UpdateOneID(h._ent.OrderStateBaseID()),
		&orderstatebasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) deletePowerRental(ctx context.Context, tx *ent.Tx) error {
	_, err := powerrentalcrud.UpdateSet(
		tx.PowerRental.UpdateOneID(h._ent.PowerRentalID()),
		&powerrentalcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) deletePowerRentalState(ctx context.Context, tx *ent.Tx) error {
	_, err := powerrentalstatecrud.UpdateSet(
		tx.PowerRentalState.UpdateOneID(h._ent.PowerRentalStateID()),
		&powerrentalstatecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) deletePoolOrderUser(ctx context.Context, tx *ent.Tx) error {
	id := h._ent.PoolOrderUserRecordID()
	if id == nil {
		return nil
	}

	_, err := poolorderusercrud.UpdateSet(
		tx.PoolOrderUser.UpdateOneID(*id),
		&poolorderusercrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) DeletePowerRentalWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &deleteHandler{
		powerRentalQueryHandler: &powerRentalQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getPowerRentalWithTx(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if !handler._ent.Exist() {
		return nil
	}

	if err := handler.deleteOrderBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.deleteOrderStateBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.deletePoolOrderUser(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.deletePowerRental(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return handler.deletePowerRentalState(ctx, tx)
}

func (h *Handler) DeletePowerRental(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		// TODO: also delete child orders
		return h.DeletePowerRentalWithTx(_ctx, tx)
	})
}
