package currencyfeed

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency/feed"
	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"

	fiat1 "github.com/NpoolPlatform/kunman/middleware/chain/fiat"

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
	FiatName:     "My BTC1",
	FiatLogo:     uuid.NewString(),
	FiatUnit:     "BTC",
	FeedType:     basetypes.CurrencyFeedType_CoinGecko,
	FeedTypeStr:  basetypes.CurrencyFeedType_CoinGecko.String(),
	FeedFiatName: "BTC123",
}

var req = &npool.FeedReq{
	FeedType:     &ret.FeedType,
	FeedFiatName: &ret.FeedFiatName,
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

	return func(*testing.T) {}
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithFiatID(req.FiatID, true),
		WithFeedType(req.FeedType, true),
		WithFeedFiatName(req.FeedFiatName, true),
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
	ret.FeedFiatName = "BTCWWWW"
	ret.Disabled = true

	req.FeedFiatName = &ret.FeedFiatName
	req.Disabled = &ret.Disabled

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithFeedFiatName(req.FeedFiatName, true),
		WithDisabled(req.Disabled, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateFeed(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, ret)
	}
}

func TestFiat(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupFiat(t)
	defer teardown(t)

	t.Run("create", create)
	t.Run("update", update)
}
