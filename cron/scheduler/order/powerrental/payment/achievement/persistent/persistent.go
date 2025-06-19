package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/payment/achievement/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	orderstatementmw "github.com/NpoolPlatform/kunman/middleware/inspire/achievement/statement/order"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if len(_order.OrderStatements) > 0 {
		multiHandler := &orderstatementmw.MultiHandler{}
		for _, req := range _order.OrderStatements {
			handler, err := orderstatementmw.NewHandler(
				ctx,
				orderstatementmw.WithEntID(req.EntID, false),
				orderstatementmw.WithAppID(req.AppID, true),
				orderstatementmw.WithUserID(req.UserID, true),
				orderstatementmw.WithGoodID(req.GoodID, true),
				orderstatementmw.WithAppGoodID(req.AppGoodID, true),
				orderstatementmw.WithOrderID(req.OrderID, true),
				orderstatementmw.WithOrderUserID(req.OrderUserID, true),
				orderstatementmw.WithDirectContributorID(req.DirectContributorID, true),
				orderstatementmw.WithGoodCoinTypeID(req.GoodCoinTypeID, true),
				orderstatementmw.WithUnits(req.Units, true),
				orderstatementmw.WithGoodValueUSD(req.GoodValueUSD, true),
				orderstatementmw.WithPaymentAmountUSD(req.PaymentAmountUSD, true),
				orderstatementmw.WithCommissionAmountUSD(req.CommissionAmountUSD, true),
				orderstatementmw.WithAppConfigID(req.AppConfigID, true),
				orderstatementmw.WithCommissionConfigID(req.CommissionConfigID, false),
				orderstatementmw.WithCommissionConfigType(req.CommissionConfigType, true),
				orderstatementmw.WithPaymentStatements(req.PaymentStatements, true),
			)
			if err != nil {
				return err
			}
			multiHandler.AppendHandler(handler)
		}
		if err := multiHandler.CreateStatements(ctx); err != nil {
			return err
		}
	}

	state := ordertypes.OrderState_OrderStateAddCommission

	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithID(&_order.ID, true),
		powerrentalordermw.WithOrderState(&state, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdatePowerRental(ctx)
}
