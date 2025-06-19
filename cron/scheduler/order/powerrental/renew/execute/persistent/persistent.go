package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/renew/execute/types"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	ledgermw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withLockBalances(ctx context.Context, order *types.PersistentOrder) error {
	balances := func() (_balances []*ledgermwpb.LockBalance) {
		for _, req := range order.FeeOrderReqs {
			for _, balance := range req.PaymentBalances {
				_balances = append(_balances, &ledgermwpb.LockBalance{
					CoinTypeID: *balance.CoinTypeID,
					Amount:     *balance.Amount,
				})
			}
		}
		return
	}()

	handler, err := ledgermw.NewHandler(
		ctx,
		ledgermw.WithAppID(&order.AppID, true),
		ledgermw.WithUserID(&order.UserID, true),
		ledgermw.WithLockID(&order.LedgerLockID, true),
		ledgermw.WithBalances(balances, true),
	)
	if err != nil {
		return err
	}

	_, err = handler.LockBalances(ctx)
	return err
}

func (p *handler) withCreateFeeOrders(ctx context.Context, order *types.PersistentOrder) error {
	multiHandler := &feeordermw.MultiHandler{}

	for _, req := range order.FeeOrderReqs {
		handler, err := feeordermw.NewHandler(
			ctx,
			feeordermw.WithEntID(req.EntID, false),
			feeordermw.WithAppID(req.AppID, true),
			feeordermw.WithUserID(req.UserID, true),
			feeordermw.WithGoodID(req.GoodID, true),
			feeordermw.WithGoodType(req.GoodType, true),
			feeordermw.WithAppGoodID(req.AppGoodID, true),
			feeordermw.WithOrderID(req.OrderID, false),
			feeordermw.WithParentOrderID(req.ParentOrderID, true),
			feeordermw.WithOrderType(req.OrderType, true),
			feeordermw.WithPaymentType(req.PaymentType, false),
			feeordermw.WithCreateMethod(req.CreateMethod, true),

			feeordermw.WithGoodValueUSD(req.GoodValueUSD, true),
			feeordermw.WithPaymentAmountUSD(req.PaymentAmountUSD, false),
			feeordermw.WithDiscountAmountUSD(req.DiscountAmountUSD, false),
			feeordermw.WithPromotionID(req.PromotionID, false),
			feeordermw.WithDurationSeconds(req.DurationSeconds, true),
			feeordermw.WithLedgerLockID(req.LedgerLockID, false),
			feeordermw.WithPaymentID(req.PaymentID, false),
			feeordermw.WithCouponIDs(req.CouponIDs, false),
			feeordermw.WithPaymentBalances(req.PaymentBalances, false),
			feeordermw.WithPaymentTransfers(req.PaymentTransfers, false),
		)
		if err != nil {
			return err
		}
		multiHandler.AppendHandler(handler)
	}

	return multiHandler.CreateFeeOrders(ctx)
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithID(&_order.ID, true),
		powerrentalordermw.WithRenewState(&_order.NewRenewState, true),
	)
	if err != nil {
		return err
	}

	if err := handler.UpdatePowerRental(ctx); err != nil {
		return err
	}
	if _order.InsufficientBalance {
		return nil
	}

	if err := p.withCreateFeeOrders(ctx, _order); err != nil {
		return err
	}
	if err := p.withLockBalances(ctx, _order); err != nil {
		return err
	}

	return nil
}
