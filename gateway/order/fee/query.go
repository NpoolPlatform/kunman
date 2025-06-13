package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	topmostmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost"
	allocatedcouponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	ordercoupongwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order/coupon"
	paymentgwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/payment"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	feeordermwcli "github.com/NpoolPlatform/kunman/middleware/order/fee"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	fees             []*feeordermwpb.FeeOrder
	infos            []*npool.FeeOrder
	apps             map[string]*appmwpb.App
	users            map[string]*usermwpb.User
	appFees          map[string]*appfeemwpb.Fee
	parentAppGoods   map[string]*appgoodmwpb.Good
	topMosts         map[string]*topmostmwpb.TopMost
	allocatedCoupons map[string]*allocatedcouponmwpb.Coupon
	coins            map[string]*coinmwpb.Coin
	paymentAccounts  map[string]*paymentaccountmwpb.Account
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = ordergwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, fee := range h.fees {
			appIDs = append(appIDs, fee.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = ordergwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, fee := range h.fees {
			userIDs = append(userIDs, fee.UserID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getParentAppGoods(ctx context.Context) (err error) {
	h.parentAppGoods, err = ordergwcommon.GetAppGoods(ctx, func() (appGoodIDs []string) {
		for _, fee := range h.fees {
			appGoodIDs = append(appGoodIDs, fee.ParentAppGoodID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getAppFees(ctx context.Context) (err error) {
	h.appFees, err = ordergwcommon.GetAppFees(ctx, func() (appGoodIDs []string) {
		for _, fee := range h.fees {
			appGoodIDs = append(appGoodIDs, fee.AppGoodID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getTopMosts(ctx context.Context) (err error) {
	h.topMosts, err = ordergwcommon.GetTopMosts(ctx, func() (topMostIDs []string) {
		for _, fee := range h.fees {
			if _, err := uuid.Parse(fee.PromotionID); err != nil {
				continue
			}
			topMostIDs = append(topMostIDs, fee.PromotionID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getAllocatedCoupons(ctx context.Context) (err error) {
	h.allocatedCoupons, err = ordergwcommon.GetAllocatedCoupons(ctx, func() (allocatedCouponIDs []string) {
		for _, fee := range h.fees {
			for _, coupon := range fee.Coupons {
				allocatedCouponIDs = append(allocatedCouponIDs, coupon.CouponID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.coins, err = ordergwcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, fee := range h.fees {
			for _, balance := range fee.PaymentBalances {
				coinTypeIDs = append(coinTypeIDs, balance.CoinTypeID)
			}
			for _, transfer := range fee.PaymentTransfers {
				coinTypeIDs = append(coinTypeIDs, transfer.CoinTypeID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getPaymentAccounts(ctx context.Context) (err error) {
	h.paymentAccounts, err = ordergwcommon.GetPaymentAccounts(ctx, func() (accountIDs []string) {
		for _, fee := range h.fees {
			for _, transfer := range fee.PaymentTransfers {
				accountIDs = append(accountIDs, transfer.AccountID)
			}
		}
		return
	}())
	return err
}

//nolint:funlen
func (h *queryHandler) formalize() {
	for _, fee := range h.fees {
		info := &npool.FeeOrder{
			ID:                fee.ID,
			EntID:             fee.EntID,
			AppID:             fee.AppID,
			UserID:            fee.UserID,
			GoodID:            fee.GoodID,
			GoodType:          fee.GoodType,
			AppGoodID:         fee.AppGoodID,
			OrderID:           fee.OrderID,
			ParentOrderID:     fee.ParentOrderID,
			ParentAppGoodID:   fee.ParentAppGoodID,
			ParentGoodType:    fee.ParentGoodType,
			OrderType:         fee.OrderType,
			PaymentType:       fee.PaymentType,
			CreateMethod:      fee.CreateMethod,
			OrderState:        fee.OrderState,
			GoodValueUSD:      fee.GoodValueUSD,
			PaymentAmountUSD:  fee.PaymentAmountUSD,
			DiscountAmountUSD: fee.DiscountAmountUSD,
			PromotionID:       fee.PromotionID,
			DurationSeconds:   fee.DurationSeconds,
			CancelState:       fee.CancelState,
			CanceledAt:        fee.CanceledAt,
			PaidAt:            fee.PaidAt,
			UserSetPaid:       fee.UserSetPaid,
			UserSetCanceled:   fee.UserSetCanceled,
			AdminSetCanceled:  fee.AdminSetCanceled,
			PaymentState:      fee.PaymentState,
			CreatedAt:         fee.CreatedAt,
			UpdatedAt:         fee.UpdatedAt,
		}
		app, ok := h.apps[fee.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[fee.UserID]
		if ok {
			info.EmailAddress = user.EmailAddress
			info.PhoneNO = user.PhoneNO
		}
		appFee, ok := h.appFees[fee.AppGoodID]
		if ok {
			info.GoodName = appFee.GoodName
			info.AppGoodName = appFee.AppGoodName
			info.DurationDisplayType = appFee.DurationDisplayType
			info.Durations, info.DurationUnit = ordergwcommon.GoodDurationDisplayType2Unit(
				appFee.DurationDisplayType, info.DurationSeconds,
			)
		}
		parentAppGood, ok := h.parentAppGoods[fee.ParentAppGoodID]
		if ok {
			info.ParentAppGoodName = parentAppGood.AppGoodName
		}
		topMost, ok := h.topMosts[fee.PromotionID]
		if ok {
			info.TopMostTitle = topMost.Title
			info.TopMostTargetUrl = topMost.TargetUrl
		}
		for _, coupon := range fee.Coupons {
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
		for _, balance := range fee.PaymentBalances {
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
		for _, transfer := range fee.PaymentTransfers {
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
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetFeeOrder(ctx context.Context) (*npool.FeeOrder, error) {
	if err := h.CheckOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := feeordermwcli.GetFeeOrder(ctx, *h.OrderID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid feeorder")
	}

	handler := &queryHandler{
		Handler: h,
		fees:    []*feeordermwpb.FeeOrder{info},
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getParentAppGoods(ctx); err != nil {
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

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, wlog.Errorf("invalid order")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetFeeOrders(ctx context.Context) ([]*npool.FeeOrder, uint32, error) { //nolint:gocyclo
	conds := &feeordermwpb.Conds{}
	if h.OrderCheckHandler.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.AppID}
	}
	if h.OrderCheckHandler.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.OrderCheckHandler.UserID}
	}
	if h.GoodID != nil {
		conds.GoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.GoodID}
	}
	if h.AppGoodID != nil {
		conds.AppGoodID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodID}
	}
	if len(h.OrderIDs) > 0 {
		conds.OrderIDs = &basetypes.StringSliceVal{Op: cruder.IN, Value: h.OrderIDs}
	}
	infos, total, err := feeordermwcli.GetFeeOrders(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler: h,
		fees:    infos,
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getParentAppGoods(ctx); err != nil {
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

	handler.formalize()

	return handler.infos, total, nil
}
