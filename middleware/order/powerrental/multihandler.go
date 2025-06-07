package powerrental

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

func (h *MultiHandler) UpdatePowerRentalsWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.UpdatePowerRentalWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) UpdatePowerRentals(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.UpdatePowerRentalsWithTx(_ctx, tx)
	})
}
