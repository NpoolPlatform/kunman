package types

import (
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type PersistentFeeOrder struct {
	*feeordermwpb.FeeOrder
	NewPaymentState *ordertypes.PaymentState
}
