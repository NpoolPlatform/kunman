package ordercoupon

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

func (h *createHandler) constructSQL() {
	h.sql = h.ConstructCreateSQL()
}

func (h *createHandler) createOrderCoupon(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail create ordercoupon: %v", err)
	}
	return nil
}

func (h *Handler) CreateOrderCoupon(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	handler := &createHandler{
		Handler: h,
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return handler.createOrderCoupon(ctx, tx)
	})
}
