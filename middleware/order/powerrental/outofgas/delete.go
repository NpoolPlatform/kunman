package outofgas

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	outofgascrud "github.com/NpoolPlatform/kunman/middleware/order/crud/outofgas"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	outofgas1 "github.com/NpoolPlatform/kunman/middleware/order/outofgas"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*powerRentalStateQueryHandler
	outOfGasSeconds uint32
	now             uint32
}

func (h *deleteHandler) deleteOutOfGas(ctx context.Context, tx *ent.Tx) error {
	_, err := outofgascrud.UpdateSet(
		tx.OutOfGas.UpdateOneID(*h.ID),
		&outofgascrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) updatePowerRentalState(ctx context.Context, tx *ent.Tx) error {
	_, err := tx.
		PowerRentalState.
		UpdateOneID(h._ent.PowerRentalStateID()).
		AddOutofgasSeconds(0 - int32(h.outOfGasSeconds)).
		Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) DeleteOutOfGas(ctx context.Context) error {
	h1, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithID(h.ID, false),
		outofgas1.WithEntID(func() *string {
			if h.EntID == nil {
				return nil
			}
			s := h.EntID.String()
			return &s
		}(), false),
		outofgas1.WithOrderID(func() *string {
			if h.OrderID == nil {
				return nil
			}
			s := h.OrderID.String()
			return &s
		}(), false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	info, err := h1.GetOutOfGas(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	handler := &deleteHandler{
		powerRentalStateQueryHandler: &powerRentalStateQueryHandler{
			Handler: h,
		},
		outOfGasSeconds: info.EndAt - info.StartAt,
		now:             uint32(time.Now().Unix()),
	}

	h.ID = &info.ID
	h.EntID = func() *uuid.UUID { uid := uuid.MustParse(info.EntID); return &uid }()
	h.OrderID = func() *uuid.UUID { uid := uuid.MustParse(info.OrderID); return &uid }()

	if err := handler.requirePowerRentalState(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteOutOfGas(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if info.EndAt == 0 {
			return nil
		}
		return handler.updatePowerRentalState(_ctx, tx)
	})
}
