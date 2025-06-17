package subscription

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	ledgerstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	ordercoupongwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order/coupon"
	paymentgwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/payment"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"
	appmw "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	fiatmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
	appsubscription1 "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"
	subscription1 "github.com/NpoolPlatform/kunman/middleware/good/subscription"
	couponmw "github.com/NpoolPlatform/kunman/middleware/inspire/coupon"
	allocatedcouponmw "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/allocated"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/gateway/order/testinit"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	inspiretypes "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
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
	AppName:             uuid.NewString(),
	UserID:              uuid.NewString(),
	PhoneNO:             fmt.Sprintf("+86%v", rand.Intn(100000000)+rand.Intn(1000000)),
	EmailAddress:        fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+rand.Intn(4000000)),
	GoodID:              uuid.NewString(),
	GoodName:            uuid.NewString(),
	GoodType:            goodtypes.GoodType_Subscription,
	AppGoodID:           uuid.NewString(),
	AppGoodName:         uuid.NewString(),
	OrderID:             uuid.NewString(),
	OrderType:           types.OrderType_Normal,
	CreateMethod:        types.OrderCreateMethod_OrderCreatedByPurchase,
	GoodValueUSD:        decimal.NewFromInt(120).String(),
	PaymentGoodValueUSD: decimal.NewFromInt(120).String(),
	PaymentAmountUSD:    decimal.RequireFromString("107.8").String(),
	DiscountAmountUSD:   decimal.RequireFromString("12.2").String(),
	PromotionID:         "",
	LedgerLockID:        uuid.NewString(),
	PaymentID:           uuid.NewString(),
	Coupons: []*ordercoupongwpb.OrderCouponInfo{
		{
			AllocatedCouponID: uuid.NewString(),
			CouponType:        inspiretypes.CouponType_FixAmount,
			Denomination:      "12.2",
		},
	},
	PaymentBalances: []*paymentgwpb.PaymentBalanceInfo{
		{
			CoinTypeID:      uuid.NewString(),
			Amount:          decimal.RequireFromString("103.8").String(),
			CoinUSDCurrency: decimal.NewFromInt(1).String(),
		},
	},
	PaymentTransfers: []*paymentgwpb.PaymentTransferInfo{},
	PaymentFiats: []*paymentgwpb.PaymentFiatInfo{
		{
			FiatID:      uuid.NewString(),
			Amount:      decimal.NewFromInt(4).String(),
			USDCurrency: decimal.NewFromInt(1).String(),
		},
	},
	OrderState:          types.OrderState_OrderStateCreated,
	PaymentState:        types.PaymentState_PaymentStateWait,
	DurationDisplayType: goodtypes.GoodDurationType_GoodDurationByWeek,
	DurationUnit:        "MSG_WEEK",
	Durations:           1,
	DurationSeconds:     timedef.SecondsPerWeek,
	PaymentType:         types.PaymentType_PayWithFiatAndBalance,
}

func setup(t *testing.T) func(*testing.T) {
	durationUnits := uint32(1)
	durationQuota := uint32(2000)

	subscriptionEntID := uuid.NewString()
	h1, err := subscription1.NewHandler(
		context.Background(),
		subscription1.WithEntID(&subscriptionEntID, true),
		subscription1.WithGoodID(&ret.GoodID, true),
		subscription1.WithGoodType(&ret.GoodType, true),
		subscription1.WithName(&ret.GoodName, true),
		subscription1.WithUSDPrice(&ret.GoodValueUSD, true),
		subscription1.WithDurationDisplayType(ret.DurationDisplayType.Enum(), true),
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
		appsubscription1.WithName(&ret.AppGoodName, true),
		appsubscription1.WithUSDPrice(&ret.GoodValueUSD, true),
	)
	assert.Nil(t, err)

	err = h2.CreateSubscription(context.Background())
	assert.Nil(t, err)

	h3, err := appmw.NewHandler(
		context.Background(),
		appmw.WithEntID(&ret.AppID, true),
		appmw.WithName(&ret.AppName, true),
	)
	assert.Nil(t, err)

	_, err = h3.CreateApp(context.Background())
	assert.Nil(t, err)

	h4, err := usermw.NewHandler(
		context.Background(),
		usermw.WithEntID(&ret.UserID, true),
		usermw.WithAppID(&ret.AppID, true),
		usermw.WithPhoneNO(&ret.PhoneNO, true),
		usermw.WithEmailAddress(&ret.EmailAddress, true),
		usermw.WithPasswordHash(&ret.AppID, true),
	)
	assert.Nil(t, err)

	_, err = h4.CreateUser(context.Background())
	assert.Nil(t, err)

	h5s := []*coinmw.Handler{}
	h51s := []*appcoinmw.Handler{}

	for _, balance := range ret.PaymentBalances {
		h5, err := coinmw.NewHandler(
			context.Background(),
			coinmw.WithEntID(&balance.CoinTypeID, true),
			coinmw.WithName(&balance.CoinTypeID, true),
			coinmw.WithUnit(&balance.CoinTypeID, true),
			coinmw.WithENV(func() *string { s := "test"; return &s }(), true),
		)
		assert.Nil(t, err)

		_, err = h5.CreateCoin(context.Background())
		assert.Nil(t, err)

		h5s = append(h5s, h5)

		h51, err := appcoinmw.NewHandler(
			context.Background(),
			appcoinmw.WithAppID(&ret.AppID, true),
			appcoinmw.WithCoinTypeID(&balance.CoinTypeID, true),
		)
		assert.Nil(t, err)

		_, err = h51.CreateCoin(context.Background())
		assert.Nil(t, err)

		h51s = append(h51s, h51)
	}

	for _, balance := range ret.PaymentTransfers {
		h5, err := coinmw.NewHandler(
			context.Background(),
			coinmw.WithEntID(&balance.CoinTypeID, true),
			coinmw.WithName(&balance.CoinTypeID, true),
			coinmw.WithUnit(&balance.CoinTypeID, true),
			coinmw.WithENV(func() *string { s := "test"; return &s }(), true),
		)
		assert.Nil(t, err)

		_, err = h5.CreateCoin(context.Background())
		assert.Nil(t, err)

		h5s = append(h5s, h5)

		h51, err := appcoinmw.NewHandler(
			context.Background(),
			appcoinmw.WithAppID(&ret.AppID, true),
			appcoinmw.WithCoinTypeID(&balance.CoinTypeID, true),
		)
		assert.Nil(t, err)

		_, err = h51.CreateCoin(context.Background())
		assert.Nil(t, err)

		h51s = append(h51s, h51)
	}

	h6s := []*fiatmw.Handler{}

	for _, balance := range ret.PaymentFiats {
		h6, err := fiatmw.NewHandler(
			context.Background(),
			fiatmw.WithEntID(&balance.FiatID, true),
			fiatmw.WithName(&balance.FiatID, true),
			fiatmw.WithUnit(&balance.FiatID, true),
		)
		assert.Nil(t, err)

		_, err = h6.CreateFiat(context.Background())
		assert.Nil(t, err)

		h6s = append(h6s, h6)
	}

	couponID := uuid.NewString()
	startAt := uint32(time.Now().Unix())
	endAt := uint32(time.Now().Unix()) + 1000
	circulation := "9999"

	h7, err := couponmw.NewHandler(
		context.Background(),
		couponmw.WithEntID(&couponID, true),
		couponmw.WithAppID(&ret.AppID, true),
		couponmw.WithCouponType(&ret.Coupons[0].CouponType, true),
		couponmw.WithStartAt(&startAt, true),
		couponmw.WithEndAt(&endAt, true),
		couponmw.WithDenomination(&ret.Coupons[0].Denomination, true),
		couponmw.WithCirculation(&circulation, true),
		couponmw.WithCouponScope(inspiretypes.CouponScope_AllGood.Enum(), true),
	)
	assert.Nil(t, err)

	_, err = h7.CreateCoupon(context.Background())
	assert.Nil(t, err)

	h8s := []*allocatedcouponmw.Handler{}

	for _, coupon := range ret.Coupons {
		h8, err := allocatedcouponmw.NewHandler(
			context.Background(),
			allocatedcouponmw.WithEntID(&coupon.AllocatedCouponID, true),
			allocatedcouponmw.WithAppID(&ret.AppID, true),
			allocatedcouponmw.WithUserID(&ret.UserID, true),
			allocatedcouponmw.WithCouponID(&couponID, true),
		)
		assert.Nil(t, err)

		err = h8.CreateCoupon(context.Background())
		assert.Nil(t, err)
	}

	ledgerStatementReqs := []*ledgerstatementmwpb.StatementReq{}
	userExtra := fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString())

	for _, balance := range ret.PaymentBalances {
		ledgerStatementReqs = append(ledgerStatementReqs, &ledgerstatementmwpb.StatementReq{
			AppID:      &ret.AppID,
			UserID:     &ret.UserID,
			CoinTypeID: &balance.CoinTypeID,
			Amount:     &circulation,
			IOType:     ledgertypes.IOType_Incoming.Enum(),
			IOSubType:  ledgertypes.IOSubType_Deposit.Enum(),
			IOExtra:    &userExtra,
		})
	}

	h9, err := ledgerstatementmw.NewHandler(
		context.Background(),
		ledgerstatementmw.WithReqs(ledgerStatementReqs, true),
	)
	assert.Nil(t, err)

	_, err = h9.CreateStatements(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		for _, h8 := range h8s {
			_ = h8.DeleteCoupon(context.Background())
		}
		_, _ = h7.DeleteCoupon(context.Background())

		for _, h51 := range h51s {
			_, _ = h51.DeleteCoin(context.Background())
		}
		for _, h5 := range h5s {
			_, _ = h5.DeleteCoin(context.Background())
		}
		for _, h6 := range h6s {
			_, _ = h6.DeleteFiat(context.Background())
		}

		_, _ = h4.DeleteUser(context.Background())
		_, _ = h3.DeleteApp(context.Background())
		_ = h2.DeleteSubscription(context.Background())
		_ = h1.DeleteSubscription(context.Background())
	}
}

func createSubscription(t *testing.T) {
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
		// WithPaymentTransferCoinTypeID(&ret.PaymentTransfers[0].CoinTypeID, true),
		WithPaymentFiatID(&ret.PaymentFiats[0].FiatID, true),
		WithCouponIDs(func() (couponIDs []string) {
			for _, coupon := range ret.Coupons {
				couponIDs = append(couponIDs, coupon.AllocatedCouponID)
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
		ret.LedgerLockID = info.LedgerLockID
		ret.PaymentID = info.PaymentID
		ret.ID = info.ID

		for i, coupon := range info.Coupons {
			ret.Coupons[i].CreatedAt = coupon.CreatedAt
		}

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
