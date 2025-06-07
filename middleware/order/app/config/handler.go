//nolint:dupl
package appconfig

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/app/config"
	appconfigcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/app/config"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	appconfigcrud.Req
	AppConfigConds *appconfigcrud.Conds
	Offset         int32
	Limit          int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		AppConfigConds: &appconfigcrud.Conds{},
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppID = &_id
		return nil
	}
}

func WithEnableSimulateOrder(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EnableSimulateOrder = b
		return nil
	}
}

func WithSimulateOrderCouponMode(e *types.SimulateOrderCouponMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid simulateordercouponmode")
			}
			return nil
		}
		switch *e {
		case types.SimulateOrderCouponMode_WithoutCoupon:
		case types.SimulateOrderCouponMode_FirstBenifit:
		case types.SimulateOrderCouponMode_RandomBenifit:
		case types.SimulateOrderCouponMode_FirstAndRandomBenifit:
		default:
			return wlog.Errorf("invalid simulateordercouponmode")
		}
		h.SimulateOrderCouponMode = e
		return nil
	}
}

func WithSimulateOrderCouponProbability(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid simulateordercouponprobability")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid simulateordercouponprobability")
		}
		if amount.Cmp(decimal.NewFromInt(1)) > 0 {
			return wlog.Errorf("invalid simulateordercouponprobability")
		}
		h.SimulateOrderCouponProbability = &amount
		return nil
	}
}

func WithSimulateOrderCashableProfitProbability(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid simulateordercashableprofitprobability")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid simulateordercashableprofitprobability")
		}
		if amount.Cmp(decimal.NewFromInt(1)) > 0 {
			return wlog.Errorf("invalid simulateordercashableprofitprobability")
		}
		h.SimulateOrderCashableProfitProbability = &amount
		return nil
	}
}

func WithMaxUnpaidOrders(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MaxUnpaidOrders = u
		return nil
	}
}

func WithMaxTypedCouponsPerOrder(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MaxTypedCouponsPerOrder = u
		return nil
	}
}

func (h *Handler) withAppConfigConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.AppConfigConds.ID = &cruder.Cond{
			Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppConfigConds.EntID = &cruder.Cond{
			Op: conds.GetEntID().GetOp(), Val: id,
		}
	}
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppConfigConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		return h.withAppConfigConds(conds)
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
