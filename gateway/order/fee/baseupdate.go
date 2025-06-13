package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
)

type baseUpdateHandler struct {
	*checkHandler
	*ordercommon.OrderOpHandler
	feeOrder    *npool.FeeOrder
	feeOrderReq *feeordermwpb.FeeOrderReq
	appFee      *appfeemwpb.Fee
}

func (h *baseUpdateHandler) getFeeOrder(ctx context.Context) (err error) {
	h.feeOrder, err = h.GetFeeOrder(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.feeOrder == nil {
		return wlog.Errorf("invalid feeorder")
	}
	return nil
}

func (h *baseUpdateHandler) getAppFee(ctx context.Context) (err error) {
	handler, err := appfeemw.NewHandler(
		ctx,
		appfeemw.WithAppGoodID(&h.feeOrder.AppGoodID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.appFee, err = handler.GetFee(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.appFee == nil {
		return wlog.Errorf("invalid appfee")
	}
	return nil
}

func (h *baseUpdateHandler) validateCancelParam() error {
	if err := h.ValidateCancelParam(); err != nil {
		return wlog.WrapError(err)
	}
	if h.feeOrder.AdminSetCanceled || h.feeOrder.UserSetCanceled {
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *baseUpdateHandler) constructFeeOrderReq() {
	req := &feeordermwpb.FeeOrderReq{
		ID:               &h.feeOrder.ID,
		EntID:            &h.feeOrder.EntID,
		OrderID:          &h.feeOrder.OrderID,
		PaymentType:      h.PaymentType,
		LedgerLockID:     h.BalanceLockID,
		PaymentID:        h.PaymentID,
		UserSetPaid:      h.UserSetPaid,
		UserSetCanceled:  h.Handler.UserSetCanceled,
		AdminSetCanceled: h.Handler.AdminSetCanceled,
	}
	req.PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		req.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	h.OrderIDs = append(h.OrderIDs, *req.OrderID)
	h.feeOrderReq = req
}

func (h *baseUpdateHandler) withUpdateFeeOrder(ctx context.Context) error {
	handler, err := feeordermw.NewHandler(
		ctx,
		feeordermw.WithID(h.feeOrderReq.ID, false),
		feeordermw.WithEntID(h.feeOrderReq.EntID, false),
		feeordermw.WithOrderID(h.feeOrderReq.OrderID, false),
		feeordermw.WithPaymentType(h.feeOrderReq.PaymentType, false),

		feeordermw.WithOrderState(h.feeOrderReq.OrderState, false),
		feeordermw.WithUserSetPaid(h.feeOrderReq.UserSetPaid, false),
		feeordermw.WithUserSetCanceled(h.feeOrderReq.UserSetCanceled, false),
		feeordermw.WithAdminSetCanceled(h.feeOrderReq.AdminSetCanceled, false),
		feeordermw.WithPaymentState(h.feeOrderReq.PaymentState, false),
		feeordermw.WithRollback(h.feeOrderReq.Rollback, false),
		feeordermw.WithLedgerLockID(h.feeOrderReq.LedgerLockID, false),
		feeordermw.WithPaymentID(h.feeOrderReq.PaymentID, false),
		feeordermw.WithPaymentBalances(h.feeOrderReq.PaymentBalances, false),
		feeordermw.WithPaymentTransfers(h.feeOrderReq.PaymentTransfers, false),

		feeordermw.WithMainOrder(func() *bool { b := true; return &b }(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return wlog.WrapError(handler.UpdateFeeOrder(ctx))
}

func (h *baseUpdateHandler) formalizePayment() {
	h.feeOrderReq.PaymentType = h.PaymentType
	h.feeOrderReq.PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		h.feeOrderReq.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	h.feeOrderReq.LedgerLockID = h.BalanceLockID
	h.feeOrderReq.PaymentID = h.PaymentID
}

func (h *baseUpdateHandler) updateFeeOrder(ctx context.Context) error {
	if len(h.CommissionLockIDs) > 0 {
		if err := h.WithCreateOrderCommissionLocks(ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.WithLockCommissions(ctx); err != nil {
			return wlog.WrapError(err)
		}
	}
	if err := h.WithLockBalances(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err := h.WithLockPaymentTransferAccount(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return wlog.WrapError(h.withUpdateFeeOrder(ctx))
}
