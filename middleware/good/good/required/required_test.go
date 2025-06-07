//nolint:dupl
package required

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/good/goodbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"

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

var ret = npool.Required{
	EntID:            uuid.NewString(),
	MainGoodID:       uuid.NewString(),
	MainGoodName:     uuid.NewString(),
	RequiredGoodID:   uuid.NewString(),
	RequiredGoodName: uuid.NewString(),
	Must:             true,
}

func setup(t *testing.T) func(*testing.T) {
	h1, err := goodbase1.NewHandler(
		context.Background(),
		goodbase1.WithEntID(&ret.MainGoodID, true),
		goodbase1.WithGoodType(func() *types.GoodType { e := types.GoodType_PowerRental; return &e }(), true),
		goodbase1.WithName(&ret.MainGoodName, true),
		goodbase1.WithBenefitType(func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }(), true),
		goodbase1.WithStartMode(func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(), true),
		goodbase1.WithServiceStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		goodbase1.WithBenefitIntervalHours(func() *uint32 { u := uint32(24); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	h2, err := goodbase1.NewHandler(
		context.Background(),
		goodbase1.WithEntID(&ret.RequiredGoodID, true),
		goodbase1.WithGoodType(func() *types.GoodType { e := types.GoodType_TechniqueServiceFee; return &e }(), true),
		goodbase1.WithName(&ret.RequiredGoodName, true),
		goodbase1.WithBenefitType(func() *types.BenefitType { e := types.BenefitType_BenefitTypeNone; return &e }(), true),
		goodbase1.WithStartMode(func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(), true),
		goodbase1.WithServiceStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		goodbase1.WithBenefitIntervalHours(func() *uint32 { u := uint32(24); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h2.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h2.DeleteGoodBase(context.Background())
		_ = h1.DeleteGoodBase(context.Background())
	}
}

func createRequired(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithMainGoodID(&ret.MainGoodID, true),
		WithRequiredGoodID(&ret.RequiredGoodID, true),
		WithMust(&ret.Must, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateRequired(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetRequired(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func updateRequired(t *testing.T) {
	ret.Must = false

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMust(&ret.Must, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateRequired(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetRequired(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getRequired(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetRequired(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getRequireds(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:             &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			MainGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainGoodID},
			RequiredGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RequiredGoodID},
			GoodID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainGoodID},
			GoodIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.MainGoodID, ret.RequiredGoodID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetRequireds(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteRequired(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteRequired(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetRequired(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestRequired(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createRequired", createRequired)
	t.Run("updateRequired", updateRequired)
	t.Run("getRequired", getRequired)
	t.Run("getRequireds", getRequireds)
	t.Run("deleteRequired", deleteRequired)
}
