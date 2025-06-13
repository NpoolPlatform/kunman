package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	ordercommon "github.com/NpoolPlatform/kunman/gateway/order/order/common"

	"github.com/google/uuid"
)

type Handler struct {
	ID    *uint32
	EntID *string
	ordercommon.OrderCheckHandler
	CompensateFromID *string
	CompensateType   *types.CompensateType
	Offset           int32
	Limit            int32
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
		if _, err := uuid.Parse(*id); err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = id
		return nil
	}
}

func WithAppID(appID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appID == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		if err := h.CheckAppWithAppID(ctx, *appID); err != nil {
			return wlog.WrapError(err)
		}
		h.AppID = appID
		return nil
	}
}

func WithUserID(userID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			if must {
				return wlog.Errorf("invalid userid")
			}
			return nil
		}
		if err := h.CheckUserWithUserID(ctx, *userID); err != nil {
			return wlog.WrapError(err)
		}
		h.UserID = userID
		return nil
	}
}

func WithGoodID(goodID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if goodID == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		if err := h.CheckGoodWithGoodID(ctx, *goodID); err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = goodID
		return nil
	}
}

func WithOrderID(orderID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if orderID == nil {
			if must {
				return wlog.Errorf("invalid orderid")
			}
			return nil
		}
		if err := h.CheckOrderWithOrderID(ctx, *orderID); err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = orderID
		return nil
	}
}

func WithCompensateFromID(compensateFromID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if compensateFromID == nil {
			if must {
				return wlog.Errorf("invalid compensateFromid")
			}
			return nil
		}
		if _, err := uuid.Parse(*compensateFromID); err != nil {
			return wlog.WrapError(err)
		}
		h.CompensateFromID = compensateFromID
		return nil
	}
}

func WithCompensateType(e *types.CompensateType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid compensatetype")
			}
			return nil
		}
		switch *e {
		case types.CompensateType_CompensateMalfunction:
		case types.CompensateType_CompensateWalfare:
		case types.CompensateType_CompensateStarterDelay:
		default:
			return wlog.Errorf("invalid compensatetype")
		}
		h.CompensateType = e
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
