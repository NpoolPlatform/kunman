package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/bookkeeping/good/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/fee"
	requiredappgoodmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/required"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	powerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/powerrental"
	goodstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/statement"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"
	requiredappgoodmw "github.com/NpoolPlatform/kunman/middleware/good/app/good/required"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	goodstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/good/ledger/statement"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type coinReward struct {
	types.CoinReward
	totalRewardAmount       decimal.Decimal
	unsoldRewardAmount      decimal.Decimal
	userRewardAmount        decimal.Decimal
	totalTechniqueFeeAmount decimal.Decimal
}

type goodHandler struct {
	*powerrentalmwpb.PowerRental
	persistent             chan interface{}
	notif                  chan interface{}
	done                   chan interface{}
	totalOrderUnits        decimal.Decimal
	appOrderUnits          map[string]map[string]decimal.Decimal
	coinRewards            []*coinReward
	appPowerRentals        map[string]map[string]*apppowerrentalmwpb.PowerRental
	requiredAppFees        []*requiredappgoodmwpb.Required
	techniqueFees          map[string]map[string]*appfeemwpb.Fee
	totalBenefitOrderUnits decimal.Decimal
}

func (h *goodHandler) getOrderUnits(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	h.appOrderUnits = map[string]map[string]decimal.Decimal{}

	conds := &powerrentalordermwpb.Conds{
		GoodID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
		LastBenefitAt: &basetypes.Uint32Val{Op: cruder.EQ, Value: h.LastRewardAt},
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
			h.totalBenefitOrderUnits = h.totalBenefitOrderUnits.Add(units)
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

func (h *goodHandler) calculateTechniqueFeeLegacy(reward *coinReward) {
	for appID, appGoodUnits := range h.appOrderUnits {
		goods, ok := h.appPowerRentals[appID]
		if !ok {
			continue
		}
		for appGoodID, units := range appGoodUnits {
			good, ok := goods[appGoodID]
			if !ok {
				continue
			}
			userRewardAmount := reward.userRewardAmount.
				Mul(units).
				Div(h.totalOrderUnits)
			techniqueFee := userRewardAmount.
				Mul(decimal.RequireFromString(good.TechniqueFeeRatio)).
				Div(decimal.NewFromInt(100))
			reward.totalTechniqueFeeAmount = reward.totalTechniqueFeeAmount.
				Add(techniqueFee)
		}
	}
}

func (h *goodHandler) _calculateTechniqueFee(reward *coinReward) error {
	logger.Sugar().Infow(
		"_calculateTechniqueFee",
		"AppOrderUnits", h.appOrderUnits,
		"TechniqueFees", h.techniqueFees,
	)
	for appID, appGoodUnits := range h.appOrderUnits {
		// For one good, event it's assign to multiple app goods,
		// we'll use the same technique fee app good due to good only can bind to one technique fee good
		techniqueFees, ok := h.techniqueFees[appID]
		if !ok {
			continue
		}
		for appGoodID, units := range appGoodUnits {
			techniqueFee, ok := techniqueFees[appGoodID]
			if !ok {
				continue
			}
			if techniqueFee.SettlementType != goodtypes.GoodSettlementType_GoodSettledByProfitPercent {
				continue
			}
			feePercent, err := decimal.NewFromString(techniqueFee.UnitValue)
			if err != nil {
				return err
			}
			feeAmount := reward.userRewardAmount.
				Mul(units).
				Div(h.totalBenefitOrderUnits).
				Mul(feePercent).
				Div(decimal.NewFromInt(100))
			reward.totalTechniqueFeeAmount = reward.totalTechniqueFeeAmount.Add(feeAmount)
		}
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

func (h *goodHandler) constructCoinRewards() error {
	totalUnits, err := decimal.NewFromString(h.GoodTotal)
	if err != nil {
		return err
	}
	for _, reward := range h.Rewards {
		totalRewardAmount, err := decimal.NewFromString(reward.LastRewardAmount)
		if err != nil {
			return err
		}
		userRewardAmount := totalRewardAmount.
			Mul(h.totalOrderUnits).
			Div(totalUnits)
		unsoldRewardAmount := totalRewardAmount.
			Sub(userRewardAmount)
		coinReward := &coinReward{
			CoinReward: types.CoinReward{
				CoinTypeID:         reward.CoinTypeID,
				TotalRewardAmount:  totalRewardAmount.String(),
				UnsoldRewardAmount: unsoldRewardAmount.String(),
				TechniqueFeeAmount: decimal.NewFromInt(0).String(),
			},
			totalRewardAmount:       totalRewardAmount,
			unsoldRewardAmount:      unsoldRewardAmount,
			userRewardAmount:        userRewardAmount,
			totalTechniqueFeeAmount: decimal.NewFromInt(0),
		}
		if reward.MainCoin {
			if err := h.calculateTechniqueFee(coinReward); err != nil {
				return err
			}
			coinReward.TechniqueFeeAmount = coinReward.totalTechniqueFeeAmount.String()
		}
		h.coinRewards = append(h.coinRewards, coinReward)
	}
	return nil
}

func (h *goodHandler) checkGoodStatement(ctx context.Context) (bool, error) {
	conds := &goodstatementmwpb.Conds{
		GoodID:      &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
		BenefitDate: &basetypes.Uint32Val{Op: cruder.EQ, Value: h.LastRewardAt},
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
		return false, err
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
			"AppOrderUnits", h.appOrderUnits,
			"CoinRewards", h.coinRewards,
			"Error", *err,
		)
	}
	persistentGood := &types.PersistentGood{
		PowerRental: h.PowerRental,
		CoinRewards: func() (rewards []*types.CoinReward) {
			for _, reward := range h.coinRewards {
				rewards = append(rewards, &reward.CoinReward)
			}
			return
		}(),
		Error: *err,
	}

	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentGood, h.persistent)
		return
	}

	persistentGood.BenefitResult = basetypes.Result_Fail
	for _, reward := range persistentGood.CoinRewards {
		reward.BenefitMessage = wlog.Unwrap(*err).Error()
	}

	asyncfeed.AsyncFeed(ctx, persistentGood, h.notif)
	asyncfeed.AsyncFeed(ctx, persistentGood, h.done)
}

//nolint:gocritic
func (h *goodHandler) exec(ctx context.Context) error {
	var err error
	var exist bool

	defer h.final(ctx, &err)

	if exist, err = h.checkGoodStatement(ctx); err != nil || exist {
		return err
	}
	if err = h.getAppPowerRentals(ctx); err != nil {
		return err
	}
	if err = h.getOrderUnits(ctx); err != nil {
		return err
	}
	if err = h.getRequiredTechniqueFees(ctx); err != nil {
		return err
	}
	if err = h.getAppTechniqueFees(ctx); err != nil {
		return err
	}
	if err = h.constructCoinRewards(); err != nil {
		return err
	}

	return nil
}
