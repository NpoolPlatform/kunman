package outofgas

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	outofgascrud "github.com/NpoolPlatform/kunman/middleware/order/crud/outofgas"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	outofgas1 "github.com/NpoolPlatform/kunman/middleware/order/outofgas"

	"github.com/google/uuid"
)

type createHandler struct {
	*powerRentalStateQueryHandler
	sqlOutOfGas string
}

func (h *createHandler) constructOutOfGasSQL(ctx context.Context, req *outofgascrud.Req) {
	handler, _ := outofgas1.NewHandler(ctx)
	handler.Req = *req
	h.sqlOutOfGas = handler.ConstructCreateSQL()
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create outofgas: %v", err)
	}
	return nil
}

func (h *createHandler) createOutOfGas(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOutOfGas)
}

func (h *Handler) CreateOutOfGas(ctx context.Context) error {
	handler := &createHandler{
		powerRentalStateQueryHandler: &powerRentalStateQueryHandler{
			Handler: h,
		},
	}
	if err := handler.requirePowerRentalState(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	handler.constructOutOfGasSQL(ctx, &h.Req)
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createOutOfGas(_ctx, tx)
	})
}
