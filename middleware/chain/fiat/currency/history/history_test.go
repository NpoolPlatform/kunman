package currencyhistory

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	currencymwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency/history"
	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"

	fiat1 "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
	currency1 "github.com/NpoolPlatform/kunman/middleware/chain/fiat/currency"
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
	FiatName:        "My BTC1",
	FiatLogo:        uuid.NewString(),
	FiatUnit:        "BTC",
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

func setupFiat(t *testing.T) func(*testing.T) {
	ret.FiatID = uuid.NewString()
	req.FiatID = &ret.FiatID
	ret.FiatName = uuid.NewString()

	h1, err := fiat1.NewHandler(
		context.Background(),
		fiat1.WithEntID(&ret.FiatID, true),
		fiat1.WithName(&ret.FiatName, true),
		fiat1.WithLogo(&ret.FiatLogo, true),
		fiat1.WithUnit(&ret.FiatUnit, true),
	)
	assert.Nil(t, err)

	_, err = h1.CreateFiat(context.Background())
	assert.Nil(t, err)

	h2, err := currency1.NewHandler(
		context.Background(),
		currency1.WithFiatID(req.FiatID, true),
		currency1.WithMarketValueHigh(req.MarketValueHigh, true),
		currency1.WithMarketValueLow(req.MarketValueLow, true),
		currency1.WithFeedType(req.FeedType, true),
	)
	assert.Nil(t, err)

	_, err = h2.CreateCurrency(context.Background())
	assert.Nil(t, err)

	_, err = h2.CreateCurrency(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {}
}

func getMany(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			FiatID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.FiatID},
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

func TestFiat(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupFiat(t)
	defer teardown(t)

	t.Run("getMany", getMany)
}
