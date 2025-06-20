package profit

import (
	"context"
	"encoding/json"

	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
	npool "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/profit"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	ordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	statementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	ordermw "github.com/NpoolPlatform/kunman/middleware/order/order"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type rewardHandler struct {
	*Handler
	statements        []*statementmwpb.Statement
	appCoins          map[string]*appcoinmwpb.Coin
	orders            map[string]*ordermwpb.Order
	infos             []*npool.MiningReward
	powerRentalOrders map[string]*powerrentalordermwpb.PowerRentalOrder
	appPowerRentals   map[string]*apppowerrentalmwpb.PowerRental
	total             uint32
}

func (h *rewardHandler) getAppCoins(ctx context.Context) error {
	ids := []string{}
	for _, val := range h.statements {
		ids = append(ids, val.CurrencyID)
	}

	conds := &appcoinmwpb.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: ids},
	}
	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithConds(conds),
		appcoinmw.WithOffset(0),
		appcoinmw.WithLimit(int32(len(ids))),
	)
	if err != nil {
		return err
	}

	coins, _, err := handler.GetCoins(ctx)
	if err != nil {
		return err
	}
	for _, coin := range coins {
		h.appCoins[coin.CoinTypeID] = coin
	}
	return nil
}

func (h *rewardHandler) getAppPowerRentals(ctx context.Context) error {
	conds := &apppowerrentalmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: func() (appGoodIDs []string) {
			for _, statement := range h.statements {
				e := struct {
					OrderID   string
					AppGoodID string
				}{}
				if err := json.Unmarshal([]byte(statement.IOExtra), &e); err != nil {
					continue
				}
				appGoodIDs = append(appGoodIDs, e.AppGoodID)
			}
			return
		}()},
	}
	handler, err := apppowerrentalmw.NewHandler(
		ctx,
		apppowerrentalmw.WithConds(conds),
		apppowerrentalmw.WithOffset(0),
		apppowerrentalmw.WithLimit(int32(len(h.statements))),
	)
	if err != nil {
		return err
	}

	appPowerRentals, _, err := handler.GetPowerRentals(ctx)
	if err != nil {
		return err
	}
	h.appPowerRentals = map[string]*apppowerrentalmwpb.PowerRental{}
	for _, appPowerRental := range appPowerRentals {
		h.appPowerRentals[appPowerRental.AppGoodID] = appPowerRental
	}
	return nil
}

// TODO: here we should get orders which is in statement extra
//nolint
func (h *rewardHandler) getOrders(ctx context.Context) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &ordermwpb.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
		OrderStates: &basetypes.Uint32SliceVal{
			Op: cruder.NIN, Value: []uint32{
				uint32(ordertypes.OrderState_OrderStateCreated),
				uint32(ordertypes.OrderState_OrderStateWaitPayment),
				uint32(ordertypes.OrderState_OrderStatePaymentTimeout),
				uint32(ordertypes.OrderState_OrderStatePreCancel),
				uint32(ordertypes.OrderState_OrderStateRestoreCanceledStock),
				uint32(ordertypes.OrderState_OrderStateCancelAchievement),
				uint32(ordertypes.OrderState_OrderStateDeductLockedCommission),
				uint32(ordertypes.OrderState_OrderStateReturnCanceledBalance),
				uint32(ordertypes.OrderState_OrderStateCanceledTransferBookKeeping),
				uint32(ordertypes.OrderState_OrderStateCancelUnlockPaymentAccount),
				uint32(ordertypes.OrderState_OrderStateCanceled),
			},
		},
	}

	for {
		handler, err := ordermw.NewHandler(
			ctx,
			ordermw.WithConds(conds),
			ordermw.WithOffset(offset),
			ordermw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		orders, _, err := handler.GetOrders(ctx)
		if err != nil {
			return err
		}
		if len(orders) == 0 {
			break
		}

		for _, order := range orders {
			h.orders[order.EntID] = order
		}
		offset += limit
	}
	return nil
}

func (h *rewardHandler) getPowerRentalOrders(ctx context.Context) error {
	orderIDs := func() (uids []string) {
		for orderID, order := range h.orders {
			switch order.GoodType {
			case goodtypes.GoodType_PowerRental:
			case goodtypes.GoodType_LegacyPowerRental:
			default:
				continue
			}
			uids = append(uids, orderID)
		}
		return
	}()

	conds := &powerrentalordermwpb.Conds{
		AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID:   &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
		OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: orderIDs},
	}
	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithConds(conds),
		powerrentalordermw.WithOffset(0),
		powerrentalordermw.WithLimit(int32(len(orderIDs))),
	)
	if err != nil {
		return err
	}

	powerRentalOrders, _, err := handler.GetPowerRentals(ctx)
	if err != nil {
		return err
	}
	h.powerRentalOrders = map[string]*powerrentalordermwpb.PowerRentalOrder{}
	for _, powerRentalOrder := range powerRentalOrders {
		h.powerRentalOrders[powerRentalOrder.OrderID] = powerRentalOrder
	}
	return nil
}

func (h *rewardHandler) getStatements(ctx context.Context) error {
	conds := &statementmwpb.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
		IOType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(types.IOType_Incoming)},
	}
	if h.StartAt != nil {
		conds.StartAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.StartAt}
	}
	if h.EndAt != nil {
		conds.EndAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.EndAt}
	}
	if h.SimulateOnly != nil && *h.SimulateOnly {
		conds.IOSubType = &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(types.IOSubType_SimulateMiningBenefit)}
	} else {
		conds.IOSubTypes = &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{
			uint32(types.IOSubType_MiningBenefit),
			uint32(types.IOSubType_SimulateMiningBenefit),
		}}
	}

	handler, err := statementmw.NewHandler(
		ctx,
		statementmw.WithConds(conds),
		statementmw.WithOffset(h.Offset),
		statementmw.WithLimit(h.Limit),
	)
	if err != nil {
		return err
	}

	statements, total, err := handler.GetStatements(ctx)
	if err != nil {
		return err
	}
	h.statements = statements
	h.total = total
	return nil
}

func (h *rewardHandler) formalize() {
	for _, statement := range h.statements {
		coin, ok := h.appCoins[statement.CurrencyID]
		if !ok {
			continue
		}
		e := struct {
			OrderID   string
			AppGoodID string
		}{}
		if err := json.Unmarshal([]byte(statement.IOExtra), &e); err != nil {
			continue
		}
		powerRentalOrder, ok := h.powerRentalOrders[e.OrderID]
		if !ok {
			continue
		}
		if powerRentalOrder.AppGoodID != e.AppGoodID {
			continue
		}

		rewardAmount, err := decimal.NewFromString(statement.Amount)
		if err != nil {
			break
		}
		units, err := decimal.NewFromString(powerRentalOrder.Units)
		if err != nil {
			continue
		}

		appPowerRental, ok := h.appPowerRentals[e.AppGoodID]
		if !ok {
			continue
		}

		h.infos = append(h.infos, &npool.MiningReward{
			ID:                  statement.ID,
			EntID:               statement.EntID,
			AppID:               statement.AppID,
			AppGoodName:         appPowerRental.AppGoodName,
			UserID:              statement.UserID,
			CoinTypeID:          statement.CurrencyID,
			CoinName:            coin.Name,
			CoinLogo:            coin.Logo,
			CoinUnit:            coin.Unit,
			IOType:              statement.IOType,
			IOSubType:           statement.IOSubType,
			RewardAmount:        statement.Amount,
			RewardAmountPerUnit: rewardAmount.Div(units).String(),
			Units:               powerRentalOrder.Units,
			Extra:               statement.IOExtra,
			AppGoodID:           e.AppGoodID,
			OrderID:             e.OrderID,
			CreatedAt:           statement.CreatedAt,
		})
	}
}

func (h *Handler) GetMiningRewards(ctx context.Context) ([]*npool.MiningReward, uint32, error) {
	handler := &rewardHandler{
		Handler:           h,
		appCoins:          map[string]*appcoinmwpb.Coin{},
		orders:            map[string]*ordermwpb.Order{},
		statements:        []*statementmwpb.Statement{},
		powerRentalOrders: map[string]*powerrentalordermwpb.PowerRentalOrder{},
	}
	if err := h.CheckStartEndAt(); err != nil {
		return nil, 0, err
	}
	if err := handler.getStatements(ctx); err != nil {
		return nil, 0, err
	}
	if len(handler.statements) == 0 {
		return nil, handler.total, nil
	}
	if err := handler.getOrders(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getPowerRentalOrders(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getAppCoins(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getAppPowerRentals(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
