package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/created/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	pltfaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
	reviewtypes "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	pltfaccmw "github.com/NpoolPlatform/kunman/middleware/account/platform"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/shopspring/decimal"
)

type withdrawHandler struct {
	*withdrawmwpb.Withdraw
	persistent                chan interface{}
	notif                     chan interface{}
	done                      chan interface{}
	withdrawAmount            decimal.Decimal
	reviewTrigger             reviewtypes.ReviewTriggerType
	userBenefitHotAccount     *pltfaccmwpb.Account
	userBenefitHotBalance     decimal.Decimal
	userBenefitHotFeeBalance  decimal.Decimal
	appCoin                   *appcoinmwpb.Coin
	feeCoin                   *coinmwpb.Coin
	autoReviewThresholdAmount decimal.Decimal
	coinReservedAmount        decimal.Decimal
	lowFeeAmount              decimal.Decimal
}

func (h *withdrawHandler) getAppCoin(ctx context.Context) error {
	conds := &appcoinmwpb.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.CoinTypeID},
		Disabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}
	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	coin, err := handler.GetCoinOnly(ctx)
	if err != nil {
		return err
	}
	if coin == nil {
		return fmt.Errorf("invalid coin")
	}
	h.appCoin = coin
	return nil
}

func (h *withdrawHandler) getFeeCoin(ctx context.Context) error {
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithEntID(&h.appCoin.FeeCoinTypeID, true),
	)
	if err != nil {
		return err
	}

	coin, err := handler.GetCoin(ctx)
	if err != nil {
		return err
	}
	if coin == nil {
		return fmt.Errorf("invalid coin")
	}
	h.feeCoin = coin
	return nil
}

func (h *withdrawHandler) getUserBenefitHotAccount(ctx context.Context) error {
	conds := &pltfaccmwpb.Conds{
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.CoinTypeID},
		UsedFor:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(basetypes.AccountUsedFor_UserBenefitHot)},
		Active:     &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		Backup:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Blocked:    &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}
	handler, err := pltfaccmw.NewHandler(
		ctx,
		pltfaccmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	account, err := handler.GetAccountOnly(ctx)
	if err != nil {
		return err
	}
	if account == nil {
		return fmt.Errorf("invalid account")
	}
	h.userBenefitHotAccount = account
	return nil
}

func (h *withdrawHandler) checkUserBenefitHotBalance(ctx context.Context) error {
	bal, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.appCoin.Name,
		Address: h.userBenefitHotAccount.Address,
	})
	if err != nil {
		return err
	}
	if bal == nil {
		return fmt.Errorf("invalid balance")
	}
	h.userBenefitHotBalance, err = decimal.NewFromString(bal.BalanceStr)
	if err != nil {
		return err
	}
	return nil
}

func (h *withdrawHandler) checkUserBenefitHotFeeBalance(ctx context.Context) error {
	bal, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    h.appCoin.FeeCoinName,
		Address: h.userBenefitHotAccount.Address,
	})
	if err != nil {
		return err
	}
	if bal == nil {
		return fmt.Errorf("invalid fee balance")
	}
	h.userBenefitHotFeeBalance, err = decimal.NewFromString(bal.BalanceStr)
	if err != nil {
		return err
	}
	return nil
}

func (h *withdrawHandler) resolveReviewTrigger() {
	h.reviewTrigger = reviewtypes.ReviewTriggerType_AutoReviewed
	if h.userBenefitHotBalance.Cmp(h.withdrawAmount.Add(h.coinReservedAmount)) <= 0 {
		h.reviewTrigger = reviewtypes.ReviewTriggerType_InsufficientFunds
	}
	if h.userBenefitHotFeeBalance.Cmp(h.lowFeeAmount) < 0 {
		switch h.reviewTrigger {
		case reviewtypes.ReviewTriggerType_InsufficientFunds:
			h.reviewTrigger = reviewtypes.ReviewTriggerType_InsufficientFundsGas
		case reviewtypes.ReviewTriggerType_AutoReviewed:
			h.reviewTrigger = reviewtypes.ReviewTriggerType_InsufficientGas
		}
		return
	}
	if h.autoReviewThresholdAmount.Cmp(h.withdrawAmount) < 0 {
		h.reviewTrigger = reviewtypes.ReviewTriggerType_LargeAmount
	}
}

//nolint:gocritic
func (h *withdrawHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Withdraw", h.Withdraw,
			"ReviewTrigger", h.reviewTrigger,
			"Error", *err,
		)
	}
	persistentWithdraw := &types.PersistentWithdraw{
		Withdraw:      h.Withdraw,
		ReviewTrigger: h.reviewTrigger,
		Error:         *err,
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.persistent)
		return
	}
	// TODO: notif to administrator
	asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.done)
}

//nolint:gocritic
func (h *withdrawHandler) exec(ctx context.Context) error {
	var err error
	defer h.final(ctx, &err)

	h.withdrawAmount, err = decimal.NewFromString(h.Amount)
	if err != nil {
		return err
	}
	if err = h.getAppCoin(ctx); err != nil {
		return err
	}
	if err = h.getFeeCoin(ctx); err != nil {
		return err
	}
	h.autoReviewThresholdAmount, err = decimal.NewFromString(h.appCoin.WithdrawAutoReviewAmount)
	if err != nil {
		return err
	}
	h.coinReservedAmount, err = decimal.NewFromString(h.appCoin.ReservedAmount)
	if err != nil {
		return err
	}
	h.lowFeeAmount, err = decimal.NewFromString(h.feeCoin.LowFeeAmount)
	if err != nil {
		return err
	}
	if err = h.getUserBenefitHotAccount(ctx); err != nil {
		return err
	}
	if err = h.checkUserBenefitHotBalance(ctx); err != nil {
		return err
	}
	if err = h.checkUserBenefitHotFeeBalance(ctx); err != nil {
		return err
	}
	h.resolveReviewTrigger()

	return nil
}
