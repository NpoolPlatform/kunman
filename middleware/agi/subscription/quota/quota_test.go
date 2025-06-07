package quota

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/subscription/quota"
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

var ret = npool.Quota{
	EntID:         uuid.NewString(),
	AppID:         uuid.NewString(),
	UserID:        uuid.NewString(),
	Quota:         1,
	ConsumedQuota: 0,
	ExpiredAt:     10,
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createQuota(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithQuota(&ret.Quota, true),
		WithConsumedQuota(&ret.ConsumedQuota, true),
		WithExpiredAt(&ret.ExpiredAt, true),
	)
	assert.Nil(t, err)

	err = handler.CreateQuota(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetQuota(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateQuota(t *testing.T) {
	ret.Quota = 10
	ret.ConsumedQuota = 5
	ret.ExpiredAt = 20

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithQuota(&ret.Quota, true),
		WithConsumedQuota(&ret.ConsumedQuota, true),
		WithExpiredAt(&ret.ExpiredAt, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateQuota(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetQuota(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getQuota(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetQuota(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getQuotas(t *testing.T) {
	conds := &npool.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		IDs:     &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		EntIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		AppIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppID}},
		UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		UserIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, err := handler.GetQuotas(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteQuota(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteQuota(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetQuota(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestQuota(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createQuota", createQuota)
	t.Run("updateQuota", updateQuota)
	t.Run("getQuota", getQuota)
	t.Run("getQuotas", getQuotas)
	t.Run("deleteQuota", deleteQuota)
}
