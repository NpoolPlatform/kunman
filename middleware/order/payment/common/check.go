package paymentcommon

import (
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	paymentbalancecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/balance"
	paymenttransfercrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/transfer"

	"github.com/shopspring/decimal"
)

type PaymentCheckHandler struct {
	PaymentType         *types.PaymentType
	PaymentBalanceReqs  []*paymentbalancecrud.Req
	PaymentTransferReqs []*paymenttransfercrud.Req
	PaymentAmountUSD    *decimal.Decimal
	DiscountAmountUSD   *decimal.Decimal
	Simulate            *bool
}

//nolint:gocyclo
func (h *PaymentCheckHandler) ValidatePayment() error {
	totalAmount := decimal.NewFromInt(0)
	for _, balance := range h.PaymentBalanceReqs {
		handler := PaymentCommonHandler{
			LocalCoinUSDCurrency: balance.LocalCoinUSDCurrency,
			LiveCoinUSDCurrency:  balance.LiveCoinUSDCurrency,
		}
		totalAmount = totalAmount.Add(balance.Amount.Mul(*handler.FormalizeCoinUSDCurrency()))
	}
	for _, transfer := range h.PaymentTransferReqs {
		handler := PaymentCommonHandler{
			LocalCoinUSDCurrency: transfer.LocalCoinUSDCurrency,
			LiveCoinUSDCurrency:  transfer.LiveCoinUSDCurrency,
		}
		totalAmount = totalAmount.Add(transfer.Amount.Mul(*handler.FormalizeCoinUSDCurrency()))
	}
	if h.PaymentAmountUSD != nil && h.PaymentAmountUSD.Sub(totalAmount).Abs().GreaterThan(decimal.RequireFromString("0.00000001")) {
		return wlog.Errorf("invalid paymentamount")
	}

	switch *h.PaymentType {
	case types.PaymentType_PayWithBalanceOnly:
		if len(h.PaymentBalanceReqs) == 0 {
			return wlog.Errorf("invalid paymentbalances")
		}
	case types.PaymentType_PayWithTransferOnly:
		if len(h.PaymentTransferReqs) == 0 {
			return wlog.Errorf("invalid paymenttransfers")
		}
	case types.PaymentType_PayWithTransferAndBalance:
		if len(h.PaymentBalanceReqs) == 0 {
			return wlog.Errorf("invalid paymentbalances")
		}
		if len(h.PaymentTransferReqs) == 0 {
			return wlog.Errorf("invalid paymenttransfers")
		}
	default:
		if len(h.PaymentBalanceReqs) > 0 {
			return wlog.Errorf("invalid paymentbalances")
		}
		if len(h.PaymentTransferReqs) > 0 {
			return wlog.Errorf("invalid paymenttransfers")
		}
	}

	switch *h.PaymentType {
	case types.PaymentType_PayWithBalanceOnly:
		fallthrough // nolint
	case types.PaymentType_PayWithTransferOnly:
		fallthrough // nolint
	case types.PaymentType_PayWithTransferAndBalance:
		if h.PaymentAmountUSD == nil || h.PaymentAmountUSD.Equal(decimal.NewFromInt(0)) {
			return wlog.Errorf("invalid paymentamount")
		}
	default:
		if h.PaymentAmountUSD != nil && !h.PaymentAmountUSD.Equal(decimal.NewFromInt(0)) {
			return wlog.Errorf("invalid paymentamount")
		}
		if h.DiscountAmountUSD != nil && !h.DiscountAmountUSD.Equal(decimal.NewFromInt(0)) {
			return wlog.Errorf("invalid paymentamount")
		}
	}
	return nil
}

func (h *PaymentCheckHandler) Payable() bool {
	if h.PaymentType == nil {
		return false
	}
	if h.Simulate != nil && *h.Simulate {
		return false
	}
	switch *h.PaymentType {
	case types.PaymentType_PayWithParentOrder:
		fallthrough //nolint
	case types.PaymentType_PayWithContract:
		fallthrough //nolint
	case types.PaymentType_PayWithOffline:
		fallthrough //nolint
	case types.PaymentType_PayWithNoPayment:
		return false
	}
	return true
}
