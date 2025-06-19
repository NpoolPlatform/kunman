package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/unlockbalance/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgermw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger"
	paymentmw "github.com/NpoolPlatform/kunman/middleware/order/payment"
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

	if _payment.XLedgerLockID != nil {
		handler, err := ledgermw.NewHandler(
			ctx,
			ledgermw.WithLockID(_payment.XLedgerLockID, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.UnlockBalances(ctx); err != nil {
			return err
		}
	}

	handler, err := paymentmw.NewHandler(
		ctx,
		paymentmw.WithID(&_payment.ID, true),
		paymentmw.WithObseleteState(ordertypes.PaymentObseleteState_PaymentObseleteTransferBookKeeping.Enum(), true),
	)
	if err != nil {
		return err
	}

	// TODO: here state is not atomic but Ledger Lock can not be unlock twice, so it may stuck here
	return handler.UpdatePayment(ctx)
}
