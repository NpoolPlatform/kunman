package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/subscription"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	subscriptioncrud "github.com/NpoolPlatform/kunman/middleware/good/crud/subscription"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	subscriptioncrud.Req
	GoodBaseReq       *goodbasecrud.Req
	SubscriptionConds *subscriptioncrud.Conds
	GoodBaseConds     *goodbasecrud.Conds
	Offset            int32
	Limit             int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		GoodBaseReq: &goodbasecrud.Req{
			BenefitType: types.BenefitType_BenefitTypeNone.Enum(),
			StartMode:   types.GoodStartMode_GoodStartModeInstantly.Enum(),
			State:       types.GoodState_GoodStateReady.Enum(),
			GoodType:    types.GoodType_Subscription.Enum(),
		},
		SubscriptionConds: &subscriptioncrud.Conds{},
		GoodBaseConds:     &goodbasecrud.Conds{},
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

func WithEntID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &id
		return nil
	}
}

func WithGoodID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = &id
		h.GoodBaseReq.EntID = &id
		return nil
	}
}

func WithDurationDisplayType(e *types.GoodDurationType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid durationtype")
			}
			return nil
		}
		switch *e {
		case types.GoodDurationType_GoodDurationByHour:
		case types.GoodDurationType_GoodDurationByDay:
		case types.GoodDurationType_GoodDurationByMonth:
		case types.GoodDurationType_GoodDurationByYear:
		default:
			return wlog.Errorf("invalid durationtype")
		}
		h.DurationDisplayType = e
		return nil
	}
}

func WithGoodType(e *types.GoodType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid goodtype")
			}
			return nil
		}
		switch *e {
		case types.GoodType_Subscription:
		default:
			return wlog.Errorf("invalid goodtype")
		}
		h.GoodBaseReq.GoodType = e
		return nil
	}
}

func WithName(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid name")
			}
			return nil
		}
		if *s == "" {
			return wlog.Errorf("invalid name")
		}
		h.GoodBaseReq.Name = s
		return nil
	}
}

func WithDurationUnits(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid durationunits")
			}
			return nil
		}
		h.DurationUnits = u
		return nil
	}
}

func WithDurationQuota(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil || *u == 0 {
			if must {
				return wlog.Errorf("invalid durationquota")
			}
			return nil
		}
		h.DurationQuota = u
		return nil
	}
}

func WithDailyBonusQuota(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid dailybonusquota")
			}
			return nil
		}
		h.DailyBonusQuota = u
		return nil
	}
}

func WithUSDPrice(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid usdprice")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return wlog.Errorf("invalid usdprice")
		}
		h.USDPrice = &amount
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
	if conds.IDs != nil {
		h.SubscriptionConds.IDs = &cruder.Cond{
			Op:  conds.GetIDs().GetOp(),
			Val: conds.GetIDs().GetValue(),
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
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.SubscriptionConds.GoodID = &cruder.Cond{
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
		h.SubscriptionConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withGoodBaseConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodBaseConds.EntID = &cruder.Cond{
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
		h.GoodBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
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
		if err := h.withSubscriptionConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withGoodBaseConds(conds); err != nil {
			return wlog.WrapError(err)
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
