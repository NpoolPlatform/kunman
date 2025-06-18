package types

import (
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type PersistentOrder struct {
	*powerrentalordermwpb.PowerRentalOrder
	NewRenewState      ordertypes.OrderRenewState
	NextRenewNotifyAt  uint32
	CreateOutOfGas     bool
	FeeEndAt           uint32
	OutOfGasEntID      string
	FinishOutOfGas     bool
	OutOfGasFinishedAt uint32
}
