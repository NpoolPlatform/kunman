package chain

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/google/uuid"
)

type Handler struct {
	ID         *uint32
	EntID      *string
	ChainType  *string
	Logo       *string
	ChainID    *string
	NativeUnit *string
	AtomicUnit *string
	UnitExp    *uint32
	GasType    *basetypes.GasType
	Nickname   *string
	ENV        *string
	Offset     int32
	Limit      int32
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
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = id
		return nil
	}
}

func WithLogo(logo *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Logo = logo
		return nil
	}
}

func WithENV(env *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if env == nil {
			if must {
				return fmt.Errorf("invalid coinenv")
			}
			return nil
		}
		switch *env {
		case "main":
		case "test":
		case "local":
		default:
			return fmt.Errorf("invalid coinenv")
		}
		h.ENV = env
		return nil
	}
}

func WithChainType(chainType *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if chainType == nil {
			if must {
				return fmt.Errorf("invalid chaintype")
			}
			return nil
		}
		if *chainType == "" {
			return fmt.Errorf("invalid chaintype")
		}
		h.ChainType = chainType
		return nil
	}
}

func WithNativeUnit(unit *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if unit == nil {
			if must {
				return fmt.Errorf("invalid nativeunit")
			}
			return nil
		}
		if *unit == "" {
			return fmt.Errorf("invalid nativeunit")
		}
		h.NativeUnit = unit
		return nil
	}
}

func WithAtomicUnit(unit *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if unit == nil {
			if must {
				return fmt.Errorf("invalid atomicunit")
			}
			return nil
		}
		if *unit == "" {
			return fmt.Errorf("invalid atomicunit")
		}
		h.AtomicUnit = unit
		return nil
	}
}

func WithUnitExp(exp *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.UnitExp = exp
		return nil
	}
}

func WithGasType(gasType *basetypes.GasType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if gasType == nil {
			if must {
				return fmt.Errorf("invalid gastype")
			}
			return nil
		}
		switch *gasType {
		case basetypes.GasType_FixedGas:
		case basetypes.GasType_DynamicGas:
		case basetypes.GasType_GasUnsupported:
		default:
			return fmt.Errorf("invalid gastype")
		}
		h.GasType = gasType
		return nil
	}
}

func WithChainID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid chainid")
			}
			return nil
		}
		if *id == "" {
			return fmt.Errorf("invalid chainid")
		}
		h.ChainID = id
		return nil
	}
}

func WithNickname(nickname *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if nickname == nil {
			if must {
				return fmt.Errorf("invalid nickname")
			}
			return nil
		}
		if *nickname == "" {
			return fmt.Errorf("invalid nickname")
		}
		h.Nickname = nickname
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
