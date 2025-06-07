package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entfeeorderstate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/feeorderstate"
	entorderbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderbase"
	entordercoupon "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/ordercoupon"
	entorderlock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderlock"
	entorderstatebase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
	entpaymentbalance "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalance"
	entpaymentbalancelock "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbalancelock"
	entpaymentbase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymentbase"
	entpaymenttransfer "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymenttransfer"
	entpoolorderuser "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/poolorderuser"
	entpowerrental "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/powerrental"
	entpowerrentalstate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/powerrentalstate"
)

type powerRentalQueryHandler struct {
	*Handler
	_ent powerRental
}

func (h *powerRentalQueryHandler) getPowerRentalEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.PowerRental.Query()
	if h.ID != nil {
		stm.Where(entpowerrental.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entpowerrental.EntID(*h.EntID))
	}
	if h.OrderID != nil {
		stm.Where(entpowerrental.OrderID(*h.OrderID))
	}
	if h._ent.entPowerRental, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalQueryHandler) getPowerRentalState(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPowerRentalState, err = cli.
		PowerRentalState.
		Query().
		Where(
			entpowerrentalstate.OrderID(h._ent.entPowerRental.OrderID),
			entpowerrentalstate.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *powerRentalQueryHandler) getOrderBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderBase, err = cli.
		OrderBase.
		Query().
		Where(
			entorderbase.EntID(h._ent.entPowerRental.OrderID),
			entorderbase.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *powerRentalQueryHandler) getOrderStateBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderStateBase, err = cli.
		OrderStateBase.
		Query().
		Where(
			entorderstatebase.OrderID(h._ent.entPowerRental.OrderID),
			entorderstatebase.DeletedAt(0),
		).
		Only(ctx)
	return wlog.WrapError(err)
}

func (h *powerRentalQueryHandler) getPaymentBase(ctx context.Context, cli *ent.Client) (err error) {
	if h._ent.entPaymentBase, err = cli.
		PaymentBase.
		Query().
		Where(
			entpaymentbase.OrderID(h._ent.entPowerRental.OrderID),
			entpaymentbase.EntID(h._ent.entPowerRentalState.PaymentID),
			entpaymentbase.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
	}
	return wlog.WrapError(err)
}

func (h *powerRentalQueryHandler) getLedgerLock(ctx context.Context, cli *ent.Client) (err error) {
	// TODO: should get ID from payment balance lock firstly
	paymentBalanceLock, err := cli.
		PaymentBalanceLock.
		Query().
		Where(
			entpaymentbalancelock.PaymentID(h._ent.entPowerRentalState.PaymentID),
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
			entorderlock.OrderID(h._ent.entPowerRental.OrderID),
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

func (h *powerRentalQueryHandler) getStockLock(ctx context.Context, cli *ent.Client) (err error) {
	if h._ent.entStockLock, err = cli.
		OrderLock.
		Query().
		Where(
			entorderlock.OrderID(h._ent.entPowerRental.OrderID),
			entorderlock.LockType(types.OrderLockType_LockStock.String()),
			entorderlock.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
	}
	return wlog.WrapError(err)
}

func (h *powerRentalQueryHandler) getPaymentBalances(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPaymentBalances, err = cli.
		PaymentBalance.
		Query().
		Where(
			entpaymentbalance.PaymentID(h._ent.entPowerRentalState.PaymentID),
			entpaymentbalance.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *powerRentalQueryHandler) getPaymentTransfers(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entPaymentTransfers, err = cli.
		PaymentTransfer.
		Query().
		Where(
			entpaymenttransfer.PaymentID(h._ent.entPowerRentalState.PaymentID),
			entpaymenttransfer.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *powerRentalQueryHandler) getOrderCoupons(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderCoupons, err = cli.
		OrderCoupon.
		Query().
		Where(
			entordercoupon.OrderID(h._ent.entPowerRental.OrderID),
			entordercoupon.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *powerRentalQueryHandler) getPayWithMeOrders(ctx context.Context, cli *ent.Client) error {
	infos, err := cli.
		FeeOrderState.
		Query().
		Where(
			entfeeorderstate.PaymentID(h._ent.entPowerRentalState.PaymentID),
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

func (h *powerRentalQueryHandler) getChildOrders(ctx context.Context, cli *ent.Client) error {
	infos, err := cli.
		OrderBase.
		Query().
		Where(
			entorderbase.ParentOrderID(h._ent.entPowerRental.OrderID),
			entorderbase.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, info := range infos {
		h._ent.childOrderIDs = append(h._ent.childOrderIDs, info.EntID)
	}
	return nil
}

func (h *powerRentalQueryHandler) getPoolOrderUserID(ctx context.Context, cli *ent.Client) (err error) {
	if h._ent.poolOrderUser, err = cli.
		PoolOrderUser.
		Query().
		Where(
			entpoolorderuser.OrderID(h._ent.entPowerRental.OrderID),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
	}
	return wlog.WrapError(err)
}

//nolint:gocyclo
func (h *powerRentalQueryHandler) _getPowerRental(ctx context.Context, cli *ent.Client, must bool) error {
	if h.ID == nil && h.EntID == nil && h.OrderID == nil {
		return wlog.Errorf("invalid id")
	}
	if err := h.getPowerRentalEnt(ctx, cli, must); err != nil {
		return wlog.WrapError(err)
	}
	if h._ent.entPowerRental == nil {
		return nil
	}
	if err := h.getPowerRentalState(ctx, cli); err != nil {
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
	if err := h.getStockLock(ctx, cli); err != nil {
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
	if err := h.getChildOrders(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.getPoolOrderUserID(ctx, cli); err != nil {
		return wlog.WrapError(err)
	}
	return h.getOrderCoupons(ctx, cli)
}

func (h *powerRentalQueryHandler) getPowerRentalWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getPowerRental(ctx, tx.Client(), false)
}

func (h *powerRentalQueryHandler) requirePowerRentalWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getPowerRental(ctx, tx.Client(), true)
}

// nolint
func (h *powerRentalQueryHandler) getPowerRental(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h._getPowerRental(_ctx, cli, false)
	})
}

// nolint
func (h *powerRentalQueryHandler) requirePowerRental(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h._getPowerRental(_ctx, cli, true)
	})
}
