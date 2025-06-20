package coinfiat

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/fiat"
	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"

	coin1 "github.com/NpoolPlatform/kunman/middleware/chain/coin"
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

var ret = &npool.CoinFiat{
	CoinName:    "My BTC1",
	CoinLogo:    uuid.NewString(),
	CoinUnit:    "BTC",
	CoinENV:     "test",
	FiatLogo:    uuid.NewString(),
	FiatUnit:    "USD",
	FeedType:    basetypes.CurrencyFeedType_CoinGecko,
	FeedTypeStr: basetypes.CurrencyFeedType_CoinGecko.String(),
}

var req = &npool.CoinFiatReq{
	FeedType: &ret.FeedType,
}

func setupCoinFiat(t *testing.T) func(*testing.T) {
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

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
	}
}

func createCoinFiat(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithCoinTypeID(req.CoinTypeID, true),
		WithFiatID(req.FiatID, true),
		WithFeedType(req.FeedType, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoinFiat(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func getCoinFiats(t *testing.T) {
	const singleRowLimit = 2
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		}),
		WithOffset(0),
		WithLimit(singleRowLimit),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetCoinFiats(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		assert.Equal(t, 1, len(infos))
		assert.Equal(t, ret, infos[0])
	}
}

func deleteCoinFiat(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCoinFiat(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoinFiat(t)
	defer teardown(t)

	t.Run("createCoinFiat", createCoinFiat)
	t.Run("getCoinFiats", getCoinFiats)
	t.Run("deleteCoinFiat", deleteCoinFiat)
}
