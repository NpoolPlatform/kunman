package subscription

import (
	"context"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/agi/crud/user/subscription"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/agi/v1"
	npool "github.com/NpoolPlatform/kunman/message/agi/middleware/v1/user/subscription"

	"github.com/google/uuid"
)

type Handler struct {
	subscriptioncrud.Req
	SubscriptionConds *subscriptioncrud.Conds
	Offset            int32
	Limit             int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		SubscriptionConds: &subscriptioncrud.Conds{},
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

func WithPackageID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid package id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PackageID = &_id
		return nil
	}
}

func WithStartAt(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid StartAt")
			}
			return nil
		}
		h.StartAt = u
		return nil
	}
}

func WithEndAt(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid EndAt")
			}
			return nil
		}
		h.EndAt = u
		return nil
	}
}

func WithUsageState(t *types.UsageState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if t == nil {
			if must {
				return wlog.Errorf("invalid usagestate")
			}
			return nil
		}

		switch *t {
		case types.UsageState_Usful:
		case types.UsageState_Expire:
		case types.UsageState_Disable:
		default:
			return wlog.Errorf("invalid usagestate")
		}

		h.UsageState = t
		return nil
	}
}

func WithSubscriptionCredit(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid subscriptioncredit")
			}
			return nil
		}
		h.SubscriptionCredit = u
		return nil
	}
}

func WithAddonCredit(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid addoncredit")
			}
			return nil
		}
		h.AddonCredit = u
		return nil
	}
}

//nolint:gocyclo
func (h *Handler) withSubscriptionConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.SubscriptionConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.SubscriptionConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.SubscriptionConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	if conds.UserID != nil {
		id, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.SubscriptionConds.UserID = &cruder.Cond{
			Op:  conds.GetUserID().GetOp(),
			Val: id,
		}
	}
	if conds.PackageID != nil {
		id, err := uuid.Parse(conds.GetPackageID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.SubscriptionConds.PackageID = &cruder.Cond{
			Op:  conds.GetPackageID().GetOp(),
			Val: id,
		}
	}
	if conds.StartAt != nil {
		h.SubscriptionConds.StartAt = &cruder.Cond{
			Op:  conds.GetStartAt().GetOp(),
			Val: conds.GetStartAt().GetValue(),
		}
	}
	if conds.EndAt != nil {
		h.SubscriptionConds.EndAt = &cruder.Cond{
			Op:  conds.GetEndAt().GetOp(),
			Val: conds.GetEndAt().GetValue(),
		}
	}
	if conds.UsageState != nil {
		h.SubscriptionConds.UsageState = &cruder.Cond{
			Op:  conds.GetUsageState().GetOp(),
			Val: types.UsageState(conds.GetUsageState().GetValue()),
		}
	}
	if conds.IDs != nil {
		h.SubscriptionConds.IDs = &cruder.Cond{
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
		h.SubscriptionConds.EntIDs = &cruder.Cond{
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
		return h.withSubscriptionConds(conds)
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
