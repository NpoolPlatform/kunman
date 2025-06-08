package currency

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	currencycrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/currency"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/currency"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID              *uint32
	EntID           *uuid.UUID
	CoinTypeID      *uuid.UUID
	FeedType        *basetypes.CurrencyFeedType
	MarketValueHigh *decimal.Decimal
	MarketValueLow  *decimal.Decimal
	Reqs            []*currencycrud.Req
	Conds           *currencycrud.Conds
	Offset          int32
	Limit           int32
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

func WithReqs(reqs []*npool.CurrencyReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*currencycrud.Req{}
		for _, req := range reqs {
			_req := &currencycrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &id
			}
			if req.CoinTypeID != nil {
				id, err := uuid.Parse(*req.CoinTypeID)
				if err != nil {
					return err
				}
				_req.CoinTypeID = &id
			}
			if req.FeedType != nil {
				switch *req.FeedType {
				case basetypes.CurrencyFeedType_CoinGecko:
				case basetypes.CurrencyFeedType_CoinBase:
				case basetypes.CurrencyFeedType_StableUSDHardCode:
				default:
					return fmt.Errorf("invalid feedtype")
				}
				_req.FeedType = req.FeedType
			}
			if req.MarketValueHigh != nil {
				amount, err := decimal.NewFromString(*req.MarketValueHigh)
				if err != nil {
					return err
				}
				_req.MarketValueHigh = &amount
			}
			if req.MarketValueLow != nil {
				amount, err := decimal.NewFromString(*req.MarketValueLow)
				if err != nil {
					return err
				}
				_req.MarketValueLow = &amount
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &currencycrud.Conds{}
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
