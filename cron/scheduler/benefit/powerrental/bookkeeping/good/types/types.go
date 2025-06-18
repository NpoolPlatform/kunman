package types

import (
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
)

type CoinReward struct {
	CoinTypeID         string
	TotalRewardAmount  string
	UnsoldRewardAmount string
	TechniqueFeeAmount string
	StatementExist     bool
	BenefitMessage     string
}

type PersistentGood struct {
	*powerrentalmwpb.PowerRental
	CoinRewards   []*CoinReward
	BenefitResult basetypes.Result
	Error         error
}
