package orderstm

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entorderstatebase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
)

type orderQueryHandler struct {
	*Handler
	_ent order
}

func (h *orderQueryHandler) getOrderBaseEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.OrderBase.Query()
	if h.OrderID != nil {
		stm.Where(entorderbase.EntID(*h.OrderID))
	}
	if h._ent.entOrderBase, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *orderQueryHandler) getOrderStateBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderStateBase, err = cli.
		OrderStateBase.
		Query().
		Where(
			entorderstatebase.OrderID(h._ent.entOrderBase.EntID),
			entorderstatebase.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *orderQueryHandler) _getOrder(ctx context.Context, cli *ent.Client, must bool) error {
	if h.OrderID == nil {
		return wlog.Errorf("invalid orderid")
	}
	if err := h.getOrderBaseEnt(ctx, cli, must); err != nil {
		return wlog.WrapError(err)
	}
	if h._ent.entOrderBase == nil {
		return nil
	}
	return h.getOrderStateBase(ctx, cli)
}

// nolint
func (h *orderQueryHandler) requireOrder(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h._getOrder(_ctx, cli, true)
	})
}

func (h *orderQueryHandler) requireOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getOrder(ctx, tx.Client(), true)
}
