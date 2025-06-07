package subscriptionorder

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	orderbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/orderbase"
	orderstatebasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/statebase"
	subscriptionordercrud "github.com/NpoolPlatform/kunman/middleware/order/crud/subscription"
	subscriptionorderstatecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/subscription/state"
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

func (h *deleteHandler) deleteSubscriptionOrder(ctx context.Context, tx *ent.Tx) error {
	_, err := subscriptionordercrud.UpdateSet(
		tx.SubscriptionOrder.UpdateOneID(h._ent.SubscriptionOrderID()),
		&subscriptionordercrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *deleteHandler) deleteSubscriptionOrderState(ctx context.Context, tx *ent.Tx) error {
	_, err := subscriptionorderstatecrud.UpdateSet(
		tx.SubscriptionOrderState.UpdateOneID(h._ent.SubscriptionOrderStateID()),
		&subscriptionorderstatecrud.Req{
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
	if err := handler.deleteSubscriptionOrder(ctx, tx); err != nil {
		return wlog.WrapError(err)
	}
	return handler.deleteSubscriptionOrderState(ctx, tx)
}

func (h *Handler) DeleteSubscriptionOrder(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteSubscriptionOrderWithTx(_ctx, tx)
	})
}
