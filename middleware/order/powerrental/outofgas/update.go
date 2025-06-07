package outofgas

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	outofgascrud "github.com/NpoolPlatform/kunman/middleware/order/crud/outofgas"
	powerrentalstatecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/powerrental/state"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	outofgas1 "github.com/NpoolPlatform/kunman/middleware/order/outofgas"
	powerrentalstate1 "github.com/NpoolPlatform/kunman/middleware/order/powerrental/state"

	"github.com/google/uuid"
)

type updateHandler struct {
	*powerRentalStateQueryHandler
	outOfGasHandler     *outOfGasQueryHandler
	sqlOutOfGas         string
	sqlPowerRentalState string
}

func (h *updateHandler) constructOutOfGasSQL(ctx context.Context, req *outofgascrud.Req) (err error) {
	handler, _ := outofgas1.NewHandler(ctx)
	handler.Req = *req
	handler.ID = h.ID
	h.sqlOutOfGas, err = handler.ConstructUpdateSQL()
	return wlog.WrapError(err)
}

func (h *updateHandler) constructPowerRentalStateSQL(ctx context.Context, req *powerrentalstatecrud.Req) (err error) {
	handler, _ := powerrentalstate1.NewHandler(ctx)
	handler.Req = *req
	h.sqlPowerRentalState, err = handler.ConstructUpdateSQL()
	return wlog.WrapError(err)
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail update outofgas: %v", err)
	}
	return nil
}

func (h *updateHandler) updateOutOfGas(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOutOfGas)
}

func (h *updateHandler) updatePowerRentalState(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlPowerRentalState)
}

func (h *Handler) UpdateOutOfGas(ctx context.Context) error {
	handler := &updateHandler{
		powerRentalStateQueryHandler: &powerRentalStateQueryHandler{
			Handler: h,
		},
		outOfGasHandler: &outOfGasQueryHandler{
			Handler: h,
		},
	}

	if err := handler.outOfGasHandler.requireOutOfGas(ctx); err != nil {
		return wlog.WrapError(err)
	}
	h.OrderID = func() *uuid.UUID { uid := handler.outOfGasHandler._ent.OrderID(); return &uid }()
	if err := handler.requirePowerRentalState(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if h.EndAt != nil && *h.EndAt <= handler.outOfGasHandler._ent.StartAt() {
		return wlog.Errorf("invalid duration")
	}

	if err := handler.constructOutOfGasSQL(ctx, &h.Req); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.constructPowerRentalStateSQL(ctx, &powerrentalstatecrud.Req{
		OrderID: h.OrderID,
		OutOfGasSeconds: func() *uint32 {
			u := *h.EndAt - handler.outOfGasHandler._ent.StartAt()
			return &u
		}(),
	}); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateOutOfGas(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updatePowerRentalState(_ctx, tx)
	})
}
