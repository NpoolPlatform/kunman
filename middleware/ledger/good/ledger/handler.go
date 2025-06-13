package ledger

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/good/ledger"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	crud.Req
	Reqs   []*crud.Req
	Conds  *crud.Conds
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithGoodID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.GoodID = &_id
		return nil
	}
}

func WithCoinTypeID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithAmount(amount *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("amount is less than 0 %v", *amount)
		}
		h.Amount = &_amount
		return nil
	}
}

func WithToPlatform(toPlatform *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if toPlatform == nil {
			return nil
		}
		_toPlatform, err := decimal.NewFromString(*toPlatform)
		if err != nil {
			return err
		}
		if _toPlatform.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("to platform is less than 0 %v", *toPlatform)
		}
		h.ToPlatform = &_toPlatform
		return nil
	}
}

func WithToUser(toUser *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if toUser == nil {
			return nil
		}
		_toUser, err := decimal.NewFromString(*toUser)
		if err != nil {
			return err
		}
		if _toUser.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("to user is less than 0 %v", *toUser)
		}
		h.ToUser = &_toUser
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
