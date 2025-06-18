package types

import (
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	orderrenewpb "github.com/NpoolPlatform/kunman/message/scheduler/middleware/v1/order/renew"
)

type PersistentOrder struct {
	*powerrentalordermwpb.PowerRentalOrder
	*orderrenewpb.MsgOrderChildsRenewReq
	NewRenewState ordertypes.OrderRenewState
}
