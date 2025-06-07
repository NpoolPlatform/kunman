package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	compensate1 "github.com/NpoolPlatform/kunman/middleware/order/compensate"
	compensatecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/compensate"
	powerrentalstatecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/powerrental/state"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	powerrentalstate1 "github.com/NpoolPlatform/kunman/middleware/order/powerrental/state"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/google/uuid"
)

type createHandler struct {
	*powerRentalStateQueryHandler
	sqlCompensates       []string
	sqlPowerRentalStates []string
}

func (h *createHandler) constructCompensateSQL(ctx context.Context, req *compensatecrud.Req) {
	handler, _ := compensate1.NewHandler(ctx)
	handler.Req = *req
	h.sqlCompensates = append(h.sqlCompensates, handler.ConstructCreateSQL())
}

func (h *createHandler) constructPowerRentalStateSQL(ctx context.Context, req *powerrentalstatecrud.Req) error {
	handler, _ := powerrentalstate1.NewHandler(ctx)
	handler.Req = *req
	sql, err := handler.ConstructUpdateSQL()
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlPowerRentalStates = append(h.sqlPowerRentalStates, sql)
	return nil
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create compensate: %v", err)
	}
	return nil
}

func (h *createHandler) createCompensates(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlCompensates {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *createHandler) updatePowerRentalStates(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPowerRentalStates {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *createHandler) createGoodCompensates(ctx context.Context, tx *ent.Tx) error {
	h.limit = constant.DefaultRowLimit

	for {
		if err := h.requirePowerRentalStatesWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if h._ent.Exhausted() {
			return nil
		}

		h.sqlCompensates = []string{}
		h.sqlPowerRentalStates = []string{}
		for _, state := range h._ent.powerRentalStates {
			h.constructCompensateSQL(ctx, &compensatecrud.Req{
				EntID:             func() *uuid.UUID { uid := uuid.New(); return &uid }(),
				OrderID:           &state.OrderID,
				CompensateFromID:  h.CompensateFromID,
				CompensateType:    h.CompensateType,
				CompensateSeconds: h.CompensateSeconds,
			})
			if err := h.constructPowerRentalStateSQL(ctx, &powerrentalstatecrud.Req{
				OrderID:           &state.OrderID,
				CompensateSeconds: h.CompensateSeconds,
			}); err != nil {
				return wlog.WrapError(err)
			}
		}

		if err := h.createCompensates(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.updatePowerRentalStates(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}

		h.offset += h.limit
	}
}

func (h *createHandler) createOrderCompensate(ctx context.Context, tx *ent.Tx) error {
	if err := h.requirePowerRentalStatesWithTx(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if h._ent.Exhausted() {
		return wlog.Errorf("invalid powerrentalorder")
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	h.constructCompensateSQL(ctx, &h.Req)
	if err := h.constructPowerRentalStateSQL(ctx, &powerrentalstatecrud.Req{
		OrderID:           h.OrderID,
		CompensateSeconds: h.CompensateSeconds,
	}); err != nil {
		return wlog.WrapError(err)
	}

	if err := h.createCompensates(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return h.updatePowerRentalStates(ctx, tx)
}

func (h *Handler) CreateCompensate(ctx context.Context) error {
	if h.OrderID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid compensate id")
	}

	handler := &createHandler{
		powerRentalStateQueryHandler: &powerRentalStateQueryHandler{
			Handler: h,
		},
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if h.OrderID != nil {
			return handler.createOrderCompensate(_ctx, tx)
		}
		return handler.createGoodCompensates(_ctx, tx)
	})
}
