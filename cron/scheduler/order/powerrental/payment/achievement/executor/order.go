package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/payment/achievement/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	inspiretypes "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	orderstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	calculatemwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/calculate"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	goodcoinmw "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	orderstatementmw "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/statement/order"
	calculatemw "github.com/NpoolPlatform/kunman/middleware/inspire/calculate"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type orderHandler struct {
	*powerrentalordermwpb.PowerRentalOrder
	persistent      chan interface{}
	notif           chan interface{}
	done            chan interface{}
	orderStatements []*orderstatementmwpb.StatementReq
	goodMainCoin    *goodcoinmwpb.GoodCoin
}

func (h *orderHandler) checkAchievementExist(ctx context.Context) (bool, error) {
	conds := &orderstatementmwpb.Conds{
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: h.OrderID},
	}
	handler, err := orderstatementmw.NewHandler(
		ctx,
		orderstatementmw.WithConds(conds),
	)
	if err != nil {
		return false, err
	}

	return handler.ExistStatementConds(ctx)
}

func (h *orderHandler) getGoodMainCoin(ctx context.Context) (err error) {
	conds := &goodcoinmwpb.Conds{
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: h.GoodID},
		Main:   &basetypes.BoolVal{Op: cruder.EQ, Value: true},
	}
	handler, err := goodcoinmw.NewHandler(
		ctx,
		goodcoinmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	h.goodMainCoin, err = handler.GetGoodCoinOnly(ctx)
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

	handler, err := calculatemw.NewHandler(
		ctx,
		calculatemw.WithAppID(h.AppID),
		calculatemw.WithUserID(h.UserID),
		calculatemw.WithGoodID(h.GoodID),
		calculatemw.WithAppGoodID(h.AppGoodID),
		calculatemw.WithOrderID(h.OrderID),
		calculatemw.WithGoodCoinTypeID(h.goodMainCoin.CoinTypeID),
		calculatemw.WithUnits(h.Units),
		calculatemw.WithPaymentAmountUSD(h.PaymentAmountUSD),
		calculatemw.WithGoodValueUSD(h.PaymentGoodValueUSD),
		calculatemw.WithSettleType(inspiretypes.SettleType_GoodOrderPayment),
		calculatemw.WithHasCommission(hasCommission),
		calculatemw.WithOrderCreatedAt(h.CreatedAt),
		calculatemw.WithPayments(func() (payments []*calculatemwpb.Payment) {
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
		}()),
	)
	if err != nil {
		return err
	}

	h.orderStatements, err = handler.Calculate(ctx)
	return wlog.WrapError(err)
}

//nolint:gocritic
func (h *orderHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"PowerRentalOrder", h.PowerRentalOrder,
			"OrderStatements", h.orderStatements,
			"Error", *err,
		)
	}
	persistentOrder := &types.PersistentOrder{
		PowerRentalOrder: h.PowerRentalOrder,
		OrderStatements:  h.orderStatements,
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
