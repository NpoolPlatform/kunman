package orderlock

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
)

type MultiHandler struct {
	Handlers []*Handler
}

func (h *MultiHandler) AppendHandler(handler *Handler) {
	h.Handlers = append(h.Handlers, handler)
}

func (h *MultiHandler) GetHandlers() []*Handler {
	return h.Handlers
}

func (h *MultiHandler) CreateOrderLocksWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.CreateOrderLockWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) CreateOrderLocks(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateOrderLocksWithTx(_ctx, tx)
	})
}

func (h *MultiHandler) DeleteOrderLocksWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.DeleteOrderLockWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) DeleteOrderLocks(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteOrderLocksWithTx(_ctx, tx)
	})
}
