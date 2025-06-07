package constraint

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	appgoodbase1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/app/good/goodbase"
	topmost1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/app/good/topmost"
	topmostgood1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/app/good/topmost/good"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/good/goodbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good/constraint"

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

var ret = npool.TopMostGoodConstraint{
	EntID:            uuid.NewString(),
	AppID:            uuid.NewString(),
	TopMostID:        uuid.NewString(),
	TopMostType:      types.GoodTopMostType_TopMostBestOffer,
	TopMostTitle:     uuid.NewString(),
	TopMostMessage:   uuid.NewString(),
	TopMostTargetUrl: uuid.NewString(),
	Constraint:       types.GoodTopMostConstraint_TopMostCreditThreshold,
	TopMostGoodID:    uuid.NewString(),
	AppGoodID:        uuid.NewString(),
	AppGoodName:      uuid.NewString(),
	TargetValue:      decimal.NewFromInt(800).String(),
	Index:            1,
}

func setup(t *testing.T) func(*testing.T) {
	ret.TopMostTypeStr = ret.TopMostType.String()
	ret.ConstraintStr = ret.Constraint.String()

	goodID := uuid.NewString()
	h1, err := goodbase1.NewHandler(
		context.Background(),
		goodbase1.WithEntID(&goodID, true),
		goodbase1.WithGoodType(func() *types.GoodType { e := types.GoodType_PowerRental; return &e }(), true),
		goodbase1.WithName(&goodID, true),
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
		appgoodbase1.WithName(&ret.AppGoodName, true),
	)
	assert.Nil(t, err)

	err = h2.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	h3, err := topmost1.NewHandler(
		context.Background(),
		topmost1.WithEntID(&ret.TopMostID, true),
		topmost1.WithAppID(&ret.AppID, true),
		topmost1.WithTopMostType(&ret.TopMostType, true),
		topmost1.WithTitle(&ret.TopMostTitle, true),
		topmost1.WithMessage(&ret.TopMostMessage, true),
		topmost1.WithTargetURL(&ret.TopMostTargetUrl, true),
		topmost1.WithStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		topmost1.WithEndAt(func() *uint32 { u := uint32(time.Now().Unix() + 10000); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h3.CreateTopMost(context.Background())
	assert.Nil(t, err)

	h4, err := topmostgood1.NewHandler(
		context.Background(),
		topmostgood1.WithEntID(&ret.TopMostGoodID, true),
		topmostgood1.WithAppGoodID(&ret.AppGoodID, true),
		topmostgood1.WithTopMostID(&ret.TopMostID, true),
		topmostgood1.WithUnitPrice(func() *string { s := decimal.NewFromInt(200).String(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h4.CreateTopMostGood(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h4.DeleteTopMostGood(context.Background())
		_ = h3.DeleteTopMost(context.Background())
		_ = h2.DeleteGoodBase(context.Background())
		_ = h1.DeleteGoodBase(context.Background())
	}
}

func createConstraint(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithTopMostGoodID(&ret.TopMostGoodID, true),
		WithConstraint(&ret.Constraint, true),
		WithTargetValue(&ret.TargetValue, true),
		WithIndex(&ret.Index, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateConstraint(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetConstraint(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func updateConstraint(t *testing.T) {
	ret.Index = 2
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithTargetValue(&ret.TargetValue, true),
		WithIndex(&ret.Index, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateConstraint(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetConstraint(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getConstraint(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetConstraint(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getConstraints(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			TopMostID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.TopMostID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetConstraints(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteConstraint(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteConstraint(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetConstraint(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestConstraint(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createConstraint", createConstraint)
	t.Run("updateConstraint", updateConstraint)
	t.Run("getConstraint", getConstraint)
	t.Run("getConstraints", getConstraints)
	t.Run("deleteConstraint", deleteConstraint)
}
