package subscription

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/subscription"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/agi/testinit"
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
	EntID:          uuid.NewString(),
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	AppGoodID:      uuid.NewString(),
	NextExtendAt:   1,
	PermanentQuota: 1,
	ConsumedQuota:  0,
	AutoExtend:     false,
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithNextExtendAt(&ret.NextExtendAt, true),
		WithPermanentQuota(&ret.PermanentQuota, true),
		WithConsumedQuota(&ret.ConsumedQuota, true),
		WithAutoExtend(&ret.AutoExtend, true),
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
	ret.NextExtendAt = 10
	ret.PermanentQuota = 10
	ret.ConsumedQuota = 5
	ret.AutoExtend = true

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithNextExtendAt(&ret.NextExtendAt, true),
		WithPermanentQuota(&ret.PermanentQuota, true),
		WithConsumedQuota(&ret.ConsumedQuota, true),
		WithAutoExtend(&ret.AutoExtend, true),
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
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetSubscription(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSubscriptions(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		IDs:        &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		AppIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppID}},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
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

	infos, err := handler.GetSubscriptions(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
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
