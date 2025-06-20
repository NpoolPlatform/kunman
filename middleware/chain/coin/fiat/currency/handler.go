package currency

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	currencycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/fiat/currency"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID              *uint32
	EntID           *uuid.UUID
	CoinTypeID      *uuid.UUID
	FiatID          *uuid.UUID
	FeedType        *basetypes.CurrencyFeedType
	MarketValueHigh *decimal.Decimal
	MarketValueLow  *decimal.Decimal
	Conds           *currencycrud.Conds
	Offset          int32
	Limit           int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		Conds: &currencycrud.Conds{},
	}
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

func WithMarketValueHigh(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid marketvaluehigh")
			}
			return nil
		}
		_value, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.MarketValueHigh = &_value
		return nil
	}
}

func WithMarketValueLow(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid marketvaluelow")
			}
			return nil
		}
		_value, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.MarketValueLow = &_value
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
