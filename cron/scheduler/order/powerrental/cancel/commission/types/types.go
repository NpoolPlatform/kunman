package types

import (
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type CommissionRevoke struct {
	LockID       string
	IOExtra      string
	StatementIDs []string
}

type PersistentPowerRentalOrder struct {
	*powerrentalordermwpb.PowerRentalOrder
	CommissionRevokes []*CommissionRevoke
}
