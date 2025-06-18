package persistent

import (
	"context"
	"fmt"

	ledgermwcli "github.com/NpoolPlatform/kunman/middleware/ledger/ledger"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/unlockbalance/types"
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

	if _payment.XLedgerLockID != nil {
		if _, err := ledgermwcli.UnlockBalances(ctx, &ledgermwpb.UnlockBalancesRequest{
			LockID: *_payment.XLedgerLockID,
		}); err != nil {
			return err
		}
	}

	// TODO: here state is not atomic but Ledger Lock can not be unlock twice, so it may stuck here
	return paymentmwcli.UpdatePayment(ctx, &paymentmwpb.PaymentReq{
		ID:            &_payment.ID,
		ObseleteState: ordertypes.PaymentObseleteState_PaymentObseleteTransferBookKeeping.Enum(),
	})
}
