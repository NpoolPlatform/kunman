package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/done/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	schedcommon "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

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
	goodCoins       map[string]*coinmwpb.Coin
	rewardTxs       map[string]*txmwpb.Tx
	benefitOrderIDs []uint32
	coinNextRewards []*coinNextReward
}

const resultMinimumReward = "Mining reward not transferred"

func (h *goodHandler) getGoodCoins(ctx context.Context) (err error) {
	h.goodCoins, err = schedcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, goodCoin := range h.GoodCoins {
			coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
		}
		return
	}())
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, goodCoin := range h.GoodCoins {
		if _, ok := h.goodCoins[goodCoin.CoinTypeID]; !ok {
			return wlog.Errorf("invalid goodcoin")
		}
	}
	return nil
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
		coinNextReward := &coinNextReward{
			CoinNextReward: types.CoinNextReward{
				CoinTypeID:            reward.CoinTypeID,
				NextRewardStartAmount: reward.NextRewardStartAmount,
			},
			lastRewardAmount: lastRewardAmount,
		}
		transferred, err := h.checkLeastTransferAmount(coinNextReward)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !transferred {
			h.coinNextRewards = append(h.coinNextRewards, coinNextReward)
			continue
		}
		rewardTx, ok := h.rewardTxs[reward.RewardTID]
		if !ok {
			return wlog.Errorf("invalid rewardtx")
		}
		coinNextReward.BenefitMessage = fmt.Sprintf(
			"%v@%v(%v)",
			rewardTx.ChainTxID,
			h.LastRewardAt,
			reward.RewardTID,
		)
		nextRewardStartAmount, err := decimal.NewFromString(reward.NextRewardStartAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		coinNextReward.NextRewardStartAmount = nextRewardStartAmount.Sub(lastRewardAmount).String()
		h.coinNextRewards = append(h.coinNextRewards, coinNextReward)
	}
	return nil
}

func (h *goodHandler) checkLeastTransferAmount(reward *coinNextReward) (bool, error) {
	coin, ok := h.goodCoins[reward.CoinTypeID]
	if !ok {
		return false, wlog.Errorf("invalid coin")
	}
	least, err := decimal.NewFromString(coin.LeastTransferAmount)
	if err != nil {
		return false, err
	}
	if least.Cmp(decimal.NewFromInt(0)) <= 0 {
		return false, wlog.Errorf("invalid leasttransferamount")
	}
	if reward.lastRewardAmount.Cmp(least) <= 0 {
		reward.BenefitMessage = fmt.Sprintf(
			"%v (coin %v, reward amount %v, least transfer amount %v [#%v])",
			resultMinimumReward,
			coin.Name,
			reward.lastRewardAmount,
			least,
			h.LastRewardAt,
		)
		return false, nil
	}
	return true, nil
}

func (h *goodHandler) getBenefitOrders(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &powerrentalordermwpb.Conds{
		GoodID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
		LastBenefitAt: &basetypes.Uint32Val{Op: cruder.EQ, Value: h.LastRewardAt},
	}

	for {
		handler, err := powerrentalordermw.NewHandler(
			ctx,
			powerrentalordermw.WithConds(conds),
			powerrentalordermw.WithOffset(offset),
			powerrentalordermw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		orders, _, err := handler.GetPowerRentals(ctx)
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
		Error:           *err,
	}
	if *err == nil {
		persistentGood.BenefitResult = basetypes.Result_Success
		asyncfeed.AsyncFeed(ctx, persistentGood, h.persistent)
		return
	}
	persistentGood.BenefitResult = basetypes.Result_Fail
	asyncfeed.AsyncFeed(ctx, persistentGood, h.notif)
	asyncfeed.AsyncFeed(ctx, persistentGood, h.done)
}

//nolint
func (h *goodHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	if err = h.getGoodCoins(ctx); err != nil {
		return err
	}
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
