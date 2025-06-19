package persistent

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/miningpool/checkproportion/types"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, req *powerrentalordermwpb.PowerRentalOrderReq) error {
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

	if _order.PowerRentalOrderReq != nil {
		return p.withUpdateOrderState(ctx, _order.PowerRentalOrderReq)
	}

	return nil
}
