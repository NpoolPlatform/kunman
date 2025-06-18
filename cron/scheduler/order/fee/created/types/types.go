package types

import (
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type PersistentFeeOrder struct {
	*feeordermwpb.FeeOrder
}
