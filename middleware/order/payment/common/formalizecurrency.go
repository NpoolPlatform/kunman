package paymentcommon

import (
	"github.com/shopspring/decimal"
)

type PaymentCommonHandler struct {
	LocalCoinUSDCurrency *decimal.Decimal
	LiveCoinUSDCurrency  *decimal.Decimal
}

func (h *PaymentCommonHandler) FormalizeCoinUSDCurrency() *decimal.Decimal {
	if h.LocalCoinUSDCurrency != nil && h.LocalCoinUSDCurrency.Cmp(decimal.NewFromInt(0)) > 0 {
		return h.LocalCoinUSDCurrency
	}
	return h.LiveCoinUSDCurrency
}
