package goodbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/notif/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	yesterday = uint32(time.Now().AddDate(0, 0, -1).Unix())
	ret       = npool.GoodBenefit{
		GoodID:      uuid.NewString(),
		GoodType:    goodtypes.GoodType_PowerRental,
		GoodTypeStr: goodtypes.GoodType_PowerRental.String(),
		GoodName:    uuid.NewString(),
		CoinTypeID:  uuid.NewString(),
		Amount:      "100",
		State:       basetypes.Result_Success,
		StateStr:    basetypes.Result_Success.String(),
		Message:     uuid.NewString(),
		BenefitDate: yesterday,
		TxID:        uuid.NewString(),
		Generated:   false,
	}

	ret2 = npool.GoodBenefit{
		GoodID:      uuid.NewString(),
		GoodType:    goodtypes.GoodType_PowerRental,
		GoodTypeStr: goodtypes.GoodType_PowerRental.String(),
		GoodName:    uuid.NewString(),
		CoinTypeID:  uuid.NewString(),
		Amount:      "10",
		State:       basetypes.Result_Success,
		StateStr:    basetypes.Result_Success.String(),
		Message:     uuid.NewString(),
		BenefitDate: uint32(time.Now().Add(-3 * time.Minute).Unix()),
		TxID:        uuid.NewString(),
		Generated:   false,
	}
)

func createGoodBenefit(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithGoodName(&ret.GoodName, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithAmount(&ret.Amount, false),
		WithState(&ret.State, true),
		WithMessage(&ret.Message, false),
		WithBenefitDate(&ret.BenefitDate, true),
		WithTxID(&ret.TxID, false),
		WithGenerated(&ret.Generated, false),
	)
	assert.Nil(t, err)

	info, err := handler.CreateGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		ret.EntID = info.EntID
		assert.Equal(t, info, &ret)
	}

	handler2, err := NewHandler(
		context.Background(),
		WithGoodID(&ret2.GoodID, true),
		WithGoodType(&ret2.GoodType, true),
		WithGoodName(&ret2.GoodName, true),
		WithCoinTypeID(&ret2.CoinTypeID, true),
		WithAmount(&ret2.Amount, false),
		WithState(&ret2.State, true),
		WithMessage(&ret2.Message, false),
		WithBenefitDate(&ret2.BenefitDate, true),
		WithTxID(&ret2.TxID, false),
		WithGenerated(&ret2.Generated, false),
	)
	assert.Nil(t, err)

	_info, err := handler2.CreateGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret2.CreatedAt = _info.CreatedAt
		ret2.UpdatedAt = _info.UpdatedAt
		ret2.ID = _info.ID
		ret2.EntID = _info.EntID
		assert.Equal(t, _info, &ret2)
	}
}

func updateGoodBenefit(t *testing.T) {
	ret.Generated = true
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithGenerated(&ret.Generated, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getGoodBenefit(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getGoodBenefits(t *testing.T) {
	conds := &npool.Conds{
		Generated:        &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Generated},
		BenefitDateStart: &basetypes.Uint32Val{Op: cruder.LTE, Value: yesterday},
		BenefitDateEnd:   &basetypes.Uint32Val{Op: cruder.GTE, Value: uint32(time.Now().Unix())},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetGoodBenefits(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteGoodBenefit(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetGoodBenefit(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)

	handler2, err := NewHandler(
		context.Background(),
		WithID(&ret2.ID, true),
	)
	assert.Nil(t, err)

	_info, err := handler2.DeleteGoodBenefit(context.Background())
	if assert.Nil(t, err) {
		ret2.UpdatedAt = _info.UpdatedAt
		assert.Equal(t, _info, &ret2)
	}

	_info, err = handler2.GetGoodBenefit(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, _info)
}

func TestGoodBenefit(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createGoodBenefit", createGoodBenefit)
	t.Run("updateGoodBenefit", updateGoodBenefit)
	t.Run("getGoodBenefit", getGoodBenefit)
	t.Run("getGoodBenefits", getGoodBenefits)
	t.Run("deleteGoodBenefit", deleteGoodBenefit)
}
