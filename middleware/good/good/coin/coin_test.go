package coin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/good/goodbase"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	"github.com/google/uuid"
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

var ret = npool.GoodCoin{
	EntID:      uuid.NewString(),
	GoodID:     uuid.NewString(),
	GoodType:   types.GoodType_PowerRental,
	GoodName:   uuid.NewString(),
	CoinTypeID: uuid.NewString(),
	Main:       true,
	Index:      5,
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

func createGoodCoin(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithMain(&ret.Main, true),
		WithIndex(&ret.Index, true),
	)
	assert.Nil(t, err)

	err = h1.CreateGoodCoin(context.Background())
	if assert.Nil(t, err) {
		info, err := h1.GetGoodCoin(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}

	h2, err := NewHandler(
		context.Background(),
		WithGoodID(&ret.GoodID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
	)
	assert.Nil(t, err)

	err = h2.CreateGoodCoin(context.Background())
	assert.NotNil(t, err)
}

func updateGoodCoin(t *testing.T) {
	ret.Main = false
	h1, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithMain(&ret.Main, true),
		WithIndex(&ret.Index, true),
	)
	assert.Nil(t, err)

	err = h1.UpdateGoodCoin(context.Background())
	if assert.Nil(t, err) {
		info, err := h1.GetGoodCoin(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getGoodCoin(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetGoodCoin(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getGoodCoins(t *testing.T) {
	conds := &npool.Conds{
		ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		GoodIDs: &basetypes.StringSliceVal{Op: cruder.EQ, Value: []string{ret.GoodID}},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetGoodCoins(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteGoodCoin(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteGoodCoin(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetGoodCoin(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestGoodCoin(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createGoodCoin", createGoodCoin)
	t.Run("updateGoodCoin", updateGoodCoin)
	t.Run("getGoodCoin", getGoodCoin)
	t.Run("getGoodCoins", getGoodCoins)
	t.Run("deleteGoodCoin", deleteGoodCoin)
}
