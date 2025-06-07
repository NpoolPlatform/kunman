package payment

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	paymentbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment"
	"github.com/NpoolPlatform/kunman/middleware/order/db"
	ent "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated"
	entpaymenttransfer "github.com/NpoolPlatform/kunman/middleware/order/db/ent/generated/paymenttransfer"
)

type updateHandler struct {
	*Handler
	obseleteState types.PaymentObseleteState
}

//nolint:gocyclo
func (h *updateHandler) validateObseleteState() error {
	if h.ObseleteState == nil {
		return nil
	}
	switch h.obseleteState {
	case types.PaymentObseleteState_PaymentObseleteNone:
		return wlog.Errorf("permission denied")
	case types.PaymentObseleteState_PaymentObseleteWait:
		if *h.ObseleteState != types.PaymentObseleteState_PaymentObseleteUnlockBalance {
			return wlog.Errorf("permission denied")
		}
	case types.PaymentObseleteState_PaymentObseleteUnlockBalance:
		switch *h.ObseleteState {
		case types.PaymentObseleteState_PaymentObseleteTransferBookKeeping:
			fallthrough //nolint
		case types.PaymentObseleteState_PaymentObseleteFail:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.PaymentObseleteState_PaymentObseleteTransferBookKeeping:
		switch *h.ObseleteState {
		case types.PaymentObseleteState_PaymentObseleteTransferUnlockAccount:
			fallthrough //nolint
		case types.PaymentObseleteState_PaymentObseleteFail:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.PaymentObseleteState_PaymentObseleteTransferUnlockAccount:
		switch *h.ObseleteState {
		case types.PaymentObseleteState_PaymentObseleted:
			fallthrough //nolint
		case types.PaymentObseleteState_PaymentObseleteFail:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.PaymentObseleteState_PaymentObseleted:
		return wlog.Errorf("permission denied")
	case types.PaymentObseleteState_PaymentObseleteFail:
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *Handler) UpdatePayment(ctx context.Context) error {
	info, err := h.GetPayment(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid payment")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler:       h,
		obseleteState: info.ObseleteState,
	}
	if err := handler.validateObseleteState(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if _, err := paymentbasecrud.UpdateSet(
			tx.PaymentBase.UpdateOneID(*h.ID),
			&h.Req,
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		for _, paymentTransfer := range h.PaymentTransferReqs {
			if _, err := tx.
				PaymentTransfer.
				Update().
				Where(
					entpaymenttransfer.EntID(*paymentTransfer.EntID),
					entpaymenttransfer.DeletedAt(0),
				).
				SetFinishAmount(*paymentTransfer.FinishAmount).
				Save(_ctx); err != nil {
				return wlog.WrapError(err)
			}
		}
		return nil
	})
}
