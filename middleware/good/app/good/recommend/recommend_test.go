package recommend

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/recommend"
	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/app/good/goodbase"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/good/goodbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

var ret = npool.Recommend{
	EntID:          uuid.NewString(),
	AppID:          uuid.NewString(),
	RecommenderID:  uuid.NewString(),
	AppGoodID:      uuid.NewString(),
	GoodName:       uuid.NewString(),
	Message:        uuid.NewString(),
	RecommendIndex: decimal.RequireFromString("4.99").String(),
}

func setup(t *testing.T) func(*testing.T) {
	goodType := types.GoodType_PowerRental
	goodID := uuid.NewString()
	ret.HideReasonStr = ret.HideReason.String()

	h1, err := goodbase1.NewHandler(
		context.Background(),
		goodbase1.WithEntID(&goodID, true),
		goodbase1.WithGoodType(&goodType, true),
		goodbase1.WithName(&ret.GoodName, true),
		goodbase1.WithBenefitType(func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }(), true),
		goodbase1.WithStartMode(func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(), true),
		goodbase1.WithServiceStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		goodbase1.WithBenefitIntervalHours(func() *uint32 { u := uint32(24); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	h2, err := appgoodbase1.NewHandler(
		context.Background(),
		appgoodbase1.WithEntID(&ret.AppGoodID, true),
		appgoodbase1.WithAppID(&ret.AppID, true),
		appgoodbase1.WithGoodID(&goodID, true),
		appgoodbase1.WithName(&ret.GoodName, true),
	)
	assert.Nil(t, err)

	err = h2.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h2.DeleteGoodBase(context.Background())
		_ = h1.DeleteGoodBase(context.Background())
	}
}

func createRecommend(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithRecommenderID(&ret.RecommenderID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithMessage(&ret.Message, true),
		WithRecommendIndex(&ret.RecommendIndex, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateRecommend(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetRecommend(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
		}
	}

	h1, err := appgoodbase1.NewHandler(
		context.Background(),
		appgoodbase1.WithEntID(&ret.AppGoodID, true),
	)
	if assert.Nil(t, err) {
		info, err := h1.GetGoodBase(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, uint32(1), info.RecommendCount())
		}
	}
}

func updateRecommend(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMessage(&ret.Message, true),
		WithRecommendIndex(&ret.RecommendIndex, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateRecommend(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetRecommend(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getRecommend(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetRecommend(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getRecommends(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RecommenderID},
			AppGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
			AppGoodIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppGoodID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetRecommends(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteRecommend(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteRecommend(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetRecommend(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestRecommend(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createRecommend", createRecommend)
	t.Run("updateRecommend", updateRecommend)
	t.Run("getRecommend", getRecommend)
	t.Run("getRecommends", getRecommends)
	t.Run("deleteRecommend", deleteRecommend)
}
