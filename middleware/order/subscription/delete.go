package subscriptionorder

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	orderbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/orderbase"
	orderstatebasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/statebase"
	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/order/crud/subscription"
	subscriptionstatecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/subscription/state"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
)

type deleteHandler struct {
	*subscriptionOrderQueryHandler
	now uint32
}

func (h *deleteHandler) deleteOrderBase(ctx context.Context, tx *ent.Tx) error {
	_, err := orderbasecrud.UpdateSet(
		tx.OrderBase.UpdateOneID(h._ent.OrderBaseID()),
		&orderbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) deleteOrderStateBase(ctx context.Context, tx *ent.Tx) error {
	_, err := orderstatebasecrud.UpdateSet(
		tx.OrderStateBase.UpdateOneID(h._ent.OrderStateBaseID()),
		&orderstatebasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) deleteSubscription(ctx context.Context, tx *ent.Tx) error {
	_, err := subscriptioncrud.UpdateSet(
		tx.SubscriptionOrder.UpdateOneID(h._ent.SubscriptionOrderID()),
		&subscriptioncrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) deleteSubscriptionState(ctx context.Context, tx *ent.Tx) error {
	_, err := subscriptionstatecrud.UpdateSet(
		tx.SubscriptionOrderState.UpdateOneID(h._ent.SubscriptionOrderStateID()),
		&subscriptionstatecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) DeleteSubscriptionOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &deleteHandler{
		subscriptionOrderQueryHandler: &subscriptionOrderQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getSubscriptionOrderWithTx(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if !handler._ent.Exist() {
		return nil
	}

	if err := handler.deleteOrderBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.deleteOrderStateBase(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	if err := handler.deleteSubscription(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return handler.deleteSubscriptionState(ctx, tx)
}

func (h *Handler) DeleteSubscriptionOrder(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		// TODO: also delete child orders
		return h.DeleteSubscriptionOrderWithTx(_ctx, tx)
	})
}
