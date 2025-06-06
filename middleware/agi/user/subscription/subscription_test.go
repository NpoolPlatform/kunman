package subscription

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/user/subscription"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/agi/testinit"
	types "github.com/NpoolPlatform/kunman/message/basetypes/agi/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
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
	EntID:              uuid.NewString(),
	AppID:              uuid.NewString(),
	UserID:             uuid.NewString(),
	PackageID:          uuid.NewString(),
	StartAt:            uint32(10),
	EndAt:              uint32(20),
	UsageState:         types.UsageState_Usful,
	UsageStateStr:      types.UsageState_Usful.String(),
	SubscriptionCredit: uint32(10),
	AddonCredit:        uint32(20),
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
		WithPackageID(&ret.PackageID, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
		WithUsageState(&ret.UsageState, true),
		WithSubscriptionCredit(&ret.SubscriptionCredit, true),
		WithAddonCredit(&ret.AddonCredit, true),
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
	ret.StartAt = uint32(20)
	ret.EndAt = uint32(30)
	ret.UsageState = types.UsageState_Expire
	ret.UsageStateStr = types.UsageState_Expire.String()
	ret.SubscriptionCredit = uint32(20)
	ret.AddonCredit = uint32(30)
	ret.PackageID = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
		WithUsageState(&ret.UsageState, true),
		WithSubscriptionCredit(&ret.SubscriptionCredit, true),
		WithAddonCredit(&ret.AddonCredit, true),
		WithPackageID(&ret.PackageID, true),
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
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		PackageID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.PackageID},
		StartAt:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.StartAt},
		EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.EndAt},
		UsageState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.UsageState)},
		IDs:        &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
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
