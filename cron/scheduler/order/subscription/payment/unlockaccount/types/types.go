package types

import (
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
)

type PersistentOrder struct {
	*subscriptionordermwpb.SubscriptionOrder
	PaymentAccountIDs          []uint32
	ExistOrderCompletedHistory bool
}
