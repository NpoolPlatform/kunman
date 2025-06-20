package constraint

import (
	"context"
	"fmt"

	topmostcommon "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/common"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID    *uint32
	EntID *string
	topmostcommon.CheckHandler
	Constraint  *types.GoodTopMostConstraint
	TargetValue *string
	Index       *uint32
	Offset      int32
	Limit       int32
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
				return fmt.Errorf("invalid id")
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
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.EntID = id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		if err := h.CheckAppWithAppID(ctx, *id); err != nil {
			return err
		}
		h.AppID = id
		return nil
	}
}

func WithTopMostID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid topmostid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.TopMostID = id
		return nil
	}
}

func WithConstraint(e *types.GoodTopMostConstraint, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid constraint")
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
			return fmt.Errorf("invalid constraint")
		}
		h.Constraint = e
		return nil
	}
}

func WithIndex(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Index = n
		return nil
	}
}

func WithTargetValue(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid targetvalue")
			}
			return nil
		}
		if _, err := decimal.NewFromString(*s); err != nil {
			return err
		}
		h.TargetValue = s
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
