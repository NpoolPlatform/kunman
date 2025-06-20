package currency

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency"
	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"

	coin1 "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

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

var ret = &npool.Currency{
	CoinName:        "My BTC1",
	CoinLogo:        uuid.NewString(),
	CoinUnit:        "BTC",
	CoinENV:         "test",
	MarketValueHigh: "0.000000000000000000",
	MarketValueLow:  "0.000000000000000000",
	FeedType:        basetypes.CurrencyFeedType_CoinGecko,
	FeedTypeStr:     basetypes.CurrencyFeedType_CoinGecko.String(),
}

var req = &npool.CurrencyReq{
	FeedType:        &ret.FeedType,
	MarketValueHigh: &ret.MarketValueHigh,
	MarketValueLow:  &ret.MarketValueLow,
}

func setupCoin(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.CoinName = uuid.NewString()

	h1, err := coin1.NewHandler(
		context.Background(),
		coin1.WithEntID(&ret.CoinTypeID, true),
		coin1.WithName(&ret.CoinName, true),
		coin1.WithLogo(&ret.CoinLogo, true),
		coin1.WithUnit(&ret.CoinUnit, true),
		coin1.WithENV(&ret.CoinENV, true),
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
		WithCoinTypeID(req.CoinTypeID, true),
		WithMarketValueHigh(req.MarketValueHigh, true),
		WithMarketValueLow(req.MarketValueLow, true),
		WithFeedType(req.FeedType, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCurrency(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func update(t *testing.T) {
	amount := "123.700000000000000000"

	ret.MarketValueHigh = amount
	ret.MarketValueLow = amount

	req.MarketValueHigh = &amount
	req.MarketValueLow = &amount

	handler, err := NewHandler(
		context.Background(),
		WithCoinTypeID(req.CoinTypeID, true),
		WithMarketValueHigh(req.MarketValueHigh, true),
		WithMarketValueLow(req.MarketValueLow, true),
		WithFeedType(req.FeedType, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCurrency(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func get(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCurrency(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getMany(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		}),
		WithOffset(0),
		WithLimit(100),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetCurrencies(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(infos))
		assert.Equal(t, uint32(1), total)
		assert.Equal(t, ret, infos[0])
	}
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoin(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
	t.Run("get", get)
	t.Run("getMany", getMany)
}
