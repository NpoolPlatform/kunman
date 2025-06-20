package coupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"
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

var ret = npool.Coupon{
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

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithCouponType(&ret.CouponType, true),
		WithAppID(&ret.AppID, true),
		WithDenomination(&ret.Denomination, true),
		WithCirculation(&ret.Circulation, true),
		WithIssuedBy(&ret.IssuedBy, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
		WithDurationDays(&ret.DurationDays, true),
		WithMessage(&ret.Message, true),
		WithName(&ret.Name, true),
		WithCouponScope(&ret.CouponScope, true),
		WithCashableProbability(&ret.CashableProbability, false),
	)
	assert.Nil(t, err)

	info, err := handler.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateCoupon(t *testing.T) {
	ret.Denomination = "10.02"
	ret.Circulation = "200.4"
	ret.CouponScope = types.CouponScope_AllGood
	ret.CouponScopeStr = types.CouponScope_AllGood.String()
	ret.EndAt = uint32(time.Now().Add(24 * 30 * time.Hour).Unix())
	ret.CashableProbability = decimal.RequireFromString("0.00001").String()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithDenomination(&ret.Denomination, true),
		WithCirculation(&ret.Circulation, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
		WithDurationDays(&ret.DurationDays, true),
		WithMessage(&ret.Message, true),
		WithName(&ret.Name, true),
		WithCouponScope(&ret.CouponScope, true),
		WithCashableProbability(&ret.CashableProbability, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateCoupon(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCoupon(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCoupons(t *testing.T) {
	conds := &npool.Conds{
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		EntIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		CouponType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponType)},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetCoupons(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteCoupon(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetCoupon(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestCoupon(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCoupon", createCoupon)
	t.Run("updateCoupon", updateCoupon)
	t.Run("getCoupon", getCoupon)
	t.Run("getCoupons", getCoupons)
	t.Run("deleteCoupon", deleteCoupon)
}
