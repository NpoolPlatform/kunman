package ledger

import (
	"context"
	"encoding/json"
	"fmt"

	constant "github.com/NpoolPlatform/kunman/pkg/const"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/simulate/ledger"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/simulate/ledger"

	"github.com/google/uuid"
)

type Handler struct {
	crud.Req
	Reqs            []*crud.Req
	Conds           *crud.Conds
	Offset          int32
	Limit           int32
	IOSubType       *types.IOSubType
	IOExtra         *string
	StatementID     *uuid.UUID
	LedgerLockState *types.LedgerLockState
	Rollback        *bool
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid app id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid user id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithStatementID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid statement id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}

		h.StatementID = &_id
		return nil
	}
}

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid coin type id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithIOSubType(_type *types.IOSubType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid io sub type")
			}
			return nil
		}
		switch *_type {
		case types.IOSubType_Withdrawal:
		case types.IOSubType_Payment:
		case types.IOSubType_CommissionRevoke:
		default:
			return fmt.Errorf("invalid io sub type")
		}
		h.IOSubType = _type
		return nil
	}
}

func WithIOExtra(extra *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if extra == nil {
			if must {
				return fmt.Errorf("invalid extra")
			}
			return nil
		}
		if !json.Valid([]byte(*extra)) {
			return fmt.Errorf("io extra is invalid json str %v", *extra)
		}

		h.IOExtra = extra
		return nil
	}
}

func WithLedgerLockState(e *types.LedgerLockState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid ledgerlockstate")
			}
			return nil
		}
		switch *e {
		case types.LedgerLockState_LedgerLockLocked:
		case types.LedgerLockState_LedgerLockSettle:
		case types.LedgerLockState_LedgerLockRollback:
		default:
			return fmt.Errorf("invalid ledgerlockstate")
		}
		h.LedgerLockState = e
		return nil
	}
}

func WithRollback(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Rollback = b
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
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
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
				Val: id,
			}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{
				Op:  conds.GetCoinTypeID().GetOp(),
				Val: id,
			}
		}
		if len(conds.GetCoinTypeIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, val := range conds.GetCoinTypeIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return err
				}
				ids = append(ids, id)
			}
			h.Conds.CoinTypeIDs = &cruder.Cond{
				Op:  conds.GetCoinTypeIDs().GetOp(),
				Val: ids,
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
