package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	statement1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/statement/order"
	common1 "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/user/common"
	commission1 "github.com/NpoolPlatform/kunman/middleware/inspire/app/config"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	orderstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/user"
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

var appconfig = appconfigmwpb.AppConfig{
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

var userID = uuid.NewString()
var statement = &orderstatementmwpb.Statement{
	EntID:                uuid.NewString(),
	AppID:                appconfig.AppID,
	UserID:               userID,
	GoodID:               uuid.NewString(),
	AppGoodID:            uuid.NewString(),
	OrderID:              uuid.NewString(),
	OrderUserID:          userID,
	DirectContributorID:  userID,
	GoodCoinTypeID:       uuid.NewString(),
	Units:                decimal.NewFromInt(10).String(),
	GoodValueUSD:         decimal.NewFromInt(120).String(),
	PaymentAmountUSD:     decimal.NewFromInt(120).String(),
	CommissionAmountUSD:  decimal.NewFromInt(0).String(),
	AppConfigID:          appconfig.EntID,
	CommissionConfigID:   uuid.Nil.String(),
	CommissionConfigType: types.CommissionConfigType_LegacyCommissionConfig,
}

func setup(t *testing.T) func(*testing.T) {
	statement.CommissionConfigTypeStr = statement.CommissionConfigType.String()
	handler, err := commission1.NewHandler(
		context.Background(),
		commission1.WithEntID(&statement.AppConfigID, true),
		commission1.WithAppID(&statement.AppID, true),
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

	// statement
	statementHandler, err := statement1.NewHandler(
		context.Background(),
		statement1.WithEntID(&statement.EntID, true),
		statement1.WithAppID(&statement.AppID, true),
		statement1.WithUserID(&statement.UserID, true),
		statement1.WithGoodID(&statement.GoodID, true),
		statement1.WithAppGoodID(&statement.AppGoodID, true),
		statement1.WithOrderID(&statement.OrderID, true),
		statement1.WithOrderUserID(&statement.OrderUserID, true),
		statement1.WithDirectContributorID(&statement.UserID, true),
		statement1.WithOrderUserID(&statement.DirectContributorID, true),
		statement1.WithGoodCoinTypeID(&statement.GoodCoinTypeID, true),
		statement1.WithUnits(&statement.Units, true),
		statement1.WithGoodValueUSD(&statement.GoodValueUSD, true),
		statement1.WithPaymentAmountUSD(&statement.PaymentAmountUSD, true),
		statement1.WithCommissionAmountUSD(&statement.CommissionAmountUSD, true),
		statement1.WithAppConfigID(&statement.AppConfigID, true),
		statement1.WithCommissionConfigID(&statement.CommissionConfigID, true),
		statement1.WithCommissionConfigType(&statement.CommissionConfigType, true),
	)
	assert.Nil(t, err)

	err = statementHandler.CreateStatement(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = handler.DeleteAppConfig(context.Background())
		_ = statementHandler.DeleteStatement(context.Background())
	}
}

var ret = npool.AchievementUser{
	AppID:                statement.AppID,
	UserID:               statement.UserID,
	TotalCommission:      decimal.Zero.String(),
	SelfCommission:       decimal.Zero.String(),
	DirectInvitees:       0,
	IndirectInvitees:     0,
	DirectConsumeAmount:  statement.GoodValueUSD,
	InviteeConsumeAmount: decimal.Zero.String(),
}

func getAchievementUsers(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		common1.WithConds(&npool.Conds{
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: statement.AppID},
			UserID:  &basetypes.StringVal{Op: cruder.EQ, Value: statement.UserID},
			UserIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{statement.UserID}},
		}),
		common1.WithLimit(1),
	)
	assert.Nil(t, err)

	infos, total, err := handler.GetAchievementUsers(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, uint32(1), total)
		if assert.Equal(t, 1, len(infos)) {
			ret.ID = infos[0].ID
			ret.EntID = infos[0].EntID
			ret.CreatedAt = infos[0].CreatedAt
			ret.UpdatedAt = infos[0].UpdatedAt
			assert.Equal(t, &ret, infos[0])
		}
	}
}

func TestAchievementUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("getAchievementUsers", getAchievementUsers)
}
