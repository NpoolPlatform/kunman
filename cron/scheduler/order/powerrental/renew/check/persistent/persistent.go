package persistent

import (
	"context"
	"fmt"

	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	powerrentaloutofgasmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental/outofgas"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/renew/check/types"
	powerrentalordermwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
	powerrentaloutofgasmwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental/outofgas"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, order interface{}, reward, notif, done chan interface{}) error {
	_order, ok := order.(*types.PersistentOrder)
	if !ok {
		return fmt.Errorf("invalid order")
	}

	defer asyncfeed.AsyncFeed(ctx, _order, done)

	if _order.CreateOutOfGas {
		if err := powerrentaloutofgasmwcli.CreateOutOfGas(ctx, &powerrentaloutofgasmwpb.OutOfGasReq{
			OrderID: &_order.OrderID,
			StartAt: &_order.FeeEndAt,
		}); err != nil {
			return err
		}
	}

	if _order.FinishOutOfGas {
		if err := powerrentaloutofgasmwcli.UpdateOutOfGas(ctx, &powerrentaloutofgasmwpb.OutOfGasReq{
			EntID: &_order.OutOfGasEntID,
			EndAt: &_order.OutOfGasFinishedAt,
		}); err != nil {
			return nil
		}
	}

	return powerrentalordermwcli.UpdatePowerRentalOrder(ctx, &powerrentalordermwpb.PowerRentalOrderReq{
		ID:            &_order.ID,
		RenewState:    &_order.NewRenewState,
		RenewNotifyAt: &_order.NextRenewNotifyAt,
	})
}
