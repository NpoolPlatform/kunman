package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	orderbenefitmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/orderbenefit"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	topmostmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	allocatedcouponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
	orderusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	feeordergwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	ordercoupongwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order/coupon"
	paymentgwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/payment"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	powerrentalordermwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	powerRentalOrders []*powerrentalordermwpb.PowerRentalOrder
	infos             []*npool.PowerRentalOrder
	apps              map[string]*appmwpb.App
	appFees           map[string]*appfeemwpb.Fee
	users             map[string]*usermwpb.User
	topMosts          map[string]*topmostmwpb.TopMost
	allocatedCoupons  map[string]*allocatedcouponmwpb.Coupon
	coins             map[string]*coinmwpb.Coin
	paymentAccounts   map[string]*paymentaccountmwpb.Account
	appPowerRentals   map[string]*apppowerrentalmwpb.PowerRental
	poolOrderUsers    map[string]*orderusermwpb.OrderUser
	orderBenefits     map[string][]*orderbenefitmwpb.Account
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = ordergwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, powerRentalOrder := range h.powerRentalOrders {
			appIDs = append(appIDs, powerRentalOrder.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = ordergwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, powerRentalOrder := range h.powerRentalOrders {
			userIDs = append(userIDs, powerRentalOrder.UserID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getAppPowerRentals(ctx context.Context) (err error) {
	h.appPowerRentals, err = ordergwcommon.GetAppPowerRentals(ctx, func() (appGoodIDs []string) {
		for _, powerRentalOrder := range h.powerRentalOrders {
			appGoodIDs = append(appGoodIDs, powerRentalOrder.AppGoodID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getAppFees(ctx context.Context) (err error) {
	h.appFees, err = ordergwcommon.GetAppFees(ctx, func() (appGoodIDs []string) {
		for _, powerRentalOrder := range h.powerRentalOrders {
			for _, feeDuration := range powerRentalOrder.FeeDurations {
				appGoodIDs = append(appGoodIDs, feeDuration.AppGoodID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getTopMosts(ctx context.Context) (err error) {
	h.topMosts, err = ordergwcommon.GetTopMosts(ctx, func() (topMostIDs []string) {
		for _, powerRentalOrder := range h.powerRentalOrders {
			if _, err := uuid.Parse(powerRentalOrder.PromotionID); err != nil {
				continue
			}
			topMostIDs = append(topMostIDs, powerRentalOrder.PromotionID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getAllocatedCoupons(ctx context.Context) (err error) {
	h.allocatedCoupons, err = ordergwcommon.GetAllocatedCoupons(ctx, func() (allocatedCouponIDs []string) {
		for _, powerRentalOrder := range h.powerRentalOrders {
			for _, coupon := range powerRentalOrder.Coupons {
				allocatedCouponIDs = append(allocatedCouponIDs, coupon.CouponID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.coins, err = ordergwcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, powerRentalOrder := range h.powerRentalOrders {
			for _, balance := range powerRentalOrder.PaymentBalances {
				coinTypeIDs = append(coinTypeIDs, balance.CoinTypeID)
			}
			for _, transfer := range powerRentalOrder.PaymentTransfers {
				coinTypeIDs = append(coinTypeIDs, transfer.CoinTypeID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getPaymentAccounts(ctx context.Context) (err error) {
	h.paymentAccounts, err = ordergwcommon.GetPaymentAccounts(ctx, func() (accountIDs []string) {
		for _, powerRentalOrder := range h.powerRentalOrders {
			for _, transfer := range powerRentalOrder.PaymentTransfers {
				accountIDs = append(accountIDs, transfer.AccountID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getOrderBenefits(ctx context.Context) (err error) {
	h.orderBenefits, err = ordergwcommon.GetOrderBenefits(ctx, func() (orderIDs []string) {
		for _, powerrentalOrder := range h.powerRentalOrders {
			orderIDs = append(orderIDs, powerrentalOrder.OrderID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getMiningPoolOrderUsers(ctx context.Context) (err error) {
	h.poolOrderUsers, err = ordergwcommon.GetMiningPoolOrderUsers(ctx, func() (orderuserIDs []string) {
		for _, powerrentalOrder := range h.powerRentalOrders {
			if powerrentalOrder.PoolOrderUserID != nil {
				orderuserIDs = append(orderuserIDs, *powerrentalOrder.PoolOrderUserID)
			}
		}
		return
	}())
	return err
}

//nolint:funlen,gocyclo
func (h *queryHandler) formalize() {
	for _, powerRentalOrder := range h.powerRentalOrders {
		info := &npool.PowerRentalOrder{
			ID:             powerRentalOrder.ID,
			EntID:          powerRentalOrder.EntID,
			AppID:          powerRentalOrder.AppID,
			UserID:         powerRentalOrder.UserID,
			GoodID:         powerRentalOrder.GoodID,
			GoodType:       powerRentalOrder.GoodType,
			AppGoodID:      powerRentalOrder.AppGoodID,
			OrderID:        powerRentalOrder.OrderID,
			OrderType:      powerRentalOrder.OrderType,
			PaymentType:    powerRentalOrder.PaymentType,
			CreateMethod:   powerRentalOrder.CreateMethod,
			Simulate:       powerRentalOrder.Simulate,
			OrderState:     powerRentalOrder.OrderState,
			StartMode:      powerRentalOrder.StartMode,
			StartAt:        powerRentalOrder.StartAt,
			LastBenefitAt:  powerRentalOrder.LastBenefitAt,
			BenefitState:   powerRentalOrder.BenefitState,
			AppGoodStockID: powerRentalOrder.AppGoodStockID,
			// TODO: mining pool information
			Units:             powerRentalOrder.Units,
			GoodValueUSD:      powerRentalOrder.GoodValueUSD,
			PaymentAmountUSD:  powerRentalOrder.PaymentAmountUSD,
			DiscountAmountUSD: powerRentalOrder.DiscountAmountUSD,
			PromotionID:       powerRentalOrder.PromotionID,
			DurationSeconds:   powerRentalOrder.DurationSeconds,
			InvestmentType:    powerRentalOrder.InvestmentType,
			GoodStockMode:     powerRentalOrder.GoodStockMode,
			CancelState:       powerRentalOrder.CancelState,
			CanceledAt:        powerRentalOrder.CanceledAt,
			EndAt:             powerRentalOrder.EndAt,
			PaidAt:            powerRentalOrder.PaidAt,
			UserSetPaid:       powerRentalOrder.UserSetPaid,
			UserSetCanceled:   powerRentalOrder.UserSetCanceled,
			AdminSetCanceled:  powerRentalOrder.AdminSetCanceled,
			PaymentState:      powerRentalOrder.PaymentState,
			OutOfGasSeconds:   powerRentalOrder.OutOfGasSeconds,
			CompensateSeconds: powerRentalOrder.CompensateSeconds,
			// TODO: fee durations
			CreatedAt: powerRentalOrder.CreatedAt,
			UpdatedAt: powerRentalOrder.UpdatedAt,
		}
		app, ok := h.apps[powerRentalOrder.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[powerRentalOrder.UserID]
		if ok {
			info.EmailAddress = user.EmailAddress
			info.PhoneNO = user.PhoneNO
		}

		appPowerRental, ok := h.appPowerRentals[powerRentalOrder.AppGoodID]
		if ok {
			info.GoodName = appPowerRental.GoodName
			info.GoodQuantityUnit = appPowerRental.QuantityUnit
			info.AppGoodName = appPowerRental.AppGoodName
			info.DurationDisplayType = appPowerRental.DurationDisplayType
			info.Durations, info.DurationUnit = ordergwcommon.GoodDurationDisplayType2Unit(
				appPowerRental.DurationDisplayType, info.DurationSeconds,
			)
			info.BenefitType = appPowerRental.BenefitType
		}
		topMost, ok := h.topMosts[powerRentalOrder.PromotionID]
		if ok {
			info.TopMostTitle = topMost.Title
			info.TopMostTargetUrl = topMost.TargetUrl
		}

		orderBenefits, ok := h.orderBenefits[powerRentalOrder.OrderID]
		if ok {
			ret := []*npool.OrderBenefitAccount{}
			for _, orderBenefit := range orderBenefits {
				ret = append(ret, &npool.OrderBenefitAccount{
					AccountID:  orderBenefit.AccountID,
					CoinTypeID: orderBenefit.CoinTypeID,
					Address:    orderBenefit.Address,
				})
			}
			info.OrderBenefitAccounts = ret
		}

		if powerRentalOrder.PoolOrderUserID != nil {
			poolOrderUser, ok := h.poolOrderUsers[*powerRentalOrder.PoolOrderUserID]
			if ok {
				info.MiningPoolName = &poolOrderUser.MiningPoolName
				info.MiningPoolLogo = &poolOrderUser.MiningPoolLogo
				info.MiningPoolOrderUserID = &poolOrderUser.EntID
				info.MiningPoolOrderUserName = &poolOrderUser.Name
				info.MiningPoolReadPageLink = &poolOrderUser.ReadPageLink
			}
		}

		for _, coupon := range powerRentalOrder.Coupons {
			orderCoupon := &ordercoupongwpb.OrderCouponInfo{
				AllocatedCouponID: coupon.CouponID,
				CreatedAt:         coupon.CreatedAt,
			}
			allocatedCoupon, ok := h.allocatedCoupons[coupon.CouponID]
			if ok {
				orderCoupon.CouponType = allocatedCoupon.CouponType
				orderCoupon.Denomination = allocatedCoupon.Denomination
				orderCoupon.CouponName = allocatedCoupon.CouponName
			}
			info.Coupons = append(info.Coupons, orderCoupon)
		}
		for _, balance := range powerRentalOrder.PaymentBalances {
			paymentBalance := &paymentgwpb.PaymentBalanceInfo{
				CoinTypeID:      balance.CoinTypeID,
				Amount:          balance.Amount,
				CoinUSDCurrency: balance.CoinUSDCurrency,
				CreatedAt:       balance.CreatedAt,
			}
			coin, ok := h.coins[balance.CoinTypeID]
			if ok {
				paymentBalance.CoinName = coin.Name
				paymentBalance.CoinUnit = coin.Unit
				paymentBalance.CoinLogo = coin.Logo
				paymentBalance.CoinENV = coin.ENV
			}
			info.PaymentBalances = append(info.PaymentBalances, paymentBalance)
		}
		for _, transfer := range powerRentalOrder.PaymentTransfers {
			paymentTransfer := &paymentgwpb.PaymentTransferInfo{
				CoinTypeID:      transfer.CoinTypeID,
				Amount:          transfer.Amount,
				AccountID:       transfer.AccountID,
				CoinUSDCurrency: transfer.CoinUSDCurrency,
				CreatedAt:       transfer.CreatedAt,
			}
			coin, ok := h.coins[transfer.CoinTypeID]
			if ok {
				paymentTransfer.CoinName = coin.Name
				paymentTransfer.CoinUnit = coin.Unit
				paymentTransfer.CoinLogo = coin.Logo
				paymentTransfer.CoinENV = coin.ENV
			}
			account, ok := h.paymentAccounts[transfer.AccountID]
			if ok {
				paymentTransfer.Address = account.Address
			}
			info.PaymentTransfers = append(info.PaymentTransfers, paymentTransfer)
		}
		for _, feeDuration := range powerRentalOrder.FeeDurations {
			appFee, ok := h.appFees[feeDuration.AppGoodID]
			if !ok {
				continue
			}
			info.FeeDurations = append(info.FeeDurations, &feeordergwpb.FeeDuration{
				AppGoodID:            feeDuration.AppGoodID,
				AppGoodName:          appFee.AppGoodName,
				TotalDurationSeconds: feeDuration.TotalDurationSeconds,
			})
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetPowerRentalOrder(ctx context.Context) (*npool.PowerRentalOrder, error) {
	if err := h.CheckOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := powerrentalordermwcli.GetPowerRentalOrder(ctx, *h.OrderID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid powerrentalorder")
	}

	handler := &queryHandler{
		Handler:           h,
		powerRentalOrders: []*powerrentalordermwpb.PowerRentalOrder{info},
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAppPowerRentals(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAppFees(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getTopMosts(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getPaymentAccounts(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAllocatedCoupons(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getOrderBenefits(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getMiningPoolOrderUsers(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, wlog.Errorf("invalid order")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetPowerRentalOrders(ctx context.Context) ([]*npool.PowerRentalOrder, uint32, error) { //nolint:gocyclo
	conds := &powerrentalordermwpb.Conds{}
	if h.OrderCheckHandler.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID}
	}
	if h.OrderCheckHandler.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.UserID}
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	if len(h.OrderIDs) > 0 {
		conds.OrderIDs = &basetypes.StringSliceVal{Op: cruder.IN, Value: h.OrderIDs}
	}
	infos, total, err := powerrentalordermwcli.GetPowerRentalOrders(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:           h,
		powerRentalOrders: infos,
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getAppPowerRentals(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getAppFees(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getTopMosts(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getPaymentAccounts(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getAllocatedCoupons(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getOrderBenefits(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getMiningPoolOrderUsers(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, total, nil
}
