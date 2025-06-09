package topmost

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/good-gateway/pkg/const"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
)

type Handler struct {
	ID    *uint32
	EntID *string
	goodgwcommon.AppUserCheckHandler
	TopMostType *types.GoodTopMostType
	Title       *string
	Message     *string
	TargetURL   *string
	StartAt     *uint32
	EndAt       *uint32
	Offset      int32
	Limit       int32
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
		if err := h.CheckAppWithAppID(ctx, *id); err != nil {
			return err
		}
		h.AppID = id
		return nil
	}
}

func WithTopMostType(e *types.GoodTopMostType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid topmosttype")
			}
			return nil
		}
		switch *e {
		case types.GoodTopMostType_TopMostPromotion:
		case types.GoodTopMostType_TopMostNoviceExclusive:
		case types.GoodTopMostType_TopMostBestOffer:
		case types.GoodTopMostType_TopMostInnovationStarter:
		case types.GoodTopMostType_TopMostLoyaltyExclusive:
		default:
			return fmt.Errorf("invalid topmosttype")
		}
		h.TopMostType = e
		return nil
	}
}

func WithTitle(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Title = s
		return nil
	}
}

func WithMessage(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Message = s
		return nil
	}
}

func WithTargetURL(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.TargetURL = s
		return nil
	}
}

func WithStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.StartAt = n
		return nil
	}
}

func WithEndAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EndAt = n
		return nil
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}
