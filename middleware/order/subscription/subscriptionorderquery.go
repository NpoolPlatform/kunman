package subscriptionorder

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entordercoupon "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/ordercoupon"
	entorderlock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderlock"
	entorderstatebase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
	entpaymentbalance "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalance"
	entpaymentbalancelock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalancelock"
	entpaymentbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbase"
	entpaymenttransfer "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymenttransfer"
	entsubscriptionorder "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/subscriptionorder"
	entsubscriptionorderstate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/subscriptionorderstate"
)

type subscriptionOrderQueryHandler struct {
	*Handler
	_ent subscriptionOrder
}

func (h *subscriptionOrderQueryHandler) getSubscriptionOrderEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.SubscriptionOrder.Query()
	if h.ID != nil {
		stm.Where(entsubscriptionorder.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entsubscriptionorder.EntID(*h.EntID))
	}
	if h.OrderID != nil {
		stm.Where(entsubscriptionorder.OrderID(*h.OrderID))
	}
	if h._ent.entSubscriptionOrder, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *subscriptionOrderQueryHandler) getSubscriptionOrderState(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entSubscriptionOrderState, err = cli.
		SubscriptionOrderState.
		Query().
		Where(
			entsubscriptionorderstate.OrderID(h._ent.entSubscriptionOrder.OrderID),
			entsubscriptionorderstate.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *subscriptionOrderQueryHandler) getOrderBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderBase, err = cli.
		OrderBase.
		Query().
		Where(
			entorderbase.EntID(h._ent.entSubscriptionOrder.OrderID),
			entorderbase.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *subscriptionOrderQueryHandler) getOrderStateBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderStateBase, err = cli.
		OrderStateBase.
		Query().
		Where(
			entorderstatebase.OrderID(h._ent.entSubscriptionOrder.OrderID),
			entorderstatebase.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *subscriptionOrderQueryHandler) getPaymentBase(ctx context.Context, cli *ent.Client) (err error) {
	if h._ent.entPaymentBase, err = cli.
		PaymentBase.
		Query().
		Where(
			entpaymentbase.OrderID(h._ent.entSubscriptionOrder.OrderID),
			entpaymentbase.EntID(h._ent.entSubscriptionOrderState.PaymentID),
			entpaymentbase.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
	}
	return wlog.WrapError(err)
}

func (h *subscriptionOrderQueryHandler) getLedgerLock(ctx context.Context, cli *ent.Client) (err error) {
	paymentBalanceLock, err := cli.
		PaymentBalanceLock.
		Query().
		Where(
			entpaymentbalancelock.PaymentID(h._ent.entSubscriptionOrderState.PaymentID),
			entpaymentbalancelock.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return wlog.WrapError(err)
	}
	if h._ent.entLedgerLock, err = cli.
		OrderLock.
		Query().
		Where(
			entorderlock.EntID(paymentBalanceLock.LedgerLockID),
			entorderlock.OrderID(h._ent.entSubscriptionOrder.OrderID),
			entorderlock.LockType(types.OrderLockType_LockBalance.String()),
			entorderlock.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
	}
	return wlog.WrapError(err)
}

func (h *subscriptionOrderQueryHandler) getPaymentBalances(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPaymentBalances, err = cli.
		PaymentBalance.
		Query().
		Where(
			entpaymentbalance.PaymentID(h._ent.entSubscriptionOrderState.PaymentID),
			entpaymentbalance.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *subscriptionOrderQueryHandler) getPaymentTransfers(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPaymentTransfers, err = cli.
		PaymentTransfer.
		Query().
		Where(
			entpaymenttransfer.PaymentID(h._ent.entSubscriptionOrderState.PaymentID),
			entpaymenttransfer.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *subscriptionOrderQueryHandler) getOrderCoupons(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderCoupons, err = cli.
		OrderCoupon.
		Query().
		Where(
			entordercoupon.OrderID(h._ent.entSubscriptionOrder.OrderID),
			entordercoupon.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *subscriptionOrderQueryHandler) getPayWithMeOrders(ctx context.Context, cli *ent.Client) error {
	infos, err := cli.
		SubscriptionOrderState.
		Query().
		Where(
			entsubscriptionorderstate.PaymentID(h._ent.entSubscriptionOrderState.PaymentID),
			entsubscriptionorderstate.OrderIDNEQ(h._ent.entSubscriptionOrder.OrderID),
			entsubscriptionorderstate.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, info := range infos {
		h._ent.payWithMeOrderIDs = append(h._ent.payWithMeOrderIDs, info.OrderID)
	}
	return nil
}

func (h *subscriptionOrderQueryHandler) _getSubscriptionOrder(ctx context.Context, cli *ent.Client, must bool) error {
	if h.ID == nil && h.EntID == nil && h.OrderID == nil {
		return wlog.Errorf("invalid id")
	}
	if err := h.getSubscriptionOrderEnt(ctx, cli, must); err != nil {
		return wlog.WrapError(err)
	}
	if h._ent.entSubscriptionOrder == nil {
		return nil
	}
	if err := h.getSubscriptionOrderState(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.getOrderBase(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.getOrderStateBase(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.getPaymentBase(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.getLedgerLock(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.getPaymentBalances(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.getPaymentTransfers(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.getPayWithMeOrders(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	return h.getOrderCoupons(ctx, cli)
}

func (h *subscriptionOrderQueryHandler) getSubscriptionOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getSubscriptionOrder(ctx, tx.Client(), false)
}

func (h *subscriptionOrderQueryHandler) requireSubscriptionOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getSubscriptionOrder(ctx, tx.Client(), true)
}

// nolint
func (h *subscriptionOrderQueryHandler) getSubscriptionOrder(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h._getSubscriptionOrder(_ctx, cli, false)
	})
}

// nolint
func (h *subscriptionOrderQueryHandler) requireSubscriptionOrder(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h._getSubscriptionOrder(_ctx, cli, true)
	})
}
