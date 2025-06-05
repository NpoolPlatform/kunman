package exchange

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/middleware/billing/const"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"

	"github.com/google/uuid"
)

type Handler struct {
	ID                *uint32
	EntID             *string
	AppID             *string
	UsageType         *types.UsageType
	Credit            *uint32
	ExchangeThreshold *uint32
	Path              *string
	Offset            int32
	Limit             int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
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
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.EntID = id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.AppID = id
		return nil
	}
}

func WithUsageType(e *types.UsageType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid usagetype")
			}
			return nil
		}
		switch *e {
		case types.UsageType_TextToken:
		case types.UsageType_ChatCount:
		case types.UsageType_ImageCount:
		case types.UsageType_VideoCount:
		case types.UsageType_FilePageCount:
		default:
			return fmt.Errorf("invalid usagetype")
		}
		h.UsageType = e
		return nil
	}
}

func WithCredit(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid credit")
			}
			return nil
		}
		h.Credit = n
		return nil
	}
}

func WithPath(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid credit")
			}
			return nil
		}
		h.Path = s
		return nil
	}
}

func WithExchangeThreshold(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid exchangethreshold")
			}
			return nil
		}
		h.ExchangeThreshold = n
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
