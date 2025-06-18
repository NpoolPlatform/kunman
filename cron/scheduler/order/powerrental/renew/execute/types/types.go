package types

import (
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type PersistentOrder struct {
	*powerrentalordermwpb.PowerRentalOrder
	InsufficientBalance bool
	FeeOrderReqs        []*feeordermwpb.FeeOrderReq
	NewRenewState       ordertypes.OrderRenewState
	LedgerLockID        string
}
