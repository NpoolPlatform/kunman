package user

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	usercrud "github.com/NpoolPlatform/kunman/middleware/account/crud/user"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID         *uint32
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	UserID     *uuid.UUID
	CoinTypeID *uuid.UUID
	AccountID  *uuid.UUID
	Address    *string
	UsedFor    *basetypes.AccountUsedFor
	Labels     []string
	Active     *bool
	Blocked    *bool
	Locked     *bool
	Memo       *string
	Conds      *usercrud.Conds
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
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
				return fmt.Errorf("invalid userid")
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

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
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

func WithAccountID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid accountid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AccountID = &_id
		return nil
	}
}

func WithAddress(addr *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if addr == nil {
			if must {
				return fmt.Errorf("invalid address")
			}
			return nil
		}
		if *addr == "" {
			return fmt.Errorf("invalid address")
		}
		h.Address = addr
		return nil
	}
}

func WithUsedFor(usedFor *basetypes.AccountUsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if usedFor == nil {
			if must {
				return fmt.Errorf("invalid usedfor")
			}
			return nil
		}
		switch *usedFor {
		case basetypes.AccountUsedFor_UserWithdraw:
		case basetypes.AccountUsedFor_UserDirectBenefit:
		default:
			return fmt.Errorf("invalid usedfor")
		}
		h.UsedFor = usedFor
		return nil
	}
}

func WithLabels(labels []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if labels == nil {
			if must {
				return fmt.Errorf("invalid labels")
			}
			return nil
		}
		if len(labels) == 0 {
			return fmt.Errorf("invalid labels")
		}

		h.Labels = labels
		return nil
	}
}

func WithActive(active *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Active = active
		return nil
	}
}

func WithBlocked(blocked *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Blocked = blocked
		return nil
	}
}

func WithLocked(locked *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Locked = locked
		return nil
	}
}

func WithMemo(memo *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if memo == nil {
			if must {
				return fmt.Errorf("invalid memo")
			}
			return nil
		}
		if *memo == "" {
			return fmt.Errorf("invalid memo")
		}
		h.Memo = memo
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &usercrud.Conds{}
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
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
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
		if conds.UsedFor != nil {
			h.Conds.UsedFor = &cruder.Cond{
				Op:  conds.GetUsedFor().GetOp(),
				Val: basetypes.AccountUsedFor(conds.GetUsedFor().GetValue()),
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
			h.Conds.EntIDs = &cruder.Cond{Op: conds.GetEntIDs().GetOp(), Val: ids}
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
		if conds.Active != nil {
			h.Conds.Active = &cruder.Cond{
				Op:  conds.GetActive().GetOp(),
				Val: conds.GetActive().GetValue(),
			}
		}
		if conds.Address != nil {
			h.Conds.Address = &cruder.Cond{
				Op:  conds.GetAddress().GetOp(),
				Val: conds.GetAddress().GetValue(),
			}
		}
		if conds.Blocked != nil {
			h.Conds.Blocked = &cruder.Cond{
				Op:  conds.GetBlocked().GetOp(),
				Val: conds.GetBlocked().GetValue(),
			}
		}
		if conds.Address != nil {
			h.Conds.Address = &cruder.Cond{
				Op:  conds.GetAddress().GetOp(),
				Val: conds.GetAddress().GetValue(),
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
