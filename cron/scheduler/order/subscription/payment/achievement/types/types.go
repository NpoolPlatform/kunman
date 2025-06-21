package types

import (
	orderstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
)

type PersistentOrder struct {
	*subscriptionordermwpb.SubscriptionOrder
	OrderStatements []*orderstatementmwpb.StatementReq
}
