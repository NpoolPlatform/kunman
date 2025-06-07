package orderstm

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
)

type validateHandler struct {
	*Handler
	*orderQueryHandler
	*rollbackHandler
	*forwardHandler
}

func (h *validateHandler) validateFinaled() error {
	switch h._ent.OrderState() {
	case types.OrderState_OrderStateExpired:
		fallthrough //nolint
	case types.OrderState_OrderStateCanceled:
		if h.Rollback == nil || !*h.Rollback {
			return wlog.Errorf("permission denied")
		}
	}
	return nil
}

func (h *validateHandler) validateUserCancelable() error {
	if h.UserSetCanceled == nil {
		return nil
	}
	if !*h.UserSetCanceled {
		return wlog.Errorf("permission denied")
	}
	if *h.UserCanceled {
		return wlog.Errorf("permission denied")
	}
	if h._ent.PaymentType() == types.PaymentType_PayWithParentOrder {
		return wlog.Errorf("permission denied")
	}
	switch h._ent.OrderType() {
	case types.OrderType_Normal:
	default:
		return wlog.Errorf("permission denied")
	}
	switch h._ent.OrderState() {
	case types.OrderState_OrderStateWaitPayment:
	case types.OrderState_OrderStatePaid:
	case types.OrderState_OrderStateInService:
	default:
		return wlog.Errorf("invalid cancelstate")
	}
	return nil
}

func (h *validateHandler) validateAdminCancelable() error {
	if h.AdminSetCanceled == nil {
		return nil
	}
	if !*h.AdminSetCanceled {
		return wlog.Errorf("invalid cancelstate")
	}
	if *h.AdminCanceled {
		return wlog.Errorf("invalid cancelstate")
	}
	switch h._ent.OrderState() {
	case types.OrderState_OrderStatePaid:
	case types.OrderState_OrderStateInService:
	default:
		return wlog.Errorf("permission denied")
	}
	switch h._ent.OrderType() {
	case types.OrderType_Offline:
	case types.OrderType_Airdrop:
	case types.OrderType_Normal:
	default:
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *validateHandler) validateUserPayable() error {
	if h.UserSetPaid == nil {
		return nil
	}
	if !*h.UserSetPaid {
		return wlog.Errorf("permission denied")
	}
	switch h._ent.OrderType() {
	case types.OrderType_Normal:
	default:
		return wlog.Errorf("permissioned denied")
	}
	return nil
}

func (h *validateHandler) validatePaymentState() error {
	if h.NewPaymentState == nil {
		return nil
	}
	switch *h.CurrentPaymentState {
	case types.PaymentState_PaymentStateWait:
		switch *h.NewPaymentState {
		case types.PaymentState_PaymentStateDone:
		case types.PaymentState_PaymentStateCanceled:
		case types.PaymentState_PaymentStateTimeout:
		case types.PaymentState_PaymentStateNoPayment:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.PaymentState_PaymentStateDone:
		fallthrough //nolint
	case types.PaymentState_PaymentStateCanceled:
		fallthrough //nolint
	case types.PaymentState_PaymentStateTimeout:
		fallthrough //nolint
	case types.PaymentState_PaymentStateNoPayment:
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *validateHandler) validateOrderState() (*types.OrderState, error) {
	if h.OrderState == nil {
		return nil, nil
	}
	if h.Rollback != nil && *h.Rollback && *h.OrderState == h._ent.OrderState() {
		return h.rollback()
	}
	return h.forward()
}

func (h *Handler) ValidateUpdateForNewState(ctx context.Context, tx *ent.Tx) (*types.OrderState, error) {
	handler := &validateHandler{
		Handler: h,
		orderQueryHandler: &orderQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireOrderWithTx(ctx, tx); err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.rollbackHandler = &rollbackHandler{
		orderQueryHandler: handler.orderQueryHandler,
	}
	handler.forwardHandler = &forwardHandler{
		orderQueryHandler: handler.orderQueryHandler,
	}

	if err := handler.validateFinaled(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateUserCancelable(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateAdminCancelable(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validateUserPayable(); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := handler.validatePaymentState(); err != nil {
		return nil, wlog.WrapError(err)
	}
	return handler.validateOrderState()
}
