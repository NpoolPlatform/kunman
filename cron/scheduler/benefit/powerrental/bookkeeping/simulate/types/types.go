package types

import (
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
)

type CoinReward struct {
	CoinTypeID string
	Amount     string
	SendCoupon bool
	Cashable   bool
}

type OrderReward struct {
	AppID       string
	UserID      string
	OrderID     string
	Extra       string
	CoinRewards []*CoinReward
}

type PersistentGood struct {
	*powerrentalmwpb.PowerRental
	OrderRewards    []*OrderReward
	BenefitResult   basetypes.Result
	BenefitMessage  string
	BenefitOrderIDs []uint32
	Error           error
}
