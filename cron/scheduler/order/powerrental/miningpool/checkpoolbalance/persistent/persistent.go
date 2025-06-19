package persistent

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/miningpool/checkpoolbalance/types"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	fractionwithdrawalmwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawal"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	fractionwithdrawalmw "github.com/NpoolPlatform/kunman/middleware/miningpool/fractionwithdrawal"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withCreateFractionWithdrawal(ctx context.Context, reqs []*fractionwithdrawalmwpb.FractionWithdrawalReq) error {
	for _, req := range reqs {
		handler, err := fractionwithdrawalmw.NewHandler(
			ctx,
			fractionwithdrawalmw.WithEntID(req.EntID, false),
			fractionwithdrawalmw.WithAppID(req.AppID, true),
			fractionwithdrawalmw.WithUserID(req.UserID, true),
			fractionwithdrawalmw.WithOrderUserID(req.OrderUserID, true),
			fractionwithdrawalmw.WithCoinTypeID(req.CoinTypeID, true),
		)
		if err != nil {
			return err
		}

		if err := handler.CreateFractionWithdrawal(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (p *handler) withUpdateOrder(ctx context.Context, req *powerrentalordermwpb.PowerRentalOrderReq) error {
	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithEntID(req.EntID, false),
		powerrentalordermw.WithAppID(req.AppID, true),
		powerrentalordermw.WithUserID(req.UserID, true),
		powerrentalordermw.WithGoodID(req.GoodID, true),
		powerrentalordermw.WithGoodType(req.GoodType, true),
		powerrentalordermw.WithAppGoodID(req.AppGoodID, true),
		powerrentalordermw.WithOrderID(req.OrderID, false),
		powerrentalordermw.WithOrderType(req.OrderType, true),
		powerrentalordermw.WithPaymentType(req.PaymentType, false),
		powerrentalordermw.WithSimulate(req.Simulate, false),
		powerrentalordermw.WithCreateMethod(req.CreateMethod, true),

		powerrentalordermw.WithAppGoodStockID(req.AppGoodStockID, false),
		powerrentalordermw.WithUnits(req.Units, true),
		powerrentalordermw.WithGoodValueUSD(req.GoodValueUSD, true),
		powerrentalordermw.WithPaymentAmountUSD(req.PaymentAmountUSD, false),
		powerrentalordermw.WithDiscountAmountUSD(req.DiscountAmountUSD, false),
		powerrentalordermw.WithPromotionID(req.PromotionID, false),
		powerrentalordermw.WithDurationSeconds(req.DurationSeconds, true),
		powerrentalordermw.WithInvestmentType(req.InvestmentType, false),
		powerrentalordermw.WithGoodStockMode(req.GoodStockMode, true),

		powerrentalordermw.WithStartMode(req.StartMode, true),
		powerrentalordermw.WithStartAt(req.StartAt, true),
		powerrentalordermw.WithAppGoodStockLockID(req.AppGoodStockLockID, false),
		powerrentalordermw.WithLedgerLockID(req.LedgerLockID, false),
		powerrentalordermw.WithPaymentID(req.PaymentID, false),
		powerrentalordermw.WithCouponIDs(req.CouponIDs, false),
		powerrentalordermw.WithPaymentBalances(req.PaymentBalances, false),
		powerrentalordermw.WithPaymentTransfers(req.PaymentTransfers, false),
	)
	if err != nil {
		return err
	}

	return handler.UpdatePowerRental(ctx)
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return wlog.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if len(_order.FractionWithdrawalReqs) > 0 {
		if err := p.withCreateFractionWithdrawal(ctx, _order.FractionWithdrawalReqs); err != nil {
			return err
		}
	}
	if _order.PowerRentalOrderReq != nil {
		if err := p.withUpdateOrder(ctx, _order.PowerRentalOrderReq); err != nil {
			return err
		}
	}

	return nil
}
