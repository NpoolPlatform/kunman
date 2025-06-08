package description

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin/description"

	appcoin1 "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
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

var ret = &npool.CoinDescription{
	EntID:      uuid.NewString(),
	AppID:      uuid.NewString(),
	CoinLogo:   uuid.NewString(),
	CoinUnit:   "BTC",
	CoinENV:    "test",
	UsedForStr: basetypes.UsedFor_ProductPage.String(),
	UsedFor:    basetypes.UsedFor_ProductPage,
	Title:      uuid.NewString(),
	Message:    uuid.NewString(),
}

var req = &npool.CoinDescriptionReq{
	AppID:   &ret.AppID,
	UsedFor: &ret.UsedFor,
	Title:   &ret.Title,
	Message: &ret.Message,
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

	info1, err := h1.CreateCoin(context.Background())
	assert.Nil(t, err)
	h1.ID = &info1.ID

	h2, err := appcoin1.NewHandler(
		context.Background(),
		appcoin1.WithAppID(&ret.AppID, true),
		appcoin1.WithCoinTypeID(&ret.CoinTypeID, true),
	)
	assert.Nil(t, err)

	info2, err := h2.CreateCoin(context.Background())
	assert.Nil(t, err)
	h2.ID = &info2.ID

	return func(*testing.T) {
		_, _ = h1.DeleteCoin(context.Background())
		_, _ = h2.DeleteCoin(context.Background())
	}
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(req.EntID, false),
		WithAppID(req.AppID, true),
		WithCoinTypeID(req.CoinTypeID, true),
		WithTitle(req.Title, true),
		WithMessage(req.Message, true),
		WithUsedFor(req.UsedFor, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoinDescription(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, ret)
	}
}

func update(t *testing.T) {
	ret.Title = uuid.NewString()
	ret.Message = uuid.NewString()

	req.Title = &ret.Title
	req.Message = &ret.Message

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithTitle(req.Title, true),
		WithMessage(req.Message, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoinDescription(context.Background())
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
