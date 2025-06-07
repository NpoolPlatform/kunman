package capacity

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/capacity"
	types "github.com/NpoolPlatform/kunman/message/basetypes/agi/v1"
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

var ret = npool.Capacity{
	EntID:          uuid.NewString(),
	AppGoodID:      uuid.NewString(),
	CapacityKey:    types.CapacityKey_CapacityQPS,
	CapacityKeyStr: types.CapacityKey_CapacityQPS.String(),
	Value:          "10",
	Description:    "10 Requests per second",
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCapacity(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithCapacityKey(&ret.CapacityKey, true),
		WithValue(&ret.Value, true),
		WithDescription(&ret.Description, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCapacity(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCapacity(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateCapacity(t *testing.T) {
	ret.Value = "20"

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithValue(&ret.Value, true),
		WithDescription(&ret.Description, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateCapacity(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCapacity(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getCapacity(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCapacity(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCapacities(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		IDs:        &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
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

	infos, err := handler.GetCapacities(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCapacity(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteCapacity(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetCapacity(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCapacity(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCapacity", createCapacity)
	t.Run("updateCapacity", updateCapacity)
	t.Run("getCapacity", getCapacity)
	t.Run("getCapacities", getCapacities)
	t.Run("deleteCapacity", deleteCapacity)
}
