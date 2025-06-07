package constraint

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	constant "github.com/NpoolPlatform/kunman/middleware/good/const"
	topmostcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost"
	constraintcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/topmost/constraint"
	topmost1 "github.com/NpoolPlatform/kunman/middleware/good/middleware/app/good/topmost"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/good/topmost/constraint"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	constraintcrud.Req
	ConstraintConds *constraintcrud.Conds
	TopMostConds    *topmostcrud.Conds
	Offset          int32
	Limit           int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		ConstraintConds: &constraintcrud.Conds{},
		TopMostConds:    &topmostcrud.Conds{},
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

func WithTopMostID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := topmost1.NewHandler(
			ctx,
			topmost1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistTopMost(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid topmost")
		}
		h.TopMostID = handler.EntID
		return nil
	}
}

func WithConstraint(e *types.GoodTopMostConstraint, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid constraint")
			}
			return nil
		}
		switch *e {
		case types.GoodTopMostConstraint_TopMostCreditThreshold:
		case types.GoodTopMostConstraint_TopMostRegisterBefore:
		case types.GoodTopMostConstraint_TopMostOrderThreshold:
		case types.GoodTopMostConstraint_TopMostPaymentAmount:
		case types.GoodTopMostConstraint_TopMostKycMust:
		default:
			return wlog.Errorf("invalid constraint")
		}
		h.Constraint = e
		return nil
	}
}

func WithTargetValue(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid targetvalue")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return wlog.Errorf("invalid targetvalue")
		}
		h.TargetValue = &amount
		return nil
	}
}

func WithIndex(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid index")
			}
			return nil
		}
		h.Index = func() *uint8 { u1 := uint8(*u); return &u1 }()
		return nil
	}
}

func (h *Handler) withConstraintConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.ConstraintConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.ConstraintConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.TopMostID != nil {
		id, err := uuid.Parse(conds.GetTopMostID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.ConstraintConds.TopMostID = &cruder.Cond{
			Op:  conds.GetTopMostID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func (h *Handler) withTopMostConds(conds *npool.Conds) error {
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.TopMostConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	if conds.TopMostID != nil {
		id, err := uuid.Parse(conds.GetTopMostID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.TopMostConds.EntID = &cruder.Cond{
			Op:  conds.GetTopMostID().GetOp(),
			Val: id,
		}
	}
	if conds.TopMostType != nil {
		h.TopMostConds.TopMostType = &cruder.Cond{
			Op:  conds.GetTopMostType().GetOp(),
			Val: types.GoodTopMostType(conds.GetTopMostType().GetValue()),
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withTopMostConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withConstraintConds(conds)
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
