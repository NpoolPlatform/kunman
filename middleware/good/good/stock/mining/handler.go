package mining

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	stockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock"
	mininggoodstockcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/stock/mining"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	mininggoodstockcrud.Req
	StockReq *stockcrud.Req
	Rollback *bool
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
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
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

func WithGoodStockID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid goodstockid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodStockID = &_id
		return nil
	}
}

func WithPoolRootUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid poolrootuserid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PoolRootUserID = &_id
		return nil
	}
}

func WithPoolGoodUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid poolgooduserid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PoolGoodUserID = &_id
		return nil
	}
}

func WithTotal(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid total")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return wlog.Errorf("invalid total")
		}
		h.Total = &amount
		return nil
	}
}

func WithState(e *types.MiningGoodStockState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid mininggoodstockstate")
			}
			return nil
		}
		switch *e {
		case types.MiningGoodStockState_MiningGoodStockStatePreWait:
		case types.MiningGoodStockState_MiningGoodStockStateWait:
		case types.MiningGoodStockState_MiningGoodStockStateCreateGoodUser:
		case types.MiningGoodStockState_MiningGoodStockStateCheckHashRate:
		case types.MiningGoodStockState_MiningGoodStockStateReady:
		case types.MiningGoodStockState_MiningGoodStockStateFail:
		default:
			return wlog.Errorf("invalid mininggoodstockstate")
		}
		h.State = e
		return nil
	}
}
