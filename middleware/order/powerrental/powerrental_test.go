package powerrental

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
	feeordermiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	ordercouponmiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/coupon"
	paymentmiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"

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

var ret = npool.PowerRentalOrder{
	EntID:               uuid.NewString(),
	AppID:               uuid.NewString(),
	UserID:              uuid.NewString(),
	GoodID:              uuid.NewString(),
	GoodType:            goodtypes.GoodType_PowerRental,
	AppGoodID:           uuid.NewString(),
	OrderID:             uuid.NewString(),
	OrderType:           types.OrderType_Normal,
	AppGoodStockID:      uuid.NewString(),
	Units:               decimal.NewFromInt(10).String(),
	PaymentType:         types.PaymentType_PayWithBalanceOnly,
	CreateMethod:        types.OrderCreateMethod_OrderCreatedByPurchase,
	GoodValueUSD:        decimal.NewFromInt(120).String(),
	PaymentGoodValueUSD: decimal.NewFromInt(220).String(),
	PaymentAmountUSD:    decimal.NewFromInt(110).String(),
	DiscountAmountUSD:   decimal.NewFromInt(10).String(),
	PromotionID:         uuid.NewString(),
	DurationSeconds:     100000,
	AppGoodStockLockID:  uuid.NewString(),
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
			Amount:               decimal.NewFromInt(10).String(),
			CoinUSDCurrency:      decimal.NewFromInt(1).String(),
			LocalCoinUSDCurrency: decimal.NewFromInt(1).String(),
			LiveCoinUSDCurrency:  decimal.NewFromInt(1).String(),
		},
		{
			CoinTypeID:           uuid.NewString(),
			Amount:               decimal.NewFromInt(100).String(),
			CoinUSDCurrency:      decimal.NewFromInt(1).String(),
			LocalCoinUSDCurrency: decimal.NewFromInt(1).String(),
			LiveCoinUSDCurrency:  decimal.NewFromInt(1).String(),
		},
	},
	OrderState:   types.OrderState_OrderStateCreated,
	PaymentState: types.PaymentState_PaymentStateWait,
	RenewState:   types.OrderRenewState_OrderRenewWait,
	StartMode:    types.OrderStartMode_OrderStartInstantly,
	FeeDurations: []*feeordermiddlewarepb.FeeDuration{
		{
			AppGoodID:            uuid.NewString(),
			TotalDurationSeconds: 16000,
		},
	},
	InvestmentType: types.InvestmentType_FullPayment,
	GoodStockMode:  goodtypes.GoodStockMode_GoodStockByUnique,
	BenefitState:   types.BenefitState_BenefitWait,
}

//nolint:unparam
func setup(t *testing.T) func(*testing.T) {
	for _, paymentBalance := range ret.PaymentBalances {
		paymentBalance.PaymentID = ret.PaymentID
	}
	for _, orderCoupon := range ret.Coupons {
		orderCoupon.OrderID = ret.OrderID
	}
	for _, feeDuration := range ret.FeeDurations {
		feeDuration.ParentOrderID = ret.OrderID
	}

	ret.EndAt = ret.StartAt + ret.DurationSeconds
	ret.GoodTypeStr = ret.GoodType.String()
	ret.OrderTypeStr = ret.OrderType.String()
	ret.PaymentTypeStr = ret.PaymentType.String()
	ret.CreateMethodStr = ret.CreateMethod.String()
	ret.OrderStateStr = ret.OrderState.String()
	ret.PaymentStateStr = ret.PaymentState.String()
	ret.CancelStateStr = ret.CancelState.String()
	ret.RenewStateStr = ret.RenewState.String()
	ret.StartModeStr = ret.StartMode.String()
	ret.InvestmentTypeStr = ret.InvestmentType.String()
	ret.GoodStockModeStr = ret.GoodStockMode.String()
	ret.BenefitStateStr = ret.BenefitState.String()

	return func(*testing.T) {}
}

func createPowerRental(t *testing.T) {
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
		WithAppGoodStockID(&ret.AppGoodStockID, true),
		WithUnits(&ret.Units, true),
		WithPaymentType(&ret.PaymentType, true),
		WithCreateMethod(&ret.CreateMethod, true),
		WithGoodValueUSD(&ret.GoodValueUSD, true),
		WithPaymentAmountUSD(&ret.PaymentAmountUSD, true),
		WithDiscountAmountUSD(&ret.DiscountAmountUSD, true),
		WithPromotionID(&ret.PromotionID, true),
		WithDurationSeconds(&ret.DurationSeconds, true),
		WithStartMode(&ret.StartMode, true),
		WithAppGoodStockLockID(&ret.AppGoodStockLockID, true),
		WithLedgerLockID(&ret.LedgerLockID, true),
		WithPaymentID(&ret.PaymentID, true),
		WithGoodStockMode(&ret.GoodStockMode, true),
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
		WithFeeOrders(func() (_reqs []*feeordermiddlewarepb.FeeOrderReq) {
			for _, req := range ret.FeeDurations {
				_reqs = append(_reqs, &feeordermiddlewarepb.FeeOrderReq{
					EntID:           func() *string { s := uuid.NewString(); return &s }(),
					GoodID:          func() *string { s := uuid.NewString(); return &s }(),
					GoodType:        func() *goodtypes.GoodType { e := goodtypes.GoodType_TechniqueServiceFee; return &e }(),
					AppGoodID:       &req.AppGoodID,
					OrderID:         func() *string { s := uuid.NewString(); return &s }(),
					GoodValueUSD:    func() *string { s := decimal.NewFromInt(100).String(); return &s }(),
					DurationSeconds: &req.TotalDurationSeconds,
				})
			}
			return
		}(), true),
	)
	if assert.Nil(t, err) {
		err = handler.CreatePowerRental(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetPowerRental(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.StartAt = info.StartAt
				ret.EndAt = info.EndAt
				ret.ID = info.ID
				ret.FeeDurations = info.FeeDurations // Here fee order is not paid
				ret.PaymentBalances = info.PaymentBalances
				for _, orderCoupon := range ret.Coupons {
					orderCoupon.CreatedAt = ret.CreatedAt
				}
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func updatePowerRental(t *testing.T) {
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
		WithDurationSeconds(&ret.DurationSeconds, true),
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
		err = handler.UpdatePowerRental(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetPowerRental(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				ret.PaymentBalances = info.PaymentBalances
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getPowerRental(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithOrderID(&ret.OrderID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetPowerRental(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getPowerRentals(t *testing.T) {
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
		infos, total, err := handler.GetPowerRentals(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deletePowerRental(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithOrderID(&ret.OrderID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeletePowerRental(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetPowerRental(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestPowerRental(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createPowerRental", createPowerRental)
	t.Run("updatePowerRental", updatePowerRental)
	t.Run("getPowerRental", getPowerRental)
	t.Run("getPowerRentals", getPowerRentals)
	t.Run("deletePowerRental", deletePowerRental)
}
