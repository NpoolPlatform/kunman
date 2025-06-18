package types

import (
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type PersistentOrder struct {
	*feeordermwpb.FeeOrder
	// ID of payment table but not account table
	PaymentAccountIDs []uint32
}
