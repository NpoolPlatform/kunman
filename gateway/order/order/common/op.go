package common

import (
	"context"
	"os"
	"strconv"
	"time"

	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	paymentaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	inspiretypes "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	currencymwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency"
	appgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	topmostgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/good"
	allocatedcouponmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/allocated"
	appgoodscopemwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon/app/scope"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	ledgerstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	orderappconfigmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/app/config"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	paymentaccountmw "github.com/NpoolPlatform/kunman/middleware/account/payment"
	appmw "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	currencymw "github.com/NpoolPlatform/kunman/middleware/chain/coin/currency"
	appgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good"
	requiredappgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
	topmostgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/topmost/good"
	allocatedcouponmw "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/allocated"
	appgoodscopemw "github.com/NpoolPlatform/kunman/middleware/inspire/coupon/app/scope"
	ledgermw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	orderappconfigmw "github.com/NpoolPlatform/kunman/middleware/order/app/config"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
	orderlockmw "github.com/NpoolPlatform/kunman/middleware/order/order/lock"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	ordergwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderOpHandler struct {
	ordergwcommon.AppGoodCheckHandler
	ordergwcommon.CoinCheckHandler
	ordergwcommon.AllocatedCouponCheckHandler
	DurationSeconds           *uint32
	PaymentTransferCoinTypeID *string
	PaymentFiatID             *string
	AllocatedCouponIDs        []string
	AppGoodIDs                []string
	OrderType                 types.OrderType
	Simulate                  bool

	allocatedCoupons  map[string]*allocatedcouponmwpb.Coupon
	coinUSDCurrencies map[string]*currencymwpb.Currency
	AppGoods          map[string]*appgoodmwpb.Good

	PaymentBalanceReqs         []*paymentmwpb.PaymentBalanceReq
	PaymentTransferReq         *paymentmwpb.PaymentTransferReq
	PaymentFiatReq             *paymentmwpb.PaymentFiatReq
	PaymentType                *types.PaymentType
	PaymentTransferAccount     *paymentaccountmwpb.Account
	PaymentTransferStartAmount decimal.Decimal
	BalanceLockID              *string
	PaymentID                  *string

	OrderID                    *string
	OrderState                 types.OrderState
	AdminSetCanceled           *bool
	UserSetCanceled            *bool
	GoodCancelMode             goodtypes.CancelMode
	CommissionLedgerStatements []*ledgerstatementmwpb.Statement
	CommissionLockIDs          map[string]string

	DeductAmountUSD   decimal.Decimal
	PaymentAmountUSD  decimal.Decimal
	TotalGoodValueUSD decimal.Decimal

	OrderConfig      *orderappconfigmwpb.AppConfig
	App              *appmwpb.App
	User             *usermwpb.User
	AppCoins         map[string]*appcoinmwpb.Coin
	RequiredAppGoods map[string]map[string]*requiredappgoodmwpb.Required
	TopMostAppGoods  map[string]*topmostgoodmwpb.TopMostGood
}

var runInUnitTest = false

const testCoinName = "tusdttrc20"

func init() {
	runInUnitTest, _ = strconv.ParseBool(os.Getenv("RUN_IN_UNIT_TEST"))
}

func (h *OrderOpHandler) GetAppConfig(ctx context.Context) (err error) {
	handler, err := orderappconfigmw.NewHandler(
		ctx,
		orderappconfigmw.WithAppID(h.AllocatedCouponCheckHandler.AppID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	h.OrderConfig, err = handler.GetAppConfig(ctx)
	return wlog.WrapError(err)
}

func (h *OrderOpHandler) GetAllocatedCoupons(ctx context.Context) error {
	conds := &allocatedcouponmwpb.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AllocatedCouponCheckHandler.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AllocatedCouponCheckHandler.UserID},
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: h.AllocatedCouponIDs},
	}
	handler, err := allocatedcouponmw.NewHandler(
		ctx,
		allocatedcouponmw.WithConds(conds),
		allocatedcouponmw.WithOffset(0),
		allocatedcouponmw.WithLimit(int32(len(h.AllocatedCouponIDs))),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	infos, _, err := handler.GetCoupons(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(infos) != len(h.AllocatedCouponIDs) {
		return wlog.Errorf("invalid allocatedcoupons")
	}
	h.allocatedCoupons = map[string]*allocatedcouponmwpb.Coupon{}
	for _, info := range infos {
		h.allocatedCoupons[info.EntID] = info
	}
	return nil
}

func (h *OrderOpHandler) GetAppCoins(ctx context.Context, parentGoodCoinTypeIDs []string) error {
	coinTypeIDs := func() (_coinTypeIDs []string) {
		for _, balance := range h.PaymentBalanceReqs {
			_coinTypeIDs = append(_coinTypeIDs, *balance.CoinTypeID)
		}
		return
	}()
	coinTypeIDs = append(coinTypeIDs, parentGoodCoinTypeIDs...)
	if h.PaymentTransferCoinTypeID != nil {
		coinTypeIDs = append(coinTypeIDs, *h.PaymentTransferCoinTypeID)
	}

	conds := &appcoinmwpb.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.AppID},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}
	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithConds(conds),
		appcoinmw.WithOffset(0),
		appcoinmw.WithLimit(int32(len(coinTypeIDs))),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	coins, _, err := handler.GetCoins(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.AppCoins = map[string]*appcoinmwpb.Coin{}
	coinENV := ""
	for _, coin := range coins {
		if coinENV != "" && coin.ENV != coinENV {
			return wlog.Errorf("invalid appcoins")
		}
		h.AppCoins[coin.CoinTypeID] = coin
		coinENV = coin.ENV
	}
	return nil
}

func (h *OrderOpHandler) GetCoinUSDCurrencies(ctx context.Context) error {
	coinTypeIDs := func() (_coinTypeIDs []string) {
		for _, balance := range h.PaymentBalanceReqs {
			_coinTypeIDs = append(_coinTypeIDs, *balance.CoinTypeID)
		}
		return
	}()
	if h.PaymentTransferCoinTypeID != nil {
		coinTypeIDs = append(coinTypeIDs, *h.PaymentTransferCoinTypeID)
	}

	conds := &currencymwpb.Conds{
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}
	handler, err := currencymw.NewHandler(
		ctx,
		currencymw.WithConds(conds),
		currencymw.WithOffset(0),
		currencymw.WithLimit(int32(len(coinTypeIDs)*2)), // Work around for multi currency channel
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	infos, _, err := handler.GetCurrencies(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.coinUSDCurrencies = map[string]*currencymwpb.Currency{}
	now := uint32(time.Now().Unix())
	for _, info := range infos {
		if info.UpdatedAt+timedef.SecondsPerMinute*10 < now && !runInUnitTest {
			return wlog.Errorf("stale coincurrency")
		}
		h.coinUSDCurrencies[info.CoinTypeID] = info
	}
	return nil
}

func (h *OrderOpHandler) GetAppGoods(ctx context.Context) error {
	conds := &appgoodmwpb.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.AppID},
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: h.AppGoodIDs},
	}

	handler, err := appgoodmw.NewHandler(
		ctx,
		appgoodmw.WithConds(conds),
		appgoodmw.WithOffset(0),
		appgoodmw.WithLimit(int32(len(h.AppGoodIDs))),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	appGoods, _, err := handler.GetGoods(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(appGoods) != len(h.AppGoodIDs) {
		return wlog.Errorf("invalid appgoods")
	}
	h.AppGoods = map[string]*appgoodmwpb.Good{}
	for _, appGood := range appGoods {
		h.AppGoods[appGood.EntID] = appGood
	}
	return nil
}

func (h *OrderOpHandler) GetApp(ctx context.Context) error {
	handler, err := appmw.NewHandler(
		ctx,
		appmw.WithEntID(h.AppGoodCheckHandler.AppID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	app, err := handler.GetApp(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if app == nil {
		return wlog.Errorf("invalid app")
	}
	h.App = app
	return nil
}

func (h *OrderOpHandler) GetUser(ctx context.Context) error {
	handler, err := usermw.NewHandler(
		ctx,
		usermw.WithAppID(h.AppGoodCheckHandler.AppID, true),
		usermw.WithEntID(h.AppGoodCheckHandler.UserID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	user, err := handler.GetUser(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if user == nil {
		return wlog.Errorf("invalid user")
	}
	h.User = user
	return nil
}

func (h *OrderOpHandler) ValidateCouponScope(ctx context.Context, parentAppGoodID *string) error {
	if len(h.allocatedCoupons) == 0 {
		return nil
	}
	reqs := []*appgoodscopemwpb.ScopeReq{}
	for _, allocatedCoupon := range h.allocatedCoupons {
		for appGoodID, appGood := range h.AppGoods {
			if parentAppGoodID != nil && *parentAppGoodID != appGoodID {
				continue
			}
			_appGoodID := appGoodID
			reqs = append(reqs, &appgoodscopemwpb.ScopeReq{
				AppID:       h.AppGoodCheckHandler.AppID,
				AppGoodID:   &_appGoodID,
				GoodID:      &appGood.GoodID,
				CouponID:    &allocatedCoupon.CouponID,
				CouponScope: &allocatedCoupon.CouponScope,
			})
		}
	}

	handler, err := appgoodscopemw.NewHandler(
		ctx,
		appgoodscopemw.WithReqs(reqs, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return handler.VerifyCouponScopes(ctx)
}

func (h *OrderOpHandler) ValidateCouponCount() error {
	discountCoupons := 0
	fixAmountCoupons := uint32(0)
	for _, coupon := range h.allocatedCoupons {
		switch coupon.CouponType {
		case inspiretypes.CouponType_Discount:
			discountCoupons++
			if discountCoupons > 1 {
				return wlog.Errorf("invalid discountcoupon")
			}
		case inspiretypes.CouponType_FixAmount:
			fixAmountCoupons++
			if h.OrderConfig == nil || h.OrderConfig.MaxTypedCouponsPerOrder == 0 {
				continue
			}
			if fixAmountCoupons > h.OrderConfig.MaxTypedCouponsPerOrder {
				return wlog.Errorf("invalid fixamountcoupon")
			}
		}
	}
	return nil
}

func (h *OrderOpHandler) ValidateMaxUnpaidOrders(ctx context.Context) error {
	if h.OrderConfig == nil || h.OrderConfig.MaxUnpaidOrders == 0 {
		return nil
	}

	prConds := &powerrentalmwpb.Conds{
		AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.AppID},
		UserID:       &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.UserID},
		OrderType:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(types.OrderType_Normal)},
		PaymentState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(types.PaymentState_PaymentStateWait)},
	}
	prHandler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithConds(prConds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	powerRentals, err := prHandler.CountPowerRentals(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	feeConds := &feeordermwpb.Conds{
		AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.AppID},
		UserID:       &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.UserID},
		OrderType:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(types.OrderType_Normal)},
		PaymentState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(types.PaymentState_PaymentStateWait)},
	}
	feeHandler, err := feeordermw.NewHandler(
		ctx,
		feeordermw.WithConds(feeConds),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	feeOrders, err := feeHandler.CountFeeOrders(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if powerRentals+feeOrders >= h.OrderConfig.MaxUnpaidOrders {
		return wlog.Errorf("too many unpaid orders")
	}
	return nil
}

func (h *OrderOpHandler) GetRequiredAppGoods(ctx context.Context, mainAppGoodID string) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	h.RequiredAppGoods = map[string]map[string]*requiredappgoodmwpb.Required{}

	for {
		conds := &requiredappgoodmwpb.Conds{
			AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.AppID},
			MainAppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: mainAppGoodID},
			AppGoodIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: h.AppGoodIDs},
		}
		handler, err := requiredappgoodmw.NewHandler(
			ctx,
			requiredappgoodmw.WithConds(conds),
			requiredappgoodmw.WithOffset(offset),
			requiredappgoodmw.WithLimit(limit),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		requiredAppGoods, _, err := handler.GetRequireds(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(requiredAppGoods) == 0 {
			return nil
		}
		for _, requiredAppGood := range requiredAppGoods {
			requireds, ok := h.RequiredAppGoods[requiredAppGood.MainAppGoodID]
			if !ok {
				requireds = map[string]*requiredappgoodmwpb.Required{}
			}
			requireds[requiredAppGood.RequiredAppGoodID] = requiredAppGood
			h.RequiredAppGoods[requiredAppGood.MainAppGoodID] = requireds
		}
		offset += limit
	}
}

func (h *OrderOpHandler) GetTopMostAppGoods(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	h.TopMostAppGoods = map[string]*topmostgoodmwpb.TopMostGood{}

	for {
		conds := &topmostgoodmwpb.Conds{
			AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppGoodCheckHandler.AppID},
			AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: h.AppGoodIDs},
		}
		handler, err := topmostgoodmw.NewHandler(
			ctx,
			topmostgoodmw.WithConds(conds),
			topmostgoodmw.WithOffset(offset),
			topmostgoodmw.WithLimit(limit),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		topMostGoods, _, err := handler.GetTopMostGoods(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(topMostGoods) == 0 {
			return nil
		}
		for _, topMostGood := range topMostGoods {
			unitPrice, err := decimal.NewFromString(topMostGood.UnitPrice)
			if err != nil {
				return wlog.WrapError(err)
			}
			unitPrice1 := decimal.NewFromInt(0)
			existTopMostGood, ok := h.TopMostAppGoods[topMostGood.AppGoodID]
			if ok {
				unitPrice1, err = decimal.NewFromString(existTopMostGood.UnitPrice)
				if err != nil {
					return wlog.WrapError(err)
				}
			}
			if unitPrice1.Equal(decimal.NewFromInt(0)) || unitPrice.LessThan(unitPrice1) {
				h.TopMostAppGoods[topMostGood.AppGoodID] = topMostGood
			}
		}
		offset += limit
	}
}

func (h *OrderOpHandler) CalculateDeductAmountUSD() error {
	if h.TotalGoodValueUSD.Equal(decimal.NewFromInt(0)) {
		return wlog.Errorf("invalid totalgoodvalueusd")
	}
	for _, allocatedCoupon := range h.allocatedCoupons {
		switch allocatedCoupon.CouponType {
		case inspiretypes.CouponType_Discount:
			discount, err := decimal.NewFromString(allocatedCoupon.Denomination)
			if err != nil {
				return wlog.WrapError(err)
			}
			discount = discount.Div(decimal.NewFromInt(100)) //nolint
			h.DeductAmountUSD = h.DeductAmountUSD.Add(h.TotalGoodValueUSD.Mul(discount))
		case inspiretypes.CouponType_FixAmount:
			amount, err := decimal.NewFromString(allocatedCoupon.Denomination)
			if err != nil {
				return wlog.WrapError(err)
			}
			h.DeductAmountUSD = h.DeductAmountUSD.Add(amount)
		default:
			return wlog.Errorf("invalid coupontype")
		}
	}
	return nil
}

func (h *OrderOpHandler) CalculatePaymentAmountUSD() {
	h.PaymentAmountUSD = h.TotalGoodValueUSD.Sub(h.DeductAmountUSD)
	if h.PaymentAmountUSD.Cmp(decimal.NewFromInt(0)) < 0 {
		h.PaymentAmountUSD = decimal.NewFromInt(0)
	}
}

func (h *OrderOpHandler) getCoinUSDCurrency(coinTypeID string) (cur decimal.Decimal, live, local *string, err error) {
	currency, ok := h.coinUSDCurrencies[coinTypeID]
	if !ok {
		return cur, live, local, wlog.Errorf("invalid currency")
	}
	amount, err := decimal.NewFromString(currency.MarketValueLow)
	if err != nil {
		return cur, live, local, wlog.WrapError(err)
	}

	cur = amount
	live = func() *string { s := amount.String(); return &s }()

	appCoin, ok := h.AppCoins[coinTypeID]
	if !ok {
		return cur, live, local, wlog.Errorf("invalid coin")
	}

	amount, err = decimal.NewFromString(appCoin.SettleValue)
	if err != nil {
		return cur, live, local, wlog.WrapError(err)
	}
	if amount.GreaterThan(decimal.NewFromInt(0)) {
		cur = amount
	}

	amount, err = decimal.NewFromString(appCoin.MarketValue)
	if err != nil {
		return cur, live, local, wlog.WrapError(err)
	}
	if amount.GreaterThan(decimal.NewFromInt(0)) {
		local = func() *string { s := amount.String(); return &s }()
	}
	if cur.Cmp(decimal.NewFromInt(0)) <= 0 {
		return cur, live, local, wlog.Errorf("invalid currency")
	}

	return cur, live, local, nil
}

func (h *OrderOpHandler) ConstructOrderPayment() error {
	switch h.OrderType {
	case types.OrderType_Offline:
		fallthrough //nolint
	case types.OrderType_Airdrop:
		return nil
	}
	if h.Simulate {
		return nil
	}

	remainAmountUSD := h.PaymentAmountUSD
	balanceReqs := []*paymentmwpb.PaymentBalanceReq{}

	for _, balance := range h.PaymentBalanceReqs {
		_balance := balance
		cur, live, local, err := h.getCoinUSDCurrency(*_balance.CoinTypeID)
		if err != nil {
			return wlog.WrapError(err)
		}
		amount, err := decimal.NewFromString(*_balance.Amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return wlog.Errorf("invalid paymentbalanceamount")
		}
		amountUSD := amount.Mul(cur)
		if remainAmountUSD.Cmp(amountUSD) < 0 {
			amountUSD = remainAmountUSD
		}
		_balance.CoinUSDCurrency = func() *string { s := cur.String(); return &s }()
		_balance.LiveCoinUSDCurrency = live
		_balance.LocalCoinUSDCurrency = local
		if remainAmountUSD.GreaterThan(amountUSD) {
			balanceReqs = append(balanceReqs, _balance)
			remainAmountUSD = remainAmountUSD.Sub(amountUSD)
		} else {
			amountStr := remainAmountUSD.Div(cur).String()
			_balance.Amount = &amountStr
			balanceReqs = append(balanceReqs, _balance)
			h.PaymentBalanceReqs = balanceReqs
			return nil
		}
	}
	if h.PaymentTransferCoinTypeID == nil {
		return wlog.Errorf("invalid paymenttransfercointypeid")
	}
	if h.PaymentTransferAccount == nil {
		return wlog.Errorf("invalid paymenttransferaccount")
	}
	cur, live, local, err := h.getCoinUSDCurrency(*h.PaymentTransferCoinTypeID)
	if err != nil {
		return wlog.WrapError(err)
	}
	remainAmountCoin := remainAmountUSD.Div(cur)
	h.PaymentTransferReq = &paymentmwpb.PaymentTransferReq{
		CoinTypeID:           h.PaymentTransferCoinTypeID,
		Amount:               func() *string { s := remainAmountCoin.String(); return &s }(),
		AccountID:            &h.PaymentTransferAccount.AccountID,
		StartAmount:          func() *string { s := h.PaymentTransferStartAmount.String(); return &s }(),
		CoinUSDCurrency:      func() *string { s := cur.String(); return &s }(),
		LiveCoinUSDCurrency:  live,
		LocalCoinUSDCurrency: local,
	}
	return nil
}

func (h *OrderOpHandler) ValidateCouponConstraint() error {
	for _, allocatedCoupon := range h.allocatedCoupons {
		if allocatedCoupon.CouponConstraint != inspiretypes.CouponConstraint_PaymentThreshold {
			continue
		}
		thresholdAmount, err := decimal.NewFromString(allocatedCoupon.Threshold)
		if err != nil {
			return wlog.WrapError(err)
		}
		if h.PaymentAmountUSD.LessThan(thresholdAmount) {
			return wlog.Errorf("not enough payment amount")
		}
	}
	return nil
}

func (h *OrderOpHandler) ResolvePaymentType() error {
	if h.PaymentTransferReq == nil && len(h.PaymentBalanceReqs) == 0 {
		if !h.Simulate {
			switch h.OrderType {
			case types.OrderType_Offline:
			case types.OrderType_Airdrop:
			default:
				return wlog.Errorf("invalid paymenttype")
			}
		}
		h.PaymentType = types.PaymentType_PayWithNoPayment.Enum()
	}
	if h.PaymentTransferReq == nil {
		h.PaymentType = types.PaymentType_PayWithBalanceOnly.Enum()
		return nil
	}
	if len(h.PaymentBalanceReqs) == 0 {
		h.PaymentType = types.PaymentType_PayWithTransferOnly.Enum()
		return nil
	}
	h.PaymentType = types.PaymentType_PayWithTransferAndBalance.Enum()
	return nil
}

/**
 * paymentAccountID: ID of account_manager.payments
 */
func (h *OrderOpHandler) recheckPaymentAccount(ctx context.Context, paymentAccountID string) (bool, error) {
	handler, err := paymentaccountmw.NewHandler(
		ctx,
		paymentaccountmw.WithEntID(&paymentAccountID, true),
	)
	if err != nil {
		return false, wlog.WrapError(err)
	}

	account, err := handler.GetAccount(ctx)
	if err != nil {
		return false, wlog.WrapError(err)
	}
	if account == nil {
		return false, wlog.Errorf("invalid account")
	}
	if account.Locked || !account.Active || account.Blocked {
		return false, nil
	}
	if account.AvailableAt > uint32(time.Now().Unix()) {
		return false, nil
	}
	return true, nil
}

func (h *OrderOpHandler) peekExistPaymentAccount(ctx context.Context) (*paymentaccountmwpb.Account, error) {
	conds := &paymentaccountmwpb.Conds{
		CoinTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.PaymentTransferCoinTypeID},
		Active:      &basetypes.BoolVal{Op: cruder.EQ, Value: true},
		Locked:      &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		Blocked:     &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		AvailableAt: &basetypes.Uint32Val{Op: cruder.LTE, Value: uint32(time.Now().Unix())},
	}
	handler, err := paymentaccountmw.NewHandler(
		ctx,
		paymentaccountmw.WithConds(conds),
		paymentaccountmw.WithOffset(0),
		paymentaccountmw.WithLimit(5),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	// TODO: add new api to lock one account directly

	accounts, _, err := handler.GetAccounts(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	for _, account := range accounts {
		usable, err := h.recheckPaymentAccount(ctx, account.EntID)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if !usable {
			continue
		}
		return account, nil
	}
	return nil, wlog.Errorf("invalid paymentaccount")
}

func (h *OrderOpHandler) peekNewPaymentAccount(ctx context.Context) (*paymentaccountmwpb.Account, error) {
	paymentTransferCoin, ok := h.AppCoins[*h.PaymentTransferCoinTypeID]
	if !ok {
		return nil, wlog.Errorf("invalid paymenttransfercoin")
	}

	if runInUnitTest {
		paymentTransferCoin.CoinName = testCoinName
	}

	for i := 0; i < 5; i++ {
		address, err := sphinxproxycli.CreateAddress(ctx, paymentTransferCoin.CoinName)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if address == nil || address.Address == "" {
			return nil, wlog.Errorf("invalid address")
		}

		handler, err := paymentaccountmw.NewHandler(
			ctx,
			paymentaccountmw.WithCoinTypeID(&paymentTransferCoin.CoinTypeID, true),
			paymentaccountmw.WithAddress(&address.Address, true),
		)
		if err != nil {
			return nil, wlog.WrapError(err)
		}

		_, err = handler.CreateAccount(ctx)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return h.peekExistPaymentAccount(ctx)
}

func (h *OrderOpHandler) AcquirePaymentTransferAccount(ctx context.Context) error {
	if h.PaymentTransferCoinTypeID == nil {
		return nil
	}
	account, err := h.peekExistPaymentAccount(ctx)
	if err != nil {
		account, err = h.peekNewPaymentAccount(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
	}
	h.PaymentTransferAccount = account
	return nil
}

func (h *OrderOpHandler) ReleasePaymentTransferAccount() {
}

func (h *OrderOpHandler) GetPaymentTransferStartAmount(ctx context.Context) error {
	if h.PaymentTransferAccount == nil {
		return nil
	}
	paymentTransferCoin, ok := h.AppCoins[*h.PaymentTransferCoinTypeID]
	if !ok {
		return wlog.Errorf("invalid paymenttransfercoin")
	}

	if runInUnitTest {
		paymentTransferCoin.CoinName = testCoinName
	}

	balance, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
		Name:    paymentTransferCoin.CoinName,
		Address: h.PaymentTransferAccount.Address,
	})
	if err != nil {
		return wlog.WrapError(err)
	}
	if balance == nil {
		return wlog.Errorf("invalid balance")
	}
	h.PaymentTransferStartAmount, err = decimal.NewFromString(balance.BalanceStr)
	return err
}

func (h *OrderOpHandler) PrepareLedgerLockID() {
	if len(h.PaymentBalanceReqs) == 0 {
		return
	}
	h.BalanceLockID = func() *string { s := uuid.NewString(); return &s }()
}

func (h *OrderOpHandler) PreparePaymentID() {
	if h.PaymentTransferReq == nil && len(h.PaymentBalanceReqs) == 0 {
		return
	}
	h.PaymentID = func() *string { s := uuid.NewString(); return &s }()
}

func (h *OrderOpHandler) WithLockBalances(ctx context.Context) error {
	if len(h.PaymentBalanceReqs) == 0 {
		return nil
	}
	balances := []*ledgermwpb.LockBalance{}
	for _, req := range h.PaymentBalanceReqs {
		balances = append(balances, &ledgermwpb.LockBalance{
			CoinTypeID: *req.CoinTypeID,
			Amount:     *req.Amount,
		})
	}

	handler, err := ledgermw.NewHandler(
		ctx,
		ledgermw.WithAppID(h.AllocatedCouponCheckHandler.AppID, true),
		ledgermw.WithUserID(h.AllocatedCouponCheckHandler.UserID, true),
		ledgermw.WithLockID(h.BalanceLockID, true),
		ledgermw.WithBalances(balances, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	_, err = handler.LockBalances(ctx)
	return wlog.WrapError(err)
}

func (h *OrderOpHandler) WithLockPaymentTransferAccount(ctx context.Context) error {
	if h.PaymentTransferAccount == nil {
		return nil
	}

	handler, err := paymentaccountmw.NewHandler(
		ctx,
		paymentaccountmw.WithID(&h.PaymentTransferAccount.ID, true),
		paymentaccountmw.WithLockedBy(basetypes.AccountLockedBy_Payment.Enum(), false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	_, err = handler.LockAccount(ctx)
	return wlog.WrapError(err)
}

func (h *OrderOpHandler) PaymentUpdatable() error {
	switch h.OrderState {
	case types.OrderState_OrderStateCreated:
	case types.OrderState_OrderStateWaitPayment:
	default:
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *OrderOpHandler) ValidateCancelParam() error {
	if h.UserSetCanceled != nil && !*h.UserSetCanceled {
		return wlog.Errorf("permission denied")
	}
	if h.AdminSetCanceled != nil && !*h.AdminSetCanceled {
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *OrderOpHandler) UserCancelable() error {
	switch h.OrderType {
	case types.OrderType_Normal:
		switch h.OrderState {
		case types.OrderState_OrderStateWaitPayment:
			if h.AdminSetCanceled != nil {
				return wlog.Errorf("permission denied")
			}
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateInService:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.OrderType_Offline:
		fallthrough //nolint
	case types.OrderType_Airdrop:
		if h.UserSetCanceled != nil {
			return wlog.Errorf("permission denied")
		}
		switch h.OrderState {
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateInService:
		default:
			return wlog.Errorf("permission denied")
		}
	default:
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *OrderOpHandler) GoodCancelable() error {
	switch h.GoodCancelMode {
	case goodtypes.CancelMode_Uncancellable:
		return wlog.Errorf("permission denied")
	case goodtypes.CancelMode_CancellableBeforeStart:
		switch h.OrderState {
		case types.OrderState_OrderStateWaitPayment:
		case types.OrderState_OrderStatePaid:
		default:
			return wlog.Errorf("permission denied")
		}
	case goodtypes.CancelMode_CancellableBeforeBenefit:
		switch h.OrderState {
		case types.OrderState_OrderStateWaitPayment:
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateInService:
			// This should be checked by upper layer
		default:
			return wlog.Errorf("permission denied")
		}
	case goodtypes.CancelMode_CancellableBeforeUsed:
	default:
		return wlog.Errorf("invalid cancelmode")
	}
	return nil
}

func (h *OrderOpHandler) GetOrderCommissions(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	for {
		conds := &ledgerstatementmwpb.Conds{
			AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.AllocatedCouponCheckHandler.AppID},
			IOType:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ledgertypes.IOType_Incoming)},
			IOSubType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ledgertypes.IOSubType_Commission)},
			IOExtra:   &basetypes.StringVal{Op: cruder.LIKE, Value: *h.OrderID},
		}
		handler, err := ledgerstatementmw.NewHandler(
			ctx,
			ledgerstatementmw.WithConds(conds),
			ledgerstatementmw.WithOffset(offset),
			ledgerstatementmw.WithLimit(limit),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		infos, _, err := handler.GetStatements(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(infos) == 0 {
			return nil
		}
		h.CommissionLedgerStatements = append(h.CommissionLedgerStatements, infos...)
		offset += limit
	}
}

func (h *OrderOpHandler) PrepareCommissionLockIDs() {
	h.CommissionLockIDs = map[string]string{}
	for _, statement := range h.CommissionLedgerStatements {
		if _, ok := h.CommissionLockIDs[statement.UserID]; ok {
			continue
		}
		h.CommissionLockIDs[statement.UserID] = uuid.NewString()
	}
}

func (h *OrderOpHandler) WithCreateOrderCommissionLocks(ctx context.Context) error {
	multiHandler := &orderlockmw.MultiHandler{}

	for userID, commissionLockID := range h.CommissionLockIDs {
		_userID := userID
		_commissionLockID := commissionLockID

		handler, err := orderlockmw.NewHandler(
			ctx,
			orderlockmw.WithEntID(&_commissionLockID, false),
			orderlockmw.WithUserID(&_userID, true),
			orderlockmw.WithOrderID(h.OrderID, true),
			orderlockmw.WithLockType(types.OrderLockType_LockCommission.Enum(), true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		multiHandler.AppendHandler(handler)
	}

	return wlog.WrapError(multiHandler.CreateOrderLocks(ctx))
}

func (h *OrderOpHandler) WithLockCommissions(ctx context.Context) error {
	balances := map[string][]*ledgermwpb.LockBalance{}
	for _, statement := range h.CommissionLedgerStatements {
		balances[statement.UserID] = append(balances[statement.UserID], &ledgermwpb.LockBalance{
			CoinTypeID: statement.CoinTypeID,
			Amount:     statement.Amount,
		})
	}

	multiHandler := ledgermw.MultiHandler{}

	for userID, userBalances := range balances {
		_userID := userID
		_userBalances := userBalances
		commisionLockID := h.CommissionLockIDs[_userID]

		handler, err := ledgermw.NewHandler(
			ctx,
			ledgermw.WithAppID(h.AllocatedCouponCheckHandler.AppID, true),
			ledgermw.WithUserID(&_userID, true),
			ledgermw.WithLockID(&commisionLockID, true),
			ledgermw.WithBalances(_userBalances, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		multiHandler.AppendHandler(handler)
	}

	return multiHandler.LockBalances(ctx)
}
