package subscription

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	paymentgwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/payment"
	ordercouponmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/coupon"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	appsubscription1 "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"
	subscription1 "github.com/NpoolPlatform/kunman/middleware/good/subscription"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/gateway/order/testinit"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
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
	DurationSeconds:     100000,
	LedgerLockID:        uuid.NewString(),
	PaymentID:           uuid.NewString(),
	Coupons: []*ordercouponmwpb.OrderCouponInfo{
		{
			CouponID: uuid.NewString(),
		},
	},
	PaymentBalances: []*paymentmwpb.PaymentBalanceInfo{
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
	ret.GoodTypeStr = ret.GoodType.String()

	goodName := uuid.NewString()
	appGoodName := uuid.NewString()
	usdPrice := "12.99"
	durationUnits := uint32(1)
	durationQuota := uint32(2000)

	subscriptionEntID := uuid.NewString()
	h1, err := subscription1.NewHandler(
		context.Background(),
		subscription1.WithEntID(&subscriptionEntID, true),
		subscription1.WithGoodID(&ret.GoodID, true),
		subscription1.WithGoodType(&ret.GoodType, true),
		subscription1.WithName(&goodName, true),
		subscription1.WithUSDPrice(&usdPrice, true),
		subscription1.WithDurationDisplayType(goodtypes.GoodDurationType_GoodDurationByWeek.Enum(), true),
		subscription1.WithDurationUnits(&durationUnits, true),
		subscription1.WithDurationQuota(&durationQuota, true),
	)
	assert.Nil(t, err)

	err = h1.CreateSubscription(context.Background())
	assert.Nil(t, err)

	appSubscriptionEntID := uuid.NewString()
	h2, err := appsubscription1.NewHandler(
		context.Background(),
		appsubscription1.WithEntID(&appSubscriptionEntID, true),
		appsubscription1.WithAppID(&ret.AppID, true),
		appsubscription1.WithGoodID(&ret.GoodID, true),
		appsubscription1.WithAppGoodID(&ret.AppGoodID, true),
		appsubscription1.WithName(&appGoodName, true),
		appsubscription1.WithUSDPrice(&usdPrice, true),
	)
	assert.Nil(t, err)

	err = h2.CreateSubscription(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h2.DeleteSubscription(context.Background())
		_ = h1.DeleteSubscription(context.Background())
	}
}

func createSubscription(t *testing.T) {
	paymentTransferCoinTypeID := uuid.NewString()
	paymentFiatID := uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithOrderID(&ret.OrderID, true),
		WithPaymentBalances(func() (balances []*paymentgwpb.PaymentBalance) {
			for _, balance := range ret.PaymentBalances {
				balances = append(balances, &paymentgwpb.PaymentBalance{
					CoinTypeID: balance.CoinTypeID,
					Amount:     balance.Amount,
				})
			}
			return
		}(), true),
		WithPaymentTransferCoinTypeID(&paymentTransferCoinTypeID, true),
		WithPaymentFiatID(&paymentFiatID, true),
		WithCouponIDs(func() (couponIDs []string) {
			for _, coupon := range ret.Coupons {
				couponIDs = append(couponIDs, coupon.CouponID)
			}
			return
		}(), true),
		WithUserSetCanceled(&ret.UserSetCanceled, true),
		WithAdminSetCanceled(&ret.AdminSetCanceled, true),
		WithCreateMethod(&ret.CreateMethod, true),
		WithOrderType(&ret.OrderType, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateSubscriptionOrder(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateSubscription(t *testing.T) {
}

func getSubscription(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetSubscriptionOrder(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func existSubscription(t *testing.T) {
}

func getSubscriptions(t *testing.T) {
}

func deleteSubscription(t *testing.T) {
}

func TestSubscription(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createSubscription", createSubscription)
	t.Run("updateSubscription", updateSubscription)
	t.Run("getSubscription", getSubscription)
	t.Run("existSubscription", existSubscription)
	t.Run("getSubscriptions", getSubscriptions)
	t.Run("deleteSubscription", deleteSubscription)
}
