package persistent

import (
	"context"
	"fmt"

	paymentaccountmwcli "github.com/NpoolPlatform/kunman/middleware/account/payment"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/transfer/unlockaccount/types"
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

	// Every time we update one account, until all accounts are unlocked
	if _payment.UnlockAccountID != nil {
		if _, err := paymentaccountmwcli.UnlockAccount(ctx, *_payment.UnlockAccountID); err != nil {
			return err
		}
		return nil
	}

	return paymentmwcli.UpdatePayment(ctx, &paymentmwpb.PaymentReq{
		ID:            &_payment.ID,
		ObseleteState: ordertypes.PaymentObseleteState_PaymentObseleted.Enum(),
	})
}
