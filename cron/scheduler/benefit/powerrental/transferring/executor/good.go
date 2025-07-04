package executor

import (
	"context"
	"encoding/json"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/transferring/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	platformaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	txmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/tx"
	goodcoinrewardmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	schedcommon "github.com/NpoolPlatform/kunman/pkg/common"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type coinReward struct {
	types.CoinReward
	toPlatformAmount decimal.Decimal
}

type goodHandler struct {
	*powerrentalmwpb.PowerRental
	persistent             chan interface{}
	notif                  chan interface{}
	done                   chan interface{}
	goodCoins              map[string]*coinmwpb.Coin
	newBenefitState        goodtypes.BenefitState
	coinRewards            []*coinReward
	rewardTxs              map[string]*txmwpb.Tx
	userBenefitHotAccounts map[string]*platformaccountmwpb.Account
	platformColdAccounts   map[string]*platformaccountmwpb.Account
	benefitResult          basetypes.Result
}

const (
	errorInvalidTx = "Invalid transaction"
	errorTxFail    = "Transaction fail"
)

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
			if reward.RewardTID == uuid.Nil.String() {
				continue
			}
			txIDs = append(txIDs, reward.RewardTID)
		}
		return
	}())
	return wlog.WrapError(err)
}

func (h *goodHandler) constructCoinRewards(ctx context.Context) error { //nolint:gocognit
	h.newBenefitState = goodtypes.BenefitState_BenefitBookKeeping

	successes := 0
	fails := 0

	for _, reward := range h.Rewards {
		able, err := h.checkTransferred(reward)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !able {
			continue
		}

		coinReward := &coinReward{
			CoinReward: types.CoinReward{
				CoinTypeID: reward.CoinTypeID,
			},
		}

		tx, ok := h.rewardTxs[reward.RewardTID]
		if !ok {
			h.benefitResult = basetypes.Result_Fail
			h.newBenefitState = goodtypes.BenefitState_BenefitFail
			coinReward.BenefitMessage = fmt.Sprintf("%v (%v)", errorInvalidTx, reward.RewardTID)
			h.coinRewards = append(h.coinRewards, coinReward)
			fails++
			continue
		}
		switch tx.State {
		case basetypes.TxState_TxStateCreated:
			fallthrough //nolint
		case basetypes.TxState_TxStateCreatedCheck:
			fallthrough //nolint
		case basetypes.TxState_TxStateWait:
			fallthrough //nolint
		case basetypes.TxState_TxStateWaitCheck:
			fallthrough //nolint
		case basetypes.TxState_TxStateTransferring:
			h.newBenefitState = h.RewardState
			return nil
		case basetypes.TxState_TxStateFail:
			// If we have some transaction fail, we just go ahead with some notification
			// Following steps should check tx state when they update benefit info
			h.benefitResult = basetypes.Result_Fail
			h.newBenefitState = goodtypes.BenefitState_BenefitFail
			coinReward.BenefitMessage = fmt.Sprintf(
				"%v %v@%v(%v)",
				errorTxFail,
				tx.ChainTxID,
				h.LastRewardAt,
				reward.RewardTID,
			)
			fails++
		case basetypes.TxState_TxStateSuccessful:
			coinReward.Transferrable = true
			successes++
		}

		p := struct {
			PlatformReward      decimal.Decimal
			TechniqueServiceFee decimal.Decimal
		}{}
		if err := json.Unmarshal([]byte(tx.Extra), &p); err != nil {
			return err
		}

		coinReward.Extra = tx.Extra
		coinReward.toPlatformAmount = p.PlatformReward.Add(p.TechniqueServiceFee)
		coinReward.ToPlatformAmount = coinReward.toPlatformAmount.String()

		userBenefitHotAccount, ok := h.userBenefitHotAccounts[reward.CoinTypeID]
		if ok {
			coinReward.UserBenefitHotAccountID = userBenefitHotAccount.AccountID
			coinReward.UserBenefitHotAddress = userBenefitHotAccount.Address
		}
		platformColdAccount, ok := h.platformColdAccounts[reward.CoinTypeID]
		if ok {
			coinReward.PlatformColdAccountID = platformColdAccount.AccountID
			coinReward.PlatformColdAddress = platformColdAccount.Address
		}

		able, err = h.checkTransferrableToPlatform(ctx, coinReward)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !able {
			h.coinRewards = append(h.coinRewards, coinReward)
			continue
		}

		h.coinRewards = append(h.coinRewards, coinReward)
	}
	if fails > 0 && successes > 0 {
		h.newBenefitState = h.RewardState
	}
	return nil
}

func (h *goodHandler) checkTransferrableToPlatform(ctx context.Context, reward *coinReward) (bool, error) {
	coin, ok := h.goodCoins[reward.CoinTypeID]
	if !ok {
		return false, wlog.Errorf("invalid goodcoin")
	}
	least, err := decimal.NewFromString(coin.LeastTransferAmount)
	if err != nil {
		return false, err
	}
	if least.Cmp(decimal.NewFromInt(0)) <= 0 {
		return false, wlog.Errorf("invalid leasttransferamount")
	}
	if reward.toPlatformAmount.Cmp(least) <= 0 {
		return false, nil
	}

	bal, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    coin.Name,
		Address: reward.UserBenefitHotAddress,
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	if bal == nil {
		return false, wlog.Errorf("invalid balance")
	}

	balance, err := decimal.NewFromString(bal.BalanceStr)
	if err != nil {
		return false, err
	}
	reserved, err := decimal.NewFromString(coin.ReservedAmount)
	if err != nil {
		return false, err
	}
	if balance.Cmp(reward.toPlatformAmount.Add(reserved)) < 0 {
		return false, nil
	}
	return true, nil
}

func (h *goodHandler) getUserBenefitHotAccounts(ctx context.Context) (err error) {
	h.userBenefitHotAccounts, err = schedcommon.GetCoinPlatformAccounts(
		ctx,
		basetypes.AccountUsedFor_UserBenefitHot,
		func() (coinTypeIDs []string) {
			for coinTypeID := range h.goodCoins {
				coinTypeIDs = append(coinTypeIDs, coinTypeID)
			}
			return
		}(),
	)
	return wlog.WrapError(err)
}

func (h *goodHandler) getPlatformColdAccounts(ctx context.Context) (err error) {
	h.platformColdAccounts, err = schedcommon.GetCoinPlatformAccounts(
		ctx,
		basetypes.AccountUsedFor_PlatformBenefitCold,
		func() (coinTypeIDs []string) {
			for coinTypeID := range h.goodCoins {
				coinTypeIDs = append(coinTypeIDs, coinTypeID)
			}
			return
		}(),
	)
	return wlog.WrapError(err)
}

func (h *goodHandler) checkTransferred(reward *goodcoinrewardmwpb.RewardInfo) (bool, error) {
	coin, ok := h.goodCoins[reward.CoinTypeID]
	if !ok {
		return false, wlog.Errorf("invalid goodcoin")
	}
	least, err := decimal.NewFromString(coin.LeastTransferAmount)
	if err != nil {
		return false, err
	}
	if least.Cmp(decimal.NewFromInt(0)) <= 0 {
		return false, wlog.Errorf("invalid leasttransferamount")
	}
	todayRewardAmount, err := decimal.NewFromString(reward.LastRewardAmount)
	if err != nil {
		return false, err
	}
	if todayRewardAmount.Cmp(least) <= 0 {
		return false, nil
	}
	return true, nil
}

//nolint:gocritic
func (h *goodHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRental", h.PowerRental,
			"NewBenefitState", h.newBenefitState,
			"BenefitResult", h.benefitResult,
			"CoinReward", h.coinRewards,
			"Error", *err,
		)
	}
	persistentGood := &types.PersistentPowerRental{
		PowerRental:     h.PowerRental,
		NewBenefitState: h.newBenefitState,
		BenefitResult:   h.benefitResult,
		CoinRewards: func() (rewards []*types.CoinReward) {
			for _, reward := range h.coinRewards {
				rewards = append(rewards, &reward.CoinReward)
			}
			return
		}(),
		Error: *err,
	}

	if persistentGood.BenefitResult == basetypes.Result_Fail {
		asyncfeed.AsyncFeed(ctx, persistentGood, h.notif)
	}
	if h.newBenefitState == h.RewardState && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentGood, h.done)
		return
	}
	if *err != nil {
		persistentGood.BenefitResult = basetypes.Result_Fail
		asyncfeed.AsyncFeed(ctx, persistentGood, h.notif)
	}
	if h.newBenefitState != h.RewardState {
		asyncfeed.AsyncFeed(ctx, persistentGood, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentGood, h.done)
}

//nolint:gocritic
func (h *goodHandler) exec(ctx context.Context) error {
	h.newBenefitState = h.RewardState
	h.benefitResult = basetypes.Result_Success

	var err error
	defer h.final(ctx, &err)

	if err = h.getGoodCoins(ctx); err != nil {
		return err
	}
	if err := h.getUserBenefitHotAccounts(ctx); err != nil {
		return err
	}
	if err := h.getPlatformColdAccounts(ctx); err != nil {
		return err
	}
	if err := h.getRewardTxs(ctx); err != nil {
		return err
	}
	if err = h.constructCoinRewards(ctx); err != nil {
		return err
	}

	return nil
}
