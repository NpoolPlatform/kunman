package label

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/label"
	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/goodbase"
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

var ret = npool.Label{
	EntID:        uuid.NewString(),
	AppID:        uuid.NewString(),
	GoodID:       uuid.NewString(),
	GoodType:     types.GoodType_PowerRental,
	GoodName:     uuid.NewString(),
	AppGoodID:    uuid.NewString(),
	AppGoodName:  uuid.NewString(),
	Icon:         uuid.NewString(),
	IconBgColor:  uuid.NewString(),
	Label:        types.GoodLabel_GoodLabelPromotion,
	LabelBgColor: uuid.NewString(),
	Index:        5,
}

func setup(t *testing.T) func(*testing.T) {
	ret.LabelStr = ret.Label.String()
	ret.GoodTypeStr = ret.GoodType.String()

	h1, err := goodbase1.NewHandler(
		context.Background(),
		goodbase1.WithEntID(&ret.GoodID, true),
		goodbase1.WithName(&ret.GoodName, true),
		goodbase1.WithBenefitType(func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }(), true),
		goodbase1.WithGoodType(&ret.GoodType, true),
		goodbase1.WithTestOnly(func() *bool { b := true; return &b }(), false),
		goodbase1.WithBenefitIntervalHours(func() *uint32 { u := uint32(24); return &u }(), true),
		goodbase1.WithServiceStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		goodbase1.WithStartMode(func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	h2, err := appgoodbase1.NewHandler(
		context.Background(),
		appgoodbase1.WithEntID(&ret.AppGoodID, true),
		appgoodbase1.WithAppID(&ret.AppID, true),
		appgoodbase1.WithGoodID(&ret.GoodID, true),
		appgoodbase1.WithName(&ret.AppGoodName, true),
	)
	assert.Nil(t, err)

	err = h2.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h2.DeleteGoodBase(context.Background())
		_ = h1.DeleteGoodBase(context.Background())
	}
}

func createLabel(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithIcon(&ret.Icon, true),
		WithIconBgColor(&ret.IconBgColor, true),
		WithLabel(&ret.Label, true),
		WithLabelBgColor(&ret.LabelBgColor, true),
		WithIndex(func() *uint8 { u := uint8(ret.Index); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateLabel(context.Background())
	if assert.Nil(t, err) {
		info, err := h1.GetLabel(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}

	h2, err := NewHandler(
		context.Background(),
		WithAppGoodID(&ret.AppGoodID, true),
		WithLabel(&ret.Label, true),
	)
	assert.Nil(t, err)

	err = h2.CreateLabel(context.Background())
	assert.NotNil(t, err)
}

func updateLabel(t *testing.T) {
	ret.Index = uint32(10)
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithLabel(&ret.Label, true),
		WithIndex(func() *uint8 { u := uint8(ret.Index); return &u }(), true),
	)
	assert.Nil(t, err)

	err = handler.UpdateLabel(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetLabel(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getLabel(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetLabel(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getLabels(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			GoodIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
			AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
			AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppGoodID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetLabels(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteLabel(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteLabel(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetLabel(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestLabel(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createLabel", createLabel)
	t.Run("updateLabel", updateLabel)
	t.Run("getLabel", getLabel)
	t.Run("getLabels", getLabels)
	t.Run("deleteLabel", deleteLabel)
}
