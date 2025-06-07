package fee

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	fee1 "github.com/NpoolPlatform/kunman/middleware/good/fee"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/good/testinit"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
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

var ret = npool.Fee{
	EntID:                   uuid.NewString(),
	AppID:                   uuid.NewString(),
	GoodID:                  uuid.NewString(),
	AppGoodID:               uuid.NewString(),
	GoodType:                types.GoodType_TechniqueServiceFee,
	GoodName:                uuid.NewString(),
	AppGoodName:             uuid.NewString(),
	SettlementType:          types.GoodSettlementType_GoodSettledByProfitPercent,
	UnitValue:               decimal.NewFromInt(20).String(),
	CancelMode:              types.CancelMode_Uncancellable,
	DurationDisplayType:     types.GoodDurationType_GoodDurationByDay,
	MinOrderDurationSeconds: 86400,
}

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.SettlementTypeStr = ret.SettlementType.String()
	ret.DurationDisplayTypeStr = ret.DurationDisplayType.String()
	ret.CancelModeStr = ret.CancelMode.String()

	feeEntID := uuid.NewString()
	h1, err := fee1.NewHandler(
		context.Background(),
		fee1.WithEntID(&feeEntID, true),
		fee1.WithGoodID(&ret.GoodID, true),
		fee1.WithGoodType(&ret.GoodType, true),
		fee1.WithName(&ret.GoodName, true),
		fee1.WithSettlementType(&ret.SettlementType, true),
		fee1.WithUnitValue(&ret.UnitValue, true),
		fee1.WithDurationDisplayType(&ret.DurationDisplayType, true),
	)
	assert.Nil(t, err)

	err = h1.CreateFee(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteFee(context.Background())
	}
}

func createFee(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithName(&ret.AppGoodName, true),
		WithUnitValue(&ret.UnitValue, true),
		WithCancelMode(&ret.CancelMode, true),
		WithMinOrderDurationSeconds(&ret.MinOrderDurationSeconds, true),
	)
	assert.Nil(t, err)

	err = handler.CreateFee(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetFee(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateFee(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithName(&ret.AppGoodName, true),
		WithUnitValue(&ret.UnitValue, true),
		WithMinOrderDurationSeconds(&ret.MinOrderDurationSeconds, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateFee(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetFee(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getFee(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetFee(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getFees(t *testing.T) {
	conds := &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		IDs:        &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{ret.ID}},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		AppIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppID, ret.AppID}},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID, ret.EntID}},
		GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		GoodIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
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

	infos, _, err := handler.GetFees(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(infos))
		assert.Equal(t, &ret, infos[0])
	}
}

func deleteFee(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteFee(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetFee(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestFee(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createFee", createFee)
	t.Run("updateFee", updateFee)
	t.Run("getFee", getFee)
	t.Run("getFees", getFees)
	t.Run("deleteFee", deleteFee)
}
