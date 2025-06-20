package reward

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/user/coin/reward"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/inspire/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.UserCoinReward{
	EntID:       uuid.NewString(),
	AppID:       uuid.NewString(),
	UserID:      uuid.NewString(),
	CoinTypeID:  uuid.NewString(),
	CoinRewards: decimal.RequireFromString("11.25").String(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createUserCoinReward(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithCoinRewards(&ret.CoinRewards, true),
	)
	assert.Nil(t, err)

	err = handler.CreateUserCoinReward(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetUserCoinReward(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateUserCoinReward(t *testing.T) {
	ret.CoinRewards = "22.5"

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithCoinRewards(&ret.CoinRewards, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateUserCoinReward(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetUserCoinReward(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getUserCoinReward(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetUserCoinReward(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getUserCoinRewards(t *testing.T) {
	conds := &npool.Conds{
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetUserCoinRewards(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteUserCoinReward(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteUserCoinReward(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetUserCoinReward(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestUserCoinReward(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createUserCoinReward", createUserCoinReward)
	t.Run("updateUserCoinReward", updateUserCoinReward)
	t.Run("getUserCoinReward", getUserCoinReward)
	t.Run("getUserCoinRewards", getUserCoinRewards)
	t.Run("deleteUserCoinReward", deleteUserCoinReward)
}
