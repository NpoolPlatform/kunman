package coupon

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/coupon"
	couponcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/coupon"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	couponcrud.Req
	Conds  *couponcrud.Conds
	Offset int32
	Limit  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
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

func WithCouponType(couponType *types.CouponType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if couponType == nil {
			if must {
				return wlog.Errorf("invalid coupontype")
			}
			return nil
		}
		switch *couponType {
		case types.CouponType_FixAmount:
		case types.CouponType_Discount:
		default:
			return wlog.Errorf("invalid coupontype")
		}
		h.CouponType = couponType
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

func WithDenomination(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid denomination")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.Denomination = &_amount
		return nil
	}
}

func WithCirculation(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid circulation")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.Circulation = &_amount
		return nil
	}
}

func WithIssuedBy(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid issuedby")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.IssuedBy = &_id
		return nil
	}
}

func WithStartAt(startAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if startAt == nil {
			if must {
				return wlog.Errorf("invalid startat")
			}
			return nil
		}
		if *startAt == 0 {
			return wlog.Errorf("invalid startat")
		}
		h.StartAt = startAt
		return nil
	}
}

func WithEndAt(endAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if endAt == nil {
			if must {
				return wlog.Errorf("invalid endat")
			}
			return nil
		}
		if *endAt == 0 {
			return wlog.Errorf("invalid endat")
		}
		h.EndAt = endAt
		return nil
	}
}

func WithDurationDays(days *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if days == nil {
			if must {
				return wlog.Errorf("invalid durationdays")
			}
			return nil
		}
		h.DurationDays = days
		return nil
	}
}

func WithMessage(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid message")
			}
			return nil
		}
		if *value == "" {
			return wlog.Errorf("invalid message")
		}
		h.Message = value
		return nil
	}
}

func WithName(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid name")
			}
			return nil
		}
		if *value == "" {
			return wlog.Errorf("invalid name")
		}
		h.Name = value
		return nil
	}
}

func WithThreshold(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid threshold")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.Threshold = &_amount
		return nil
	}
}

func WithAllocated(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return wlog.Errorf("invalid allocated")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.Allocated = &_amount
		return nil
	}
}

func WithCouponConstraint(constraint *types.CouponConstraint, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if constraint == nil {
			if must {
				return wlog.Errorf("invalid constraint")
			}
			return nil
		}
		switch *constraint {
		case types.CouponConstraint_Normal:
		case types.CouponConstraint_PaymentThreshold:
		default:
			return wlog.Errorf("invalid constraint")
		}
		h.CouponConstraint = constraint
		return nil
	}
}

func WithCouponScope(couponScope *types.CouponScope, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if couponScope == nil {
			if must {
				return wlog.Errorf("invalid couponscope")
			}
			return nil
		}
		switch *couponScope {
		case types.CouponScope_AllGood:
		case types.CouponScope_Blacklist:
		case types.CouponScope_Whitelist:
		default:
			return wlog.Errorf("invalid couponscope")
		}
		h.CouponScope = couponScope
		return nil
	}
}

func WithRandom(random *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if random == nil {
			if must {
				return wlog.Errorf("invalid random")
			}
			return nil
		}
		h.Random = random
		return nil
	}
}

func WithCashableProbability(probability *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if probability == nil {
			if must {
				return wlog.Errorf("invalid probability")
			}
			return nil
		}
		_probability, err := decimal.NewFromString(*probability)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _probability.Cmp(decimal.NewFromInt(0)) < 0 || _probability.Cmp(decimal.NewFromInt(1)) > 0 {
			return wlog.Errorf("invalid probability")
		}
		h.CashableProbability = &_probability
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &couponcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.EntID = &cruder.Cond{
				Op: conds.GetEntID().GetOp(), Val: id,
			}
		}
		if conds.CouponType != nil {
			h.Conds.CouponType = &cruder.Cond{
				Op: conds.GetCouponType().GetOp(), Val: types.CouponType(conds.GetCouponType().GetValue()),
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.AppID = &cruder.Cond{
				Op: conds.GetAppID().GetOp(), Val: id,
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
			h.Conds.EntIDs = &cruder.Cond{
				Op: conds.GetEntIDs().GetOp(), Val: ids,
			}
		}
		return nil
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}
