package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/inspire/testinit"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
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
	ret = npool.AppConfig{
		EntID:               uuid.NewString(),
		AppID:               uuid.NewString(),
		SettleMode:          types.SettleMode_SettleWithGoodValue,
		SettleModeStr:       types.SettleMode_SettleWithGoodValue.String(),
		SettleAmountType:    types.SettleAmountType_SettleByPercent,
		SettleAmountTypeStr: types.SettleAmountType_SettleByPercent.String(),
		SettleInterval:      types.SettleInterval_SettleYearly,
		SettleIntervalStr:   types.SettleInterval_SettleYearly.String(),
		CommissionType:      types.CommissionType_LayeredCommission,
		CommissionTypeStr:   types.CommissionType_LayeredCommission.String(),
		SettleBenefit:       false,
		StartAt:             uint32(time.Now().Unix()),
		MaxLevel:            uint32(5),
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createAppConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithCommissionType(&ret.CommissionType, true),
		WithSettleMode(&ret.SettleMode, true),
		WithSettleAmountType(&ret.SettleAmountType, true),
		WithSettleInterval(&ret.SettleInterval, true),
		WithSettleBenefit(&ret.SettleBenefit, true),
		WithStartAt(&ret.StartAt, true),
		WithMaxLevel(&ret.MaxLevel, true),
	)
	assert.Nil(t, err)

	err = handler.CreateAppConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetAppConfig(context.Background())
		if assert.Nil(t, err) {
			ret.ID = info.ID
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func updateAppConfig(t *testing.T) {
	ret.StartAt += 10000

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithStartAt(&ret.StartAt, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateAppConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetAppConfig(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getAppConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetAppConfig(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAppConfigs(t *testing.T) {
	conds := &npool.Conds{
		EntID:            &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:            &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		SettleMode:       &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleMode)},
		SettleAmountType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleAmountType)},
		SettleInterval:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleInterval)},
		CommissionType:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CommissionType)},
		SettleBenefit:    &basetypes.BoolVal{Op: cruder.EQ, Value: ret.SettleBenefit},
		EndAt:            &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.EndAt},
		StartAt:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.StartAt},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetAppConfigs(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteAppConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteAppConfig(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetAppConfig(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAppConfig(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createAppConfig", createAppConfig)
	t.Run("updateAppConfig", updateAppConfig)
	t.Run("getAppConfig", getAppConfig)
	t.Run("getAppConfigs", getAppConfigs)
	t.Run("deleteAppConfig", deleteAppConfig)
}
