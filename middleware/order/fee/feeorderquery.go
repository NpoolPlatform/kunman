package feeorder

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entfeeorder "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/feeorder"
	entfeeorderstate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/feeorderstate"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entordercoupon "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/ordercoupon"
	entorderlock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderlock"
	entorderstatebase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
	entpaymentbalance "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalance"
	entpaymentbalancelock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalancelock"
	entpaymentbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbase"
	entpaymenttransfer "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymenttransfer"
)

type feeOrderQueryHandler struct {
	*Handler
	_ent feeOrder
}

func (h *feeOrderQueryHandler) getFeeOrderEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.FeeOrder.Query()
	if h.ID != nil {
		stm.Where(entfeeorder.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entfeeorder.EntID(*h.EntID))
	}
	if h.OrderID != nil {
		stm.Where(entfeeorder.OrderID(*h.OrderID))
	}
	if h._ent.entFeeOrder, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *feeOrderQueryHandler) getFeeOrderState(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entFeeOrderState, err = cli.
		FeeOrderState.
		Query().
		Where(
			entfeeorderstate.OrderID(h._ent.entFeeOrder.OrderID),
			entfeeorderstate.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *feeOrderQueryHandler) getOrderBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderBase, err = cli.
		OrderBase.
		Query().
		Where(
			entorderbase.EntID(h._ent.entFeeOrder.OrderID),
			entorderbase.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *feeOrderQueryHandler) getOrderStateBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderStateBase, err = cli.
		OrderStateBase.
		Query().
		Where(
			entorderstatebase.OrderID(h._ent.entFeeOrder.OrderID),
			entorderstatebase.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *feeOrderQueryHandler) getPaymentBase(ctx context.Context, cli *ent.Client) (err error) {
	if h._ent.entPaymentBase, err = cli.
		PaymentBase.
		Query().
		Where(
			entpaymentbase.OrderID(h._ent.entFeeOrder.OrderID),
			entpaymentbase.EntID(h._ent.entFeeOrderState.PaymentID),
			entpaymentbase.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
	}
	return wlog.WrapError(err)
}

func (h *feeOrderQueryHandler) getLedgerLock(ctx context.Context, cli *ent.Client) (err error) {
	paymentBalanceLock, err := cli.
		PaymentBalanceLock.
		Query().
		Where(
			entpaymentbalancelock.PaymentID(h._ent.entFeeOrderState.PaymentID),
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
			entorderlock.OrderID(h._ent.entFeeOrder.OrderID),
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

func (h *feeOrderQueryHandler) getPaymentBalances(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPaymentBalances, err = cli.
		PaymentBalance.
		Query().
		Where(
			entpaymentbalance.PaymentID(h._ent.entFeeOrderState.PaymentID),
			entpaymentbalance.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *feeOrderQueryHandler) getPaymentTransfers(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPaymentTransfers, err = cli.
		PaymentTransfer.
		Query().
		Where(
			entpaymenttransfer.PaymentID(h._ent.entFeeOrderState.PaymentID),
			entpaymenttransfer.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *feeOrderQueryHandler) getOrderCoupons(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderCoupons, err = cli.
		OrderCoupon.
		Query().
		Where(
			entordercoupon.OrderID(h._ent.entFeeOrder.OrderID),
			entordercoupon.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *feeOrderQueryHandler) getPayWithMeOrders(ctx context.Context, cli *ent.Client) error {
	infos, err := cli.
		FeeOrderState.
		Query().
		Where(
			entfeeorderstate.PaymentID(h._ent.entFeeOrderState.PaymentID),
			entfeeorderstate.OrderIDNEQ(h._ent.entFeeOrder.OrderID),
			entfeeorderstate.DeletedAt(0),
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

func (h *feeOrderQueryHandler) _getFeeOrder(ctx context.Context, cli *ent.Client, must bool) error {
	if h.ID == nil && h.EntID == nil && h.OrderID == nil {
		return wlog.Errorf("invalid id")
	}
	if err := h.getFeeOrderEnt(ctx, cli, must); err != nil {
		return wlog.WrapError(err)
	}
	if h._ent.entFeeOrder == nil {
		return nil
	}
	if err := h.getFeeOrderState(ctx, cli); err != nil {
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

func (h *feeOrderQueryHandler) getFeeOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getFeeOrder(ctx, tx.Client(), false)
}

func (h *feeOrderQueryHandler) requireFeeOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getFeeOrder(ctx, tx.Client(), true)
}

// nolint
func (h *feeOrderQueryHandler) getFeeOrder(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h._getFeeOrder(_ctx, cli, false)
	})
}

// nolint
func (h *feeOrderQueryHandler) requireFeeOrder(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h._getFeeOrder(_ctx, cli, true)
	})
}
