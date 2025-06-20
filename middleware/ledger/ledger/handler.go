package ledger

import (
	"context"
	"encoding/json"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/ledger"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	crud.Req
	Reqs            []*crud.Req
	Conds           *crud.Conds
	Offset          int32
	Limit           int32
	IOSubType       *types.IOSubType
	IOExtra         *string
	LockID          *uuid.UUID
	StatementID     *uuid.UUID
	LedgerLockState *types.LedgerLockState
	Rollback        *bool
	Balances        []*LockBalance
	StatementIDs    []uuid.UUID
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

func WithLockID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid lock id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}

		h.LockID = &_id
		return nil
	}
}

func WithCurrencyID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid currency id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CurrencyID = &_id
		return nil
	}
}

func WithCurrencyType(e *types.CurrencyType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid currencytype")
			}
			return nil
		}
		h.CurrencyType = e
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

func WithLocked(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid locked")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("amount is less than 0 %v", *amount)
		}
		h.Locked = &_amount
		return nil
	}
}

func WithSpendable(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid spendable")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("amount is less than 0 %v", _amount.String())
		}
		h.Spendable = &_amount
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

func WithBalances(balances []*npool.LockBalance, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, balance := range balances {
			coinTypeID, err := uuid.Parse(balance.CurrencyID)
			if err != nil {
				return err
			}
			amount, err := decimal.NewFromString(balance.Amount)
			if err != nil {
				return err
			}
			h.Balances = append(h.Balances, &LockBalance{
				CurrencyID: coinTypeID,
				Amount:     amount,
			})
		}
		return nil
	}
}

func WithStatementIDs(ids []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, id := range ids {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			h.StatementIDs = append(h.StatementIDs, _id)
		}
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
		if conds.CurrencyID != nil {
			id, err := uuid.Parse(conds.GetCurrencyID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CurrencyID = &cruder.Cond{
				Op:  conds.GetCurrencyID().GetOp(),
				Val: id,
			}
		}
		if len(conds.GetCurrencyIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, val := range conds.GetCurrencyIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return err
				}
				ids = append(ids, id)
			}
			h.Conds.CurrencyIDs = &cruder.Cond{
				Op:  conds.GetCurrencyIDs().GetOp(),
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
