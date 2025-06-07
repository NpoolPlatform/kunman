package ordercoupon

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/coupon"
	orderbase1 "github.com/NpoolPlatform/kunman/middleware/order/order/orderbase"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/middleware/order/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.OrderCoupon{
	EntID:     uuid.NewString(),
	AppID:     uuid.NewString(),
	UserID:    uuid.NewString(),
	GoodID:    uuid.NewString(),
	GoodType:  goodtypes.GoodType_PowerRental,
	AppGoodID: uuid.NewString(),
	OrderID:   uuid.NewString(),
	CouponID:  uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()

	h1, err := orderbase1.NewHandler(
		context.Background(),
		orderbase1.WithEntID(&ret.OrderID, false),
		orderbase1.WithAppID(&ret.AppID, true),
		orderbase1.WithUserID(&ret.UserID, true),
		orderbase1.WithGoodID(&ret.GoodID, true),
		orderbase1.WithGoodType(&ret.GoodType, true),
		orderbase1.WithAppGoodID(&ret.AppGoodID, true),
		orderbase1.WithOrderType(func() *types.OrderType { e := types.OrderType_Offline; return &e }(), true),
		orderbase1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByAdmin; return &e }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateOrderBase(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteOrderBase(context.Background())
	}
}

func createOrderCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithOrderID(&ret.OrderID, true),
		WithCouponID(&ret.CouponID, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateOrderCoupon(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetOrderCoupon(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getOrderCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetOrderCoupon(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getOrderCoupons(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetOrderCoupons(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteOrderCoupon(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteOrderCoupon(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetOrderCoupon(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestOrderCoupon(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createOrderCoupon", createOrderCoupon)
	t.Run("getOrderCoupon", getOrderCoupon)
	t.Run("getOrderCoupons", getOrderCoupons)
	t.Run("deleteOrderCoupon", deleteOrderCoupon)
}
