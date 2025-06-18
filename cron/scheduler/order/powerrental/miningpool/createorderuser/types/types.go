package types

import (
	orderusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type PersistentOrder struct {
	*powerrentalordermwpb.PowerRentalOrder
	PowerRentalOrderReq *powerrentalordermwpb.PowerRentalOrderReq
	OrderUserReq        *orderusermwpb.OrderUserReq
}
