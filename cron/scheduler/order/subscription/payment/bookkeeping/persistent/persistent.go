package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/subscription/payment/bookkeeping/types"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgerstatementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	subscriptionordermw "github.com/NpoolPlatform/kunman/middleware/order/subscription"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateOrderState(ctx context.Context, order *types.PersistentOrder) error {
	state := ordertypes.OrderState_OrderStatePaymentSpendBalance

	handler, err := subscriptionordermw.NewHandler(
		ctx,
		subscriptionordermw.WithID(&order.ID, true),
		subscriptionordermw.WithOrderState(&state, true),
		subscriptionordermw.WithPaymentTransfers(func() (paymentTransfers []*paymentmwpb.PaymentTransferReq) {
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

	return handler.UpdateSubscriptionOrder(ctx)
}

func (p *handler) withCreateStatements(ctx context.Context, order *types.PersistentOrder) error {
	reqs := []*ledgerstatementmwpb.StatementReq{}

	for _, paymentTransfer := range order.XPaymentTransfers {
		if paymentTransfer.IncomingAmount == nil {
			continue
		}
		reqs = append(reqs, &ledgerstatementmwpb.StatementReq{
			AppID:        &order.AppID,
			UserID:       &order.UserID,
			CurrencyID:   &paymentTransfer.CoinTypeID,
			CurrencyType: ledgertypes.CurrencyType_CurrencyCrypto.Enum(),
			IOType:       func() *ledgertypes.IOType { e := ledgertypes.IOType_Incoming; return &e }(),
			IOSubType:    func() *ledgertypes.IOSubType { e := ledgertypes.IOSubType_Payment; return &e }(),
			Amount:       paymentTransfer.IncomingAmount,
			IOExtra:      &paymentTransfer.IncomingExtra,
		}, &ledgerstatementmwpb.StatementReq{
			AppID:        &order.AppID,
			UserID:       &order.UserID,
			CurrencyID:   &paymentTransfer.CoinTypeID,
			CurrencyType: ledgertypes.CurrencyType_CurrencyCrypto.Enum(),
			IOType:       func() *ledgertypes.IOType { e := ledgertypes.IOType_Outcoming; return &e }(),
			IOSubType:    func() *ledgertypes.IOSubType { e := ledgertypes.IOSubType_Payment; return &e }(),
			Amount:       &paymentTransfer.Amount,
			IOExtra:      &paymentTransfer.OutcomingExtra,
		})
	}
	for _, paymentFiat := range order.XPaymentFiats {
		reqs = append(reqs, &ledgerstatementmwpb.StatementReq{
			AppID:        &order.AppID,
			UserID:       &order.UserID,
			CurrencyID:   &paymentFiat.FiatID,
			CurrencyType: ledgertypes.CurrencyType_CurrencyFiat.Enum(),
			IOType:       func() *ledgertypes.IOType { e := ledgertypes.IOType_Incoming; return &e }(),
			IOSubType:    func() *ledgertypes.IOSubType { e := ledgertypes.IOSubType_Payment; return &e }(),
			Amount:       &paymentFiat.Amount,
			IOExtra:      &paymentFiat.Extra,
		}, &ledgerstatementmwpb.StatementReq{
			AppID:        &order.AppID,
			UserID:       &order.UserID,
			CurrencyID:   &paymentFiat.FiatID,
			CurrencyType: ledgertypes.CurrencyType_CurrencyFiat.Enum(),
			IOType:       func() *ledgertypes.IOType { e := ledgertypes.IOType_Outcoming; return &e }(),
			IOSubType:    func() *ledgertypes.IOSubType { e := ledgertypes.IOSubType_Payment; return &e }(),
			Amount:       &paymentFiat.Amount,
			IOExtra:      &paymentFiat.Extra,
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
	if err := p.withCreateStatements(ctx, _order); err != nil {
		return err
	}

	return nil
}
