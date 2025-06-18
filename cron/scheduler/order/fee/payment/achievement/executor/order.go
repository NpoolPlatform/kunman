package executor

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	goodcoinmwcli "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	orderstatementmwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/achievement/statement/order"
	calculatemwcli "github.com/NpoolPlatform/inspire-middleware/pkg/client/calculate"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	inspiretypes "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	orderstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	calculatemwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/calculate"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/payment/achievement/types"

	"github.com/shopspring/decimal"
)

type orderHandler struct {
	*feeordermwpb.FeeOrder
	persistent      chan interface{}
	notif           chan interface{}
	done            chan interface{}
	orderStatements []*orderstatementmwpb.StatementReq
	goodMainCoin    *goodcoinmwpb.GoodCoin
}

func (h *orderHandler) checkAchievementExist(ctx context.Context) (bool, error) {
	return orderstatementmwcli.ExistStatementConds(ctx, &orderstatementmwpb.Conds{
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: h.OrderID},
	})
}

func (h *orderHandler) getGoodMainCoin(ctx context.Context) (err error) {
	h.goodMainCoin, err = goodcoinmwcli.GetGoodCoinOnly(ctx, &goodcoinmwpb.Conds{
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: h.ParentGoodID},
		Main:   &basetypes.BoolVal{Op: cruder.EQ, Value: true},
	})
	return wlog.WrapError(err)
}

func (h *orderHandler) calculateOrderStatements(ctx context.Context) (err error) {
	hasCommission := false
	switch h.OrderType {
	case ordertypes.OrderType_Normal:
		hasCommission = true
	case ordertypes.OrderType_Offline:
	case ordertypes.OrderType_Airdrop:
		return nil
	}
	h.orderStatements, err = calculatemwcli.Calculate(ctx, &calculatemwpb.CalculateRequest{
		AppID:            h.AppID,
		UserID:           h.UserID,
		GoodID:           h.GoodID,
		AppGoodID:        h.AppGoodID,
		OrderID:          h.OrderID,
		GoodCoinTypeID:   h.goodMainCoin.CoinTypeID,
		Units:            decimal.NewFromInt(0).String(),
		PaymentAmountUSD: h.PaymentAmountUSD,
		GoodValueUSD:     h.PaymentGoodValueUSD,
		SettleType:       inspiretypes.SettleType_GoodOrderPayment,
		HasCommission:    hasCommission,
		OrderCreatedAt:   h.CreatedAt,
		Payments: func() (payments []*calculatemwpb.Payment) {
			for _, _payment := range h.PaymentBalances {
				payments = append(payments, &calculatemwpb.Payment{
					CoinTypeID: _payment.CoinTypeID,
					Amount:     _payment.Amount,
				})
			}
			for _, _payment := range h.PaymentTransfers {
				payments = append(payments, &calculatemwpb.Payment{
					CoinTypeID: _payment.CoinTypeID,
					Amount:     _payment.Amount,
				})
			}
			return
		}(),
	})
	return wlog.WrapError(err)
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"FeeOrder", h.FeeOrder,
			"OrderStatements", h.orderStatements,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		FeeOrder:        h.FeeOrder,
		OrderStatements: h.orderStatements,
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentOrder, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.notif)
	asyncfeed.AsyncFeed(ctx, persistentOrder, h.done)
}

//nolint:gocritic
func (h *orderHandler) exec(ctx context.Context) error {
	var err error
	var exist bool

	defer h.final(ctx, &err)

	if exist, err = h.checkAchievementExist(ctx); err != nil || exist {
		return err
	}
	if err = h.getGoodMainCoin(ctx); err != nil {
		return err
	}
	if err = h.calculateOrderStatements(ctx); err != nil {
		return err
	}

	return nil
}
