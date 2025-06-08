package chain

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	chaincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/chain"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/chain"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID         *uint32
	EntID      *uuid.UUID
	ChainType  *string
	Logo       *string
	ChainID    *string
	NativeUnit *string
	AtomicUnit *string
	UnitExp    *uint32
	ENV        *string
	GasType    *basetypes.GasType
	Nickname   *string
	Conds      *chaincrud.Conds
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

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
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
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
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

func WithLogo(logo *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Logo = logo
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
		if exp == nil {
			if must {
				return fmt.Errorf("invalid unitexp")
			}
			return nil
		}
		h.UnitExp = exp
		return nil
	}
}

func WithENV(env *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if env == nil {
			if must {
				return fmt.Errorf("invalid ent")
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &chaincrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.ENV != nil {
			h.Conds.ENV = &cruder.Cond{
				Op:  conds.GetENV().GetOp(),
				Val: conds.GetENV().GetValue(),
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.ChainType != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetChainType().GetOp(),
				Val: conds.GetChainType().GetValue(),
			}
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
