package powerrentalstate

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	powerrentalstatecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/powerrental/state"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	powerrentalstatecrud.Req
	Rollback *bool
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = &_id
		return nil
	}
}

//nolint:gocyclo
func WithCancelState(state *types.OrderState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return wlog.Errorf("invalid cancelstate")
			}
			return nil
		}
		switch *state {
		case types.OrderState_OrderStateCreated:
		case types.OrderState_OrderStateWaitPayment:
		case types.OrderState_OrderStatePaymentTransferReceived:
		case types.OrderState_OrderStatePaymentTransferBookKeeping:
		case types.OrderState_OrderStatePaymentSpendBalance:
		case types.OrderState_OrderStateTransferGoodStockLocked:
		case types.OrderState_OrderStateAddCommission:
		case types.OrderState_OrderStateAchievementBookKeeping:
		case types.OrderState_OrderStatePaymentUnlockAccount:
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateTransferGoodStockWaitStart:
		case types.OrderState_OrderStateInService:
		case types.OrderState_OrderStatePaymentTimeout:
		case types.OrderState_OrderStatePreCancel:
		case types.OrderState_OrderStatePreExpired:
		case types.OrderState_OrderStateRestoreExpiredStock:
		case types.OrderState_OrderStateRestoreCanceledStock:
		case types.OrderState_OrderStateCancelAchievement:
		case types.OrderState_OrderStateDeductLockedCommission:
		case types.OrderState_OrderStateReturnCanceledBalance:
		case types.OrderState_OrderStateCanceledTransferBookKeeping:
		case types.OrderState_OrderStateCancelUnlockPaymentAccount:
		case types.OrderState_OrderStateCanceled:
		case types.OrderState_OrderStateExpired:
		default:
			return wlog.Errorf("invalid cancelstate")
		}
		h.CancelState = state
		return nil
	}
}

func WithPaidAt(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid paidat")
			}
			return nil
		}
		if *u < uint32(time.Now().Unix()) {
			return wlog.Errorf("invalid paidat")
		}
		h.PaidAt = u
		return nil
	}
}

func WithUserSetPaid(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.UserSetPaid = b
		return nil
	}
}

func WithUserSetCanceled(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.UserSetCanceled = b
		return nil
	}
}

func WithAdminSetCanceled(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AdminSetCanceled = b
		return nil
	}
}

func WithPaymentState(state *types.PaymentState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return wlog.Errorf("invalid paymentstate")
			}
			return nil
		}
		switch *state {
		case types.PaymentState_PaymentStateWait:
		case types.PaymentState_PaymentStateCanceled:
		case types.PaymentState_PaymentStateTimeout:
		case types.PaymentState_PaymentStateDone:
		case types.PaymentState_PaymentStateNoPayment:
		default:
			return wlog.Errorf("invalid paymentstate")
		}
		h.PaymentState = state
		return nil
	}
}

func WithRenewState(e *types.OrderRenewState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid renewstate")
			}
			return nil
		}
		switch *e {
		case types.OrderRenewState_OrderRenewWait:
		case types.OrderRenewState_OrderRenewCheck:
		case types.OrderRenewState_OrderRenewNotify:
		case types.OrderRenewState_OrderRenewExecute:
		case types.OrderRenewState_OrderRenewFail:
		default:
			return wlog.Errorf("invalid renewstate")
		}
		h.RenewState = e
		return nil
	}
}

func WithRenewNotifyAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RenewNotifyAt = n
		return nil
	}
}

func WithRollback(rollback *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Rollback = rollback
		return nil
	}
}

func WithOutOfGasSeconds(outOfGasSeconds *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if outOfGasSeconds == nil {
			if must {
				return wlog.Errorf("invalid outofgasseconds")
			}
			return nil
		}
		h.OutOfGasSeconds = outOfGasSeconds
		return nil
	}
}

func WithCompensateSeconds(compensateSeconds *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if compensateSeconds == nil {
			if must {
				return wlog.Errorf("invalid compensateseconds")
			}
			return nil
		}
		h.CompensateSeconds = compensateSeconds
		return nil
	}
}
