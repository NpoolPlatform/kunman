package reward

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/user/reward"
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

var ret = npool.UserReward{
	EntID:                uuid.NewString(),
	AppID:                uuid.NewString(),
	UserID:               uuid.NewString(),
	ActionCredits:        decimal.RequireFromString("11.25").String(),
	CouponAmount:         decimal.RequireFromString("11.25").String(),
	CouponCashableAmount: decimal.RequireFromString("11.25").String(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createUserReward(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithActionCredits(&ret.ActionCredits, true),
		WithCouponAmount(&ret.CouponAmount, true),
		WithCouponCashableAmount(&ret.CouponCashableAmount, true),
	)
	assert.Nil(t, err)

	err = handler.AddUserReward(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetUserReward(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

//nolint:dupl
func addUserReward(t *testing.T) {
	addActionCredits := decimal.RequireFromString("22.25").String()
	addCouponAmount := decimal.RequireFromString("22.25").String()
	addCouponCashableAmount := decimal.RequireFromString("22.25").String()

	ret.ActionCredits = decimal.RequireFromString("33.5").String()
	ret.CouponAmount = decimal.RequireFromString("33.5").String()
	ret.CouponCashableAmount = decimal.RequireFromString("33.5").String()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithActionCredits(&addActionCredits, true),
		WithCouponAmount(&addCouponAmount, true),
		WithCouponCashableAmount(&addCouponCashableAmount, true),
	)
	assert.Nil(t, err)

	err = handler.AddUserReward(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetUserReward(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

//nolint:dupl
func subUserReward(t *testing.T) {
	subActionCredits := decimal.RequireFromString("22.25").String()
	subCouponAmount := decimal.RequireFromString("22.25").String()
	subCouponCashableAmount := decimal.RequireFromString("22.25").String()

	ret.ActionCredits = decimal.RequireFromString("11.25").String()
	ret.CouponAmount = decimal.RequireFromString("11.25").String()
	ret.CouponCashableAmount = decimal.RequireFromString("11.25").String()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithActionCredits(&subActionCredits, true),
		WithCouponAmount(&subCouponAmount, true),
		WithCouponCashableAmount(&subCouponCashableAmount, true),
	)
	assert.Nil(t, err)

	err = handler.SubUserReward(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetUserReward(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getUserReward(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetUserReward(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getUserRewards(t *testing.T) {
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

	infos, _, err := handler.GetUserRewards(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteUserReward(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteUserReward(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetUserReward(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestUserReward(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createUserReward", createUserReward)
	t.Run("addUserReward", addUserReward)
	t.Run("subUserReward", subUserReward)
	t.Run("getUserReward", getUserReward)
	t.Run("getUserRewards", getUserRewards)
	t.Run("deleteUserReward", deleteUserReward)
}
