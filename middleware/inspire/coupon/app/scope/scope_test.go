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

	coupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
	scope1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/scope"
	couponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/app/scope"
	scopemwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/scope"
	"github.com/google/uuid"
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

	scope = scopemwpb.Scope{
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

	ret = npool.Scope{
		EntID:              uuid.NewString(),
		AppID:              coupon.AppID,
		AppGoodID:          uuid.NewString(),
		CouponID:           scope.CouponID,
		CouponName:         coupon.Name,
		CouponType:         scope.CouponType,
		CouponTypeStr:      scope.CouponType.String(),
		CouponScope:        scope.CouponScope,
		CouponScopeStr:     scope.CouponScope.String(),
		CouponDenomination: scope.CouponDenomination,
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithEntID(&coupon.EntID, true),
		coupon1.WithAppID(&coupon.AppID, true),
		coupon1.WithName(&scope.CouponName, true),
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

	coup, err := h1.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		coupon.ID = coup.ID
		coupon.CreatedAt = coup.CreatedAt
		coupon.UpdatedAt = coup.UpdatedAt
		assert.Equal(t, &coupon, coup)
		h1.ID = &coup.ID
	}

	h2, err := scope1.NewHandler(
		context.Background(),
		scope1.WithEntID(&scope.EntID, true),
		scope1.WithGoodID(&scope.GoodID, true),
		scope1.WithCouponID(&coupon.EntID, true),
		scope1.WithCouponScope(&coupon.CouponScope, true),
	)
	assert.Nil(t, err)

	info, err := h2.CreateScope(context.Background())
	if assert.Nil(t, err) {
		scope.ID = info.ID
		scope.CreatedAt = info.CreatedAt
		scope.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &scope, info)
		h2.ID = &info.ID
	}

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
		_, _ = h2.DeleteScope(context.Background())
	}
}

func createAppGoodScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithCouponID(&ret.CouponID, true),
		WithCouponScope(&ret.CouponScope, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateAppGoodScope(context.Background())
	if assert.Nil(t, err) {
		ret.ID = info.ID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getAppGoodScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetAppGoodScope(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAppGoodScopes(t *testing.T) {
	conds := &npool.Conds{
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		AppGoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
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

	infos, _, err := handler.GetAppGoodScopes(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, &ret, infos[0])
	}
}

func existAppGoodScopeConds(t *testing.T) {
	conds := &npool.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		CouponID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		AppGoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		CouponScope: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponScope)},
	}
	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistAppGoodScopeConds(context.Background())
	assert.Nil(t, err)
	assert.True(t, exist)
}

func verifyCouponScope(t *testing.T) {
	reqs := []*npool.ScopeReq{{
		AppID:       &ret.AppID,
		GoodID:      &scope.GoodID,
		AppGoodID:   &ret.AppGoodID,
		CouponID:    &ret.CouponID,
		CouponScope: &ret.CouponScope,
	}}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(reqs, true),
	)
	assert.Nil(t, err)

	err = handler.VerifyCouponScopes(context.Background())
	assert.Nil(t, err)
}

func deleteAppGoodScope(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteAppGoodScope(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetAppGoodScope(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestScope(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createAppGoodScope", createAppGoodScope)
	t.Run("getAppGoodScope", getAppGoodScope)
	t.Run("getAppGoodScopes", getAppGoodScopes)
	t.Run("existAppGoodScope", existAppGoodScopeConds)
	t.Run("verifyCouponScope", verifyCouponScope)
	t.Run("deleteAppGoodScope", deleteAppGoodScope)
}
