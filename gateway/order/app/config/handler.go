package appconfig

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	common "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID    *uint32
	EntID *string
	common.AppUserCheckHandler
	Units                                  *string
	Duration                               *uint32
	EnableSimulateOrder                    *bool
	SimulateOrderUnits                     *string
	SimulateOrderDurationSeconds           *uint32
	SimulateOrderCouponMode                *types.SimulateOrderCouponMode
	SimulateOrderCouponProbability         *string
	SimulateOrderCashableProfitProbability *string
	MaxUnpaidOrders                        *uint32
	MaxTypedCouponsPerOrder                *uint32
	Offset                                 int32
	Limit                                  int32
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

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
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
		_, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = id
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
		if err := h.CheckAppWithAppID(ctx, *id); err != nil {
			return wlog.WrapError(err)
		}
		h.AppID = id
		return nil
	}
}

func WithEnableSimulateOrder(enabled *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EnableSimulateOrder = enabled
		return nil
	}
}

func WithSimulateOrderUnits(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid simulateorderunits")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _amount.Cmp(decimal.NewFromInt32(0)) <= 0 {
			return wlog.Errorf("invalid simulateorderunits")
		}
		h.SimulateOrderUnits = amount
		return nil
	}
}

func WithSimulateOrderDurationSeconds(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			if must {
				return wlog.Errorf("invalid simulateorderdurationseconds")
			}
			return nil
		}
		if *duration <= 0 {
			return wlog.Errorf("invalid simulateorderdurationseconds")
		}
		h.SimulateOrderDurationSeconds = duration
		return nil
	}
}

func WithMaxUnpaidOrders(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			if must {
				return wlog.Errorf("invalid maxunpaidorders")
			}
			return nil
		}
		if *duration <= 0 {
			return wlog.Errorf("invalid maxunpaidorders")
		}
		h.MaxUnpaidOrders = duration
		return nil
	}
}

func WithMaxTypedCouponsPerOrder(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			if must {
				return wlog.Errorf("invalid maxtypedcouponsperorder")
			}
			return nil
		}
		if *duration <= 0 {
			return wlog.Errorf("invalid maxtypedcouponsperorder")
		}
		h.MaxTypedCouponsPerOrder = duration
		return nil
	}
}

//nolint:dupl
func WithSimulateOrderCouponProbability(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid simulateordercouponprobability")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid simulateordercouponprobability")
		}
		if _amount.Cmp(decimal.NewFromInt(1)) > 0 {
			return wlog.Errorf("invalid simulateordercouponprobability")
		}
		h.SimulateOrderCouponProbability = amount
		return nil
	}
}

func WithSimulateOrderCouponMode(value *types.SimulateOrderCouponMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid simulateordercouponmode")
			}
			return nil
		}
		switch *value {
		case types.SimulateOrderCouponMode_WithoutCoupon:
		case types.SimulateOrderCouponMode_FirstBenifit:
		case types.SimulateOrderCouponMode_RandomBenifit:
		case types.SimulateOrderCouponMode_FirstAndRandomBenifit:
		default:
			return wlog.Errorf("invalid simulateordercouponmode")
		}
		h.SimulateOrderCouponMode = value
		return nil
	}
}

//nolint:dupl
func WithSimulateOrderCashableProfitProbability(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid simulateordercashableprofitprobability")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid simulateordercashableprofitprobability")
		}
		if _amount.Cmp(decimal.NewFromInt(1)) > 0 {
			return wlog.Errorf("invalid simulateordercashableprofitprobability")
		}
		h.SimulateOrderCashableProfitProbability = amount
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
