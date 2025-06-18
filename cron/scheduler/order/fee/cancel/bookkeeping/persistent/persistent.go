package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/bookkeeping/types"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentFeeOrder) error {
	handler, err := feeordermw.NewHandler(
		ctx,
		feeordermw.WithID(&order.ID, true),
		feeordermw.WithOrderState(ordertypes.OrderState_OrderStateCancelUnlockPaymentAccount.Enum(), true),
		feeordermw.WithPaymentTransfers(func() (paymentTransfers []*paymentmwpb.PaymentTransferReq) {
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
		return err
	}

	return handler.UpdateFeeOrder(ctx)
}

func (p *handler) withCreateIncomingStatements(ctx context.Context, order *types.PersistentFeeOrder) error {
	reqs := []*statementmwpb.StatementReq{}
	ioType := ledgertypes.IOType_Incoming
	ioSubType := ledgertypes.IOSubType_Payment

	for _, paymentTransfer := range order.XPaymentTransfers {
		if paymentTransfer.IncomingAmount == nil {
			continue
		}
		reqs = append(reqs, &statementmwpb.StatementReq{
			AppID:      &order.AppID,
			UserID:     &order.UserID,
			CoinTypeID: &paymentTransfer.CoinTypeID,
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
		return err
	}

	_, err = handler.CreateStatements(ctx)
	return err
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentFeeOrder)
	if !ok {
		return fmt.Errorf("invalid feeorder")
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
