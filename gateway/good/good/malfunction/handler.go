package malfunction

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	goodcommon "github.com/NpoolPlatform/good-gateway/pkg/good/common"
	constant "github.com/NpoolPlatform/good-middleware/pkg/const"

	"github.com/google/uuid"
)

type Handler struct {
	ID    *uint32
	EntID *string
	goodcommon.GoodCheckHandler
	Title             *string
	Message           *string
	StartAt           *uint32
	DurationSeconds   *uint32
	CompensateSeconds *uint32
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
				return wlog.Errorf("invalid id")
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

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		if err := h.CheckGoodWithGoodID(ctx, *id); err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = id
		return nil
	}
}

func WithTitle(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid title")
			}
			return nil
		}
		if len(*s) < 3 {
			return wlog.Errorf("invalid title")
		}
		h.Title = s
		return nil
	}
}

func WithMessage(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid message")
			}
			return nil
		}
		if len(*s) < 3 {
			return wlog.Errorf("invalid message")
		}
		h.Message = s
		return nil
	}
}

func WithStartAt(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid startat")
			}
			return nil
		}
		if *u == 0 || *u >= uint32(time.Now().Unix()) {
			return wlog.Errorf("invalid startat")
		}
		h.StartAt = u
		return nil
	}
}

func WithDurationSeconds(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid durationseconds")
			}
			return nil
		}
		if *u == 0 || *u >= uint32(time.Now().Unix()) {
			return wlog.Errorf("invalid durationseconds")
		}
		h.DurationSeconds = u
		return nil
	}
}

func WithCompensateSeconds(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid compensateseconds")
			}
			return nil
		}
		if *u == 0 || *u >= uint32(time.Now().Unix()) {
			return wlog.Errorf("invalid compensateseconds")
		}
		h.CompensateSeconds = u
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
