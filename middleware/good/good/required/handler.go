//nolint:dupl
package required

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	requiredcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/required"
	goodbase1 "github.com/NpoolPlatform/kunman/middleware/good/good/goodbase"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/required"

	"github.com/google/uuid"
)

type Handler struct {
	requiredcrud.Req
	Conds  *requiredcrud.Conds
	Offset int32
	Limit  int32
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

func WithMainGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := goodbase1.NewHandler(
			ctx,
			goodbase1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistGoodBase(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid good")
		}
		h.MainGoodID = handler.EntID
		return nil
	}
}

func WithRequiredGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := goodbase1.NewHandler(
			ctx,
			goodbase1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistGoodBase(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid good")
		}
		h.RequiredGoodID = handler.EntID
		return nil
	}
}

func WithMust(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Must = b
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &requiredcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
		}
		if conds.MainGoodID != nil {
			id, err := uuid.Parse(conds.GetMainGoodID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.MainGoodID = &cruder.Cond{
				Op:  conds.GetMainGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.RequiredGoodID != nil {
			id, err := uuid.Parse(conds.GetRequiredGoodID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.RequiredGoodID = &cruder.Cond{
				Op:  conds.GetRequiredGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.GoodID = &cruder.Cond{
				Op:  conds.GetGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.GoodIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetGoodIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return wlog.WrapError(err)
				}
				ids = append(ids, _id)
			}
			h.Conds.GoodIDs = &cruder.Cond{
				Op:  conds.GetGoodIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.Must != nil {
			h.Conds.Must = &cruder.Cond{
				Op:  conds.GetMust().GetOp(),
				Val: conds.GetMust().GetValue(),
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
