package types

import (
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type XPaymentTransfer struct {
	PaymentTransferID     string
	CoinTypeID            string
	AccountID             string
	PaymentAccountBalance string
	IncomingAmount        *string
	Amount                string
	StartAmount           string
	FinishAmount          string
	IncomingExtra         string
	OutcomingExtra        string
}

type PersistentOrder struct {
	*feeordermwpb.FeeOrder
	XPaymentTransfers []*XPaymentTransfer
	Error             error
}
