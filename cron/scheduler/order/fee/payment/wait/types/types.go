package types

import (
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type PersistentOrder struct {
	*feeordermwpb.FeeOrder
	NewOrderState   ordertypes.OrderState
	NewPaymentState *ordertypes.PaymentState
	Error           error
}
