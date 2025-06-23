package types

import (
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
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

type XPaymentFiat struct {
	PaymentFiatID string
	FiatID        string
	Amount        string
	Extra         string
}

type PersistentOrder struct {
	*subscriptionordermwpb.SubscriptionOrder
	XPaymentTransfers []*XPaymentTransfer
	XPaymentFiats     []*XPaymentFiat
	Error             error
}
