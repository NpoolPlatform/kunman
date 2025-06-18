package types

import (
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
)

type CoinReward struct {
	CoinTypeID              string
	ToPlatformAmount        string
	UserBenefitHotAccountID string
	UserBenefitHotAddress   string
	PlatformColdAccountID   string
	PlatformColdAddress     string
	Extra                   string
	BenefitMessage          string
	Transferrable           bool
}

type PersistentPowerRental struct {
	*powerrentalmwpb.PowerRental
	NewBenefitState goodtypes.BenefitState
	CoinRewards     []*CoinReward
	BenefitResult   basetypes.Result
	Error           error
}
