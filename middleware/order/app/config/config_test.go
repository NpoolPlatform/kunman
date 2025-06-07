package appconfig

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/app/config"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/order/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.AppConfig{
	EntID:                                  uuid.NewString(),
	AppID:                                  uuid.NewString(),
	EnableSimulateOrder:                    false,
	SimulateOrderCouponMode:                ordertypes.SimulateOrderCouponMode_WithoutCoupon,
	SimulateOrderCouponProbability:         "1",
	SimulateOrderCashableProfitProbability: "0.1",
	MaxUnpaidOrders:                        10,
	MaxTypedCouponsPerOrder:                10,
}

//nolint:unparam
func setup(t *testing.T) func(*testing.T) {
	ret.SimulateOrderCouponModeStr = ret.SimulateOrderCouponMode.String()
	return func(*testing.T) {}
}

func createAppConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithEnableSimulateOrder(&ret.EnableSimulateOrder, true),
		WithSimulateOrderCouponMode(&ret.SimulateOrderCouponMode, true),
		WithSimulateOrderCouponProbability(&ret.SimulateOrderCouponProbability, true),
		WithSimulateOrderCashableProfitProbability(&ret.SimulateOrderCashableProfitProbability, true),
		WithMaxUnpaidOrders(&ret.MaxUnpaidOrders, true),
		WithMaxTypedCouponsPerOrder(&ret.MaxTypedCouponsPerOrder, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateAppConfig(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetAppConfig(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func updateAppConfig(t *testing.T) {
	ret.EnableSimulateOrder = true
	ret.SimulateOrderCouponMode = ordertypes.SimulateOrderCouponMode_RandomBenifit
	ret.SimulateOrderCouponModeStr = ordertypes.SimulateOrderCouponMode_RandomBenifit.String()
	ret.SimulateOrderCouponProbability = "0.5"
	ret.SimulateOrderCashableProfitProbability = "0.5"

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithEnableSimulateOrder(&ret.EnableSimulateOrder, true),
		WithSimulateOrderCouponMode(&ret.SimulateOrderCouponMode, true),
		WithSimulateOrderCouponProbability(&ret.SimulateOrderCouponProbability, true),
		WithSimulateOrderCashableProfitProbability(&ret.SimulateOrderCashableProfitProbability, true),
		WithMaxUnpaidOrders(&ret.MaxUnpaidOrders, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateAppConfig(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetAppConfig(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getAppConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetAppConfig(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getAppConfigs(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetAppConfigs(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteAppConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteAppConfig(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetAppConfig(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
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
