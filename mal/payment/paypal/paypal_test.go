package paypal

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
	ordercouponmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/coupon"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
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
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/kunman/mal/payment/testinit"
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

var subscriptionOrder = npool.SubscriptionOrder{
	EntID:               uuid.NewString(),
	AppID:               uuid.NewString(),
	UserID:              uuid.NewString(),
	GoodID:              uuid.NewString(),
	GoodType:            goodtypes.GoodType_Subscription,
	AppGoodID:           uuid.NewString(),
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
	Coupons: []*ordercouponmwpb.OrderCouponInfo{
		{
			CouponID: uuid.NewString(),
		},
	},
	PaymentBalances: []*paymentmwpb.PaymentBalanceInfo{
		{
			CoinTypeID:           uuid.NewString(),
			Amount:               decimal.RequireFromString("103.8").String(),
			CoinUSDCurrency:      decimal.NewFromInt(1).String(),
			LiveCoinUSDCurrency:  decimal.NewFromInt(1).String(),
			LocalCoinUSDCurrency: decimal.NewFromInt(1).String(),
		},
	},
	PaymentTransfers: nil,
	PaymentFiats: []*paymentmwpb.PaymentFiatInfo{
		{
			FiatID:           uuid.NewString(),
			PaymentChannel:   types.FiatPaymentChannel_PaymentChannelPaypal,
			Amount:           decimal.NewFromInt(4).String(),
			USDCurrency:      decimal.NewFromInt(1).String(),
			ChannelPaymentID: uuid.NewString(),
		},
	},
	OrderState:   types.OrderState_OrderStateCreated,
	PaymentState: types.PaymentState_PaymentStateWait,
	LifeSeconds:  timedef.SecondsPerWeek,
	PaymentType:  types.PaymentType_PayWithFiatAndBalance,
}

var paypalPlanID = ""
var paypalSubscriptionID = ""

func setup(t *testing.T) func(*testing.T) {
	durationUnits := uint32(1)
	durationQuota := uint32(2000)
	goodName := uuid.NewString()

	subscriptionEntID := uuid.NewString()
	h1, err := subscription1.NewHandler(
		context.Background(),
		subscription1.WithEntID(&subscriptionEntID, true),
		subscription1.WithGoodID(&subscriptionOrder.GoodID, true),
		subscription1.WithGoodType(&subscriptionOrder.GoodType, true),
		subscription1.WithName(&goodName, true),
		subscription1.WithUSDPrice(&subscriptionOrder.GoodValueUSD, true),
		subscription1.WithDurationDisplayType(goodtypes.GoodDurationType_GoodDurationByWeek.Enum(), true),
		subscription1.WithDurationUnits(&durationUnits, true),
		subscription1.WithDurationQuota(&durationQuota, true),
	)
	assert.Nil(t, err)

	err = h1.CreateSubscription(context.Background())
	assert.Nil(t, err)

	appSubscriptionEntID := uuid.NewString()
	productID := "PROD-8VL330703E147442N"

	h2, err := appsubscription1.NewHandler(
		context.Background(),
		appsubscription1.WithEntID(&appSubscriptionEntID, true),
		appsubscription1.WithAppID(&subscriptionOrder.AppID, true),
		appsubscription1.WithGoodID(&subscriptionOrder.GoodID, true),
		appsubscription1.WithAppGoodID(&subscriptionOrder.AppGoodID, true),
		appsubscription1.WithName(&goodName, true),
		appsubscription1.WithUSDPrice(&subscriptionOrder.GoodValueUSD, true),
		appsubscription1.WithProductID(&productID, true),
	)
	assert.Nil(t, err)

	err = h2.CreateSubscription(context.Background())
	assert.Nil(t, err)

	h3, err := appmw.NewHandler(
		context.Background(),
		appmw.WithEntID(&subscriptionOrder.AppID, true),
		appmw.WithName(&subscriptionOrder.AppID, true),
	)
	assert.Nil(t, err)

	_, err = h3.CreateApp(context.Background())
	assert.Nil(t, err)

	countryCode := "+86"
	phoneNO := fmt.Sprintf("+86%v", 13900000000+rand.Intn(1000000))
	emailAddress := fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+rand.Intn(4000000))

	h4, err := usermw.NewHandler(
		context.Background(),
		usermw.WithEntID(&subscriptionOrder.UserID, true),
		usermw.WithAppID(&subscriptionOrder.AppID, true),
		usermw.WithCountryCode(&countryCode, true),
		usermw.WithPhoneNO(&phoneNO, true),
		usermw.WithEmailAddress(&emailAddress, true),
		usermw.WithPasswordHash(&subscriptionOrder.AppID, true),
	)
	assert.Nil(t, err)

	_, err = h4.CreateUser(context.Background())
	assert.Nil(t, err)

	h5s := []*coinmw.Handler{}
	h51s := []*appcoinmw.Handler{}

	for _, balance := range subscriptionOrder.PaymentBalances {
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
			appcoinmw.WithAppID(&subscriptionOrder.AppID, true),
			appcoinmw.WithCoinTypeID(&balance.CoinTypeID, true),
		)
		assert.Nil(t, err)

		_, err = h51.CreateCoin(context.Background())
		assert.Nil(t, err)

		h51s = append(h51s, h51)
	}

	for _, balance := range subscriptionOrder.PaymentTransfers {
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
			appcoinmw.WithAppID(&subscriptionOrder.AppID, true),
			appcoinmw.WithCoinTypeID(&balance.CoinTypeID, true),
		)
		assert.Nil(t, err)

		_, err = h51.CreateCoin(context.Background())
		assert.Nil(t, err)

		h51s = append(h51s, h51)
	}

	h6s := []*fiatmw.Handler{}
	fiatUnit := "USD"

	for _, balance := range subscriptionOrder.PaymentFiats {
		h6, err := fiatmw.NewHandler(
			context.Background(),
			fiatmw.WithEntID(&balance.FiatID, true),
			fiatmw.WithName(&balance.FiatID, true),
			fiatmw.WithUnit(&fiatUnit, true),
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
	denomination := "12.2"

	h7, err := couponmw.NewHandler(
		context.Background(),
		couponmw.WithEntID(&couponID, true),
		couponmw.WithAppID(&subscriptionOrder.AppID, true),
		couponmw.WithCouponType(inspiretypes.CouponType_FixAmount.Enum(), true),
		couponmw.WithStartAt(&startAt, true),
		couponmw.WithEndAt(&endAt, true),
		couponmw.WithDenomination(&denomination, true),
		couponmw.WithCirculation(&circulation, true),
		couponmw.WithCouponScope(inspiretypes.CouponScope_AllGood.Enum(), true),
	)
	assert.Nil(t, err)

	_, err = h7.CreateCoupon(context.Background())
	assert.Nil(t, err)

	h8s := []*allocatedcouponmw.Handler{}

	for _, coupon := range subscriptionOrder.Coupons {
		h8, err := allocatedcouponmw.NewHandler(
			context.Background(),
			allocatedcouponmw.WithEntID(&coupon.CouponID, true),
			allocatedcouponmw.WithAppID(&subscriptionOrder.AppID, true),
			allocatedcouponmw.WithUserID(&subscriptionOrder.UserID, true),
			allocatedcouponmw.WithCouponID(&couponID, true),
		)
		assert.Nil(t, err)

		err = h8.CreateCoupon(context.Background())
		assert.Nil(t, err)
	}

	ledgerStatementReqs := []*ledgerstatementmwpb.StatementReq{}
	userExtra := fmt.Sprintf(`{"AccountID": "%v", "UserID": "%v"}`, uuid.NewString(), uuid.NewString())

	for _, balance := range subscriptionOrder.PaymentBalances {
		ledgerStatementReqs = append(ledgerStatementReqs, &ledgerstatementmwpb.StatementReq{
			AppID:      &subscriptionOrder.AppID,
			UserID:     &subscriptionOrder.UserID,
			CurrencyID: &balance.CoinTypeID,
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

	h10, err := subscriptionordermw.NewHandler(
		context.Background(),
		subscriptionordermw.WithEntID(&subscriptionOrder.EntID, true),
		subscriptionordermw.WithAppID(&subscriptionOrder.AppID, true),
		subscriptionordermw.WithUserID(&subscriptionOrder.UserID, true),
		subscriptionordermw.WithGoodID(&subscriptionOrder.GoodID, true),
		subscriptionordermw.WithGoodType(&subscriptionOrder.GoodType, true),
		subscriptionordermw.WithAppGoodID(&subscriptionOrder.AppGoodID, true),
		subscriptionordermw.WithOrderID(&subscriptionOrder.OrderID, true),
		subscriptionordermw.WithPaymentBalances(func() (balances []*paymentmwpb.PaymentBalanceReq) {
			for _, balance := range subscriptionOrder.PaymentBalances {
				balances = append(balances, &paymentmwpb.PaymentBalanceReq{
					CoinTypeID:           &balance.CoinTypeID,
					Amount:               &balance.Amount,
					CoinUSDCurrency:      &balance.CoinUSDCurrency,
					LocalCoinUSDCurrency: &balance.LocalCoinUSDCurrency,
					LiveCoinUSDCurrency:  &balance.LiveCoinUSDCurrency,
				})
			}
			return
		}(), true),
		subscriptionordermw.WithPaymentFiats(func() (balances []*paymentmwpb.PaymentFiatReq) {
			for _, balance := range subscriptionOrder.PaymentFiats {
				balances = append(balances, &paymentmwpb.PaymentFiatReq{
					FiatID:         &balance.FiatID,
					Amount:         &balance.Amount,
					PaymentChannel: &balance.PaymentChannel,
					USDCurrency:    &balance.USDCurrency,
					// We don't have payment id here
				})
			}
			return
		}(), true),
		// WithPaymentTransferCoinTypeID(&subscriptionOrder.PaymentTransfers[0].CoinTypeID, true),
		subscriptionordermw.WithCouponIDs(func() (couponIDs []string) {
			for _, coupon := range subscriptionOrder.Coupons {
				couponIDs = append(couponIDs, coupon.CouponID)
			}
			return
		}(), true),
		subscriptionordermw.WithUserSetCanceled(&subscriptionOrder.UserSetCanceled, true),
		subscriptionordermw.WithAdminSetCanceled(&subscriptionOrder.AdminSetCanceled, true),
		subscriptionordermw.WithCreateMethod(&subscriptionOrder.CreateMethod, true),
		subscriptionordermw.WithOrderType(&subscriptionOrder.OrderType, true),
		subscriptionordermw.WithPaymentAmountUSD(&subscriptionOrder.PaymentAmountUSD, true),
		subscriptionordermw.WithLedgerLockID(&subscriptionOrder.LedgerLockID, true),
		subscriptionordermw.WithGoodValueUSD(&subscriptionOrder.GoodValueUSD, true),
	)
	assert.Nil(t, err)

	err = h10.CreateSubscriptionOrder(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h10.DeleteSubscriptionOrder(context.Background())
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

func createPayment(t *testing.T) {
	cli, err := NewPaymentClient(
		context.Background(),
		WithOrderID(subscriptionOrder.OrderID),
		WithReturnURL("http://localhost/callback"),
		WithCancelURL("http://localhost/cancel"),
	)
	if assert.Nil(t, err) {
		resp, err := cli.CreatePayment(context.Background())
		assert.Nil(t, err)

		subscriptionOrder.PaymentFiats[0].ChannelPaymentID = resp.ID
		// TODO: update channel payment id
		fmt.Printf("\n\n\nAccess %v in browser to confirm payment\n\n\n", resp.ApproveLink())
	}
}

func waitPaymentApproved(t *testing.T) {
	for {
		time.Sleep(10 * time.Second)
		cli, err := NewPaymentClient(
			context.Background(),
			WithOrderID(subscriptionOrder.OrderID),
			WithPaypalPaymentID(subscriptionOrder.PaymentFiats[0].ChannelPaymentID),
		)
		if assert.Nil(t, err) {
			resp, err := cli.GetPayment(context.Background())
			assert.Nil(t, err)

			if resp.Approved() {
				break
			}
		}
	}
}

func capturePayment(t *testing.T) {
	cli, err := NewPaymentClient(
		context.Background(),
		WithOrderID(subscriptionOrder.OrderID),
		WithPaypalPaymentID(subscriptionOrder.PaymentFiats[0].ChannelPaymentID),
	)
	if assert.Nil(t, err) {
		err = cli.CapturePayment(context.Background())
		assert.Nil(t, err)
	}
}

func getPayment(t *testing.T) {
	cli, err := NewPaymentClient(
		context.Background(),
		WithOrderID(subscriptionOrder.OrderID),
		WithPaypalPaymentID(subscriptionOrder.PaymentFiats[0].ChannelPaymentID),
	)
	if assert.Nil(t, err) {
		resp, err := cli.GetPayment(context.Background())
		assert.Nil(t, err)

		fmt.Println(resp.String())
	}
}

func createPlan(t *testing.T) {
	cli, err := NewPaymentClient(
		context.Background(),
		WithAppGoodID(subscriptionOrder.AppGoodID),
	)
	if assert.Nil(t, err) {
		resp, err := cli.CreatePlan(context.Background())
		assert.Nil(t, err)

		paypalPlanID = resp.ID
	}
}

func createSubscription(t *testing.T) {
	cli, err := NewPaymentClient(
		context.Background(),
		WithOrderID(subscriptionOrder.OrderID),
		WithAppGoodID(subscriptionOrder.AppGoodID),
		WithPaypalPlanID(paypalPlanID),
		WithReturnURL("http://localhost/callback"),
		WithCancelURL("http://localhost/cancel"),
	)
	if assert.Nil(t, err) {
		resp, err := cli.CreateSubscription(context.Background())
		assert.Nil(t, err)

		fmt.Printf("\n\n\nAccess %v in browser to confirm subscription\n\n\n", resp.ApproveLink())
		paypalSubscriptionID = resp.ID
	}
}

func waitSubscriptionApproved(t *testing.T) {
	for {
		time.Sleep(10 * time.Second)
		cli, err := NewPaymentClient(
			context.Background(),
			WithPaypalSubscriptionID(paypalSubscriptionID),
		)
		if assert.Nil(t, err) {
			resp, err := cli.GetSubscription(context.Background())
			assert.Nil(t, err)

			fmt.Println(resp.ID, resp.Status)
		}
	}
}

func getSubscription(t *testing.T) {
	cli, err := NewPaymentClient(
		context.Background(),
		WithPaypalSubscriptionID(paypalSubscriptionID),
	)
	if assert.Nil(t, err) {
		resp, err := cli.GetSubscription(context.Background())
		assert.Nil(t, err)

		fmt.Println(resp)
	}
}

func TestPaypal(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	runByCI, _ := strconv.ParseBool(os.Getenv("RUN_BY_CI"))
	os.Setenv("PAYPAL_MODE", "sandbox")

	if _, err := LoadConfig(); err != nil || runByCI {
		fmt.Printf("\n\n\n\033[1;31m         WE DO NOT RUN UNIT TEST FOR THIS MODULE DUE TO PAYPAL CONFIG \033[0m\n")
		fmt.Printf("\033[1;31m         WE DO NOT RUN UNIT TEST FOR THIS MODULE DUE TO PAYPAL CONFIG \033[0m\n")
		fmt.Printf("\033[1;31m         WE DO NOT RUN UNIT TEST FOR THIS MODULE DUE TO PAYPAL CONFIG \033[0m\n")
		fmt.Printf("\n\n\n")
		return
	}

	teardown := setup(t)
	defer teardown(t)

	// t.Run("createPayment", createPayment)
	// t.Run("waitPaymentApproved", waitPaymentApproved)
	// t.Run("capturePayment", capturePayment)
	// t.Run("getPayment", getPayment)

	t.Run("createPlan", createPlan)
	t.Run("createSubscription", createSubscription)
	t.Run("waitSubscriptionApproved", waitSubscriptionApproved)
	t.Run("getSubscription", getSubscription)
}
