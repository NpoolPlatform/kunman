package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/wait/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
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

	handler, err := paymentmw.NewHandler(
		ctx,
		paymentmw.WithID(&_payment.ID, true),
		paymentmw.WithObseleteState(ordertypes.PaymentObseleteState_PaymentObseleteUnlockBalance.Enum(), true),
	)
	if err != nil {
		return err
	}

	return handler.UpdatePayment(ctx)
}
