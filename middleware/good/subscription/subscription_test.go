package subscription

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/good/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Subscription{
	EntID:               uuid.NewString(),
	GoodID:              uuid.NewString(),
	GoodType:            types.GoodType_Subscription,
	Name:                uuid.NewString(),
	DurationDisplayType: types.GoodDurationType_GoodDurationByDay,
	DurationUnits:       10,
	DurationQuota:       10,
	DailyBonusQuota:     10,
	USDPrice:            decimal.NewFromInt(20).String(),
}

//nolint:unparam
func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.DurationDisplayTypeStr = ret.DurationDisplayType.String()
	return func(*testing.T) {}
}

func createSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithName(&ret.Name, true),
		WithDurationDisplayType(&ret.DurationDisplayType, true),
		WithDurationUnits(&ret.DurationUnits, true),
		WithDurationQuota(&ret.DurationQuota, true),
		WithDailyBonusQuota(&ret.DailyBonusQuota, true),
		WithUSDPrice(&ret.USDPrice, true),
	)
	assert.Nil(t, err)

	err = handler.CreateSubscription(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetSubscription(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateSubscription(t *testing.T) {
	ret.Name = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithName(&ret.Name, true),
		WithDurationDisplayType(&ret.DurationDisplayType, true),
		WithDurationUnits(&ret.DurationUnits, true),
		WithDurationQuota(&ret.DurationQuota, true),
		WithDailyBonusQuota(&ret.DailyBonusQuota, true),
		WithUSDPrice(&ret.USDPrice, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateSubscription(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetSubscription(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetSubscription(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSubscriptions(t *testing.T) {
	conds := &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, err := handler.GetSubscriptions(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteSubscription(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetSubscription(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestSubscription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createSubscription", createSubscription)
	t.Run("updateSubscription", updateSubscription)
	t.Run("getSubscription", getSubscription)
	t.Run("getSubscriptions", getSubscriptions)
	t.Run("deleteSubscription", deleteSubscription)
}
