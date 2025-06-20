package profit

import (
	"context"

	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	npool "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/profit"
	profitmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/profit"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	profitmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/profit"
	statementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type coinProfitHandler struct {
	*Handler
	infos       []*npool.CoinProfit
	appCoins    []*appcoinmwpb.Coin
	profits     map[string]*profitmwpb.Profit
	coinTypeIDs []string
	total       uint32
}

//nolint:dupl
func (h *coinProfitHandler) getStatements(ctx context.Context) error {
	if h.StartAt == nil && h.EndAt == nil {
		return nil
	}
	offset := int32(0)
	limit := constant.DefaultRowLimit
	conds := &statementmwpb.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
		IOType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ledgertypes.IOType_Incoming)},
		IOSubTypes: &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{
			uint32(ledgertypes.IOSubType_MiningBenefit),
			uint32(ledgertypes.IOSubType_SimulateMiningBenefit),
		}},
		CurrencyIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: h.coinTypeIDs},
	}
	if h.StartAt != nil {
		conds.StartAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.StartAt}
	}
	if h.EndAt != nil {
		conds.EndAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.EndAt}
	}

	for {
		handler, err := statementmw.NewHandler(
			ctx,
			statementmw.WithConds(conds),
			statementmw.WithOffset(offset),
			statementmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		statements, _, err := handler.GetStatements(ctx)
		if err != nil {
			return err
		}
		if len(statements) == 0 {
			break
		}
		for _, statement := range statements {
			profit, ok := h.profits[statement.CurrencyID]
			if ok {
				incoming, err := decimal.NewFromString(profit.Incoming)
				if err != nil {
					return err
				}
				amount, err := decimal.NewFromString(statement.Amount)
				if err != nil {
					return err
				}
				profit.Incoming = incoming.Add(amount).String()
				h.profits[statement.CurrencyID] = profit
				continue
			}

			newProfit := &profitmwpb.Profit{
				CoinTypeID: statement.CurrencyID,
				Incoming:   statement.Amount,
			}
			h.profits[statement.CurrencyID] = newProfit
		}
		offset += limit
	}
	return nil
}

func (h *coinProfitHandler) getAppCoins(ctx context.Context) error {
	conds := &appcoinmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}
	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithConds(conds),
		appcoinmw.WithOffset(h.Offset),
		appcoinmw.WithLimit(h.Limit),
	)
	if err != nil {
		return err
	}

	coins, total, err := handler.GetCoins(ctx)
	if err != nil {
		return err
	}

	for _, coin := range coins {
		h.coinTypeIDs = append(h.coinTypeIDs, coin.CoinTypeID)
	}

	h.total = total
	h.appCoins = coins
	return nil
}

func (h *coinProfitHandler) getProfits(ctx context.Context) error {
	if h.StartAt != nil || h.EndAt != nil {
		return nil
	}
	coinTypeIDs := func() (_coinTypeIDs []string) {
		for _, appCoin := range h.appCoins {
			_coinTypeIDs = append(_coinTypeIDs, appCoin.CoinTypeID)
		}
		return
	}()
	conds := &profitmwpb.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID:      &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}

	handler, err := profitmw.NewHandler(
		ctx,
		profitmw.WithConds(conds),
		profitmw.WithOffset(0),
		profitmw.WithLimit(int32(len(coinTypeIDs))),
	)
	if err != nil {
		return err
	}

	profits, _, err := handler.GetProfits(ctx)
	if err != nil {
		return err
	}
	for _, profit := range profits {
		h.profits[profit.CoinTypeID] = profit
	}
	return nil
}

func (h *coinProfitHandler) formalize() {
	for _, coin := range h.appCoins {
		h.infos = append(h.infos, &npool.CoinProfit{
			AppID:        coin.AppID,
			UserID:       *h.UserID,
			CoinTypeID:   coin.CoinTypeID,
			CoinName:     coin.Name,
			DisplayNames: coin.DisplayNames,
			CoinLogo:     coin.Logo,
			CoinUnit:     coin.Unit,
			Incoming: func() string {
				profit, ok := h.profits[coin.CoinTypeID]
				if !ok {
					return decimal.NewFromInt(0).String()
				}
				return profit.Incoming
			}(),
		})
	}
}

func (h *Handler) GetCoinProfits(ctx context.Context) ([]*npool.CoinProfit, uint32, error) {
	handler := &coinProfitHandler{
		Handler:  h,
		appCoins: []*appcoinmwpb.Coin{},
		profits:  map[string]*profitmwpb.Profit{},
	}
	if err := h.CheckStartEndAt(); err != nil {
		return nil, 0, err
	}
	if err := handler.getAppCoins(ctx); err != nil {
		return nil, 0, err
	}
	if len(handler.appCoins) == 0 {
		return nil, 0, nil
	}
	if err := handler.getProfits(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getStatements(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
