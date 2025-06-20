package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/cancel/returnbalance/types"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	ledgermw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	state := ordertypes.OrderState_OrderStateCanceledTransferBookKeeping

	handler, err := feeordermw.NewHandler(
		ctx,
		feeordermw.WithID(&order.ID, true),
		feeordermw.WithOrderState(&state, true),
	)
	if err != nil {
		return err
	}

	return handler.UpdateFeeOrder(ctx)
}

func (p *handler) withReturnLockedBalance(ctx context.Context, order *types.PersistentOrder) error {
	if order.PaymentOp != types.Unlock {
		return nil
	}

	handler, err := ledgermw.NewHandler(
		ctx,
		ledgermw.WithLockID(&order.LedgerLockID, true),
	)
	if err != nil {
		return err
	}

	_, err = handler.UnlockBalances(ctx)
	return err
}

func (p *handler) withReturnSpent(ctx context.Context, order *types.PersistentOrder) error {
	if order.PaymentOp != types.Unspend {
		return nil
	}

	ioType := ledgertypes.IOType_Incoming
	ioSubType := ledgertypes.IOSubType_OrderRevoke
	reqs := []*statementmwpb.StatementReq{}

	for _, payment := range order.Payments {
		reqs = append(reqs, &statementmwpb.StatementReq{
			AppID:      &order.AppID,
			UserID:     &order.UserID,
			CurrencyID: &payment.CoinTypeID,
			IOType:     &ioType,
			IOSubType:  &ioSubType,
			Amount:     &payment.Amount,
			IOExtra:    &payment.SpentExtra,
		})
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
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if err := p.withUpdateOrderState(ctx, _order); err != nil {
		return err
	}
	if err := p.withReturnLockedBalance(ctx, _order); err != nil {
		return err
	}
	if err := p.withReturnSpent(ctx, _order); err != nil {
		return err
	}

	return nil
}
