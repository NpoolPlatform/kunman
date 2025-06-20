package delegatedstaking

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/delegatedstaking"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	goodcoinrewardmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward"
	delegatedstaking1 "github.com/NpoolPlatform/kunman/middleware/good/delegatedstaking"
	goodcoin1 "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
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

var ret = npool.DelegatedStaking{
	EntID:     uuid.NewString(),
	AppID:     uuid.NewString(),
	GoodID:    uuid.NewString(),
	AppGoodID: uuid.NewString(),

	GoodType:              types.GoodType_DelegatedStaking,
	BenefitType:           types.BenefitType_BenefitTypeContract,
	GoodName:              uuid.NewString(),
	GoodServiceStartAt:    uint32(time.Now().Unix()),
	GoodStartMode:         types.GoodStartMode_GoodStartModeInstantly,
	TestOnly:              true,
	BenefitIntervalHours:  20,
	GoodPurchasable:       true,
	GoodOnline:            true,
	AppGoodPurchasable:    true,
	AppGoodOnline:         true,
	EnableProductPage:     true,
	ProductPage:           uuid.NewString(),
	Visible:               true,
	AppGoodName:           uuid.NewString(),
	DisplayIndex:          1,
	Banner:                uuid.NewString(),
	EnableSetCommission:   true,
	AppGoodServiceStartAt: uint32(time.Now().Unix()),
	AppGoodStartMode:      types.GoodStartMode_GoodStartModeInstantly,
	ContractCodeURL:       uuid.NewString(),
	ContractCodeBranch:    uuid.NewString(),
	ContractState:         types.ContractState_ContractWaitDeployment,

	GoodCoins: []*goodcoinmwpb.GoodCoinInfo{
		{
			CoinTypeID: uuid.NewString(),
			Main:       true,
		},
		{
			CoinTypeID: uuid.NewString(),
			Main:       false,
		},
		{
			CoinTypeID: uuid.NewString(),
			Main:       false,
		},
	},

	Score: decimal.NewFromInt(0).String(),
}

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.BenefitTypeStr = ret.BenefitType.String()
	ret.GoodStartModeStr = ret.GoodStartMode.String()
	ret.AppGoodStartModeStr = ret.AppGoodStartMode.String()
	ret.ContractStateStr = ret.ContractState.String()
	for _, goodCoin := range ret.GoodCoins {
		ret.Rewards = append(ret.Rewards, &goodcoinrewardmwpb.RewardInfo{
			GoodID:                ret.GoodID,
			CoinTypeID:            goodCoin.CoinTypeID,
			RewardTID:             uuid.Nil.String(),
			LastRewardAmount:      decimal.NewFromInt(0).String(),
			NextRewardStartAmount: decimal.NewFromInt(0).String(),
			LastUnitRewardAmount:  decimal.NewFromInt(0).String(),
			TotalRewardAmount:     decimal.NewFromInt(0).String(),
			MainCoin:              goodCoin.Main,
		})
	}

	delegatedstakingID := uuid.NewString()

	h5, err := delegatedstaking1.NewHandler(
		context.Background(),
		delegatedstaking1.WithEntID(&delegatedstakingID, true),
		delegatedstaking1.WithGoodID(&ret.GoodID, true),
		delegatedstaking1.WithGoodType(&ret.GoodType, true),
		delegatedstaking1.WithBenefitType(&ret.BenefitType, true),
		delegatedstaking1.WithName(&ret.GoodName, true),
		delegatedstaking1.WithServiceStartAt(&ret.GoodServiceStartAt, true),
		delegatedstaking1.WithStartMode(&ret.GoodStartMode, true),
		delegatedstaking1.WithTestOnly(&ret.TestOnly, true),
		delegatedstaking1.WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		delegatedstaking1.WithPurchasable(&ret.GoodPurchasable, true),
		delegatedstaking1.WithOnline(&ret.GoodOnline, true),
		delegatedstaking1.WithContractCodeURL(&ret.ContractCodeURL, true),
		delegatedstaking1.WithContractCodeBranch(&ret.ContractCodeBranch, true),
		delegatedstaking1.WithContractState(&ret.ContractState, true),
	)
	assert.Nil(t, err)

	err = h5.CreateDelegatedStaking(context.Background())
	assert.Nil(t, err)

	h6s := []*goodcoin1.Handler{}
	for _, goodCoin := range ret.GoodCoins {
		goodCoin.GoodID = ret.GoodID
		h6, err := goodcoin1.NewHandler(
			context.Background(),
			goodcoin1.WithGoodID(&ret.GoodID, true),
			goodcoin1.WithCoinTypeID(&goodCoin.CoinTypeID, true),
			goodcoin1.WithMain(&goodCoin.Main, true),
			goodcoin1.WithIndex(&goodCoin.Index, true),
		)
		assert.Nil(t, err)

		err = h6.CreateGoodCoin(context.Background())
		assert.Nil(t, err)

		h6s = append(h6s, h6)
	}

	return func(*testing.T) {
		for _, h6 := range h6s {
			_ = h6.DeleteGoodCoin(context.Background())
		}
		_ = h5.DeleteDelegatedStaking(context.Background())
	}
}

func createDelegatedStaking(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithPurchasable(&ret.GoodPurchasable, true),
		WithEnableProductPage(&ret.EnableProductPage, true),
		WithProductPage(&ret.ProductPage, true),
		WithName(&ret.AppGoodName, true),
		WithOnline(&ret.GoodOnline, true),
		WithVisible(&ret.Visible, true),
		WithDisplayIndex(&ret.DisplayIndex, true),
		WithBanner(&ret.Banner, true),
		WithServiceStartAt(&ret.AppGoodServiceStartAt, true),
		WithStartMode(&ret.AppGoodStartMode, true),
		WithEnableSetCommission(&ret.EnableSetCommission, true),
	)
	assert.Nil(t, err)

	err = h1.CreateDelegatedStaking(context.Background())
	if assert.Nil(t, err) {
		info, err := h1.GetDelegatedStaking(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			ret.State = info.State
			ret.StateStr = info.StateStr
			assert.Equal(t, &ret, info)
		}
	}
}

func updateDelegatedStaking(t *testing.T) {
	ret.AppGoodServiceStartAt = uint32(time.Now().Unix() + 10)
	h1, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithPurchasable(&ret.GoodPurchasable, true),
		WithEnableProductPage(&ret.EnableProductPage, true),
		WithProductPage(&ret.ProductPage, true),
		WithName(&ret.AppGoodName, true),
		WithOnline(&ret.GoodOnline, true),
		WithVisible(&ret.Visible, true),
		WithDisplayIndex(&ret.DisplayIndex, true),
		WithBanner(&ret.Banner, true),
		WithServiceStartAt(&ret.AppGoodServiceStartAt, true),
		WithStartMode(&ret.AppGoodStartMode, true),
		WithEnableSetCommission(&ret.EnableSetCommission, true),
	)
	assert.Nil(t, err)

	err = h1.UpdateDelegatedStaking(context.Background())
	if assert.Nil(t, err) {
		info, err := h1.GetDelegatedStaking(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getDelegatedStaking(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetDelegatedStaking(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getDelegatedStakings(t *testing.T) {
	conds := &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetDelegatedStakings(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteDelegatedStaking(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteDelegatedStaking(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetDelegatedStaking(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDelegatedStaking(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createDelegatedStaking", createDelegatedStaking)
	t.Run("updateDelegatedStaking", updateDelegatedStaking)
	t.Run("getDelegatedStaking", getDelegatedStaking)
	t.Run("getDelegatedStakings", getDelegatedStakings)
	t.Run("deleteDelegatedStaking", deleteDelegatedStaking)
}
