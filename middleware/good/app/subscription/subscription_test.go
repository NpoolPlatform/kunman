package subscription

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription"
	subscription1 "github.com/NpoolPlatform/kunman/middleware/good/subscription"
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
	AppID:               uuid.NewString(),
	GoodID:              uuid.NewString(),
	AppGoodID:           uuid.NewString(),
	GoodType:            types.GoodType_Subscription,
	GoodName:            uuid.NewString(),
	AppGoodName:         uuid.NewString(),
	USDPrice:            decimal.NewFromInt(20).String(),
	DurationDisplayType: types.GoodDurationType_GoodDurationByDay,
	DurationUnits:       2,
	DurationQuota:       2000,
}

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.DurationDisplayTypeStr = ret.DurationDisplayType.String()

	subscriptionEntID := uuid.NewString()
	h1, err := subscription1.NewHandler(
		context.Background(),
		subscription1.WithEntID(&subscriptionEntID, true),
		subscription1.WithGoodID(&ret.GoodID, true),
		subscription1.WithGoodType(&ret.GoodType, true),
		subscription1.WithName(&ret.GoodName, true),
		subscription1.WithUSDPrice(&ret.USDPrice, true),
		subscription1.WithDurationDisplayType(&ret.DurationDisplayType, true),
		subscription1.WithDurationUnits(&ret.DurationUnits, true),
		subscription1.WithDurationQuota(&ret.DurationQuota, true),
	)
	assert.Nil(t, err)

	err = h1.CreateSubscription(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteSubscription(context.Background())
	}
}

func createSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithName(&ret.AppGoodName, true),
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
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithName(&ret.AppGoodName, true),
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
		WithAppGoodID(&ret.AppGoodID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetSubscription(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func existSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistSubscription(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func getSubscriptions(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		IDs:        &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		AppIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppID, ret.AppID}},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID, ret.EntID}},
		GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		GoodIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppGoodID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetSubscriptions(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(infos))
		assert.Equal(t, &ret, infos[0])
	}
}

func deleteSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
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
	t.Run("existSubscription", existSubscription)
	t.Run("getSubscriptions", getSubscriptions)
	t.Run("deleteSubscription", deleteSubscription)
}
