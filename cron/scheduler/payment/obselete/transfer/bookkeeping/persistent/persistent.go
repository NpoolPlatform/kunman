package persistent

import (
	"context"
	"fmt"

	ledgerstatementmwcli "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/transfer/bookkeeping/types"
	paymentmwcli "github.com/NpoolPlatform/kunman/middleware/order/payment"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, payment interface{}, reward, notif, done chan interface{}) error {
	_payment, ok := payment.(*types.PersistentPayment)
	if !ok {
		return fmt.Errorf("invalid payment")
	}

	defer asyncfeed.AsyncFeed(ctx, _payment, done)

	if len(_payment.Statements) > 0 {
		if _, err := ledgerstatementmwcli.CreateStatements(ctx, _payment.Statements); err != nil {
			return err
		}
	}

	return paymentmwcli.UpdatePayment(ctx, &paymentmwpb.PaymentReq{
		ID:               &_payment.ID,
		ObseleteState:    ordertypes.PaymentObseleteState_PaymentObseleteTransferUnlockAccount.Enum(),
		PaymentTransfers: _payment.PaymentTransfers,
	})
}
