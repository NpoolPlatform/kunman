package executor

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/user/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"
	requiredappgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	statementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type coinReward struct {
	mainCoin          bool
	totalRewardAmount decimal.Decimal
	userRewardAmount  decimal.Decimal
}

type goodHandler struct {
	*powerrentalmwpb.PowerRental
	persistent         chan interface{}
	notif              chan interface{}
	done               chan interface{}
	totalOrderUnits    decimal.Decimal
	appOrderUnits      map[string]map[string]decimal.Decimal
	appPowerRentals    map[string]map[string]*apppowerrentalmwpb.PowerRental
	requiredAppFees    []*requiredappgoodmwpb.Required
	techniqueFees      map[string]map[string]*appfeemwpb.Fee
	appGoodUnitRewards map[string]map[string]map[string]decimal.Decimal
	orderRewards       []*types.OrderReward
	coinRewards        map[string]*coinReward
}

//nolint:dupl
func (h *goodHandler) getOrderUnits(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	h.appOrderUnits = map[string]map[string]decimal.Decimal{}

	conds := &powerrentalordermwpb.Conds{
		GoodID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
		LastBenefitAt: &basetypes.Uint32Val{Op: cruder.EQ, Value: h.LastRewardAt},
		BenefitState:  &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ordertypes.BenefitState_BenefitCalculated)},
		Simulate:      &basetypes.BoolVal{Op: cruder.EQ, Value: false},
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
			return err
		}
		if len(orders) == 0 {
			break
		}
		for _, order := range orders {
			units, err := decimal.NewFromString(order.Units)
			if err != nil {
				return err
			}
			h.totalOrderUnits = h.totalOrderUnits.Add(units)
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
			return err
		}
		if len(goods) == 0 {
			break
		}
		for _, good := range goods {
			_goods, ok := h.appPowerRentals[good.AppID]
			if !ok {
				_goods = map[string]*apppowerrentalmwpb.PowerRental{}
			}
			_goods[good.AppGoodID] = good
			h.appPowerRentals[good.AppID] = _goods
		}
		offset += limit
	}
	return nil
}

func (h *goodHandler) calculateUnitRewardsLegacy() error { //nolint:gocognit
	for appID, appGoodUnits := range h.appOrderUnits {
		goods, ok := h.appPowerRentals[appID]
		if !ok {
			continue
		}
		appUnitRewards, ok := h.appGoodUnitRewards[appID]
		if !ok {
			appUnitRewards = map[string]map[string]decimal.Decimal{}
		}
		for appGoodID, units := range appGoodUnits {
			good, ok := goods[appGoodID]
			if !ok {
				continue
			}
			appGoodUnitRewards, ok := appUnitRewards[appGoodID]
			if !ok {
				appGoodUnitRewards = map[string]decimal.Decimal{}
			}
			for coinTypeID, reward := range h.coinRewards {
				if reward.userRewardAmount.Cmp(decimal.NewFromInt(0)) < 0 {
					return wlog.Errorf("invalid userrewardamount")
				}
				userRewardAmount := reward.userRewardAmount.
					Mul(units).
					Div(h.totalOrderUnits)
				techniqueFee := decimal.NewFromInt(0)
				if reward.mainCoin {
					techniqueFee = userRewardAmount.
						Mul(decimal.RequireFromString(good.TechniqueFeeRatio)).
						Div(decimal.NewFromInt(100))
				}
				appGoodUnitRewards[coinTypeID] = userRewardAmount.
					Sub(techniqueFee).
					Div(units)
			}
			appUnitRewards[appGoodID] = appGoodUnitRewards
		}
		h.appGoodUnitRewards[appID] = appUnitRewards
	}
	return nil
}

//nolint:gocognit
func (h *goodHandler) _calculateUnitRewards() error {
	for appID, appGoodUnits := range h.appOrderUnits {
		appUnitRewards, ok := h.appGoodUnitRewards[appID]
		if !ok {
			appUnitRewards = map[string]map[string]decimal.Decimal{}
		}
		for appGoodID, units := range appGoodUnits {
			var techniqueFee *appfeemwpb.Fee
			techniqueFees, ok := h.techniqueFees[appID]
			if ok {
				techniqueFee, ok = techniqueFees[appGoodID]
			}

			feePercent := decimal.NewFromInt(0)
			var err error

			if ok && techniqueFee.SettlementType == goodtypes.GoodSettlementType_GoodSettledByProfitPercent {
				feePercent, err = decimal.NewFromString(techniqueFee.UnitValue)
				if err != nil {
					return err
				}
			}

			appGoodUnitRewards, ok := appUnitRewards[appGoodID]
			if !ok {
				appGoodUnitRewards = map[string]decimal.Decimal{}
			}
			for coinTypeID, reward := range h.coinRewards {
				if reward.userRewardAmount.Cmp(decimal.NewFromInt(0)) < 0 {
					return wlog.Errorf("invalid userrewardamount")
				}
				userRewardAmount := reward.userRewardAmount.
					Mul(units).
					Div(h.totalOrderUnits)
				techniqueFee := decimal.NewFromInt(0)
				if reward.mainCoin {
					techniqueFee = userRewardAmount.
						Mul(feePercent).
						Div(decimal.NewFromInt(100))
				}
				appGoodUnitRewards[coinTypeID] = userRewardAmount.
					Sub(techniqueFee).
					Div(units)
			}
			appUnitRewards[appGoodID] = appGoodUnitRewards
		}
		h.appGoodUnitRewards[appID] = appUnitRewards
	}
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
			return err
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

func (h *goodHandler) constructCoinRewards() error {
	totalUnits, err := decimal.NewFromString(h.GoodTotal)
	if err != nil {
		return err
	}
	h.coinRewards = map[string]*coinReward{}
	for _, reward := range h.Rewards {
		totalRewardAmount, err := decimal.NewFromString(reward.LastRewardAmount)
		if err != nil {
			return err
		}
		userRewardAmount := totalRewardAmount.
			Mul(h.totalOrderUnits).
			Div(totalUnits)
		_coinReward := &coinReward{
			totalRewardAmount: totalRewardAmount,
			userRewardAmount:  userRewardAmount,
		}
		for _, goodCoin := range h.GoodCoins {
			if goodCoin.CoinTypeID == reward.CoinTypeID && goodCoin.Main {
				_coinReward.mainCoin = true
				break
			}
		}
		h.coinRewards[reward.CoinTypeID] = _coinReward
	}
	return nil
}

func (h *goodHandler) calculateUnitRewards() error {
	h.appGoodUnitRewards = map[string]map[string]map[string]decimal.Decimal{}
	if h.totalOrderUnits.Cmp(decimal.NewFromInt(0)) <= 0 {
		return nil
	}
	if h.GoodType == goodtypes.GoodType_LegacyPowerRental {
		return h.calculateUnitRewardsLegacy()
	}
	return h._calculateUnitRewards()
}

func (h *goodHandler) checkBenefitStatement(ctx context.Context, reward *types.OrderReward) (bool, error) {
	ioType := ledgertypes.IOType_Incoming
	ioSubType := ledgertypes.IOSubType_MiningBenefit

	conds := &statementmwpb.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: reward.AppID},
		UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: reward.UserID},
		IOType:    &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ioType)},
		IOSubType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ioSubType)},
	}
	handler, err := statementmw.NewHandler(
		ctx,
		statementmw.WithConds(conds),
	)
	if err != nil {
		return false, wlog.WrapError(err)
	}

	exist, err := handler.ExistStatementConds(ctx)
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *goodHandler) calculateOrderReward(ctx context.Context, order *powerrentalordermwpb.PowerRentalOrder) error {
	appUnitRewards, ok := h.appGoodUnitRewards[order.AppID]
	if !ok {
		return nil
	}
	appGoodUnitRewards, ok := appUnitRewards[order.AppGoodID]
	if !ok {
		return nil
	}
	ioExtra := fmt.Sprintf(
		`{"GoodID":"%v","AppGoodID":"%v","OrderID":"%v","Units":"%v","BenefitDate":"%v"}`,
		h.GoodID,
		order.AppGoodID,
		order.OrderID,
		order.Units,
		h.LastRewardAt,
	)
	units, err := decimal.NewFromString(order.Units)
	if err != nil {
		return err
	}
	orderReward := &types.OrderReward{
		AppID:   order.AppID,
		UserID:  order.UserID,
		OrderID: order.OrderID,
		Extra:   ioExtra,
	}
	for coinTypeID := range h.coinRewards {
		unitReward, ok := appGoodUnitRewards[coinTypeID]
		if !ok {
			continue
		}
		amount := unitReward.Mul(units)
		if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
			continue
		}
		orderReward.CoinRewards = append(orderReward.CoinRewards, &types.CoinReward{
			CoinTypeID: coinTypeID,
			Amount:     amount.String(),
		})
	}
	exist, err := h.checkBenefitStatement(ctx, orderReward)
	if err != nil {
		return err
	}
	orderReward.FirstBenefit = !exist
	h.orderRewards = append(h.orderRewards, orderReward)
	return nil
}

func (h *goodHandler) calculateOrderRewards(ctx context.Context) error {
	// If orderRewards is not empty, we do not update good benefit state, then we get next 20 orders
	conds := &powerrentalordermwpb.Conds{
		GoodID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
		LastBenefitAt: &basetypes.Uint32Val{Op: cruder.EQ, Value: h.LastRewardAt},
		BenefitState:  &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ordertypes.BenefitState_BenefitCalculated)},
		Simulate:      &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}
	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithConds(conds),
		powerrentalordermw.WithOffset(0),
		powerrentalordermw.WithLimit(20),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	orders, _, err := handler.GetPowerRentals(ctx)
	if err != nil {
		return err
	}
	if len(orders) == 0 {
		return nil
	}

	for _, order := range orders {
		if err := h.calculateOrderReward(ctx, order); err != nil {
			return err
		}
	}
	return nil
}

//nolint:gocritic
func (h *goodHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRental", h.PowerRental,
			"OrderRewards", h.orderRewards,
			"AppOrderUnits", h.appOrderUnits,
			"LastRewardAt", h.LastRewardAt,
			"Error", *err,
		)
	}
	persistentGood := &types.PersistentGood{
		PowerRental:  h.PowerRental,
		OrderRewards: h.orderRewards,
		Error:        *err,
	}

	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentGood, h.persistent)
		return
	}

	persistentGood.BenefitResult = basetypes.Result_Fail

	asyncfeed.AsyncFeed(ctx, persistentGood, h.notif)
	asyncfeed.AsyncFeed(ctx, persistentGood, h.done)
}

//nolint:gocritic
func (h *goodHandler) exec(ctx context.Context) error {
	var err error

	defer h.final(ctx, &err)

	if err = h.getAppPowerRentals(ctx); err != nil {
		return err
	}
	if err := h.getRequiredTechniqueFees(ctx); err != nil {
		return err
	}
	if err := h.getAppTechniqueFees(ctx); err != nil {
		return err
	}
	if err = h.getOrderUnits(ctx); err != nil {
		return err
	}
	if err = h.constructCoinRewards(); err != nil {
		return err
	}
	if err := h.calculateUnitRewards(); err != nil {
		return err
	}
	if err = h.calculateOrderRewards(ctx); err != nil {
		return err
	}

	return nil
}
