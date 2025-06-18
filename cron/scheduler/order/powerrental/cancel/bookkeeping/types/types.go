package types

import (
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"

	"github.com/shopspring/decimal"
)

type XPaymentTransfer struct {
	PaymentTransferID     string
	CoinTypeID            string
	AccountID             string
	PaymentAccountBalance string
	IncomingAmount        *string
	Amount                decimal.Decimal
	StartAmount           decimal.Decimal
	FinishAmount          string
}

type PersistentPowerRentalOrder struct {
	*powerrentalordermwpb.PowerRentalOrder
	XPaymentTransfers []*XPaymentTransfer
	IncomingExtra     string
}
