package commission

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/commission"
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

var ret = npool.Commission{
	EntID:               uuid.NewString(),
	AppID:               uuid.NewString(),
	UserID:              uuid.NewString(),
	GoodID:              uuid.NewString(),
	AppGoodID:           uuid.NewString(),
	SettleType:          types.SettleType_GoodOrderPayment,
	SettleTypeStr:       types.SettleType_GoodOrderPayment.String(),
	SettleMode:          types.SettleMode_SettleWithGoodValue,
	SettleModeStr:       types.SettleMode_SettleWithGoodValue.String(),
	SettleAmountType:    types.SettleAmountType_SettleByPercent,
	SettleAmountTypeStr: types.SettleAmountType_SettleByPercent.String(),
	SettleInterval:      types.SettleInterval_SettleYearly,
	SettleIntervalStr:   types.SettleInterval_SettleYearly.String(),
	AmountOrPercent:     decimal.RequireFromString("12.25").String(),
	StartAt:             uint32(time.Now().Unix()),
	Threshold:           decimal.RequireFromString("12.26").String(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCommission(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithSettleType(&ret.SettleType, true),
		WithSettleMode(&ret.SettleMode, true),
		WithSettleAmountType(&ret.SettleAmountType, true),
		WithSettleInterval(&ret.SettleInterval, true),
		WithAmountOrPercent(&ret.AmountOrPercent, true),
		WithStartAt(&ret.StartAt, true),
		WithThreshold(&ret.Threshold, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCommission(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateCommission(t *testing.T) {
	ret.AmountOrPercent = "13"
	ret.StartAt += 10000
	ret.Threshold = decimal.NewFromInt(10).String()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAmountOrPercent(&ret.AmountOrPercent, true),
		WithStartAt(&ret.StartAt, true),
		WithThreshold(&ret.Threshold, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateCommission(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCommission(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getCommission(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCommission(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCommissions(t *testing.T) {
	conds := &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		SettleType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SettleType)},
		EndAt:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.EndAt},
		UserIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.UserID}},
		StartAt:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.StartAt},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCommissions(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCommission(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteCommission(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetCommission(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCommission(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCommission", createCommission)
	t.Run("updateCommission", updateCommission)
	t.Run("getCommission", getCommission)
	t.Run("getCommissions", getCommissions)
	t.Run("deleteCommission", deleteCommission)
}
