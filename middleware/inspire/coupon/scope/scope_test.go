package scope

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/shopspring/decimal"

	couponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/scope"
	coupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
	"github.com/google/uuid"
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
		AppID:               uuid.NewString(),
		Name:                uuid.NewString(),
		Message:             uuid.NewString(),
		CouponType:          types.CouponType_FixAmount,
		CouponTypeStr:       types.CouponType_FixAmount.String(),
		Denomination:        decimal.RequireFromString("100").String(),
		Circulation:         decimal.RequireFromString("100").String(),
		DurationDays:        365,
		IssuedBy:            uuid.NewString(),
		StartAt:             uint32(time.Now().Unix()),
		EndAt:               uint32(time.Now().Add(24 * time.Hour).Unix()),
		CouponConstraint:    types.CouponConstraint_Normal,
		CouponConstraintStr: types.CouponConstraint_Normal.String(),
		CouponScope:         types.CouponScope_Whitelist,
		CouponScopeStr:      types.CouponScope_Whitelist.String(),
		Allocated:           decimal.NewFromInt(0).String(),
		Threshold:           decimal.NewFromInt(0).String(),
		CashableProbability: decimal.RequireFromString("0.0001").String(),
	}

	ret = npool.Scope{
		EntID:              uuid.NewString(),
		GoodID:             uuid.NewString(),
		CouponID:           coupon.EntID,
		CouponType:         coupon.CouponType,
		CouponTypeStr:      coupon.CouponTypeStr,
		CouponScope:        coupon.CouponScope,
		CouponScopeStr:     coupon.CouponScopeStr,
		CouponName:         coupon.Name,
		CouponDenomination: coupon.Denomination,
		CouponCirculation:  coupon.Circulation,
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithEntID(&coupon.EntID, true),
		coupon1.WithAppID(&coupon.AppID, true),
		coupon1.WithName(&ret.CouponName, true),
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

func createScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithCouponID(&ret.CouponID, true),
		WithCouponScope(&ret.CouponScope, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateScope(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetScope(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getScopes(t *testing.T) {
	conds := &npool.Conds{
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		CouponID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		CouponScope: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponScope)},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(1),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetScopes(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, &ret, infos[0])
	}
}

func existScopeConds(t *testing.T) {
	conds := &npool.Conds{
		CouponID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		GoodID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		CouponScope: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponScope)},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistScopeConds(context.Background())
	assert.Nil(t, err)
	assert.True(t, exist)
}

func deleteScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteScope(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetScope(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestScope(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createScope", createScope)
	t.Run("getScope", getScope)
	t.Run("getScopes", getScopes)
	t.Run("existScope", existScopeConds)
	t.Run("deleteScope", deleteScope)
}
