package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coin/config"
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

var ret = npool.CoinConfig{
	EntID:      uuid.NewString(),
	AppID:      uuid.NewString(),
	CoinTypeID: uuid.NewString(),
	MaxValue:   decimal.RequireFromString("11.25").String(),
	Allocated:  decimal.RequireFromString("0").String(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCoinConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithMaxValue(&ret.MaxValue, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCoinConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCoinConfig(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateCoinConfig(t *testing.T) {
	ret.MaxValue = "22.5"
	ret.Allocated = "2.5"

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMaxValue(&ret.MaxValue, true),
		WithAllocated(&ret.Allocated, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateCoinConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCoinConfig(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getCoinConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCoinConfig(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCoinConfigs(t *testing.T) {
	conds := &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCoinConfigs(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCoinConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteCoinConfig(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetCoinConfig(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCoinConfig(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCoinConfig", createCoinConfig)
	t.Run("updateCoinConfig", updateCoinConfig)
	t.Run("getCoinConfig", getCoinConfig)
	t.Run("getCoinConfigs", getCoinConfigs)
	t.Run("deleteCoinConfig", deleteCoinConfig)
}
