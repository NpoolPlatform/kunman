package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/commission/config"
	appconfigmw "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"
	appconfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/config"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
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

var (
	ret = npool.AppCommissionConfig{
		EntID:           uuid.NewString(),
		AppID:           uuid.NewString(),
		SettleType:      types.SettleType_GoodOrderPayment,
		SettleTypeStr:   types.SettleType_GoodOrderPayment.String(),
		AmountOrPercent: decimal.RequireFromString("12.25").String(),
		ThresholdAmount: decimal.RequireFromString("12.26").String(),
		StartAt:         uint32(time.Now().Unix()),
		Invites:         uint32(1),
		Level:           uint32(1),
		Disabled:        false,
	}
	appConfigRet = appconfigmw.AppConfig{
		EntID:               uuid.NewString(),
		AppID:               ret.AppID,
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
	appConfigReq = appconfigmw.AppConfigReq{
		EntID:            &appConfigRet.EntID,
		AppID:            &appConfigRet.AppID,
		SettleMode:       &appConfigRet.SettleMode,
		SettleAmountType: &appConfigRet.SettleAmountType,
		SettleInterval:   &appConfigRet.SettleInterval,
		CommissionType:   &appConfigRet.CommissionType,
		SettleBenefit:    &appConfigRet.SettleBenefit,
		StartAt:          &appConfigRet.StartAt,
		MaxLevel:         &appConfigRet.MaxLevel,
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := appconfig1.NewHandler(
		context.Background(),
		appconfig1.WithEntID(appConfigReq.EntID, true),
		appconfig1.WithAppID(appConfigReq.AppID, true),
		appconfig1.WithSettleInterval(appConfigReq.SettleInterval, true),
		appconfig1.WithSettleMode(appConfigReq.SettleMode, true),
		appconfig1.WithSettleAmountType(appConfigReq.SettleAmountType, true),
		appconfig1.WithCommissionType(appConfigReq.CommissionType, true),
		appconfig1.WithSettleBenefit(appConfigReq.SettleBenefit, true),
		appconfig1.WithStartAt(appConfigReq.StartAt, true),
		appconfig1.WithMaxLevel(appConfigReq.MaxLevel, true),
	)
	assert.Nil(t, err)

	err = h1.CreateAppConfig(context.Background())
	if assert.Nil(t, err) {
		info1, err := h1.GetAppConfig(context.Background())
		if assert.Nil(t, err) {
			h1.ID = &info1.ID
		}
	}

	return func(*testing.T) {
		_ = h1.DeleteAppConfig(context.Background())
	}
}

func createCommissionConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithSettleType(&ret.SettleType, true),
		WithThresholdAmount(&ret.ThresholdAmount, true),
		WithAmountOrPercent(&ret.AmountOrPercent, true),
		WithStartAt(&ret.StartAt, true),
		WithInvites(&ret.Invites, true),
		WithLevel(&ret.Level, true),
		WithDisabled(&ret.Disabled, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCommissionConfig(context.Background())
		if assert.Nil(t, err) {
			ret.ID = info.ID
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func updateCommissionConfig(t *testing.T) {
	ret.AmountOrPercent = "13"
	ret.StartAt += 10000
	ret.ThresholdAmount = decimal.NewFromInt(10).String()
	ret.Level = uint32(2)

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAmountOrPercent(&ret.AmountOrPercent, true),
		WithStartAt(&ret.StartAt, true),
		WithThresholdAmount(&ret.ThresholdAmount, true),
		WithLevel(&ret.Level, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCommissionConfig(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getCommissionConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCommissionConfig(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCommissionConfigs(t *testing.T) {
	conds := &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleType)},
		EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.EndAt},
		StartAt:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.StartAt},
		Disabled:   &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Level:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.Level},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCommissionConfigs(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCommissionConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteCommissionConfig(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetCommissionConfig(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCommission(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCommissionConfig", createCommissionConfig)
	t.Run("updateCommissionConfig", updateCommissionConfig)
	t.Run("getCommissionConfig", getCommissionConfig)
	t.Run("getCommissionConfigs", getCommissionConfigs)
	t.Run("deleteCommissionConfig", deleteCommissionConfig)
}
