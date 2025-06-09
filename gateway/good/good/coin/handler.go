package goodcoin

import (
	"context"
	"fmt"

	goodgwcommon "github.com/NpoolPlatform/kunman/pkg/common"
	constant "github.com/NpoolPlatform/good-gateway/pkg/const"
	goodcommon "github.com/NpoolPlatform/good-gateway/pkg/good/common"

	"github.com/google/uuid"
)

type Handler struct {
	ID    *uint32
	EntID *string
	goodcommon.GoodCheckHandler
	goodgwcommon.CoinCheckHandler
	Main   *bool
	Index  *int32
	Offset int32
	Limit  int32
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
				return fmt.Errorf("invalid entid")
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

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid goodid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.GoodID = id
		return nil
	}
}

func WithCoinTypeID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
			return nil
		}
		if err := h.CheckCoinWithCoinTypeID(ctx, *s); err != nil {
			return err
		}
		h.CoinTypeID = s
		return nil
	}
}

func WithMain(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Main = b
		return nil
	}
}

func WithIndex(n *int32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Index = n
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
