package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/cancel/bookkeeping/types"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgerstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	powerrentalordermw "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentPowerRentalOrder) error {
	handler, err := powerrentalordermw.NewHandler(
		ctx,
		powerrentalordermw.WithID(&order.ID, true),
		powerrentalordermw.WithOrderState(ordertypes.OrderState_OrderStateCancelUnlockPaymentAccount.Enum(), true),
		powerrentalordermw.WithPaymentTransfers(func() (paymentTransfers []*paymentmwpb.PaymentTransferReq) {
			for _, paymentTransfer := range order.XPaymentTransfers {
				paymentTransfers = append(paymentTransfers, &paymentmwpb.PaymentTransferReq{
					EntID:        &paymentTransfer.PaymentTransferID,
					FinishAmount: &paymentTransfer.FinishAmount,
				})
			}
			return
		}(), true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return handler.UpdatePowerRental(ctx)
}

func (p *handler) withCreateIncomingStatements(ctx context.Context, order *types.PersistentPowerRentalOrder) error {
	reqs := []*ledgerstatementmwpb.StatementReq{}
	ioType := ledgertypes.IOType_Incoming
	ioSubType := ledgertypes.IOSubType_Payment

	for _, paymentTransfer := range order.XPaymentTransfers {
		if paymentTransfer.IncomingAmount == nil {
			continue
		}
		reqs = append(reqs, &ledgerstatementmwpb.StatementReq{
			AppID:      &order.AppID,
			UserID:     &order.UserID,
			CurrencyID: &paymentTransfer.CoinTypeID,
			IOType:     &ioType,
			IOSubType:  &ioSubType,
			Amount:     paymentTransfer.IncomingAmount,
			IOExtra:    &order.IncomingExtra,
		})
	}
	if len(reqs) == 0 {
		return nil
	}

	handler, err := ledgerstatementmw.NewHandler(
		ctx,
		ledgerstatementmw.WithReqs(reqs, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	_, err = handler.CreateStatements(ctx)
	return err
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentPowerRentalOrder)
	if !ok {
		return fmt.Errorf("invalid powerrentalorder")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if err := p.withUpdateOrderState(ctx, _order); err != nil {
		return err
	}
	if err := p.withCreateIncomingStatements(ctx, _order); err != nil {
		return err
	}

	return nil
}
