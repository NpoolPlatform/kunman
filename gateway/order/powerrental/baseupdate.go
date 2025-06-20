package powerrental

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	goodledgerstatementpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/statement"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	goodledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/good/ledger/statement"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type baseUpdateHandler struct {
	*checkHandler
	*ordercommon.OrderOpHandler
	powerRentalOrder    *npool.PowerRentalOrder
	powerRentalOrderReq *powerrentalordermwpb.PowerRentalOrderReq
	appPowerRental      *apppowerrentalmwpb.PowerRental
	goodBenefitedAt     uint32
}

func (h *baseUpdateHandler) getPowerRentalOrder(ctx context.Context) (err error) {
	h.powerRentalOrder, err = h.GetPowerRentalOrder(ctx)
	return wlog.WrapError(err)
}

func (h *baseUpdateHandler) validateOrderStateWhenCancel() error {
	if h.powerRentalOrder.GoodStockMode == goodtypes.GoodStockMode_GoodStockByMiningPool &&
		h.powerRentalOrder.OrderState == ordertypes.OrderState_OrderStateInService {
		return wlog.Errorf("cannot cancel in service order of stock by miningpool")
	}
	return nil
}

func (h *baseUpdateHandler) validateCancelParam() error {
	if err := h.ValidateCancelParam(); err != nil {
		return wlog.WrapError(err)
	}
	if h.powerRentalOrder.AdminSetCanceled || h.powerRentalOrder.UserSetCanceled {
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *baseUpdateHandler) getGoodBenefitTime(ctx context.Context) error {
	conds := &goodledgerstatementpb.Conds{
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: h.powerRentalOrder.GoodID},
	}
	handler, err := goodledgerstatementmw.NewHandler(
		ctx,
		goodledgerstatementmw.WithConds(conds),
		goodledgerstatementmw.WithOffset(0),
		goodledgerstatementmw.WithLimit(1),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	statements, _, err := handler.GetGoodStatements(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(statements) > 0 {
		h.goodBenefitedAt = statements[0].BenefitDate
	}
	return nil
}

func (h *baseUpdateHandler) getAppPowerRental(ctx context.Context) (err error) {
	handler, err := apppowerrentalmw.NewHandler(
		ctx,
		apppowerrentalmw.WithAppGoodID(&h.powerRentalOrder.AppGoodID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.appPowerRental, err = handler.GetPowerRental(ctx)
	return wlog.WrapError(err)
}

func (h *baseUpdateHandler) goodCancelable() error {
	if err := h.GoodCancelable(); err != nil {
		return wlog.WrapError(err)
	}
	now := uint32(time.Now().Unix())
	switch h.appPowerRental.CancelMode {
	case goodtypes.CancelMode_CancellableBeforeStart:
		checkOrderStartAt := h.powerRentalOrder.StartAt - h.appPowerRental.CancelableBeforeStartSeconds
		if checkOrderStartAt <= now {
			return wlog.Errorf("permission denied")
		}
	case goodtypes.CancelMode_CancellableBeforeBenefit:
		if h.goodBenefitedAt == 0 {
			return nil
		}
		if h.powerRentalOrder.LastBenefitAt != 0 {
			return wlog.Errorf("permission denied")
		}
		benefitIntervalSeconds := uint32((time.Duration(h.appPowerRental.BenefitIntervalHours) * time.Hour).Seconds())
		thisBenefitAt := uint32(time.Now().Unix()) / benefitIntervalSeconds * benefitIntervalSeconds
		nextBenefitAt := (uint32(time.Now().Unix())/benefitIntervalSeconds + 1) * benefitIntervalSeconds
		if (thisBenefitAt-h.appPowerRental.CancelableBeforeStartSeconds <= now &&
			now <= thisBenefitAt+h.appPowerRental.CancelableBeforeStartSeconds) ||
			(nextBenefitAt-h.appPowerRental.CancelableBeforeStartSeconds <= now &&
				now <= nextBenefitAt+h.appPowerRental.CancelableBeforeStartSeconds) {
			return wlog.Errorf("permission denied")
		}
	}
	return nil
}

func (h *baseUpdateHandler) constructPowerRentalOrderReq() {
	req := &powerrentalordermwpb.PowerRentalOrderReq{
		ID:               &h.powerRentalOrder.ID,
		EntID:            &h.powerRentalOrder.EntID,
		OrderID:          &h.powerRentalOrder.OrderID,
		PaymentType:      h.PaymentType,
		LedgerLockID:     h.BalanceLockID,
		PaymentID:        h.PaymentID,
		UserSetPaid:      h.UserSetPaid,
		UserSetCanceled:  h.UserSetCanceled,
		AdminSetCanceled: h.AdminSetCanceled,
	}
	req.PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		req.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	h.OrderIDs = append(h.OrderIDs, *req.OrderID)
	h.powerRentalOrderReq = req
}

func (h *baseUpdateHandler) withUpdatePowerRentalOrder(ctx context.Context) error {
	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithID(h.powerRentalOrderReq.ID, false),
		powerrentalordermw.WithEntID(h.powerRentalOrderReq.EntID, false),
		powerrentalordermw.WithOrderID(h.powerRentalOrderReq.OrderID, false),
		powerrentalordermw.WithPaymentType(h.powerRentalOrderReq.PaymentType, false),

		powerrentalordermw.WithOrderState(h.powerRentalOrderReq.OrderState, false),
		powerrentalordermw.WithStartMode(h.powerRentalOrderReq.StartMode, false),
		powerrentalordermw.WithStartAt(h.powerRentalOrderReq.StartAt, false),
		powerrentalordermw.WithLastBenefitAt(h.powerRentalOrderReq.LastBenefitAt, false),
		powerrentalordermw.WithBenefitState(h.powerRentalOrderReq.BenefitState, false),
		powerrentalordermw.WithUserSetPaid(h.powerRentalOrderReq.UserSetPaid, false),
		powerrentalordermw.WithUserSetCanceled(h.powerRentalOrderReq.UserSetCanceled, false),
		powerrentalordermw.WithAdminSetCanceled(h.powerRentalOrderReq.AdminSetCanceled, false),
		powerrentalordermw.WithPaymentState(h.powerRentalOrderReq.PaymentState, false),
		powerrentalordermw.WithRenewState(h.powerRentalOrderReq.RenewState, false),
		powerrentalordermw.WithRenewNotifyAt(h.powerRentalOrderReq.RenewNotifyAt, false),

		powerrentalordermw.WithLedgerLockID(h.powerRentalOrderReq.LedgerLockID, false),
		powerrentalordermw.WithPaymentID(h.powerRentalOrderReq.PaymentID, false),
		powerrentalordermw.WithCouponIDs(h.powerRentalOrderReq.CouponIDs, false),
		powerrentalordermw.WithPaymentBalances(h.powerRentalOrderReq.PaymentBalances, false),
		powerrentalordermw.WithPaymentTransfers(h.powerRentalOrderReq.PaymentTransfers, false),

		powerrentalordermw.WithRollback(h.powerRentalOrderReq.Rollback, false),
		powerrentalordermw.WithPoolOrderUserID(h.powerRentalOrderReq.PoolOrderUserID, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	return handler.UpdatePowerRental(ctx)
}

func (h *baseUpdateHandler) formalizePayment() {
	h.powerRentalOrderReq.PaymentType = h.PaymentType
	h.powerRentalOrderReq.PaymentBalances = h.PaymentBalanceReqs
	if h.PaymentTransferReq != nil {
		h.powerRentalOrderReq.PaymentTransfers = []*paymentmwpb.PaymentTransferReq{h.PaymentTransferReq}
	}
	h.powerRentalOrderReq.LedgerLockID = h.BalanceLockID
	h.powerRentalOrderReq.PaymentID = h.PaymentID
}

func (h *baseUpdateHandler) updatePowerRentalOrder(ctx context.Context) error {
	if !h.Simulate {
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
	}
	return h.withUpdatePowerRentalOrder(ctx)
}
