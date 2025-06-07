package payment

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/payment"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	paymentbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment"
	paymenttransfercrud "github.com/NpoolPlatform/kunman/middleware/order/crud/payment/transfer"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	paymentbasecrud.Req
	PaymentBaseConds    *paymentbasecrud.Conds
	PaymentTransferReqs []*paymenttransfercrud.Req
	Offset              int32
	Limit               int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		PaymentBaseConds: &paymentbasecrud.Conds{},
	}
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

func WithObseleteState(e *types.PaymentObseleteState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid obseletestate")
			}
			return nil
		}
		switch *e {
		case types.PaymentObseleteState_PaymentObseleteNone:
		case types.PaymentObseleteState_PaymentObseleteWait:
		case types.PaymentObseleteState_PaymentObseleteUnlockBalance:
		case types.PaymentObseleteState_PaymentObseleteTransferBookKeeping:
		case types.PaymentObseleteState_PaymentObseleteTransferUnlockAccount:
		case types.PaymentObseleteState_PaymentObseleted:
		case types.PaymentObseleteState_PaymentObseleteFail:
		default:
			return wlog.Errorf("invalid obseletestate")
		}
		h.ObseleteState = e
		return nil
	}
}

func WithPaymentTransfers(bs []*npool.PaymentTransferReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, b := range bs {
			req := &paymenttransfercrud.Req{}

			id, err := uuid.Parse(b.GetEntID())
			if err != nil {
				return wlog.WrapError(err)
			}
			req.EntID = &id

			finishAmount, err := decimal.NewFromString(b.GetFinishAmount())
			if err != nil {
				return wlog.WrapError(err)
			}
			req.FinishAmount = &finishAmount

			h.PaymentTransferReqs = append(h.PaymentTransferReqs, req)
		}
		return nil
	}
}

func (h *Handler) withPaymentConds(conds *npool.Conds) {
	if conds.ObseleteState != nil {
		h.PaymentBaseConds.ObseleteState = &cruder.Cond{
			Op:  conds.GetObseleteState().GetOp(),
			Val: types.PaymentObseleteState(conds.GetObseleteState().GetValue()),
		}
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		h.withPaymentConds(conds)
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
