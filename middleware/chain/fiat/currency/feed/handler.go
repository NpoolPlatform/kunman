package currencyfeed

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat/currency/feed"
	feedcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/fiat/currency/feed"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID           *uint32
	EntID        *uuid.UUID
	FiatID       *uuid.UUID
	FeedType     *basetypes.CurrencyFeedType
	FeedFiatName *string
	Disabled     *bool
	Conds        *feedcrud.Conds
	Offset       int32
	Limit        int32
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

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
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

func WithFiatID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid fiatid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.FiatID = &_id
		return nil
	}
}

func WithFeedType(feedType *basetypes.CurrencyFeedType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if feedType == nil {
			if must {
				return fmt.Errorf("invalid feedtype")
			}
			return nil
		}
		switch *feedType {
		case basetypes.CurrencyFeedType_CoinGecko:
		case basetypes.CurrencyFeedType_CoinBase:
		case basetypes.CurrencyFeedType_StableUSDHardCode:
		default:
			return fmt.Errorf("invalid feedtype")
		}
		h.FeedType = feedType
		return nil
	}
}

func WithFeedFiatName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid feedfiatname")
		}
		h.FeedFiatName = name
		return nil
	}
}

func WithDisabled(disabled *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Disabled = disabled
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &feedcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.FiatID != nil {
			id, err := uuid.Parse(conds.GetFiatID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.FiatID = &cruder.Cond{
				Op:  conds.GetFiatID().GetOp(),
				Val: id,
			}
		}
		if conds.FiatIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetFiatIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.FiatIDs = &cruder.Cond{
				Op:  conds.GetFiatIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.Disabled != nil {
			h.Conds.Disabled = &cruder.Cond{
				Op:  conds.GetDisabled().GetOp(),
				Val: conds.GetDisabled().GetValue(),
			}
		}
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
