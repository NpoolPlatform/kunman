package event

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	coinconfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/coin/config"
	coupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
	eventcoin1 "github.com/NpoolPlatform/kunman/middleware/inspire/event/coin"
	eventcoupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/event/coupon"
	taskconfig1 "github.com/NpoolPlatform/kunman/middleware/inspire/task/config"
	coinconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coin/config"
	couponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"
	eventcoinmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event/coin"
	eventcouponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event/coupon"
	taskconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/task/config"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/inspire/testinit"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
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

var (
	appID     = uuid.NewString()
	userID    = uuid.NewString()
	couponID1 = uuid.NewString()
	eventRet  = npool.Event{
		EntID:          uuid.NewString(),
		AppID:          appID,
		EventType:      basetypes.UsedFor_Signup,
		EventTypeStr:   basetypes.UsedFor_Signup.String(),
		CouponIDs:      []string{couponID1},
		Credits:        decimal.RequireFromString("100").String(),
		CreditsPerUSD:  decimal.RequireFromString("12.25").String(),
		MaxConsecutive: 1,
		InviterLayers:  2,
	}
	eventRet2 = npool.Event{
		EntID:          uuid.NewString(),
		AppID:          appID,
		EventType:      basetypes.UsedFor_SetWithdrawAddress,
		EventTypeStr:   basetypes.UsedFor_SetWithdrawAddress.String(),
		CouponIDs:      []string{eventcoupon.EntID},
		Credits:        decimal.RequireFromString("10").String(),
		CreditsPerUSD:  decimal.RequireFromString("1.25").String(),
		MaxConsecutive: 1,
		InviterLayers:  2,
	}
	coinConfig = coinconfigmwpb.CoinConfig{
		EntID:      uuid.NewString(),
		AppID:      appID,
		CoinTypeID: uuid.NewString(),
		MaxValue:   decimal.RequireFromString("20").String(),
		Allocated:  decimal.RequireFromString("0").String(),
	}
	coinConfig2 = coinconfigmwpb.CoinConfig{
		EntID:      uuid.NewString(),
		AppID:      appID,
		CoinTypeID: uuid.NewString(),
		MaxValue:   decimal.RequireFromString("120").String(),
		Allocated:  decimal.RequireFromString("0").String(),
	}
	taskConfig = taskconfigmwpb.TaskConfig{
		EntID:                  uuid.NewString(),
		AppID:                  appID,
		EventID:                eventRet.EntID,
		Name:                   uuid.NewString(),
		TaskDesc:               uuid.NewString(),
		StepGuide:              uuid.NewString(),
		RecommendMessage:       uuid.NewString(),
		Index:                  uint32(1),
		LastTaskID:             uuid.Nil.String(),
		MaxRewardCount:         uint32(1),
		CooldownSecond:         uint32(120),
		TaskType:               types.TaskType_BaseTask,
		TaskTypeStr:            types.TaskType_BaseTask.String(),
		IntervalReset:          true,
		IntervalResetSecond:    24 * 60 * 60,
		MaxIntervalRewardCount: 2,
	}
	taskConfig2 = taskconfigmwpb.TaskConfig{
		EntID:                  uuid.NewString(),
		AppID:                  appID,
		EventID:                eventRet2.EntID,
		Name:                   uuid.NewString(),
		TaskDesc:               uuid.NewString(),
		StepGuide:              uuid.NewString(),
		RecommendMessage:       uuid.NewString(),
		Index:                  uint32(2),
		LastTaskID:             uuid.Nil.String(),
		MaxRewardCount:         uint32(2),
		CooldownSecond:         uint32(120),
		TaskType:               types.TaskType_BaseTask,
		TaskTypeStr:            types.TaskType_BaseTask.String(),
		IntervalReset:          false,
		IntervalResetSecond:    0,
		MaxIntervalRewardCount: 0,
	}
	eventcoupon = couponmwpb.Coupon{
		EntID:               uuid.NewString(),
		CouponType:          types.CouponType_FixAmount,
		CouponTypeStr:       types.CouponType_FixAmount.String(),
		AppID:               appID,
		Denomination:        decimal.RequireFromString("12.25").String(),
		Circulation:         decimal.RequireFromString("50.25").String(),
		IssuedBy:            uuid.NewString(),
		StartAt:             uint32(time.Now().Unix()),
		EndAt:               uint32(time.Now().Add(24 * time.Hour).Unix()),
		DurationDays:        234,
		Message:             uuid.NewString(),
		Name:                uuid.NewString(),
		CouponConstraint:    types.CouponConstraint_Normal,
		CouponConstraintStr: types.CouponConstraint_Normal.String(),
		CouponScope:         types.CouponScope_Whitelist,
		CouponScopeStr:      types.CouponScope_Whitelist.String(),
		Allocated:           decimal.NewFromInt(0).String(),
		Threshold:           decimal.NewFromInt(0).String(),
		CashableProbability: decimal.RequireFromString("0.0001").String(),
	}
	eventCoupon = eventcouponmwpb.EventCoupon{
		EntID:    uuid.NewString(),
		AppID:    appID,
		EventID:  eventRet.EntID,
		CouponID: eventcoupon.EntID,
	}
	eventCoin = eventcoinmwpb.EventCoin{
		EntID:        uuid.NewString(),
		AppID:        appID,
		EventID:      eventRet.EntID,
		CoinConfigID: coinConfig.EntID,
		CoinTypeID:   coinConfig.CoinTypeID,
		CoinValue:    decimal.RequireFromString("5").String(),
		CoinPerUSD:   decimal.RequireFromString("0.1").String(),
	}
	eventCoin2 = eventcoinmwpb.EventCoin{
		EntID:        uuid.NewString(),
		AppID:        appID,
		EventID:      eventRet.EntID,
		CoinConfigID: coinConfig2.EntID,
		CoinTypeID:   coinConfig2.CoinTypeID,
		CoinValue:    decimal.RequireFromString("10").String(),
		CoinPerUSD:   decimal.RequireFromString("0.15").String(),
	}
)

//nolint:funlen,dupl
func resetup(t *testing.T) func(*testing.T) {
	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithEntID(&eventcoupon.EntID, true),
		coupon1.WithAppID(&eventcoupon.AppID, true),
		coupon1.WithName(&eventcoupon.Name, true),
		coupon1.WithMessage(&eventcoupon.Message, true),
		coupon1.WithCouponType(&eventcoupon.CouponType, true),
		coupon1.WithDenomination(&eventcoupon.Denomination, true),
		coupon1.WithCouponScope(&eventcoupon.CouponScope, true),
		coupon1.WithCirculation(&eventcoupon.Circulation, true),
		coupon1.WithDurationDays(&eventcoupon.DurationDays, true),
		coupon1.WithIssuedBy(&eventcoupon.IssuedBy, true),
		coupon1.WithStartAt(&eventcoupon.StartAt, true),
		coupon1.WithEndAt(&eventcoupon.EndAt, true),
		coupon1.WithCashableProbability(&eventcoupon.CashableProbability, true),
	)
	assert.Nil(t, err)

	info, err := h1.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		eventcoupon.ID = info.ID
		eventcoupon.CreatedAt = info.CreatedAt
		eventcoupon.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &eventcoupon, info)
		h1.ID = &info.ID
	}

	handler, err := NewHandler(
		context.Background(),
		WithEntID(&eventRet.EntID, true),
		WithAppID(&eventRet.AppID, true),
		WithEventType(&eventRet.EventType, true),
		WithCredits(&eventRet.Credits, true),
		WithCreditsPerUSD(&eventRet.CreditsPerUSD, true),
		WithMaxConsecutive(&eventRet.MaxConsecutive, true),
		WithInviterLayers(&eventRet.InviterLayers, true),
	)
	assert.Nil(t, err)

	err = handler.CreateEvent(context.Background())
	if assert.Nil(t, err) {
		info2, err := handler.GetEvent(context.Background())
		if assert.Nil(t, err) {
			eventRet.ID = info2.ID
			eventRet.CreatedAt = info2.CreatedAt
			eventRet.UpdatedAt = info2.UpdatedAt
			eventRet.CouponIDs = info2.CouponIDs
			eventRet.CouponIDsStr = info2.CouponIDsStr
			eventRet.Coins = info2.Coins
			eventRet.GoodID = info2.GoodID
			eventRet.AppGoodID = info2.AppGoodID
			assert.Equal(t, info2, &eventRet)
			handler.ID = &info2.ID
		}
	}

	h2, err := taskconfig1.NewHandler(
		context.Background(),
		taskconfig1.WithEntID(&taskConfig.EntID, true),
		taskconfig1.WithAppID(&taskConfig.AppID, true),
		taskconfig1.WithEventID(&taskConfig.EventID, true),
		taskconfig1.WithName(&taskConfig.Name, true),
		taskconfig1.WithTaskDesc(&taskConfig.TaskDesc, true),
		taskconfig1.WithStepGuide(&taskConfig.StepGuide, true),
		taskconfig1.WithRecommendMessage(&taskConfig.RecommendMessage, true),
		taskconfig1.WithIndex(&taskConfig.Index, true),
		taskconfig1.WithLastTaskID(&taskConfig.LastTaskID, true),
		taskconfig1.WithMaxRewardCount(&taskConfig.MaxRewardCount, true),
		taskconfig1.WithCooldownSecond(&taskConfig.CooldownSecond, true),
		taskconfig1.WithTaskType(&taskConfig.TaskType, true),
		taskconfig1.WithIntervalReset(&taskConfig.IntervalReset, false),
		taskconfig1.WithIntervalResetSecond(&taskConfig.IntervalResetSecond, false),
		taskconfig1.WithMaxIntervalRewardCount(&taskConfig.MaxIntervalRewardCount, false),
	)
	assert.Nil(t, err)

	err = h2.CreateTaskConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := h2.GetTaskConfig(context.Background())
		if assert.Nil(t, err) {
			taskConfig.ID = info.ID
			taskConfig.CreatedAt = info.CreatedAt
			taskConfig.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &taskConfig, info)
			h2.ID = &info.ID
		}
	}

	h3, err := coinconfig1.NewHandler(
		context.Background(),
		coinconfig1.WithEntID(&coinConfig.EntID, true),
		coinconfig1.WithAppID(&coinConfig.AppID, true),
		coinconfig1.WithCoinTypeID(&coinConfig.CoinTypeID, true),
		coinconfig1.WithMaxValue(&coinConfig.MaxValue, true),
		coinconfig1.WithAllocated(&coinConfig.Allocated, true),
	)
	assert.Nil(t, err)

	err = h3.CreateCoinConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := h3.GetCoinConfig(context.Background())
		if assert.Nil(t, err) {
			coinConfig.ID = info.ID
			coinConfig.CreatedAt = info.CreatedAt
			coinConfig.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &coinConfig, info)
			h3.ID = &info.ID
		}
	}

	h4, err := coinconfig1.NewHandler(
		context.Background(),
		coinconfig1.WithEntID(&coinConfig2.EntID, true),
		coinconfig1.WithAppID(&coinConfig2.AppID, true),
		coinconfig1.WithCoinTypeID(&coinConfig2.CoinTypeID, true),
		coinconfig1.WithMaxValue(&coinConfig2.MaxValue, true),
		coinconfig1.WithAllocated(&coinConfig2.Allocated, true),
	)
	assert.Nil(t, err)

	err = h4.CreateCoinConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := h4.GetCoinConfig(context.Background())
		if assert.Nil(t, err) {
			coinConfig2.ID = info.ID
			coinConfig2.CreatedAt = info.CreatedAt
			coinConfig2.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &coinConfig2, info)
			h4.ID = &info.ID
		}
	}

	h6, err := eventcoin1.NewHandler(
		context.Background(),
		eventcoin1.WithEntID(&eventCoin.EntID, true),
		eventcoin1.WithAppID(&eventCoin.AppID, true),
		eventcoin1.WithEventID(&eventCoin.EventID, true),
		eventcoin1.WithCoinConfigID(&eventCoin.CoinConfigID, true),
		eventcoin1.WithCoinValue(&eventCoin.CoinValue, true),
		eventcoin1.WithCoinPerUSD(&eventCoin.CoinPerUSD, true),
	)
	assert.Nil(t, err)

	err = h6.CreateEventCoin(context.Background())
	if assert.Nil(t, err) {
		info, err := h6.GetEventCoin(context.Background())
		if assert.Nil(t, err) {
			eventCoin.ID = info.ID
			eventCoin.CreatedAt = info.CreatedAt
			eventCoin.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &eventCoin, info)
			h6.ID = &info.ID
		}
	}

	h7, err := eventcoin1.NewHandler(
		context.Background(),
		eventcoin1.WithEntID(&eventCoin2.EntID, true),
		eventcoin1.WithAppID(&eventCoin2.AppID, true),
		eventcoin1.WithEventID(&eventCoin2.EventID, true),
		eventcoin1.WithCoinConfigID(&eventCoin2.CoinConfigID, true),
		eventcoin1.WithCoinValue(&eventCoin2.CoinValue, true),
		eventcoin1.WithCoinPerUSD(&eventCoin2.CoinPerUSD, true),
	)
	assert.Nil(t, err)

	err = h7.CreateEventCoin(context.Background())
	if assert.Nil(t, err) {
		info, err := h7.GetEventCoin(context.Background())
		if assert.Nil(t, err) {
			eventCoin2.ID = info.ID
			eventCoin2.CreatedAt = info.CreatedAt
			eventCoin2.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &eventCoin2, info)
			h7.ID = &info.ID
		}
	}

	h8, err := eventcoupon1.NewHandler(
		context.Background(),
		eventcoupon1.WithEntID(&eventCoupon.EntID, true),
		eventcoupon1.WithAppID(&eventCoupon.AppID, true),
		eventcoupon1.WithEventID(&eventCoupon.EventID, true),
		eventcoupon1.WithCouponID(&eventCoupon.CouponID, true),
	)
	assert.Nil(t, err)

	err = h8.CreateEventCoupon(context.Background())
	if assert.Nil(t, err) {
		info, err := h8.GetEventCoupon(context.Background())
		if assert.Nil(t, err) {
			eventCoupon.ID = info.ID
			eventCoupon.CreatedAt = info.CreatedAt
			eventCoupon.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &eventCoupon, info)
			h8.ID = &info.ID
		}
	}

	handler2, err := NewHandler(
		context.Background(),
		WithEntID(&eventRet2.EntID, true),
		WithAppID(&eventRet2.AppID, true),
		WithEventType(&eventRet2.EventType, true),
		WithCredits(&eventRet2.Credits, true),
		WithCreditsPerUSD(&eventRet2.CreditsPerUSD, true),
		WithMaxConsecutive(&eventRet2.MaxConsecutive, true),
		WithInviterLayers(&eventRet2.InviterLayers, true),
	)
	assert.Nil(t, err)

	err = handler2.CreateEvent(context.Background())
	if assert.Nil(t, err) {
		info3, err := handler2.GetEvent(context.Background())
		if assert.Nil(t, err) {
			eventRet2.ID = info3.ID
			eventRet2.CreatedAt = info3.CreatedAt
			eventRet2.UpdatedAt = info3.UpdatedAt
			eventRet2.CouponIDs = info3.CouponIDs
			eventRet2.CouponIDsStr = info3.CouponIDsStr
			eventRet2.Coins = info3.Coins
			eventRet2.GoodID = info3.GoodID
			eventRet2.AppGoodID = info3.AppGoodID
			assert.Equal(t, info3, &eventRet2)
			handler2.ID = &info3.ID
		}
	}

	h5, err := taskconfig1.NewHandler(
		context.Background(),
		taskconfig1.WithEntID(&taskConfig2.EntID, true),
		taskconfig1.WithAppID(&taskConfig2.AppID, true),
		taskconfig1.WithEventID(&taskConfig2.EventID, true),
		taskconfig1.WithName(&taskConfig2.Name, true),
		taskconfig1.WithTaskDesc(&taskConfig2.TaskDesc, true),
		taskconfig1.WithStepGuide(&taskConfig2.StepGuide, true),
		taskconfig1.WithRecommendMessage(&taskConfig2.RecommendMessage, true),
		taskconfig1.WithIndex(&taskConfig2.Index, true),
		taskconfig1.WithLastTaskID(&taskConfig2.LastTaskID, true),
		taskconfig1.WithMaxRewardCount(&taskConfig2.MaxRewardCount, true),
		taskconfig1.WithCooldownSecond(&taskConfig2.CooldownSecond, true),
		taskconfig1.WithTaskType(&taskConfig2.TaskType, true),
		taskconfig1.WithIntervalReset(&taskConfig2.IntervalReset, false),
	)
	assert.Nil(t, err)

	err = h5.CreateTaskConfig(context.Background())
	if assert.Nil(t, err) {
		info, err := h5.GetTaskConfig(context.Background())
		if assert.Nil(t, err) {
			taskConfig2.ID = info.ID
			taskConfig2.CreatedAt = info.CreatedAt
			taskConfig2.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &taskConfig2, info)
			h5.ID = &info.ID
		}
	}

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
		_ = h2.DeleteTaskConfig(context.Background())
		_ = h3.DeleteCoinConfig(context.Background())
		_ = h4.DeleteCoinConfig(context.Background())
		_ = h5.DeleteTaskConfig(context.Background())
		_ = handler.DeleteEvent(context.Background())
		_ = handler2.DeleteEvent(context.Background())
	}
}

//nolint:dupl
func rewardEvent(t *testing.T) {
	eventType := basetypes.UsedFor_Signup
	consecutive := uint32(1)
	amount := decimal.NewFromInt(10).String()
	handler, err := NewHandler(
		context.Background(),
		WithAppID(&eventRet.AppID, true),
		WithUserID(&userID, true),
		WithEventType(&eventType, true),
		WithConsecutive(&consecutive, true),
		WithAmount(&amount, false),
	)
	assert.Nil(t, err)

	reward, err := handler.CalcluateEventRewards(context.Background())
	if assert.Nil(t, err) {
		fmt.Println("reward: ", reward)
	}
}

//nolint:dupl
func rewardEvent2(t *testing.T) {
	eventType := basetypes.UsedFor_SetWithdrawAddress
	consecutive := uint32(1)
	amount := decimal.NewFromInt(1).String()
	handler, err := NewHandler(
		context.Background(),
		WithAppID(&eventRet2.AppID, true),
		WithUserID(&userID, true),
		WithEventType(&eventType, true),
		WithConsecutive(&consecutive, true),
		WithAmount(&amount, false),
	)
	assert.Nil(t, err)

	reward, err := handler.CalcluateEventRewards(context.Background())
	if assert.Nil(t, err) {
		fmt.Println("reward: ", reward)
	}
}

func TestReward(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := resetup(t)
	defer teardown(t)

	t.Run("rewardEvent", rewardEvent)
	t.Run("rewardEvent2", rewardEvent2)
}
