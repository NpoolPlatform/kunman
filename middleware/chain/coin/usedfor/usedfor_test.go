package coinusedfor

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	types "github.com/NpoolPlatform/kunman/message/basetypes/chain/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/usedfor"
	testinit "github.com/NpoolPlatform/kunman/middleware/chain/testinit"

	coin1 "github.com/NpoolPlatform/kunman/middleware/chain/coin"
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

var ret = &npool.CoinUsedFor{
	CoinName:   "My BTC1",
	CoinLogo:   uuid.NewString(),
	CoinUnit:   "BTC",
	CoinENV:    "test",
	UsedFor:    types.CoinUsedFor_CoinUsedForGoodFee,
	UsedForStr: types.CoinUsedFor_CoinUsedForGoodFee.String(),
	Priority:   1,
}

var req = &npool.CoinUsedForReq{
	UsedFor: &ret.UsedFor,
}

func setupCoinUsedFor(t *testing.T) func(*testing.T) {
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

func createCoinUsedFor(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithCoinTypeID(req.CoinTypeID, true),
		WithUsedFor(req.UsedFor, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoinUsedFor(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, ret, info)
	}
}

func getCoinUsedFors(t *testing.T) {
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

	infos, total, err := handler.GetCoinUsedFors(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		assert.Equal(t, 1, len(infos))
		assert.Equal(t, ret, infos[0])
	}
}

func deleteCoinUsedFor(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCoinUsedFor(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func TestCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupCoinUsedFor(t)
	defer teardown(t)

	t.Run("createCoinUsedFor", createCoinUsedFor)
	t.Run("getCoinUsedFors", getCoinUsedFors)
	t.Run("deleteCoinUsedFor", deleteCoinUsedFor)
}
