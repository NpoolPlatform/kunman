package allocated

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	coupon1 "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
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
	ret = npool.Coupon{
		EntID:               uuid.NewString(),
		CouponType:          types.CouponType_FixAmount,
		CouponTypeStr:       types.CouponType_FixAmount.String(),
		AppID:               uuid.NewString(),
		UserID:              uuid.NewString(),
		Denomination:        decimal.RequireFromString("10").String(),
		Circulation:         decimal.RequireFromString("100").String(),
		Allocated:           decimal.RequireFromString("10").String(),
		StartAt:             uint32(time.Now().Unix()),
		EndAt:               uint32(time.Now().Add(24 * time.Hour).Unix()),
		Threshold:           decimal.RequireFromString("0").String(),
		UsedByOrderID:       uuid.Nil.String(),
		DurationDays:        27,
		CouponID:            uuid.NewString(),
		CouponName:          uuid.NewString(),
		Message:             uuid.NewString(),
		CouponConstraint:    types.CouponConstraint_Normal,
		CouponConstraintStr: types.CouponConstraint_Normal.String(),
		CouponScope:         types.CouponScope_Whitelist,
		CouponScopeStr:      types.CouponScope_Whitelist.String(),
		Valid:               true,
		Used:                false,
	}

	ret1 = npool.Coupon{
		EntID:               uuid.NewString(),
		AppID:               ret.AppID,
		UserID:              uuid.NewString(),
		CouponID:            ret.CouponID,
		CouponName:          ret.CouponName,
		Message:             ret.Message,
		CouponType:          ret.CouponType,
		CouponTypeStr:       ret.CouponType.String(),
		Denomination:        ret.Denomination,
		Circulation:         ret.Circulation,
		Allocated:           decimal.RequireFromString("20").String(),
		StartAt:             uint32(time.Now().Unix()),
		EndAt:               ret.EndAt,
		DurationDays:        27,
		Threshold:           decimal.RequireFromString("0").String(),
		UsedByOrderID:       uuid.Nil.String(),
		CouponConstraint:    ret.CouponConstraint,
		CouponConstraintStr: ret.CouponConstraint.String(),
		CouponScope:         ret.CouponScope,
		CouponScopeStr:      ret.CouponScope.String(),
		Valid:               true,
		Used:                false,
	}
)

func setup(t *testing.T) func(*testing.T) {
	ret.CouponTypeStr = ret.CouponType.String()
	ret.CouponConstraintStr = ret.CouponConstraint.String()

	h1, err := coupon1.NewHandler(
		context.Background(),
		coupon1.WithEntID(&ret.CouponID, true),
		coupon1.WithCouponType(&ret.CouponType, true),
		coupon1.WithAppID(&ret.AppID, true),
		coupon1.WithDenomination(&ret.Denomination, true),
		coupon1.WithCirculation(&ret.Circulation, true),
		coupon1.WithIssuedBy(&ret.UserID, true),
		coupon1.WithStartAt(&ret.StartAt, true),
		coupon1.WithEndAt(&ret.EndAt, true),
		coupon1.WithDurationDays(&ret.DurationDays, true),
		coupon1.WithMessage(&ret.Message, true),
		coupon1.WithName(&ret.CouponName, true),
	)
	assert.Nil(t, err)

	coup, err := h1.CreateCoupon(context.Background())
	assert.Nil(t, err)
	h1.ID = &coup.ID

	return func(*testing.T) {
		_, _ = h1.DeleteCoupon(context.Background())
	}
}

func createCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithCouponID(&ret.CouponID, true),
		WithUserID(&ret.UserID, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCoupon(context.Background())
		if assert.Nil(t, err) {
			ret.ID = info.ID
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.StartAt = info.StartAt
			ret.EndAt = info.EndAt
			assert.Equal(t, &ret, info)
		}
	}

	handler, err = NewHandler(
		context.Background(),
		WithEntID(&ret1.EntID, true),
		WithAppID(&ret1.AppID, true),
		WithCouponID(&ret1.CouponID, true),
		WithUserID(&ret1.UserID, true),
	)
	assert.Nil(t, err)

	err = handler.CreateCoupon(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCoupon(context.Background())
		if assert.Nil(t, err) {
			ret1.ID = info.ID
			ret1.CreatedAt = info.CreatedAt
			ret1.UpdatedAt = info.UpdatedAt
			ret1.StartAt = info.StartAt
			ret1.EndAt = info.EndAt
			assert.Equal(t, &ret1, info)
			ret.Allocated = info.Allocated
		}
	}
}

func updateCoupon(t *testing.T) {
	orderID := uuid.NewString()
	ret.UsedByOrderID = orderID
	ret.Used = true

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithUsed(&ret.Used, true),
		WithUsedByOrderID(&ret.UsedByOrderID, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateCoupon(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetCoupon(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			ret.UsedAt = info.UsedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func updateCoupons(t *testing.T) {
	orderID := uuid.NewString()
	ret1.UsedByOrderID = orderID
	ret1.Used = true

	reqs := []*npool.CouponReq{{
		ID:            &ret1.ID,
		Used:          &ret1.Used,
		UsedByOrderID: &ret1.UsedByOrderID,
	}}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(reqs, true),
	)
	assert.Nil(t, err)

	infos, err := handler.UpdateCoupons(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, int(1), len(infos))
		ret1.UsedAt = infos[0].UsedAt
		ret1.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, &ret1, infos[0])
	}
}

func getCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetCoupon(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getCoupons(t *testing.T) {
	conds := &npool.Conds{
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		EntIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.EntID}},
		CouponType:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.CouponType)},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		CouponID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.CouponID},
		UsedByOrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.UsedByOrderID},
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
		assert.Equal(t, &ret, infos[0])
	}
}

func deleteCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteCoupon(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetCoupon(context.Background())
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
	t.Run("updateCoupons", updateCoupons)
	t.Run("getCoupon", getCoupon)
	t.Run("getCoupons", getCoupons)
	t.Run("deleteCoupon", deleteCoupon)
}
