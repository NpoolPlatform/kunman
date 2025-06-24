package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	fiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	topmostmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost"
	appsubscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription"
	allocatedcouponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
	ordercoupongwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order/coupon"
	paymentgwpb "github.com/NpoolPlatform/kunman/message/order/gateway/v1/payment"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"
	subscriptionordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/subscription"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
	common "github.com/NpoolPlatform/kunman/pkg/common"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	subscriptionOrders []*subscriptionordermwpb.SubscriptionOrder
	infos              []*npool.SubscriptionOrder
	apps               map[string]*appmwpb.App
	users              map[string]*usermwpb.User
	topMosts           map[string]*topmostmwpb.TopMost
	allocatedCoupons   map[string]*allocatedcouponmwpb.Coupon
	coins              map[string]*coinmwpb.Coin
	fiats              map[string]*fiatmwpb.Fiat
	paymentAccounts    map[string]*paymentaccountmwpb.Account
	appSubscriptions   map[string]*appsubscriptionmwpb.Subscription
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = common.GetApps(ctx, func() (appIDs []string) {
		for _, subscriptionOrder := range h.subscriptionOrders {
			appIDs = append(appIDs, subscriptionOrder.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = common.GetUsers(ctx, func() (userIDs []string) {
		for _, subscriptionOrder := range h.subscriptionOrders {
			userIDs = append(userIDs, subscriptionOrder.UserID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getAppSubscriptions(ctx context.Context) (err error) {
	h.appSubscriptions, err = common.GetAppSubscriptions(ctx, func() (appGoodIDs []string) {
		for _, subscriptionOrder := range h.subscriptionOrders {
			appGoodIDs = append(appGoodIDs, subscriptionOrder.AppGoodID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getTopMosts(ctx context.Context) (err error) {
	h.topMosts, err = common.GetTopMosts(ctx, func() (topMostIDs []string) {
		for _, subscriptionOrder := range h.subscriptionOrders {
			if _, err := uuid.Parse(subscriptionOrder.PromotionID); err != nil {
				continue
			}
			topMostIDs = append(topMostIDs, subscriptionOrder.PromotionID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getAllocatedCoupons(ctx context.Context) (err error) {
	h.allocatedCoupons, err = common.GetAllocatedCoupons(ctx, func() (allocatedCouponIDs []string) {
		for _, subscriptionOrder := range h.subscriptionOrders {
			for _, coupon := range subscriptionOrder.Coupons {
				allocatedCouponIDs = append(allocatedCouponIDs, coupon.CouponID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getCoins(ctx context.Context) (err error) {
	h.coins, err = common.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, subscriptionOrder := range h.subscriptionOrders {
			for _, balance := range subscriptionOrder.PaymentBalances {
				coinTypeIDs = append(coinTypeIDs, balance.CoinTypeID)
			}
			for _, transfer := range subscriptionOrder.PaymentTransfers {
				coinTypeIDs = append(coinTypeIDs, transfer.CoinTypeID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) getFiats(ctx context.Context) (err error) {
	if h.PaymentFiatID == nil {
		return nil
	}
	h.fiats, err = common.GetFiats(ctx, []string{*h.PaymentFiatID})
	return err
}

func (h *queryHandler) getPaymentAccounts(ctx context.Context) (err error) {
	h.paymentAccounts, err = common.GetPaymentAccounts(ctx, func() (accountIDs []string) {
		for _, subscriptionOrder := range h.subscriptionOrders {
			for _, transfer := range subscriptionOrder.PaymentTransfers {
				accountIDs = append(accountIDs, transfer.AccountID)
			}
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, subscriptionOrder := range h.subscriptionOrders {
		info := &npool.SubscriptionOrder{
			ID:                  subscriptionOrder.ID,
			EntID:               subscriptionOrder.EntID,
			AppID:               subscriptionOrder.AppID,
			UserID:              subscriptionOrder.UserID,
			GoodID:              subscriptionOrder.GoodID,
			GoodType:            subscriptionOrder.GoodType,
			AppGoodID:           subscriptionOrder.AppGoodID,
			OrderID:             subscriptionOrder.OrderID,
			OrderType:           subscriptionOrder.OrderType,
			PaymentType:         subscriptionOrder.PaymentType,
			CreateMethod:        subscriptionOrder.CreateMethod,
			OrderState:          subscriptionOrder.OrderState,
			GoodValueUSD:        subscriptionOrder.GoodValueUSD,
			PaymentGoodValueUSD: subscriptionOrder.PaymentGoodValueUSD,
			PaymentAmountUSD:    subscriptionOrder.PaymentAmountUSD,
			DiscountAmountUSD:   subscriptionOrder.DiscountAmountUSD,
			PromotionID:         subscriptionOrder.PromotionID,
			LifeSeconds:         subscriptionOrder.LifeSeconds,
			CancelState:         subscriptionOrder.CancelState,
			CanceledAt:          subscriptionOrder.CanceledAt,
			PaidAt:              subscriptionOrder.PaidAt,
			UserSetPaid:         subscriptionOrder.UserSetPaid,
			UserSetCanceled:     subscriptionOrder.UserSetCanceled,
			AdminSetCanceled:    subscriptionOrder.AdminSetCanceled,
			PaymentState:        subscriptionOrder.PaymentState,
			CreatedAt:           subscriptionOrder.CreatedAt,
			UpdatedAt:           subscriptionOrder.UpdatedAt,
		}
		app, ok := h.apps[subscriptionOrder.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[subscriptionOrder.UserID]
		if ok {
			info.EmailAddress = user.EmailAddress
			info.PhoneNO = user.PhoneNO
		}

		appSubscription, ok := h.appSubscriptions[subscriptionOrder.AppGoodID]
		if ok {
			info.GoodName = appSubscription.GoodName
			info.AppGoodName = appSubscription.AppGoodName
			info.DurationDisplayType = appSubscription.DurationDisplayType
			info.Durations, info.DurationUnit = common.GoodDurationDisplayType2Unit(
				appSubscription.DurationDisplayType, info.LifeSeconds,
			)
		}
		topMost, ok := h.topMosts[subscriptionOrder.PromotionID]
		if ok {
			info.TopMostTitle = topMost.Title
			info.TopMostTargetUrl = topMost.TargetUrl
		}

		for _, coupon := range subscriptionOrder.Coupons {
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
		for _, balance := range subscriptionOrder.PaymentBalances {
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
		for _, transfer := range subscriptionOrder.PaymentTransfers {
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
		for _, fiat := range subscriptionOrder.PaymentFiats {
			paymentFiat := &paymentgwpb.PaymentFiatInfo{
				FiatID:           fiat.FiatID,
				PaymentChannel:   fiat.PaymentChannel,
				Amount:           fiat.Amount,
				USDCurrency:      fiat.USDCurrency,
				ChannelPaymentID: fiat.ChannelPaymentID,
				ApproveLink:      fiat.ApproveLink,
			}
			_fiat, ok := h.fiats[fiat.FiatID]
			if ok {
				paymentFiat.FiatName = _fiat.Name
				paymentFiat.FiatUnit = _fiat.Unit
				paymentFiat.FiatLogo = _fiat.Logo
			}
			info.PaymentFiats = append(info.PaymentFiats, paymentFiat)
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetSubscriptionOrder(ctx context.Context) (*npool.SubscriptionOrder, error) {
	if err := h.CheckOrder(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	prHandler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithOrderID(h.OrderID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	info, err := prHandler.GetSubscriptionOrder(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid subscriptionorder")
	}

	handler := &queryHandler{
		Handler:            h,
		subscriptionOrders: []*subscriptionordermwpb.SubscriptionOrder{info},
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getAppSubscriptions(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getTopMosts(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.getFiats(ctx); err != nil {
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

func (h *Handler) GetSubscriptionOrders(ctx context.Context) ([]*npool.SubscriptionOrder, uint32, error) { //nolint:gocyclo
	conds := &subscriptionordermwpb.Conds{}
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

	prHandler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithConds(conds),
		subscriptionordermw.WithOffset(h.Offset),
		subscriptionordermw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	infos, total, err := prHandler.GetSubscriptionOrders(ctx)
	if err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if len(infos) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:            h,
		subscriptionOrders: infos,
	}

	if err := handler.getApps(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getAppSubscriptions(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getTopMosts(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getCoins(ctx); err != nil {
		return nil, 0, wlog.WrapError(err)
	}
	if err := handler.getFiats(ctx); err != nil {
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
