package types

import (
	orderstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type PersistentOrder struct {
	*feeordermwpb.FeeOrder
	OrderStatements []*orderstatementmwpb.StatementReq
}
