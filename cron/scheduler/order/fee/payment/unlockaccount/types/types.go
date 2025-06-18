package types

import (
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type PersistentOrder struct {
	*feeordermwpb.FeeOrder
	PaymentAccountIDs []uint32
}
