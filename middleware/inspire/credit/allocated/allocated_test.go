package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/credit/allocated"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/inspire/testinit"
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

var (
	ret = npool.CreditAllocated{
		EntID:  uuid.NewString(),
		AppID:  uuid.NewString(),
		UserID: uuid.NewString(),
		Value:  decimal.RequireFromString("2.25").String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCreditAllocated(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithValue(&ret.Value, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCreditAllocated(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCreditAllocated(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func getCreditAllocated(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCreditAllocated(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCreditAllocateds(t *testing.T) {
	conds := &npool.Conds{
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCreditAllocateds(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCreditAllocated(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteCreditAllocated(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetCreditAllocated(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCreditAllocated(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCreditAllocated", createCreditAllocated)
	t.Run("getCreditAllocated", getCreditAllocated)
	t.Run("getCreditAllocateds", getCreditAllocateds)
	t.Run("deleteCreditAllocated", deleteCreditAllocated)
}
