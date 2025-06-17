package subscriptionorder

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	ordercouponmiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/coupon"
	paymentmiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

var ret = npool.SubscriptionOrder{
	EntID:               uuid.NewString(),
	AppID:               uuid.NewString(),
	UserID:              uuid.NewString(),
	GoodID:              uuid.NewString(),
	GoodType:            goodtypes.GoodType_Subscription,
	AppGoodID:           uuid.NewString(),
	OrderID:             uuid.NewString(),
	OrderType:           types.OrderType_Normal,
	PaymentType:         types.PaymentType_PayWithBalanceOnly,
	CreateMethod:        types.OrderCreateMethod_OrderCreatedByPurchase,
	GoodValueUSD:        decimal.NewFromInt(120).String(),
	PaymentGoodValueUSD: decimal.NewFromInt(120).String(),
	PaymentAmountUSD:    decimal.NewFromInt(110).String(),
	DiscountAmountUSD:   decimal.NewFromInt(10).String(),
	PromotionID:         uuid.NewString(),
	LifeSeconds:         100000,
	LedgerLockID:        uuid.NewString(),
	PaymentID:           uuid.NewString(),
	Coupons: []*ordercouponmiddlewarepb.OrderCouponInfo{
		{
			CouponID: uuid.NewString(),
		},
	},
	PaymentBalances: []*paymentmiddlewarepb.PaymentBalanceInfo{
		{
			CoinTypeID:           uuid.NewString(),
			Amount:               decimal.NewFromInt(110).String(),
			CoinUSDCurrency:      decimal.NewFromInt(1).String(),
			LocalCoinUSDCurrency: decimal.NewFromInt(1).String(),
			LiveCoinUSDCurrency:  decimal.NewFromInt(1).String(),
		},
	},
	OrderState:   types.OrderState_OrderStateCreated,
	PaymentState: types.PaymentState_PaymentStateWait,
}

func setup(t *testing.T) func(*testing.T) {
	for _, paymentBalance := range ret.PaymentBalances {
		paymentBalance.PaymentID = ret.PaymentID
	}
	for _, orderCoupon := range ret.Coupons {
		orderCoupon.OrderID = ret.OrderID
	}

	ret.GoodTypeStr = ret.GoodType.String()
	ret.OrderTypeStr = ret.OrderType.String()
	ret.PaymentTypeStr = ret.PaymentType.String()
	ret.CreateMethodStr = ret.CreateMethod.String()
	ret.OrderStateStr = ret.OrderState.String()
	ret.PaymentStateStr = ret.PaymentState.String()
	ret.CancelStateStr = ret.CancelState.String()

	return func(*testing.T) {}
}

func createSubscriptionOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithOrderID(&ret.OrderID, true),
		WithOrderType(&ret.OrderType, true),
		WithPaymentType(&ret.PaymentType, true),
		WithCreateMethod(&ret.CreateMethod, true),
		WithGoodValueUSD(&ret.GoodValueUSD, true),
		WithPaymentAmountUSD(&ret.PaymentAmountUSD, true),
		WithDiscountAmountUSD(&ret.DiscountAmountUSD, true),
		WithPromotionID(&ret.PromotionID, true),
		WithLifeSeconds(&ret.LifeSeconds, true),
		WithLedgerLockID(&ret.LedgerLockID, true),
		WithPaymentID(&ret.PaymentID, true),
		WithCouponIDs(func() (_couponIDs []string) {
			for _, coupon := range ret.Coupons {
				_couponIDs = append(_couponIDs, coupon.CouponID)
			}
			return
		}(), true),
		WithPaymentBalances(func() (_reqs []*paymentmiddlewarepb.PaymentBalanceReq) {
			for _, req := range ret.PaymentBalances {
				_reqs = append(_reqs, &paymentmiddlewarepb.PaymentBalanceReq{
					CoinTypeID:           &req.CoinTypeID,
					Amount:               &req.Amount,
					LocalCoinUSDCurrency: &req.LocalCoinUSDCurrency,
					LiveCoinUSDCurrency:  &req.LiveCoinUSDCurrency,
				})
			}
			return
		}(), true),
		WithPaymentTransfers([]*paymentmiddlewarepb.PaymentTransferReq{}, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateSubscriptionOrder(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetSubscriptionOrder(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				ret.PaymentBalances = info.PaymentBalances
				for _, orderCoupon := range ret.Coupons {
					orderCoupon.CreatedAt = ret.CreatedAt
				}
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func updateSubscriptionOrder(t *testing.T) {
	ret.PaymentID = uuid.NewString()
	ret.LedgerLockID = uuid.NewString()
	for _, paymentBalance := range ret.PaymentBalances {
		paymentBalance.PaymentID = ret.PaymentID
	}

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, false),
		WithOrderID(&ret.OrderID, true),
		WithGoodValueUSD(&ret.GoodValueUSD, true),
		WithPaymentAmountUSD(&ret.PaymentAmountUSD, true),
		WithDiscountAmountUSD(&ret.DiscountAmountUSD, true),
		WithPromotionID(&ret.PromotionID, true),
		WithLifeSeconds(&ret.LifeSeconds, true),
		WithLedgerLockID(&ret.LedgerLockID, true),
		WithPaymentID(&ret.PaymentID, true),
		WithCouponIDs(func() (_couponIDs []string) {
			for _, coupon := range ret.Coupons {
				_couponIDs = append(_couponIDs, coupon.CouponID)
			}
			return
		}(), true),
		WithPaymentBalances(func() (_reqs []*paymentmiddlewarepb.PaymentBalanceReq) {
			for _, req := range ret.PaymentBalances {
				_reqs = append(_reqs, &paymentmiddlewarepb.PaymentBalanceReq{
					CoinTypeID:           &req.CoinTypeID,
					Amount:               &req.Amount,
					LocalCoinUSDCurrency: &req.LocalCoinUSDCurrency,
					LiveCoinUSDCurrency:  &req.LiveCoinUSDCurrency,
				})
			}
			return
		}(), true),
		WithPaymentTransfers([]*paymentmiddlewarepb.PaymentTransferReq{}, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateSubscriptionOrder(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetSubscriptionOrder(context.Background())
			if assert.Nil(t, err) {
				ret.PaymentBalances = info.PaymentBalances
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getSubscriptionOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithOrderID(&ret.OrderID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetSubscriptionOrder(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getSubscriptionOrders(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			OrderID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
			OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetSubscriptionOrders(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteSubscriptionOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithOrderID(&ret.OrderID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteSubscriptionOrder(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetSubscriptionOrder(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestSubscriptionOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createSubscriptionOrder", createSubscriptionOrder)
	t.Run("updateSubscriptionOrder", updateSubscriptionOrder)
	t.Run("getSubscriptionOrder", getSubscriptionOrder)
	t.Run("getSubscriptionOrders", getSubscriptionOrders)
	t.Run("deleteSubscriptionOrder", deleteSubscriptionOrder)
}
