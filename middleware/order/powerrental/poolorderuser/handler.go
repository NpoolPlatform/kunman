package poolorderuser

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	poolorderusermiddlewarepb "github.com/NpoolPlatform/kunman/message/order/middleware/v1/powerrental/poolorderuser"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	poolorderusercrud "github.com/NpoolPlatform/kunman/middleware/order/crud/powerrental/poolorderuser"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	poolorderusercrud.Req
	Conds  *poolorderusercrud.Conds
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

func WithReq(req *poolorderusercrud.Req, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if req == nil {
			if must {
				return wlog.Errorf("invalid req")
			}
			return nil
		}
		h.Req = *req
		return nil
	}
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid id")
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

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = &_id
		return nil
	}
}

func WithOrderUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid pool orderuserid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PoolOrderUserID = &_id
		return nil
	}
}

func WithConds(conds *poolorderusermiddlewarepb.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &poolorderusercrud.Conds{}
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
		if conds.OrderID != nil {
			id, err := uuid.Parse(conds.GetOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.OrderID = &cruder.Cond{
				Op:  conds.GetOrderID().GetOp(),
				Val: id,
			}
		}
		if conds.PoolOrderUserID != nil {
			id, err := uuid.Parse(conds.GetPoolOrderUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.PoolOrderUserID = &cruder.Cond{
				Op:  conds.GetPoolOrderUserID().GetOp(),
				Val: id,
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
