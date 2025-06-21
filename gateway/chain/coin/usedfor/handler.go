package coinusedfor

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/chain/v1"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *string
	CoinTypeID  *string
	CoinTypeIDs []string
	UsedFor     *types.CoinUsedFor
	UsedFors    []types.CoinUsedFor
	Priority    *uint32
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

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
			return nil
		}

		handler, err := coinmw.NewHandler(
			ctx,
			coinmw.WithEntID(id, true),
		)
		if err != nil {
			return err
		}

		_coin, err := handler.GetCoin(ctx)
		if err != nil {
			return err
		}
		if _coin == nil {
			return fmt.Errorf("invalid coin")
		}
		h.CoinTypeID = id
		return nil
	}
}

func WithCoinTypeIDs(ids []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, id := range ids {
			if _, err := uuid.Parse(id); err != nil {
				return err
			}
		}
		h.CoinTypeIDs = ids
		return nil
	}
}

func WithUsedFor(usedFor *types.CoinUsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if usedFor == nil {
			if must {
				return fmt.Errorf("invalid usedfor")
			}
			return nil
		}
		switch *usedFor {
		case types.CoinUsedFor_CoinUsedForCouponCash:
		case types.CoinUsedFor_CoinUsedForGoodFee:
		default:
			return fmt.Errorf("invalid usedfor")
		}
		h.UsedFor = usedFor
		return nil
	}
}

func WithUsedFors(usedFors []types.CoinUsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, usedFor := range usedFors {
			switch usedFor {
			case types.CoinUsedFor_CoinUsedForCouponCash:
			case types.CoinUsedFor_CoinUsedForGoodFee:
			default:
				return fmt.Errorf("invalid usedfors")
			}
		}
		h.UsedFors = usedFors
		return nil
	}
}

func WithPriority(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Priority = n
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
