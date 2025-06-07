package orderlock

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/middleware/v1/order/lock"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	orderlockcrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/lock"
	orderbasecrud "github.com/NpoolPlatform/kunman/middleware/order/crud/order/orderbase"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	orderlockcrud.Req
	OrderLockConds *orderlockcrud.Conds
	OrderBaseConds *orderbasecrud.Conds
	Offset         int32
	Limit          int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		OrderLockConds: &orderlockcrud.Conds{},
		OrderBaseConds: &orderbasecrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
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

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.UserID = &_id
		return nil
	}
}

func WithLockType(lockType *types.OrderLockType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lockType == nil {
			if must {
				return wlog.Errorf("invalid locktype")
			}
			return nil
		}
		switch *lockType {
		case types.OrderLockType_LockStock:
		case types.OrderLockType_LockBalance:
		case types.OrderLockType_LockCommission:
		default:
			return wlog.Errorf("invalid locktype")
		}
		h.LockType = lockType
		return nil
	}
}

func (h *Handler) withOrderBaseConds(conds *npool.Conds) error {
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
	}
	if conds.OrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetOrderIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderBaseConds.EntIDs = &cruder.Cond{Op: conds.GetOrderIDs().GetOp(), Val: ids}
	}
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.EntID = &cruder.Cond{Op: conds.GetOrderID().GetOp(), Val: id}
	}
	return nil
}

//nolint:gocyclo
func (h *Handler) withOrderLockConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.OrderLockConds.ID = &cruder.Cond{
			Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderLockConds.EntID = &cruder.Cond{
			Op: conds.GetEntID().GetOp(), Val: id,
		}
	}
	if conds.IDs != nil {
		h.OrderLockConds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: conds.GetIDs().GetValue()}
	}
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderLockConds.OrderID = &cruder.Cond{Op: conds.GetOrderID().GetOp(), Val: id}
	}
	if conds.LockType != nil {
		switch conds.GetLockType().GetValue() {
		case uint32(types.OrderLockType_LockBalance):
		case uint32(types.OrderLockType_LockStock):
		case uint32(types.OrderLockType_LockCommission):
		default:
			return wlog.Errorf("invalid locktype")
		}
		_type := conds.GetLockType().GetValue()
		h.OrderLockConds.LockType = &cruder.Cond{Op: conds.GetLockType().GetOp(), Val: types.OrderLockType(_type)}
	}
	if conds.OrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetOrderIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderLockConds.OrderIDs = &cruder.Cond{Op: conds.GetOrderIDs().GetOp(), Val: ids}
	}
	if conds.UserID != nil {
		id, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderLockConds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
	}
	if conds.UserIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetUserIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderLockConds.UserIDs = &cruder.Cond{Op: conds.GetUserIDs().GetOp(), Val: ids}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withOrderBaseConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withOrderLockConds(conds)
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
