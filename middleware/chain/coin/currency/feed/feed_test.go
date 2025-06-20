package currencyfeed

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency/feed"
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

var ret = &npool.Feed{
	CoinName:     "My BTC1",
	CoinLogo:     uuid.NewString(),
	CoinUnit:     "BTC",
	CoinENV:      "test",
	FeedType:     basetypes.CurrencyFeedType_CoinGecko,
	FeedTypeStr:  basetypes.CurrencyFeedType_CoinGecko.String(),
	FeedCoinName: "BTC123",
}

var req = &npool.FeedReq{
	FeedType:     &ret.FeedType,
	FeedCoinName: &ret.FeedCoinName,
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
		WithFeedType(req.FeedType, true),
		WithFeedCoinName(req.FeedCoinName, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateFeed(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.FeedCoinName = "BTCWWWW"
	ret.Disabled = true

	req.FeedCoinName = &ret.FeedCoinName
	req.Disabled = &ret.Disabled

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithFeedCoinName(req.FeedCoinName, true),
		WithDisabled(req.Disabled, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateFeed(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
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
}
