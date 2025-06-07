package subscriptionorder

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"

	"github.com/google/uuid"
)

type MultiHandler struct {
	Handlers []*Handler
}

func (h *MultiHandler) AppendHandler(handler *Handler) {
	h.Handlers = append(h.Handlers, handler)
}

func (h *MultiHandler) GetHandlers() []*Handler {
	return h.Handlers
}

func (h *MultiHandler) validatePaymentOrder() (bool, error) {
	paymentOrders := 0
	offlineOrder := false

	for _, handler := range h.Handlers {
		switch *handler.OrderBaseReq.OrderType {
		case types.OrderType_Offline:
			fallthrough //nolint
		case types.OrderType_Airdrop:
			offlineOrder = true
			continue
		}
		if offlineOrder {
			return false, wlog.Errorf("invalid ordertype")
		}
		if len(handler.PaymentTransferReqs) > 0 || len(handler.PaymentBalanceReqs) > 0 {
			paymentOrders += 1
		}
		if handler.OrderStateBaseReq.PaymentType == nil {
			continue
		}
	}
	return paymentOrders > 0, nil
}

//nolint:unparam
func (h *MultiHandler) formalizePaymentID() error {
	var paymentID *uuid.UUID

	for _, handler := range h.Handlers {
		if handler.PaymentBaseReq.EntID != nil {
			paymentID = handler.PaymentBaseReq.EntID
			break
		}
	}
	if paymentID == nil {
		paymentID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	for _, handler := range h.Handlers {
		handler.PaymentBaseReq.EntID = paymentID
		for _, balance := range handler.PaymentBalanceReqs {
			balance.PaymentID = paymentID
		}
		for _, transfer := range handler.PaymentTransferReqs {
			transfer.PaymentID = paymentID
		}
	}
	return nil
}

func (h *MultiHandler) CreateSubscriptionOrdersWithTx(ctx context.Context, tx *ent.Tx) error {
	shouldPay, err := h.validatePaymentOrder()
	if err != nil {
		return wlog.WrapError(err)
	}
	if shouldPay {
		if err := h.formalizePaymentID(); err != nil {
			return wlog.WrapError(err)
		}
	}
	for _, handler := range h.Handlers {
		if err := handler.CreateSubscriptionOrderWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) CreateSubscriptionOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateSubscriptionOrdersWithTx(_ctx, tx)
	})
}

func (h *MultiHandler) UpdateSubscriptionOrdersWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.UpdateSubscriptionOrderWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) UpdateSubscriptionOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.UpdateSubscriptionOrdersWithTx(_ctx, tx)
	})
}

func (h *MultiHandler) DeleteSubscriptionOrdersWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.DeleteSubscriptionOrderWithTx(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
	}
	return nil
}

func (h *MultiHandler) DeleteSubscriptionOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteSubscriptionOrdersWithTx(_ctx, tx)
	})
}
