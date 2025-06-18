package sentinel

import (
	"context"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/order/fee/paid/check/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	feeordermwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/fee"
	feeordermw "github.com/NpoolPlatform/kunman/middleware/order/fee"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) scanOrdersPayment(ctx context.Context, state ordertypes.OrderState, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &feeordermwpb.Conds{
		OrderState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(state)},
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
		AdminSetCanceled: &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		UserSetCanceled:  &basetypes.BoolVal{Op: cruder.EQ, Value: false},
	}

	for {
		handler, err := feeordermw.NewHandler(
			ctx,
			feeordermw.WithConds(conds),
			feeordermw.WithOffset(offset),
			feeordermw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		orders, _, err := handler.GetFeeOrders(ctx)
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
	return h.scanOrdersPayment(ctx, ordertypes.OrderState_OrderStatePaid, exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return nil
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if order, ok := ent.(*types.PersistentOrder); ok {
		return order.EntID
	}
	return ent.(*feeordermwpb.FeeOrder).EntID
}
