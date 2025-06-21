package types

import (
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
)

type PersistentSubscriptionOrder struct {
	*subscriptionordermwpb.SubscriptionOrder
}
