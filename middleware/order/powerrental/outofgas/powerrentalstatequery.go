package outofgas

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entorderstatebase "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/orderstatebase"
	entpowerrentalstate "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/powerrentalstate"
)

type powerRentalStateQueryHandler struct {
	*Handler
	_ent powerRentalState
}

func (h *powerRentalStateQueryHandler) getPowerRentalStateEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if _, err := cli.
		OrderStateBase.
		Query().
		Where(
			entorderstatebase.OrderID(*h.OrderID),
			entorderstatebase.OrderState(types.OrderState_OrderStateInService.String()),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}

	if h._ent.entPowerRentalState, err = cli.
		PowerRentalState.
		Query().
		Where(
			entpowerrentalstate.OrderID(*h.OrderID),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalStateQueryHandler) _getPowerRentalState(ctx context.Context, must bool) error {
	if h.OrderID == nil {
		return wlog.Errorf("invalid orderid")
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h.getPowerRentalStateEnt(ctx, cli, must)
	})
}

func (h *powerRentalStateQueryHandler) requirePowerRentalState(ctx context.Context) error {
	return h._getPowerRentalState(ctx, true)
}
