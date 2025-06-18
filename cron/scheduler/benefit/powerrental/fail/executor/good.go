package executor

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/fail/types"
	schedcommon "github.com/NpoolPlatform/kunman/cron/scheduler/common"
	constant "github.com/NpoolPlatform/kunman/cron/scheduler/const"
	powerrentalordermwcli "github.com/NpoolPlatform/order-middleware/pkg/client/powerrental"

	"github.com/shopspring/decimal"
)

type coinNextReward struct {
	types.CoinNextReward
	lastRewardAmount decimal.Decimal
}

type goodHandler struct {
	*powerrentalmwpb.PowerRental
	persistent      chan interface{}
	notif           chan interface{}
	done            chan interface{}
	rewardTxs       map[string]*txmwpb.Tx
	benefitOrderIDs []uint32
	coinNextRewards []*coinNextReward
}

func (h *goodHandler) getRewardTxs(ctx context.Context) (err error) {
	h.rewardTxs, err = schedcommon.GetTxs(ctx, func() (txIDs []string) {
		for _, reward := range h.Rewards {
			txIDs = append(txIDs, reward.RewardTID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *goodHandler) calculateCoinNextRewardStartAmounts() error {
	for _, reward := range h.Rewards {
		lastRewardAmount, err := decimal.NewFromString(reward.LastRewardAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		nextRewardStartAmount, err := decimal.NewFromString(reward.NextRewardStartAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		coinNextReward := &coinNextReward{
			CoinNextReward: types.CoinNextReward{
				CoinTypeID:            reward.CoinTypeID,
				NextRewardStartAmount: nextRewardStartAmount.Sub(lastRewardAmount).String(),
			},
			lastRewardAmount: lastRewardAmount,
		}
		h.coinNextRewards = append(h.coinNextRewards, coinNextReward)
	}
	return nil
}

func (h *goodHandler) getBenefitOrders(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		orders, _, err := powerrentalordermwcli.GetPowerRentalOrders(ctx, &powerrentalordermwpb.Conds{
			GoodID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
			LastBenefitAt: &basetypes.Uint32Val{Op: cruder.EQ, Value: h.LastRewardAt},
		}, offset, limit)
		if err != nil {
			return err
		}
		if len(orders) == 0 {
			break
		}
		for _, order := range orders {
			h.benefitOrderIDs = append(h.benefitOrderIDs, order.ID)
		}
		offset += limit
	}
	return nil
}

//nolint:gocritic
func (h *goodHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRental", h.PowerRental,
			"RewardTxs", h.rewardTxs,
			"Error", *err,
		)
	}

	persistentGood := &types.PersistentGood{
		PowerRental: h.PowerRental,
		CoinNextRewards: func() (rewards []*types.CoinNextReward) {
			for _, reward := range h.coinNextRewards {
				rewards = append(rewards, &reward.CoinNextReward)
			}
			return
		}(),
		BenefitOrderIDs: h.benefitOrderIDs,
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentGood, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentGood, h.notif)
	asyncfeed.AsyncFeed(ctx, persistentGood, h.done)
}

//nolint
func (h *goodHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	if err = h.getRewardTxs(ctx); err != nil {
		return err
	}
	if err = h.getBenefitOrders(ctx); err != nil {
		return err
	}
	if err = h.calculateCoinNextRewardStartAmounts(); err != nil {
		return err
	}

	return nil
}
