package addon

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	types "github.com/NpoolPlatform/kunman/message/basetypes/billing/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID          *uint32
	EntID       *string
	AppID       *string
	PackageName *string
	UsdPrice    *string
	Description *string
	SortOrder   *uint32
	Enabled     *bool
	PackageType *types.PackageType
	Credit      *uint32
	ResetType   *types.ResetType
	QPSLimit    *uint32
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
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.AppID = id
		return nil
	}
}

func WithPackageName(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid packagename")
			}
			return nil
		}
		if len(*s) < 3 {
			return fmt.Errorf("invalid packagename")
		}
		h.PackageName = s
		return nil
	}
}

func WithUsdPrice(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid price")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
			return fmt.Errorf("invalid price")
		}
		h.UsdPrice = s
		return nil
	}
}

func WithDescription(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid description")
			}
			return nil
		}
		if len(*s) < 3 {
			return fmt.Errorf("invalid description")
		}
		h.Description = s
		return nil
	}
}

func WithSortOrder(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid sortorder")
			}
			return nil
		}
		h.SortOrder = n
		return nil
	}
}

func WithPackageType(e *types.PackageType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid packagetype")
			}
			return nil
		}
		switch *e {
		case types.PackageType_Normal:
		case types.PackageType_Senior:
		default:
			return fmt.Errorf("invalid packagetype")
		}
		h.PackageType = e
		return nil
	}
}

func WithCredit(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid credit")
			}
			return nil
		}
		h.Credit = n
		return nil
	}
}

func WithEnabled(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if b == nil {
			if must {
				return fmt.Errorf("invalid enabled")
			}
			return nil
		}
		h.Enabled = b
		return nil
	}
}

func WithResetType(e *types.ResetType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid resettype")
			}
			return nil
		}
		switch *e {
		case types.ResetType_Weekly:
		case types.ResetType_Monthly:
		case types.ResetType_Quarterly:
		case types.ResetType_Semiyearly:
		case types.ResetType_Yearly:
		default:
			return fmt.Errorf("invalid resettype")
		}
		h.ResetType = e
		return nil
	}
}

func WithQPSLimit(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid qpslimit")
			}
			return nil
		}
		h.QPSLimit = n
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
