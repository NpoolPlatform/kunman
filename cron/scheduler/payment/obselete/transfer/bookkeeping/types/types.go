package types

import (
	ledgerstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
)

type PersistentPayment struct {
	*paymentmwpb.Payment
	Statements       []*ledgerstatementmwpb.StatementReq
	PaymentTransfers []*paymentmwpb.PaymentTransferReq
}
