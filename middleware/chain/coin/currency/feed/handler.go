package currencyfeed

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency/feed"
	feedcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/currency/feed"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID           *uint32
	EntID        *uuid.UUID
	CoinTypeID   *uuid.UUID
	FeedType     *basetypes.CurrencyFeedType
	FeedCoinName *string
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

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
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

func WithFeedCoinName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid feedcoinname")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid feedcoinname")
		}
		h.FeedCoinName = name
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
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{
				Op:  conds.GetCoinTypeID().GetOp(),
				Val: id,
			}
		}
		if conds.FeedType != nil {
			switch conds.GetFeedType().GetValue() {
			case uint32(basetypes.CurrencyFeedType_CoinGecko):
			case uint32(basetypes.CurrencyFeedType_CoinBase):
			case uint32(basetypes.CurrencyFeedType_StableUSDHardCode):
			default:
				return fmt.Errorf("invalid feedtype")
			}
			_type := conds.GetFeedType().GetValue()
			h.Conds.FeedType = &cruder.Cond{Op: conds.GetFeedType().GetOp(), Val: basetypes.CurrencyFeedType(_type)}
		}
		if conds.CoinTypeIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetCoinTypeIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.CoinTypeIDs = &cruder.Cond{
				Op:  conds.GetCoinTypeIDs().GetOp(),
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
