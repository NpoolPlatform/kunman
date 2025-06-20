package goodmalfunction

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/malfunction"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/good/goodbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

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

var ret = npool.Malfunction{
	EntID:             uuid.NewString(),
	GoodID:            uuid.NewString(),
	GoodType:          types.GoodType_PowerRental,
	GoodName:          uuid.NewString(),
	Title:             uuid.NewString(),
	Message:           uuid.NewString(),
	StartAt:           uint32(time.Now().Unix()),
	DurationSeconds:   1,
	CompensateSeconds: 1,
}

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()

	h1, err := goodbase1.NewHandler(
		context.Background(),
		goodbase1.WithEntID(&ret.GoodID, true),
		goodbase1.WithGoodType(&ret.GoodType, true),
		goodbase1.WithName(&ret.GoodName, true),
		goodbase1.WithBenefitType(func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }(), true),
		goodbase1.WithStartMode(func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(), true),
		goodbase1.WithServiceStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		goodbase1.WithBenefitIntervalHours(func() *uint32 { u := uint32(24); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteGoodBase(context.Background())
	}
}

func createMalfunction(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithTitle(&ret.Title, true),
		WithMessage(&ret.Message, true),
		WithStartAt(&ret.StartAt, true),
		WithDurationSeconds(&ret.DurationSeconds, true),
		WithCompensateSeconds(&ret.CompensateSeconds, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateMalfunction(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetMalfunction(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func updateMalfunction(t *testing.T) {
	ret.StartAt = uint32(time.Now().Unix() + 1)
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithTitle(&ret.Title, true),
		WithMessage(&ret.Message, true),
		WithStartAt(&ret.StartAt, true),
		WithDurationSeconds(&ret.DurationSeconds, true),
		WithCompensateSeconds(&ret.CompensateSeconds, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateMalfunction(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetMalfunction(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getMalfunction(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetMalfunction(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getMalfunctions(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
			GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetMalfunctions(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteMalfunction(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteMalfunction(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetMalfunction(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestMalfunction(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createMalfunction", createMalfunction)
	t.Run("updateMalfunction", updateMalfunction)
	t.Run("getMalfunction", getMalfunction)
	t.Run("getMalfunctions", getMalfunctions)
	t.Run("deleteMalfunction", deleteMalfunction)
}
