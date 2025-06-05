package subscription

import (
	"context"

	constant "github.com/NpoolPlatform/kunman/middleware/billing/const"
	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/billing/crud/subscription"
	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"
	npool "github.com/NpoolPlatform/kunman/message/billing/mw/v1/subscription"
	"github.com/shopspring/decimal"

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

func WithPackageName(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil || *s == "" {
			if must {
				return wlog.Errorf("invalid packagename")
			}
			return nil
		}
		h.PackageName = s
		return nil
	}
}

func WithUsdPrice(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid price")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.UsdPrice = &amount

		return nil
	}
}

func WithDescription(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil || *s == "" {
			if must {
				return wlog.Errorf("invalid description")
			}
			return nil
		}
		h.Description = s
		return nil
	}
}

func WithSortOrder(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid sortorder")
			}
			return nil
		}
		h.SortOrder = u
		return nil
	}
}

func WithPackageType(t *types.PackageType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if t == nil {
			if must {
				return wlog.Errorf("invalid packagetype")
			}
			return nil
		}

		switch *t {
		case types.PackageType_Normal:
		case types.PackageType_Senior:
		default:
			return wlog.Errorf("invalid packagetype")
		}

		h.PackageType = t
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

func WithResetType(t *types.ResetType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if t == nil {
			if must {
				return wlog.Errorf("invalid packagetype")
			}
			return nil
		}

		switch *t {
		case types.ResetType_Weekly:
		case types.ResetType_Monthly:
		case types.ResetType_Quarterly:
		case types.ResetType_Semiyearly:
		case types.ResetType_Yearly:
		default:
			return wlog.Errorf("invalid packagetype")
		}

		h.ResetType = t
		return nil
	}
}

func WithQPSLimit(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid qpslimit")
			}
			return nil
		}
		h.QPSLimit = u
		return nil
	}
}

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
	if conds.PackageName != nil {
		h.SubscriptionConds.PackageName = &cruder.Cond{
			Op:  conds.GetPackageName().GetOp(),
			Val: conds.GetPackageName().GetValue(),
		}
	}
	if conds.SortOrder != nil {
		h.SubscriptionConds.SortOrder = &cruder.Cond{
			Op:  conds.GetSortOrder().GetOp(),
			Val: conds.GetSortOrder().GetValue(),
		}
	}
	if conds.PackageType != nil {
		h.SubscriptionConds.PackageType = &cruder.Cond{
			Op:  conds.GetPackageType().GetOp(),
			Val: types.PackageType(conds.GetPackageType().GetValue()),
		}
	}
	if conds.ResetType != nil {
		h.SubscriptionConds.ResetType = &cruder.Cond{
			Op:  conds.GetResetType().GetOp(),
			Val: types.ResetType(conds.GetResetType().GetValue()),
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
