package types

import (
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
)

type PersistentOrder struct {
	*subscriptionordermwpb.SubscriptionOrder
	NewOrderState   ordertypes.OrderState
	NewPaymentState *ordertypes.PaymentState
	Error           error
}
