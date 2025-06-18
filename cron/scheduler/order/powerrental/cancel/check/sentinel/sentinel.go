package sentinel

import (
	"context"
	"time"

	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	powerrentalordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental"
	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/powerrental/cancel/check/types"
	powerrentalordermwcli "github.com/NpoolPlatform/kunman/middleware/order/powerrental"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) scanPowerRentalOrders(ctx context.Context, admin bool, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	for {
		updatedAt := uint32(time.Now().Unix()) - timedef.SecondsPerMinute
		conds := &powerrentalordermwpb.Conds{
			OrderStates: &basetypes.Uint32SliceVal{Op: cruder.IN, Value: []uint32{
				uint32(ordertypes.OrderState_OrderStatePaid),
				uint32(ordertypes.OrderState_OrderStateWaitPayment),
				uint32(ordertypes.OrderState_OrderStateInService),
			}},
			Simulate:  &basetypes.BoolVal{Op: cruder.EQ, Value: false},
			UpdatedAt: &basetypes.Uint32Val{Op: cruder.LT, Value: updatedAt},
			PaymentTypes: &basetypes.Uint32SliceVal{
				Op: cruder.IN,
				Value: []uint32{
					uint32(ordertypes.PaymentType_PayWithBalanceOnly),
					uint32(ordertypes.PaymentType_PayWithTransferOnly),
					uint32(ordertypes.PaymentType_PayWithTransferAndBalance),
					uint32(ordertypes.PaymentType_PayWithOffline),
					uint32(ordertypes.PaymentType_PayWithNoPayment),
				},
			},
		}
		if admin {
			conds.AdminSetCanceled = &basetypes.BoolVal{Op: cruder.EQ, Value: true}
		} else {
			conds.UserSetCanceled = &basetypes.BoolVal{Op: cruder.EQ, Value: true}
		}
		orders, _, err := powerrentalordermwcli.GetPowerRentalOrders(ctx, conds, offset, limit)
		if err != nil {
			return err
		}
		if len(orders) == 0 {
			return nil
		}

		for _, order := range orders {
			cancelablefeed.CancelableFeed(ctx, order, exec)
		}

		offset += limit
	}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	if err := h.scanPowerRentalOrders(ctx, true, exec); err != nil {
		return err
	}
	return h.scanPowerRentalOrders(ctx, false, exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if order, ok := ent.(*types.PersistentPowerRentalOrder); ok {
		return order.OrderID
	}
	return ent.(*powerrentalordermwpb.PowerRentalOrder).OrderID
}
