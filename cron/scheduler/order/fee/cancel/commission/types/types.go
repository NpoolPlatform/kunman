package types

import (
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
)

type CommissionRevoke struct {
	LockID       string
	IOExtra      string
	StatementIDs []string
}

type PersistentFeeOrder struct {
	*feeordermwpb.FeeOrder
	CommissionRevokes []*CommissionRevoke
}
