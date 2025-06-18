package types

import (
	fractionwithdrawalmwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawal"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
)

type PersistentOrder struct {
	*powerrentalordermwpb.PowerRentalOrder
	FractionWithdrawalReqs []*fractionwithdrawalmwpb.FractionWithdrawalReq
	PowerRentalOrderReq    *powerrentalordermwpb.PowerRentalOrderReq
}
