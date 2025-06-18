package types

import (
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
)

type CoinReward struct {
	CoinTypeID              string
	Amount                  string
	NextRewardStartAmount   string
	GoodBenefitAccountID    string
	GoodBenefitAddress      string
	UserBenefitHotAccountID string
	UserBenefitHotAddress   string
	Extra                   string
	BenefitMessage          string
	Transferrable           bool
}

type PersistentPowerRental struct {
	*powerrentalmwpb.PowerRental
	BenefitOrderIDs  []uint32
	CoinRewards      []*CoinReward
	BenefitTimestamp uint32
	BenefitResult    basetypes.Result
	Error            error
}

type FeedPowerRental struct {
	*powerrentalmwpb.PowerRental
	TriggerBenefitTimestamp uint32
}

type TriggerCond struct {
	GoodIDs  []string
	RewardAt uint32
}

func (c *TriggerCond) ContainGoodID(goodID string) bool {
	for _, id := range c.GoodIDs {
		if id == goodID {
			return true
		}
	}
	return false
}
