package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"

	"github.com/google/uuid"
)

type createHandler struct {
	*baseCreateHandler
}

//nolint:funlen,gocyclo
func (h *Handler) CreatePowerRentalOrder(ctx context.Context) (*npool.PowerRentalOrder, error) {
	handler := &createHandler{
		baseCreateHandler: &baseCreateHandler{
			checkHandler: &checkHandler{
				Handler: h,
			},
			OrderOpHandler: &ordercommon.OrderOpHandler{
				OrderType:                   *h.OrderType,
				AppGoodCheckHandler:         h.AppGoodCheckHandler,
				CoinCheckHandler:            h.CoinCheckHandler,
				AppGoodIDs:                  append(h.FeeAppGoodIDs, *h.AppGoodID),
				AllocatedCouponCheckHandler: h.AllocatedCouponCheckHandler,
				PaymentTransferCoinTypeID:   h.PaymentTransferCoinTypeID,
				PaymentBalanceReqs:          h.Balances,
				AllocatedCouponIDs:          h.CouponIDs,
				Simulate:                    h.Simulate != nil && *h.Simulate,
			},
			appGoodStockLockID: func() *string { s := uuid.NewString(); return &s }(),
		},
	}

	if err := handler.GetApp(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAppConfig(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetUser(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetRequiredAppGoods(ctx, *h.AppGoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAppGoods(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateRequiredAppGoods(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAllocatedCoupons(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.ValidateCouponScope(ctx, h.AppGoodID); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.ValidateCouponCount(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAppPowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if !handler.Simulate &&
		*h.OrderType != types.OrderType_Offline &&
		*h.OrderType != types.OrderType_Airdrop {
		if err := handler.ValidateMaxUnpaidOrders(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.validateOrderUnits(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if handler.Simulate {
		if err := handler.getAppPowerRentalSimulate(ctx); err != nil {
			return nil, wlog.WrapError(err)
		}
		if err := handler.formalizeSimulateOrder(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if err := handler.validateOrderDuration(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getGoodCoins(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetAppCoins(ctx, func() (coinTypeIDs []string) {
		for _, goodCoin := range handler.goodCoins {
			coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
		}
		return
	}()); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.formalizeOrderBenefitReqs(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateOrderBenefitReqs(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validatePowerRentalGoodState(); err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := handler.GetTopMostAppGoods(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.GetCoinUSDCurrencies(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.formalizeFeeAppGoodIDs()
	if err := handler.getAppFees(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.checkEnableSimulateOrder(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.AcquirePaymentTransferAccount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	defer handler.ReleasePaymentTransferAccount()
	if err := handler.GetPaymentTransferStartAmount(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.constructPowerRentalOrderReq(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.constructFeeOrderReqs(); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.constructOrderBenefitReqs()
	if err := handler.calculateTotalGoodValueUSD(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if !handler.Simulate {
		if err := handler.CalculateDeductAmountUSD(); err != nil {
			return nil, wlog.WrapError(err)
		}
		handler.CalculatePaymentAmountUSD()
		if err := handler.ConstructOrderPayment(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	if err := handler.ResolvePaymentType(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if !handler.Simulate {
		handler.PrepareLedgerLockID()
		handler.PreparePaymentID()
		handler.formalizePayment()
		if err := handler.ValidateCouponConstraint(); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	if err := handler.createPowerRentalOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	if !handler.Simulate && *h.OrderType == types.OrderType_Normal {
		existGoodID, err := handler.checkExistEventGood(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		handler.rewardPurchase(existGoodID)
		handler.rewardAffiliatePurchase(existGoodID)
	}

	return h.GetPowerRentalOrder(ctx)
}
