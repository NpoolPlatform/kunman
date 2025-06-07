package compensate

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	compensate1 "github.com/NpoolPlatform/kunman/middleware/order/compensate"
	compensatecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/compensate"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*powerRentalStateQueryHandler
	now uint32
}

func (h *deleteHandler) deleteCompensate(ctx context.Context, tx *ent.Tx) error {
	_, err := compensatecrud.UpdateSet(
		tx.Compensate.UpdateOneID(*h.ID),
		&compensatecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) updatePowerRentalState(ctx context.Context, tx *ent.Tx) error {
	_, err := tx.
		PowerRentalState.
		UpdateOneID(h._ent.powerRentalStates[0].ID).
		AddCompensateSeconds(0 - int32(*h.CompensateSeconds)).
		Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) DeleteCompensate(ctx context.Context) error {
	h1, err := compensate1.NewHandler(
		ctx,
		compensate1.WithID(h.ID, false),
		compensate1.WithEntID(func() *string {
			if h.EntID == nil {
				return nil
			}
			s := h.EntID.String()
			return &s
		}(), false),
		compensate1.WithOrderID(func() *string {
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

	info, err := h1.GetCompensate(ctx)
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
		now: uint32(time.Now().Unix()),
	}

	h.ID = &info.ID
	h.EntID = func() *uuid.UUID { uid := uuid.MustParse(info.EntID); return &uid }()
	h.OrderID = func() *uuid.UUID { uid := uuid.MustParse(info.OrderID); return &uid }()
	h.CompensateSeconds = &info.CompensateSeconds

	if err := handler.requirePowerRentalStates(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if handler._ent.Exhausted() {
		return wlog.Errorf("invalid powerrentalorder")
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteCompensate(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updatePowerRentalState(_ctx, tx)
	})
}
