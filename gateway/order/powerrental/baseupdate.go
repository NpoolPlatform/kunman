package powerrental

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	goodledgerstatementcli "github.com/NpoolPlatform/ledger-middleware/pkg/client/good/ledger/statement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	goodledgerstatementpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/statement"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	ordermwsvcname "github.com/NpoolPlatform/order-middleware/pkg/servicename"

	dtmcli "github.com/NpoolPlatform/dtm-cluster/pkg/dtm"
	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"
)

type baseUpdateHandler struct {
	*dtmHandler
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
	statements, _, err := goodledgerstatementcli.GetGoodStatements(ctx, &goodledgerstatementpb.Conds{
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: h.powerRentalOrder.GoodID},
	}, 0, 1)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(statements) > 0 {
		h.goodBenefitedAt = statements[0].BenefitDate
	}
	return nil
}

func (h *baseUpdateHandler) getAppPowerRental(ctx context.Context) (err error) {
	h.appPowerRental, err = apppowerrentalmwcli.GetPowerRental(ctx, h.powerRentalOrder.AppGoodID)
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

func (h *baseUpdateHandler) withUpdatePowerRentalOrder(dispose *dtmcli.SagaDispose) {
	dispose.Add(
		ordermwsvcname.ServiceDomain,
		"order.middleware.powerrental.v1.Middleware/UpdatePowerRentalOrder",
		"",
		&powerrentalordermwpb.UpdatePowerRentalOrderRequest{
			Info: h.powerRentalOrderReq,
		},
	)
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
	sagaDispose := dtmcli.NewSagaDispose(dtmimp.TransOptions{
		WaitResult:     true,
		RequestTimeout: 10,
		TimeoutToFail:  10,
	})

	if !h.OrderOpHandler.Simulate {
		if len(h.CommissionLockIDs) > 0 {
			h.WithCreateOrderCommissionLocks(sagaDispose)
			h.WithLockCommissions(sagaDispose)
		}
		h.WithLockBalances(sagaDispose)
		h.WithLockPaymentTransferAccount(sagaDispose)
	}
	h.withUpdatePowerRentalOrder(sagaDispose)
	return h.dtmDo(ctx, sagaDispose)
}
