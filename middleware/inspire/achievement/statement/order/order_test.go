package orderstatement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	appconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"
	commission1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/config"

	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	orderstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	orderpaymentstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order/payment"

	"github.com/NpoolPlatform/kunman/middleware/inspire/testinit"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &orderstatementmwpb.Statement{
	EntID:                uuid.NewString(),
	AppID:                uuid.NewString(),
	UserID:               uuid.NewString(),
	GoodID:               uuid.NewString(),
	AppGoodID:            uuid.NewString(),
	OrderID:              uuid.NewString(),
	OrderUserID:          uuid.NewString(),
	GoodCoinTypeID:       uuid.NewString(),
	Units:                decimal.NewFromInt(10).String(),
	GoodValueUSD:         decimal.NewFromInt(120).String(),
	PaymentAmountUSD:     decimal.NewFromInt(120).String(),
	CommissionAmountUSD:  decimal.NewFromInt(0).String(),
	AppConfigID:          uuid.NewString(),
	CommissionConfigID:   uuid.Nil.String(),
	CommissionConfigType: types.CommissionConfigType_LegacyCommissionConfig,
}

var appconfig = appconfigmwpb.AppConfig{
	EntID:               ret.AppConfigID,
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

func setup(t *testing.T) func(*testing.T) {
	ret.CommissionConfigTypeStr = ret.CommissionConfigType.String()
	handler, err := commission1.NewHandler(
		context.Background(),
		commission1.WithEntID(&ret.AppConfigID, true),
		commission1.WithAppID(&ret.AppID, true),
		commission1.WithCommissionType(&appconfig.CommissionType, true),
		commission1.WithSettleMode(&appconfig.SettleMode, true),
		commission1.WithSettleAmountType(&appconfig.SettleAmountType, true),
		commission1.WithSettleInterval(&appconfig.SettleInterval, true),
		commission1.WithSettleBenefit(&appconfig.SettleBenefit, true),
		commission1.WithStartAt(&appconfig.StartAt, true),
		commission1.WithMaxLevel(&appconfig.MaxLevel, true),
	)
	assert.Nil(t, err)

	err = handler.CreateAppConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetAppConfig(context.Background())
		if assert.Nil(t, err) {
			appconfig.ID = info.ID
			appconfig.CreatedAt = info.CreatedAt
			appconfig.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &appconfig)
		}
	}
	return func(*testing.T) {
		_ = handler.DeleteAppConfig(context.Background())
	}
}

func createStatement(t *testing.T) {
	ret.DirectContributorID = ret.UserID
	payments := []*orderpaymentstatementmwpb.StatementReq{}
	payments = append(payments, &orderpaymentstatementmwpb.StatementReq{
		PaymentCoinTypeID: &ret.GoodCoinTypeID,
		Amount:            &ret.PaymentAmountUSD,
		CommissionAmount:  &ret.CommissionAmountUSD,
	})
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithOrderID(&ret.OrderID, true),
		WithOrderUserID(&ret.OrderUserID, true),
		WithDirectContributorID(&ret.DirectContributorID, true),
		WithGoodCoinTypeID(&ret.GoodCoinTypeID, true),
		WithUnits(&ret.Units, true),
		WithGoodValueUSD(&ret.GoodValueUSD, true),
		WithPaymentAmountUSD(&ret.PaymentAmountUSD, true),
		WithCommissionAmountUSD(&ret.CommissionAmountUSD, true),
		WithAppConfigID(&ret.AppConfigID, true),
		WithCommissionConfigID(&ret.CommissionConfigID, true),
		WithCommissionConfigType(&ret.CommissionConfigType, true),
		WithPaymentStatements(payments, true),
	)
	assert.Nil(t, err)

	err = handler.CreateStatement(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetStatement(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func deleteStatement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteStatement(context.Background())
	assert.Nil(t, err)
}

func TestAchievement(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createStatement", createStatement)
	t.Run("deleteStatement", deleteStatement)
}
