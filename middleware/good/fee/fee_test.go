package fee

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/good/testinit"
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
	EntID:               uuid.NewString(),
	GoodID:              uuid.NewString(),
	GoodType:            types.GoodType_TechniqueServiceFee,
	Name:                uuid.NewString(),
	SettlementType:      types.GoodSettlementType_GoodSettledByProfitPercent,
	UnitValue:           decimal.NewFromInt(20).String(),
	DurationDisplayType: types.GoodDurationType_GoodDurationByDay,
}

//nolint:unparam
func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.SettlementTypeStr = ret.SettlementType.String()
	ret.DurationDisplayTypeStr = ret.DurationDisplayType.String()
	return func(*testing.T) {}
}

func createFee(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithName(&ret.Name, true),
		WithSettlementType(&ret.SettlementType, true),
		WithUnitValue(&ret.UnitValue, true),
		WithDurationDisplayType(&ret.DurationDisplayType, true),
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
	ret.Name = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithName(&ret.Name, true),
		WithSettlementType(&ret.SettlementType, true),
		WithUnitValue(&ret.UnitValue, true),
		WithDurationDisplayType(&ret.DurationDisplayType, true),
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
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetFee(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getFees(t *testing.T) {
	conds := &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetFees(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteFee(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
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
