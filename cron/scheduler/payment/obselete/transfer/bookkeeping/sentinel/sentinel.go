package sentinel

import (
	"context"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/payment/obselete/transfer/bookkeeping/types"
	ordertypes "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	paymentmwpb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	paymentmw "github.com/NpoolPlatform/kunman/middleware/order/payment"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type handler struct{}

func NewSentinel() basesentinel.Scanner {
	return &handler{}
}

func (h *handler) scanPayments(ctx context.Context, state ordertypes.PaymentObseleteState, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit

	conds := &paymentmwpb.Conds{
		ObseleteState: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(state)},
	}

	for {
		handler, err := paymentmw.NewHandler(
			ctx,
			paymentmw.WithConds(conds),
			paymentmw.WithOffset(offset),
			paymentmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		payments, _, err := handler.GetPayments(ctx)
		if err != nil {
			return err
		}
		if len(payments) == 0 {
			return nil
		}

		for _, payment := range payments {
			cancelablefeed.CancelableFeed(ctx, payment, exec)
		}

		offset += limit
	}
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	return h.scanPayments(ctx, ordertypes.PaymentObseleteState_PaymentObseleteTransferBookKeeping, exec)
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return h.scanPayments(ctx, ordertypes.PaymentObseleteState_PaymentObseleteTransferBookKeeping, exec)
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	if payment, ok := ent.(*types.PersistentPayment); ok {
		return payment.EntID
	}
	return ent.(*paymentmwpb.Payment).EntID
}
