package executor

import (
	"context"
	"fmt"
	"time"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	common "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/wait/common"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/wait/types"
	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodbenefitmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/goodbenefit"
	platformaccountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/platform"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	goodstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/statement"
	outofgasmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/outofgas"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"
	requiredappgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	goodstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/good/ledger/statement"
	outofgasmw "github.com/NpoolPlatform/kunman/middleware/order/outofgas"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	schedcommon "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	"github.com/shopspring/decimal"
)

type coinReward struct {
	types.CoinReward
	todayRewardAmount  decimal.Decimal
	userRewardAmount   decimal.Decimal
	techniqueFeeAmount decimal.Decimal
}

type goodHandler struct {
	*types.FeedPowerRental
	*common.Handler
	persistent             chan interface{}
	notif                  chan interface{}
	done                   chan interface{}
	totalUnits             decimal.Decimal
	coinBenefitBalances    map[string]decimal.Decimal
	totalBenefitOrderUnits decimal.Decimal
	appOrderUnits          map[string]map[string]decimal.Decimal
	goodCoins              map[string]*coinmwpb.Coin
	coinRewards            []*coinReward
	appPowerRentals        map[string]map[string]*apppowerrentalmwpb.PowerRental
	requiredAppFees        []*requiredappgoodmwpb.Required
	techniqueFees          map[string]map[string]*appfeemwpb.Fee
	userBenefitHotAccounts map[string]*platformaccountmwpb.Account
	goodBenefitAccounts    map[string]*goodbenefitmwpb.Account
	benefitOrderIDs        []uint32
	orderOutOfGases        map[string]*outofgasmwpb.OutOfGas
	benefitResult          basetypes.Result
	benefitMessage         string
	notifiable             bool
	benefitTimestamp       uint32
	benefitable            bool
}

const (
	resultNotMining     = "Mining not start"
	resultMinimumReward = "Mining reward not transferred"
	resultInvalidStock  = "Good stock not consistent"
)

func (h *goodHandler) checkBenefitable() bool {
	if h.ServiceStartAt >= uint32(time.Now().Unix()) {
		h.benefitResult = basetypes.Result_Success
		h.benefitMessage = fmt.Sprintf(
			"%v (start at %v, now %v)",
			resultNotMining,
			time.Unix(int64(h.ServiceStartAt), 0),
			time.Now(),
		)
		h.notifiable = true
		return false
	}
	h.benefitable = true
	return true
}

func (h *goodHandler) getGoodCoins(ctx context.Context) (err error) {
	h.goodCoins, err = schedcommon.GetCoins(ctx, func() (coinTypeIDs []string) {
		for _, goodCoin := range h.GoodCoins {
			coinTypeIDs = append(coinTypeIDs, goodCoin.CoinTypeID)
		}
		return
	}())
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, goodCoin := range h.GoodCoins {
		if _, ok := h.goodCoins[goodCoin.CoinTypeID]; !ok {
			return wlog.Errorf("invalid goodcoin")
		}
	}
	return nil
}

func (h *goodHandler) getBenefitBalances(ctx context.Context) error {
	h.coinBenefitBalances = map[string]decimal.Decimal{}
	for coinTypeID, goodBenefitAccount := range h.goodBenefitAccounts {
		coin, ok := h.goodCoins[coinTypeID]
		if !ok {
			return wlog.Errorf("invalid coin")
		}
		balance, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
			Name:    coin.Name,
			Address: goodBenefitAccount.Address,
		})
		if err != nil {
			return wlog.Errorf(
				"%v (coin %v, address %v)",
				err,
				coin.Name,
				goodBenefitAccount.Address,
			)
		}
		if balance == nil {
			return wlog.Errorf(
				"invalid balance (coin %v, address %v)",
				coin.Name,
				goodBenefitAccount.Address,
			)
		}
		bal, err := decimal.NewFromString(balance.BalanceStr)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.coinBenefitBalances[coinTypeID] = bal
	}
	return nil
}

func (h *goodHandler) orderBenefitable(order *powerrentalordermwpb.PowerRentalOrder) bool {
	if order.Simulate {
		return false
	}

	if _, ok := h.orderOutOfGases[order.OrderID]; ok {
		return false
	}

	now := uint32(time.Now().Unix())
	switch order.PaymentState {
	case ordertypes.PaymentState_PaymentStateDone:
	case ordertypes.PaymentState_PaymentStateNoPayment:
	default:
		return false
	}

	// Here we must use endat for compensate
	if order.EndAt < now {
		return false
	}
	if order.StartAt > now {
		return false
	}
	if now < order.StartAt+uint32(h.BenefitInterval().Seconds()) {
		return false
	}

	return true
}

func (h *goodHandler) getOutOfGasesWithOrderIDs(ctx context.Context, orderIDs []string) error {
	h.orderOutOfGases = map[string]*outofgasmwpb.OutOfGas{}

	conds := &outofgasmwpb.Conds{
		OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: orderIDs},
		EndAt:    &basetypes.Uint32Val{Op: cruder.EQ, Value: 0},
	}
	handler, err := outofgasmw.NewHandler(
		ctx,
		outofgasmw.WithConds(conds),
		outofgasmw.WithOffset(0),
		outofgasmw.WithLimit(int32(len(orderIDs))),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	infos, _, err := handler.GetOutOfGases(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	for _, info := range infos {
		h.orderOutOfGases[info.OrderID] = info
	}
	return nil
}

//nolint:gocognit
func (h *goodHandler) getOrderUnits(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	h.appOrderUnits = map[string]map[string]decimal.Decimal{}
	now := uint32(time.Now().Unix())

	conds := &powerrentalordermwpb.Conds{
		GoodID:       &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
		OrderState:   &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ordertypes.OrderState_OrderStateInService)},
		BenefitState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ordertypes.BenefitState_BenefitWait)},
		CreatedAt:    &basetypes.Uint32Val{Op: cruder.LT, Value: now - timedef.SecondsPerDay},
		StartAt:      &basetypes.Uint32Val{Op: cruder.LT, Value: now - timedef.SecondsPerDay},
	}

	for {
		handler, err := powerrentalordermw.NewHandler(
			ctx,
			powerrentalordermw.WithConds(conds),
			powerrentalordermw.WithOffset(offset),
			powerrentalordermw.WithLimit(limit),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		orders, _, err := handler.GetPowerRentals(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(orders) == 0 {
			break
		}
		if err := h.getOutOfGasesWithOrderIDs(ctx, func() (orderIDs []string) {
			for _, order := range orders {
				orderIDs = append(orderIDs, order.OrderID)
			}
			return
		}()); err != nil {
			return wlog.WrapError(err)
		}
		for _, order := range orders {
			units, err := decimal.NewFromString(order.Units)
			if err != nil {
				return wlog.WrapError(err)
			}
			if !h.orderBenefitable(order) {
				if order.Simulate {
					h.benefitOrderIDs = append(h.benefitOrderIDs, order.ID)
				}
				continue
			}
			h.benefitOrderIDs = append(h.benefitOrderIDs, order.ID)
			h.totalBenefitOrderUnits = h.totalBenefitOrderUnits.Add(units)
			appGoodUnits, ok := h.appOrderUnits[order.AppID]
			if !ok {
				appGoodUnits = map[string]decimal.Decimal{
					order.AppGoodID: decimal.NewFromInt(0),
				}
			}
			appGoodUnits[order.AppGoodID] = appGoodUnits[order.AppGoodID].Add(units)
			h.appOrderUnits[order.AppID] = appGoodUnits
		}
		offset += limit
	}
	return nil
}

//nolint:gocognit
func (h *goodHandler) constructCoinRewards() error {
	for _, reward := range h.Rewards {
		startRewardAmount, err := decimal.NewFromString(reward.NextRewardStartAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		benefitBalance, ok := h.coinBenefitBalances[reward.CoinTypeID]
		if !ok {
			return wlog.Errorf("Invalid benefit balance")
		}
		todayRewardAmount := benefitBalance.Sub(startRewardAmount)
		coin, ok := h.goodCoins[reward.CoinTypeID]
		if !ok {
			return wlog.Errorf("Invalid goodcoin")
		}
		reservedAmount, err := decimal.NewFromString(coin.ReservedAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		if startRewardAmount.Equal(decimal.NewFromInt(0)) {
			todayRewardAmount = todayRewardAmount.Sub(reservedAmount)
		}
		if todayRewardAmount.LessThan(decimal.NewFromInt(0)) {
			todayRewardAmount = decimal.NewFromInt(0)
		}

		nextRewardStartAmount := startRewardAmount
		if todayRewardAmount.GreaterThan(decimal.NewFromInt(0)) {
			nextRewardStartAmount = benefitBalance
		}
		userRewardAmount := todayRewardAmount.
			Mul(h.totalBenefitOrderUnits).
			Div(h.totalUnits)

		coinReward := &coinReward{
			CoinReward: types.CoinReward{
				CoinTypeID:            reward.CoinTypeID,
				Amount:                todayRewardAmount.String(),
				NextRewardStartAmount: nextRewardStartAmount.String(),
			},
			todayRewardAmount: todayRewardAmount,
			userRewardAmount:  userRewardAmount,
		}

		if reward.MainCoin {
			if err := h.calculateTechniqueFee(coinReward); err != nil {
				return wlog.WrapError(err)
			}
		}
		platformRewardAmount := todayRewardAmount.
			Sub(userRewardAmount)

		goodBenefitAccount, ok := h.goodBenefitAccounts[reward.CoinTypeID]
		if ok {
			coinReward.GoodBenefitAccountID = goodBenefitAccount.AccountID
			coinReward.GoodBenefitAddress = goodBenefitAccount.Address
		}
		userBenefitHotAccount, ok := h.userBenefitHotAccounts[reward.CoinTypeID]
		if ok {
			coinReward.UserBenefitHotAccountID = userBenefitHotAccount.AccountID
			coinReward.UserBenefitHotAddress = userBenefitHotAccount.Address
		}
		coinReward.Extra = fmt.Sprintf(
			`{"GoodID":"%v","Reward":"%v","UserReward":"%v","PlatformReward":"%v","TechniqueServiceFee":"%v"}`,
			h.GoodID,
			todayRewardAmount,
			userRewardAmount,
			platformRewardAmount,
			coinReward.techniqueFeeAmount,
		)
		if err := h.checkTransferrable(coinReward); err != nil {
			return wlog.WrapError(err)
		}
		h.coinRewards = append(h.coinRewards, coinReward)
	}
	return nil
}

func (h *goodHandler) getAppPowerRentals(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	h.appPowerRentals = map[string]map[string]*apppowerrentalmwpb.PowerRental{}

	conds := &apppowerrentalmwpb.Conds{
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
	}

	for {
		handler, err := apppowerrentalmw.NewHandler(
			ctx,
			apppowerrentalmw.WithConds(conds),
			apppowerrentalmw.WithOffset(offset),
			apppowerrentalmw.WithLimit(limit),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		goods, _, err := handler.GetPowerRentals(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(goods) == 0 {
			break
		}
		for _, good := range goods {
			appPowerRentals, ok := h.appPowerRentals[good.AppID]
			if !ok {
				appPowerRentals = map[string]*apppowerrentalmwpb.PowerRental{}
			}
			appPowerRentals[good.AppGoodID] = good
			h.appPowerRentals[good.AppID] = appPowerRentals
		}
		offset += limit
	}
	return nil
}

func (h *goodHandler) calculateTechniqueFeeLegacy(reward *coinReward) {
	for appID, appGoodUnits := range h.appOrderUnits {
		appPowerRentals, ok := h.appPowerRentals[appID]
		if !ok {
			continue
		}
		for appGoodID, units := range appGoodUnits {
			good, ok := appPowerRentals[appGoodID]
			if !ok {
				continue
			}

			feeAmount := reward.userRewardAmount.
				Mul(units).
				Div(h.totalBenefitOrderUnits).
				Mul(decimal.RequireFromString(good.TechniqueFeeRatio)).
				Div(decimal.NewFromInt(100))
			reward.techniqueFeeAmount = reward.techniqueFeeAmount.Add(feeAmount)
		}
	}
	reward.userRewardAmount = reward.userRewardAmount.Sub(reward.techniqueFeeAmount)
}

func (h *goodHandler) _calculateTechniqueFee(reward *coinReward) error {
	for appID, appGoodUnits := range h.appOrderUnits {
		// For one good, event it's assign to multiple app goods,
		// we'll use the same technique fee app good due to good only can bind to one technique fee good
		for appGoodID, units := range appGoodUnits {
			techniqueFees, ok := h.techniqueFees[appID]
			if !ok {
				continue
			}
			techniqueFee, ok := techniqueFees[appGoodID]
			if !ok {
				continue
			}
			if techniqueFee.SettlementType != goodtypes.GoodSettlementType_GoodSettledByProfitPercent {
				continue
			}
			feePercent, err := decimal.NewFromString(techniqueFee.UnitValue)
			if err != nil {
				return wlog.WrapError(err)
			}
			feeAmount := reward.userRewardAmount.
				Mul(units).
				Div(h.totalBenefitOrderUnits).
				Mul(feePercent).
				Div(decimal.NewFromInt(100))
			reward.techniqueFeeAmount = reward.techniqueFeeAmount.Add(feeAmount)
		}
	}
	reward.userRewardAmount = reward.userRewardAmount.Sub(reward.techniqueFeeAmount)
	return nil
}

func (h *goodHandler) getRequiredTechniqueFees(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &requiredappgoodmwpb.Conds{
		MainAppGoodIDs: &basetypes.StringSliceVal{
			Op: cruder.IN, Value: func() (appGoodIDs []string) {
				for _, appPowerRentals := range h.appPowerRentals {
					for _, appPowerRental := range appPowerRentals {
						appGoodIDs = append(appGoodIDs, appPowerRental.AppGoodID)
					}
				}
				return
			}(),
		},
		RequiredGoodType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(goodtypes.GoodType_TechniqueServiceFee)},
	}

	for {
		handler, err := requiredappgoodmw.NewHandler(
			ctx,
			requiredappgoodmw.WithConds(conds),
			requiredappgoodmw.WithOffset(offset),
			requiredappgoodmw.WithLimit(limit),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		requireds, _, err := handler.GetRequireds(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(requireds) == 0 {
			return nil
		}
		h.requiredAppFees = append(h.requiredAppFees, requireds...)
		offset += limit
	}
}

func (h *goodHandler) getMainAppGoodID(requiredAppGoodID string) (string, error) {
	for _, required := range h.requiredAppFees {
		if required.RequiredAppGoodID == requiredAppGoodID {
			return required.MainAppGoodID, nil
		}
	}
	return "", wlog.Errorf("invalid required")
}

func (h *goodHandler) getAppTechniqueFees(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	h.techniqueFees = map[string]map[string]*appfeemwpb.Fee{}

	conds := &appfeemwpb.Conds{
		AppGoodIDs: &basetypes.StringSliceVal{
			Op: cruder.IN, Value: func() (appGoodIDs []string) {
				for _, requiredAppFee := range h.requiredAppFees {
					appGoodIDs = append(appGoodIDs, requiredAppFee.RequiredAppGoodID)
				}
				return
			}(),
		},
		GoodType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(goodtypes.GoodType_TechniqueServiceFee)},
	}

	for {
		handler, err := appfeemw.NewHandler(
			ctx,
			appfeemw.WithConds(conds),
			appfeemw.WithOffset(offset),
			appfeemw.WithLimit(limit),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		goods, _, err := handler.GetFees(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if len(goods) == 0 {
			break
		}
		for _, good := range goods {
			techniqueFees, ok := h.techniqueFees[good.AppID]
			if !ok {
				techniqueFees = map[string]*appfeemwpb.Fee{}
			}
			if _, ok := techniqueFees[good.AppGoodID]; ok {
				return wlog.Errorf("duplicated techniquefee")
			}
			mainAppGoodID, err := h.getMainAppGoodID(good.AppGoodID)
			if err != nil {
				return wlog.WrapError(err)
			}
			techniqueFees[mainAppGoodID] = good
			h.techniqueFees[good.AppID] = techniqueFees
		}
		offset += limit
	}

	return nil
}

func (h *goodHandler) calculateTechniqueFee(reward *coinReward) error {
	if h.totalBenefitOrderUnits.Cmp(decimal.NewFromInt(0)) <= 0 {
		return nil
	}
	if reward.userRewardAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return nil
	}

	if h.GoodType == goodtypes.GoodType_LegacyPowerRental {
		h.calculateTechniqueFeeLegacy(reward)
		return nil
	}

	return h._calculateTechniqueFee(reward)
}

func (h *goodHandler) getUserBenefitHotAccounts(ctx context.Context) (err error) {
	h.userBenefitHotAccounts, err = schedcommon.GetCoinPlatformAccounts(
		ctx,
		basetypes.AccountUsedFor_UserBenefitHot,
		func() (coinTypeIDs []string) {
			for coinTypeID := range h.goodCoins {
				coinTypeIDs = append(coinTypeIDs, coinTypeID)
			}
			return
		}(),
	)
	return wlog.WrapError(err)
}

func (h *goodHandler) getGoodBenefitAccounts(ctx context.Context) (err error) {
	h.goodBenefitAccounts, err = schedcommon.GetGoodCoinBenefitAccounts(
		ctx,
		h.GoodID,
		func() (coinTypeIDs []string) {
			for coinTypeID := range h.goodCoins {
				coinTypeIDs = append(coinTypeIDs, coinTypeID)
			}
			return
		}(),
	)
	return wlog.WrapError(err)
}

func (h *goodHandler) checkTransferrable(reward *coinReward) error {
	coin, ok := h.goodCoins[reward.CoinTypeID]
	if !ok {
		return nil
	}
	least, err := decimal.NewFromString(coin.LeastTransferAmount)
	if err != nil {
		return err
	}
	if least.Cmp(decimal.NewFromInt(0)) <= 0 {
		return wlog.Errorf("invalid leasttransferamount")
	}
	if reward.todayRewardAmount.Cmp(least) <= 0 {
		reward.BenefitMessage = fmt.Sprintf(
			"%v (coin %v, reward amount %v [@%v], least transfer amount %v [#%v])",
			resultMinimumReward,
			coin.Name,
			reward.todayRewardAmount,
			reward.GoodBenefitAddress,
			least,
			h.LastRewardAt,
		)
		h.notifiable = true
		return nil
	}
	reward.Transferrable = true
	return nil
}

func (h *goodHandler) validateInServiceUnits() error {
	goodInService, err := decimal.NewFromString(h.GoodInService)
	if err != nil {
		return wlog.WrapError(err)
	}

	inService := decimal.NewFromInt(0)
	for _, appPowerRentals := range h.appPowerRentals {
		for _, appPowerRental := range appPowerRentals {
			_inService, err := decimal.NewFromString(appPowerRental.AppGoodInService)
			if err != nil {
				return wlog.WrapError(err)
			}
			inService = inService.Add(_inService)
		}
	}
	if inService.Cmp(goodInService) != 0 {
		h.benefitResult = basetypes.Result_Fail
		h.benefitMessage = fmt.Sprintf(
			"%v (good %v [%v], in service %v != %v)",
			resultInvalidStock,
			h.Name,
			h.GoodID,
			inService,
			goodInService,
		)
		h.notifiable = true
		return wlog.Errorf("invalid inservice")
	}
	return nil
}

func (h *goodHandler) resolveBenefitTimestamp() {
	h.benefitTimestamp = h.TriggerBenefitTimestamp
	if h.benefitTimestamp == 0 {
		h.benefitTimestamp = h.BenefitTimestamp()
	}
}

func (h *goodHandler) checkGoodStatement(ctx context.Context) (bool, error) {
	conds := &goodstatementmwpb.Conds{
		GoodID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
		BenefitDate: &basetypes.Uint32Val{Op: cruder.EQ, Value: h.benefitTimestamp},
	}
	handler, err := goodstatementmw.NewHandler(
		ctx,
		goodstatementmw.WithConds(conds),
	)
	if err != nil {
		return false, wlog.WrapError(err)
	}

	exist, err := handler.ExistGoodStatementConds(ctx)
	if err != nil {
		return false, wlog.WrapError(err)
	}
	if !exist {
		return false, nil
	}
	return true, nil
}

//nolint:gocritic
func (h *goodHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRental", h.PowerRental,
			"Notifiable", h.notifiable,
			"BenefitTimestamp", h.benefitTimestamp,
			"BenefitOrderIDs", len(h.benefitOrderIDs),
			"CoinRewards", h.coinRewards,
			"BenefitMessage", h.benefitMessage,
			"BenefitResult", h.benefitResult,
			"Benefitable", h.benefitable,
			"Error", *err,
		)
	}

	persistentGood := &types.PersistentPowerRental{
		PowerRental:     h.PowerRental,
		BenefitOrderIDs: h.benefitOrderIDs,
		CoinRewards: func() (rewards []*types.CoinReward) {
			for _, reward := range h.coinRewards {
				rewards = append(rewards, &reward.CoinReward)
			}
			return
		}(),
		BenefitTimestamp: h.benefitTimestamp,
		Error:            *err,
	}

	if h.notifiable {
		persistentGood.BenefitResult = h.benefitResult
		for _, reward := range persistentGood.CoinRewards {
			if reward.BenefitMessage == "" {
				reward.BenefitMessage = h.benefitMessage
			}
		}
		asyncfeed.AsyncFeed(ctx, persistentGood, h.notif)
	}
	if *err != nil {
		persistentGood.BenefitResult = basetypes.Result_Fail
		for _, reward := range persistentGood.CoinRewards {
			reward.BenefitMessage = wlog.Unwrap(*err).Error()
		}
		asyncfeed.AsyncFeed(ctx, persistentGood, h.notif)
		asyncfeed.AsyncFeed(ctx, persistentGood, h.done)
		return
	}
	if h.benefitable {
		asyncfeed.AsyncFeed(ctx, persistentGood, h.persistent)
	} else {
		asyncfeed.AsyncFeed(ctx, persistentGood, h.done)
	}
}

//nolint:gocritic
func (h *goodHandler) exec(ctx context.Context) error {
	h.benefitResult = basetypes.Result_Success

	var err error
	exist := false
	defer h.final(ctx, &err)

	h.resolveBenefitTimestamp()
	if exist, err = h.checkGoodStatement(ctx); err != nil || exist {
		return wlog.WrapError(err)
	}
	h.totalUnits, err = decimal.NewFromString(h.GoodTotal)
	if err != nil {
		return wlog.WrapError(err)
	}
	if h.totalUnits.Cmp(decimal.NewFromInt(0)) <= 0 {
		err = wlog.Errorf("invalid stock")
		return wlog.WrapError(err)
	}
	if benefitable := h.checkBenefitable(); !benefitable {
		return nil
	}
	if err = h.getGoodCoins(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.getUserBenefitHotAccounts(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.getGoodBenefitAccounts(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.getBenefitBalances(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.getOrderUnits(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.getAppPowerRentals(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.validateInServiceUnits(); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.getRequiredTechniqueFees(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.getAppTechniqueFees(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if err = h.constructCoinRewards(); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}
