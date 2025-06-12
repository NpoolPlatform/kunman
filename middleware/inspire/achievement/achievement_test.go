package achievement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	commission1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/config"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	orderstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order/payment"

	achievement1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/good"
	statement1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/statement/order"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/good"
	appconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"

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
	EntID:                   uuid.NewString(),
	AppID:                   uuid.NewString(),
	UserID:                  uuid.NewString(),
	GoodID:                  uuid.NewString(),
	AppGoodID:               uuid.NewString(),
	OrderID:                 uuid.NewString(),
	OrderUserID:             uuid.NewString(),
	GoodCoinTypeID:          uuid.NewString(),
	Units:                   decimal.NewFromInt(10).String(),
	GoodValueUSD:            decimal.NewFromInt(120).String(),
	PaymentAmountUSD:        decimal.NewFromInt(120).String(),
	CommissionAmountUSD:     decimal.NewFromInt(0).String(),
	AppConfigID:             uuid.NewString(),
	CommissionConfigID:      uuid.Nil.String(),
	CommissionConfigType:    types.CommissionConfigType_LegacyCommissionConfig,
	CommissionConfigTypeStr: types.CommissionConfigType_LegacyCommissionConfig.String(),
}

var ret1 = &npool.Achievement{
	AppID:              ret.AppID,
	UserID:             ret.UserID,
	GoodID:             ret.GoodID,
	AppGoodID:          ret.AppGoodID,
	TotalAmountUSD:     ret.PaymentAmountUSD,
	SelfAmountUSD:      "0",
	TotalUnits:         ret.Units,
	SelfUnits:          "0",
	TotalCommissionUSD: ret.CommissionAmountUSD,
	SelfCommissionUSD:  "0",
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

func createAppConfig(t *testing.T) {
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
}

func createStatement(t *testing.T) {
	paymentCoinTypeID := uuid.NewString()
	amount := "1000"
	commissionAmount := "0"
	ret.DirectContributorID = ret.UserID
	handler, err := statement1.NewHandler(
		context.Background(),
		statement1.WithEntID(&ret.EntID, true),
		statement1.WithAppID(&ret.AppID, true),
		statement1.WithUserID(&ret.UserID, true),
		statement1.WithGoodID(&ret.GoodID, true),
		statement1.WithAppGoodID(&ret.AppGoodID, true),
		statement1.WithOrderID(&ret.OrderID, true),
		statement1.WithOrderUserID(&ret.OrderUserID, true),
		statement1.WithDirectContributorID(&ret.DirectContributorID, true),
		statement1.WithGoodCoinTypeID(&ret.GoodCoinTypeID, true),
		statement1.WithUnits(&ret.Units, true),
		statement1.WithGoodValueUSD(&ret.GoodValueUSD, true),
		statement1.WithPaymentAmountUSD(&ret.PaymentAmountUSD, true),
		statement1.WithCommissionAmountUSD(&ret.CommissionAmountUSD, true),
		statement1.WithAppConfigID(&ret.AppConfigID, true),
		statement1.WithCommissionConfigID(&ret.CommissionConfigID, true),
		statement1.WithCommissionConfigType(&ret.CommissionConfigType, true),
		statement1.WithPaymentStatements([]*paymentmwpb.StatementReq{{
			PaymentCoinTypeID: &paymentCoinTypeID,
			Amount:            &amount,
			CommissionAmount:  &commissionAmount,
		},
		}, true),
	)
	assert.Nil(t, err)

	err = handler.CreateStatement(context.Background())
	assert.Nil(t, err)
}

func expropriateAchievement(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithOrderID(&ret.OrderID, true),
	)
	assert.Nil(t, err)

	err = handler.ExpropriateAchievement(context.Background())
	assert.Nil(t, err)

	h2, err := achievement1.NewHandler(
		context.Background(),
		achievement1.WithConds(&npool.Conds{
			AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			GoodID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
			UserIDs:   &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
		}),
		achievement1.WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := h2.GetAchievements(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		ret1.ID = infos[0].ID
		ret1.EntID = infos[0].EntID
		ret1.TotalAmountUSD = decimal.NewFromInt(0).String()
		ret1.SelfAmountUSD = decimal.NewFromInt(0).String()
		ret1.TotalUnits = decimal.NewFromInt(0).String()
		ret1.SelfUnits = decimal.NewFromInt(0).String()
		ret1.TotalCommissionUSD = decimal.NewFromInt(0).String()
		ret1.SelfCommissionUSD = decimal.NewFromInt(0).String()
		ret1.CreatedAt = infos[0].CreatedAt
		ret1.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], ret1)
	}
}

func TestAchievement(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createAppConfig", createAppConfig)
	t.Run("createStatement", createStatement)
	t.Run("expropriateAchievement", expropriateAchievement)
}
