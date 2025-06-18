package types

import (
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type PersistentPowerRentalOrder struct {
	*powerrentalordermwpb.PowerRentalOrder
	NewPaymentState *ordertypes.PaymentState
}
