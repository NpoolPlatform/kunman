package ledger

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/ledger/db"
	ent "github.com/NpoolPlatform/kunman/middleware/ledger/db/ent/generated"
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

func (h *MultiHandler) LockBalancesWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if _, err := handler.LockBalancesWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) LockBalances(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.LockBalancesWithTx(_ctx, tx)
	})
}
