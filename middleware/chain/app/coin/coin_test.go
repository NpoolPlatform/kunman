package appcoin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"

	coin1 "github.com/NpoolPlatform/kunman/middleware/chain/coin"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	coinUnit = "BTC"
	coinENV  = "test"
)

var ret = &npool.Coin{
	AppID:                       uuid.NewString(),
	Name:                        "My BTC1",
	DisplayNames:                []string{"123123", "2323"},
	Logo:                        uuid.NewString(),
	Unit:                        coinUnit,
	ENV:                         coinENV,
	ReservedAmount:              "0.000000000000000000",
	WithdrawFeeByStableUSD:      true,
	WithdrawFeeAmount:           "0.000000000000000000",
	CollectFeeAmount:            "0.000000000000000000",
	HotWalletFeeAmount:          "0.000000000000000000",
	LowFeeAmount:                "0.000000000000000000",
	HotLowFeeAmount:             "0.000000000000000000",
	HotWalletAccountAmount:      "0.000000000000000000",
	PaymentAccountCollectAmount: "0.000000000000000000",
	FeeCoinUnit:                 coinUnit,
	FeeCoinENV:                  coinENV,
	WithdrawAutoReviewAmount:    "0.000000000000000000",
	MarketValue:                 "0.000000000000000000",
	SettleValue:                 "0.000000000000000000",
	SettlePercent:               90,
	SettleTips:                  []string{"1213", "234234"},
	Setter:                      uuid.NewString(),
	Display:                     true,
	DailyRewardAmount:           "0.000000000000000000",
	MaxAmountPerWithdraw:        "0.000000000000000000",
	LeastTransferAmount:         "0.000000000000000000",
}

var req = &npool.CoinReq{
	AppID:         &ret.AppID,
	Name:          &ret.Name,
	DisplayNames:  ret.DisplayNames,
	Logo:          &ret.Logo,
	SettlePercent: &ret.SettlePercent,
	SettleTips:    ret.SettleTips,
	Setter:        &ret.Setter,
}

func setupCoin(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.CoinName = uuid.NewString()
	ret.FeeCoinName = ret.CoinName

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithEntID(&ret.CoinTypeID, true),
		coin1.WithName(&ret.CoinName, true),
		coin1.WithUnit(&coinUnit, true),
		coin1.WithENV(&coinENV, true),
		coin1.WithChainNativeUnit(&coinUnit, true),
	)
	assert.Nil(t, err)

	_, err = h1.CreateCoin(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
	}
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, false),
		WithAppID(req.AppID, true),
		WithCoinTypeID(req.CoinTypeID, true),
		WithName(req.Name, true),
		WithDisplayNames(req.DisplayNames, false),
		WithLogo(req.Logo, true),
		WithForPay(req.ForPay, false),
		WithProductPage(req.ProductPage, false),
		WithWithdrawAutoReviewAmount(req.WithdrawAutoReviewAmount, false),
		WithDailyRewardAmount(req.DailyRewardAmount, false),
		WithDisplay(req.Display, false),
		WithDisplayIndex(req.DisplayIndex, false),
		WithMaxAmountPerWithdraw(req.MaxAmountPerWithdraw, false),
		WithMarketValue(req.MarketValue, false),
		WithSettlePercent(req.SettlePercent, false),
		WithSettleTips(req.SettleTips, false),
		WithSetter(req.Setter, false),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoin(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		ret.FeeCoinTypeID = ret.CoinTypeID
		ret.DisplayNamesStr = info.DisplayNamesStr
		ret.SettleTipsStr = info.SettleTipsStr
		ret.CheckNewAddressBalance = info.CheckNewAddressBalance
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	amount := "123.700000000000000000"
	logo := uuid.NewString()

	ret.Logo = logo
	ret.WithdrawAutoReviewAmount = amount
	ret.MarketValue = amount
	ret.SettleValue = "111.330000000000000000"

	req.ID = &ret.ID
	req.Logo = &logo
	req.WithdrawAutoReviewAmount = &amount
	req.MarketValue = &amount

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID, true),
		WithName(req.Name, false),
		WithDisplayNames(req.DisplayNames, false),
		WithLogo(req.Logo, false),
		WithForPay(req.ForPay, false),
		WithProductPage(req.ProductPage, false),
		WithWithdrawAutoReviewAmount(req.WithdrawAutoReviewAmount, false),
		WithDailyRewardAmount(req.DailyRewardAmount, false),
		WithDisplay(req.Display, false),
		WithDisplayIndex(req.DisplayIndex, false),
		WithMaxAmountPerWithdraw(req.MaxAmountPerWithdraw, false),
		WithMarketValue(req.MarketValue, false),
		WithSettlePercent(req.SettlePercent, false),
		WithSettleTips(req.SettleTips, false),
		WithSetter(req.Setter, false),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoin(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func _delete(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	_, err = handler.DeleteCoin(context.Background())
	assert.Nil(t, err)
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoin(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
	t.Run("delete", _delete)
}
