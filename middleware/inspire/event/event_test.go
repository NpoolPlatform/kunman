package event

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	couponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/event"
	coupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	"github.com/NpoolPlatform/kunman/middleware/inspire/testinit"
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
	coupon = couponmwpb.Coupon{
		EntID:               uuid.NewString(),
		CouponType:          types.CouponType_FixAmount,
		CouponTypeStr:       types.CouponType_FixAmount.String(),
		AppID:               uuid.NewString(),
		Denomination:        decimal.RequireFromString("12.25").String(),
		Circulation:         decimal.RequireFromString("12.25").String(),
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
	ret = npool.Event{
		EntID:          uuid.NewString(),
		AppID:          uuid.NewString(),
		EventType:      basetypes.UsedFor_Signup,
		EventTypeStr:   basetypes.UsedFor_Signup.String(),
		Credits:        decimal.RequireFromString("12.25").String(),
		CreditsPerUSD:  decimal.RequireFromString("12.25").String(),
		MaxConsecutive: 1,
		InviterLayers:  2,
	}
)

//nolint:dupl
func setup(t *testing.T) func(*testing.T) {
	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithEntID(&coupon.EntID, true),
		coupon1.WithAppID(&coupon.AppID, true),
		coupon1.WithName(&coupon.Name, true),
		coupon1.WithMessage(&coupon.Message, true),
		coupon1.WithCouponType(&coupon.CouponType, true),
		coupon1.WithDenomination(&coupon.Denomination, true),
		coupon1.WithCouponScope(&coupon.CouponScope, true),
		coupon1.WithCirculation(&coupon.Circulation, true),
		coupon1.WithDurationDays(&coupon.DurationDays, true),
		coupon1.WithIssuedBy(&coupon.IssuedBy, true),
		coupon1.WithStartAt(&coupon.StartAt, true),
		coupon1.WithEndAt(&coupon.EndAt, true),
		coupon1.WithCashableProbability(&coupon.CashableProbability, true),
	)
	assert.Nil(t, err)

	info, err := h1.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		coupon.ID = info.ID
		coupon.CreatedAt = info.CreatedAt
		coupon.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &coupon, info)
		h1.ID = &info.ID
	}

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
	}
}

func createEvent(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithEventType(&ret.EventType, true),
		WithCredits(&ret.Credits, true),
		WithCreditsPerUSD(&ret.CreditsPerUSD, true),
		WithMaxConsecutive(&ret.MaxConsecutive, true),
		WithInviterLayers(&ret.InviterLayers, true),
	)
	assert.Nil(t, err)

	err = handler.CreateEvent(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetEvent(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.GoodID = info.GoodID
			ret.AppGoodID = info.AppGoodID
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateEvent(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithCredits(&ret.Credits, true),
		WithCreditsPerUSD(&ret.CreditsPerUSD, true),
		WithMaxConsecutive(&ret.MaxConsecutive, true),
		WithInviterLayers(&ret.InviterLayers, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateEvent(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetEvent(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getEvent(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetEvent(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getEvents(t *testing.T) {
	conds := &npool.Conds{
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		EntIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		EventType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.EventType)},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetEvents(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteEvent(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteEvent(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetEvent(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestEvent(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createEvent", createEvent)
	t.Run("updateEvent", updateEvent)
	t.Run("getEvent", getEvent)
	t.Run("getEvents", getEvents)
	t.Run("deleteEvent", deleteEvent)
}
