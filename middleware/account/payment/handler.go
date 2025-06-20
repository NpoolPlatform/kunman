package payment

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/payment"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	paymentcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/payment"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID            *uint32
	EntID         *uuid.UUID
	CoinTypeID    *uuid.UUID
	AccountID     *uuid.UUID
	Address       *string
	Active        *bool
	Locked        *bool
	LockedBy      *basetypes.AccountLockedBy
	Blocked       *bool
	CollectingTID *uuid.UUID
	AvailableAt   *uint32
	Conds         *paymentcrud.Conds
	Offset        int32
	Limit         int32
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

func WithCoinTypeID(coinTypeID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if coinTypeID == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
			return nil
		}
		_coinTypeID, err := uuid.Parse(*coinTypeID)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_coinTypeID
		return nil
	}
}

func WithAccountID(accountID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if accountID == nil {
			if must {
				return fmt.Errorf("invalid accountid")
			}
			return nil
		}
		_accountID, err := uuid.Parse(*accountID)
		if err != nil {
			return err
		}
		h.AccountID = &_accountID
		return nil
	}
}

func WithCollectingTID(collectingTID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if collectingTID == nil {
			if must {
				return fmt.Errorf("invalid collectingtid")
			}
			return nil
		}
		_collectingTID, err := uuid.Parse(*collectingTID)
		if err != nil {
			return err
		}
		h.CollectingTID = &_collectingTID
		return nil
	}
}

func WithAddress(address *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if address == nil {
			if must {
				return fmt.Errorf("invalid address")
			}
			return nil
		}
		if *address == "" {
			return fmt.Errorf("invalid address")
		}
		h.Address = address
		return nil
	}
}

func WithActive(active *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Active = active
		return nil
	}
}

func WithLocked(locked *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Locked = locked
		return nil
	}
}

func WithLockedBy(lockedBy *basetypes.AccountLockedBy, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lockedBy == nil {
			if must {
				return fmt.Errorf("invalid lockedby")
			}
			return nil
		}
		switch *lockedBy {
		case basetypes.AccountLockedBy_Payment:
		case basetypes.AccountLockedBy_Collecting:
		default:
			return fmt.Errorf("invalid lockedby")
		}
		h.LockedBy = lockedBy
		return nil
	}
}

func WithBlocked(blocked *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Blocked = blocked
		return nil
	}
}

func WithAvailableAt(availableAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AvailableAt = availableAt
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &paymentcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{Op: conds.GetCoinTypeID().GetOp(), Val: id}
		}
		if conds.AccountID != nil {
			id, err := uuid.Parse(conds.GetAccountID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AccountID = &cruder.Cond{Op: conds.GetAccountID().GetOp(), Val: id}
		}
		if conds.Address != nil {
			h.Conds.Address = &cruder.Cond{
				Op:  conds.GetAddress().GetOp(),
				Val: conds.GetAddress().GetValue(),
			}
		}
		if conds.Active != nil {
			h.Conds.Active = &cruder.Cond{
				Op:  conds.GetActive().GetOp(),
				Val: conds.GetActive().GetValue(),
			}
		}
		if conds.Locked != nil {
			h.Conds.Locked = &cruder.Cond{
				Op:  conds.GetLocked().GetOp(),
				Val: conds.GetLocked().GetValue(),
			}
		}
		if conds.LockedBy != nil {
			h.Conds.LockedBy = &cruder.Cond{
				Op:  conds.GetLockedBy().GetOp(),
				Val: basetypes.AccountLockedBy(conds.GetLockedBy().GetValue()),
			}
		}
		if conds.Blocked != nil {
			h.Conds.Blocked = &cruder.Cond{
				Op:  conds.GetBlocked().GetOp(),
				Val: conds.GetBlocked().GetValue(),
			}
		}
		if conds.AccountIDs != nil {
			accountIDs := []uuid.UUID{}
			for _, accountID := range conds.GetAccountIDs().GetValue() {
				_accountID, err := uuid.Parse(accountID)
				if err != nil {
					return err
				}
				accountIDs = append(accountIDs, _accountID)
			}
			h.Conds.AccountIDs = &cruder.Cond{Op: conds.GetAccountIDs().GetOp(), Val: accountIDs}
		}
		if conds.AvailableAt != nil {
			h.Conds.AvailableAt = &cruder.Cond{
				Op:  conds.GetAvailableAt().GetOp(),
				Val: conds.GetAvailableAt().GetValue(),
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
