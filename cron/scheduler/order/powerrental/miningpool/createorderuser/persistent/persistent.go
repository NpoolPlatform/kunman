package persistent

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/miningpool/createorderuser/types"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	orderusermwpb "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	orderusermw "github.com/NpoolPlatform/kunman/middleware/miningpool/orderuser"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withCreateOrderUser(ctx context.Context, req *orderusermwpb.OrderUserReq) error {
	handler, err := orderusermw.NewHandler(
		ctx,
		orderusermw.WithEntID(req.EntID, false),
		orderusermw.WithGoodUserID(req.GoodUserID, true),
		orderusermw.WithAppID(req.AppID, true),
		orderusermw.WithUserID(req.UserID, true),
	)
	if err != nil {
		return err
	}

	return handler.CreateOrderUser(ctx)
}

func (p *handler) withUpdateOrder(ctx context.Context, req *powerrentalordermwpb.PowerRentalOrderReq) error {
	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithID(req.ID, false),
		powerrentalordermw.WithEntID(req.EntID, false),
		powerrentalordermw.WithOrderID(req.OrderID, false),
		powerrentalordermw.WithPaymentType(req.PaymentType, false),

		powerrentalordermw.WithOrderState(req.OrderState, false),
		powerrentalordermw.WithStartMode(req.StartMode, false),
		powerrentalordermw.WithStartAt(req.StartAt, false),
		powerrentalordermw.WithLastBenefitAt(req.LastBenefitAt, false),
		powerrentalordermw.WithBenefitState(req.BenefitState, false),
		powerrentalordermw.WithUserSetPaid(req.UserSetPaid, false),
		powerrentalordermw.WithUserSetCanceled(req.UserSetCanceled, false),
		powerrentalordermw.WithAdminSetCanceled(req.AdminSetCanceled, false),
		powerrentalordermw.WithPaymentState(req.PaymentState, false),
		powerrentalordermw.WithRenewState(req.RenewState, false),
		powerrentalordermw.WithRenewNotifyAt(req.RenewNotifyAt, false),

		powerrentalordermw.WithLedgerLockID(req.LedgerLockID, false),
		powerrentalordermw.WithPaymentID(req.PaymentID, false),
		powerrentalordermw.WithCouponIDs(req.CouponIDs, false),
		powerrentalordermw.WithPaymentBalances(req.PaymentBalances, false),
		powerrentalordermw.WithPaymentTransfers(req.PaymentTransfers, false),

		powerrentalordermw.WithRollback(req.Rollback, false),
		powerrentalordermw.WithPoolOrderUserID(req.PoolOrderUserID, false),
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

	if _order.OrderUserReq != nil {
		if err := p.withCreateOrderUser(ctx, _order.OrderUserReq); err != nil {
			return err
		}
	}
	if _order.PowerRentalOrderReq != nil {
		return p.withUpdateOrder(ctx, _order.PowerRentalOrderReq)
	}

	return nil
}
