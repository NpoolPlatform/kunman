package orderlock

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

func (h *createHandler) createOrderLock(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create orderlock: %v", err)
	}
	return nil
}

func (h *Handler) CreateOrderLock(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateOrderLockWithTx(_ctx, tx)
	})
}

func (h *Handler) CreateOrderLockWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	handler.sql = handler.ConstructCreateSQL()
	return handler.createOrderLock(ctx, tx)
}
