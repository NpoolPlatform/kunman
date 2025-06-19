package notif

import (
	"context"
	"fmt"

	basenotif "github.com/NpoolPlatform/kunman/cron/scheduler/base/notif"
	retry1 "github.com/NpoolPlatform/kunman/cron/scheduler/base/retry"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/payment/unlockaccount/types"
)

type handler struct{}

func NewNotif() basenotif.Notify {
	return &handler{}
}

func (p *handler) notifyPaid(order *types.PersistentOrder) error {
	// TODO
	return nil
}

func (p *handler) Notify(ctx context.Context, order interface{}, retry chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}
	if err := p.notifyPaid(_order); err != nil {
		retry1.Retry(_order.EntID, _order, retry)
		return err
	}
	return nil
}
