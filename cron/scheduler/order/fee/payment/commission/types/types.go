package types

import (
	ledgerstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type PersistentOrder struct {
	*feeordermwpb.FeeOrder
	LedgerStatements []*ledgerstatementmwpb.StatementReq
}
