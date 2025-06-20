package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/fee"
	feecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/fee"
	goodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/goodbase"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	feecrud.Req
	GoodBaseReq   *goodbasecrud.Req
	FeeConds      *feecrud.Conds
	GoodBaseConds *goodbasecrud.Conds
	Offset        int32
	Limit         int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		GoodBaseReq: &goodbasecrud.Req{
			BenefitType: types.BenefitType_BenefitTypeNone.Enum(),
			StartMode:   types.GoodStartMode_GoodStartModeWithParent.Enum(),
			State:       types.GoodState_GoodStateReady.Enum(),
		},
		FeeConds:      &feecrud.Conds{},
		GoodBaseConds: &goodbasecrud.Conds{},
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

func WithSettlementType(e *types.GoodSettlementType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid settlementtype")
			}
			return nil
		}
		switch *e {
		case types.GoodSettlementType_GoodSettledByPaymentPercent:
		case types.GoodSettlementType_GoodSettledByPaymentAmount:
		case types.GoodSettlementType_GoodSettledByProfitPercent:
		default:
			return wlog.Errorf("invalid settlementtype")
		}
		h.SettlementType = e
		return nil
	}
}

func WithUnitValue(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid unitvalue")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
			return wlog.Errorf("invalid unitvalue")
		}
		h.UnitValue = &amount
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
		case types.GoodType_TechniqueServiceFee:
		case types.GoodType_ElectricityFee:
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

func (h *Handler) withFeeConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.FeeConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.IDs != nil {
		h.FeeConds.IDs = &cruder.Cond{
			Op:  conds.GetIDs().GetOp(),
			Val: conds.GetIDs().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.FeeConds.EntID = &cruder.Cond{
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
		h.FeeConds.EntIDs = &cruder.Cond{
			Op:  conds.GetEntIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.FeeConds.GoodID = &cruder.Cond{
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
		h.FeeConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.SettlementType != nil {
		h.FeeConds.SettlementType = &cruder.Cond{
			Op:  conds.GetSettlementType().GetOp(),
			Val: types.GoodSettlementType(conds.GetSettlementType().GetValue()),
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
		if err := h.withFeeConds(conds); err != nil {
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
