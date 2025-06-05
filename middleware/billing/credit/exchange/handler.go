package exchange

import (
	"context"

	constant "github.com/NpoolPlatform/kunman/middleware/billing/const"
	exchangecrud "github.com/NpoolPlatform/kunman/middleware/billing/crud/credit/exchange"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"
	npool "github.com/NpoolPlatform/kunman/message/billing/mw/v1/credit/exchange"

	"github.com/google/uuid"
)

type Handler struct {
	exchangecrud.Req
	ExchangeConds *exchangecrud.Conds
	Offset        int32
	Limit         int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		ExchangeConds: &exchangecrud.Conds{},
	}
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppID = &_id
		return nil
	}
}

func WithUsageType(t *types.UsageType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if t == nil {
			if must {
				return wlog.Errorf("invalid usagetype")
			}
			return nil
		}

		switch *t {
		case types.UsageType_TextToken:
		case types.UsageType_ChatCount:
		case types.UsageType_ImageCount:
		case types.UsageType_VideoCount:
		case types.UsageType_FilePageCount:
		default:
			return wlog.Errorf("invalid usagetype")
		}

		h.UsageType = t
		return nil
	}
}

func WithCredit(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid credit")
			}
			return nil
		}
		h.Credit = u
		return nil
	}
}

func WithExchangeThreshold(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid exchangethreshold")
			}
			return nil
		}
		h.ExchangeThreshold = u
		return nil
	}
}

func WithPath(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid path")
			}
			return nil
		}
		h.Path = s
		return nil
	}
}

func (h *Handler) withExchangeConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.ExchangeConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.ExchangeConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.ExchangeConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}

	if conds.Path != nil {
		h.ExchangeConds.Path = &cruder.Cond{
			Op:  conds.GetPath().GetOp(),
			Val: conds.GetPath().GetValue(),
		}
	}

	if conds.UsageType != nil {
		h.ExchangeConds.UsageType = &cruder.Cond{
			Op:  conds.GetUsageType().GetOp(),
			Val: types.UsageType(conds.GetUsageType().GetValue()),
		}
	}

	if conds.IDs != nil {
		h.ExchangeConds.IDs = &cruder.Cond{
			Op:  conds.GetIDs().GetOp(),
			Val: conds.GetIDs().GetValue(),
		}
	}
	if conds.EntIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetEntIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.ExchangeConds.EntIDs = &cruder.Cond{
			Op:  conds.GetEntIDs().GetOp(),
			Val: ids,
		}
	}

	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		return h.withExchangeConds(conds)
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
