package currencyhistory

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	currencymwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat/currency"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat/currency/history"
	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"

	coin1 "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	currency1 "github.com/NpoolPlatform/kunman/middleware/chain/coin/fiat/currency"
	fiat1 "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

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

var ret = &currencymwpb.Currency{
	CoinName:        "My BTC1",
	CoinLogo:        uuid.NewString(),
	CoinUnit:        "BTC",
	CoinENV:         "test",
	FiatLogo:        uuid.NewString(),
	FiatUnit:        "USD",
	MarketValueHigh: "0.000000000000000000",
	MarketValueLow:  "0.000000000000000000",
	FeedType:        basetypes.CurrencyFeedType_CoinGecko,
	FeedTypeStr:     basetypes.CurrencyFeedType_CoinGecko.String(),
}

var req = &currencymwpb.CurrencyReq{
	FeedType:        &ret.FeedType,
	MarketValueHigh: &ret.MarketValueHigh,
	MarketValueLow:  &ret.MarketValueLow,
}

func setupCoin(t *testing.T) func(*testing.T) {
	ret.CoinTypeID = uuid.NewString()
	req.CoinTypeID = &ret.CoinTypeID
	ret.CoinName = uuid.NewString()
	ret.FiatName = uuid.NewString()
	ret.FiatID = uuid.NewString()
	req.FiatID = &ret.FiatID

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

	h2, err := fiat1.NewHandler(
		context.Background(),
		fiat1.WithEntID(&ret.FiatID, true),
		fiat1.WithName(&ret.FiatName, true),
		fiat1.WithLogo(&ret.FiatLogo, true),
		fiat1.WithUnit(&ret.FiatUnit, true),
	)
	assert.Nil(t, err)

	_, err = h2.CreateFiat(context.Background())
	assert.Nil(t, err)

	h3, err := currency1.NewHandler(
		context.Background(),
		currency1.WithCoinTypeID(req.CoinTypeID, true),
		currency1.WithFiatID(req.FiatID, true),
		currency1.WithMarketValueHigh(req.MarketValueHigh, true),
		currency1.WithMarketValueLow(req.MarketValueLow, true),
		currency1.WithFeedType(req.FeedType, true),
	)
	assert.Nil(t, err)

	_, err = h3.CreateCurrency(context.Background())
	assert.Nil(t, err)

	_, err = h3.CreateCurrency(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
	}
}

func getMany(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			CoinTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
			CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.CoinTypeID}},
		}),
		WithOffset(0),
		WithLimit(100),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetCurrencies(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, uint32(2), total)
	assert.Equal(t, 2, len(infos))
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoin(t)
	defer teardown(t)

	t.Run("getMany", getMany)
}
