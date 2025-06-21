package types

import (
	ledgerstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
)

type PersistentOrder struct {
	*subscriptionordermwpb.SubscriptionOrder
	LedgerStatements []*ledgerstatementmwpb.StatementReq
}
