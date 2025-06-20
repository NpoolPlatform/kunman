package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/reviewing/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	pltfaccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
	useraccmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/user"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	reviewtypes "github.com/NpoolPlatform/kunman/message/basetypes/review/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	withdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	pltfaccmw "github.com/NpoolPlatform/kunman/middleware/account/platform"
	useraccmw "github.com/NpoolPlatform/kunman/middleware/account/user"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	reviewmw "github.com/NpoolPlatform/kunman/middleware/review/review"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type withdrawHandler struct {
	*withdrawmwpb.Withdraw
	persistent                chan interface{}
	notif                     chan interface{}
	done                      chan interface{}
	withdrawAmount            decimal.Decimal
	newWithdrawState          ledgertypes.WithdrawState
	newReviewState            reviewtypes.ReviewState
	userBenefitHotAccount     *pltfaccmwpb.Account
	userBenefitHotBalance     decimal.Decimal
	userBenefitHotFeeBalance  decimal.Decimal
	appCoin                   *appcoinmwpb.Coin
	feeCoin                   *coinmwpb.Coin
	autoReviewThresholdAmount decimal.Decimal
	coinReservedAmount        decimal.Decimal
	lowFeeAmount              decimal.Decimal
	needUpdateReview          bool
}

func (h *withdrawHandler) checkWithdrawReview(ctx context.Context) error {
	if _, err := uuid.Parse(h.ReviewID); err != nil {
		h.newWithdrawState = ledgertypes.WithdrawState_PreRejected
		return err
	}

	handler, err := reviewmw.NewHandler(
		ctx,
		reviewmw.WithEntID(&h.ReviewID, true),
	)
	if err != nil {
		return err
	}

	review, err := handler.GetReview(ctx)
	if err != nil {
		return err
	}
	if review == nil {
		h.newWithdrawState = ledgertypes.WithdrawState_PreRejected
		return fmt.Errorf("invalid review")
	}
	switch review.State {
	case reviewtypes.ReviewState_Approved:
		h.newWithdrawState = ledgertypes.WithdrawState_Approved
	case reviewtypes.ReviewState_Rejected:
		h.newWithdrawState = ledgertypes.WithdrawState_PreRejected
	}
	return nil
}

func (h *withdrawHandler) checkWithdrawAccount(ctx context.Context) error {
	conds := &useraccmwpb.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: h.UserID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: h.CoinTypeID},
		AccountID:  &basetypes.StringVal{Op: cruder.EQ, Value: h.AccountID},
		Address:    &basetypes.StringVal{Op: cruder.EQ, Value: h.Address},
	}
	handler, err := useraccmw.NewHandler(
		ctx,
		useraccmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistAccountConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		h.newWithdrawState = ledgertypes.WithdrawState_PreRejected
		return fmt.Errorf("invalid account")
	}
	return nil
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

func (h *withdrawHandler) checkWithdrawReviewState() error {
	if h.userBenefitHotBalance.Cmp(h.withdrawAmount.Add(h.coinReservedAmount)) <= 0 {
		return fmt.Errorf("insufficient funds")
	}
	if h.userBenefitHotFeeBalance.Cmp(h.lowFeeAmount) < 0 {
		return fmt.Errorf("insufficient gas")
	}
	if h.autoReviewThresholdAmount.Cmp(h.withdrawAmount) < 0 {
		return nil
	}
	h.newWithdrawState = ledgertypes.WithdrawState_Approved
	h.newReviewState = reviewtypes.ReviewState_Approved
	h.needUpdateReview = true
	return nil
}

//nolint:gocritic
func (h *withdrawHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Withdraw", h.Withdraw,
			"NewWithdrawState", h.newWithdrawState,
			"NewReviewState", h.newReviewState,
			"Error", *err,
		)
	}

	persistentWithdraw := &types.PersistentWithdraw{
		Withdraw:         h.Withdraw,
		NewWithdrawState: h.newWithdrawState,
		NewReviewState:   h.newReviewState,
		NeedUpdateReview: h.needUpdateReview,
		Error:            *err,
	}

	if h.newWithdrawState == h.State && *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.done)
		return
	}
	if *err != nil {
		asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.notif)
	}
	if h.newWithdrawState != h.State {
		asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentWithdraw, h.done)
}

//nolint:gocritic
func (h *withdrawHandler) exec(ctx context.Context) error {
	h.newWithdrawState = h.State

	var err error
	defer h.final(ctx, &err)

	if err = h.checkWithdrawReview(ctx); err != nil {
		return err
	}
	h.withdrawAmount, err = decimal.NewFromString(h.Amount)
	if err != nil {
		return err
	}
	if err = h.checkWithdrawAccount(ctx); err != nil {
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
	if err = h.checkWithdrawReviewState(); err != nil {
		return err
	}

	return nil
}
