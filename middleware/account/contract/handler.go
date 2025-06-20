package contract

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/account/middleware/v1/contract"
	accounttypes "github.com/NpoolPlatform/kunman/message/basetypes/account/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	contractcrud "github.com/NpoolPlatform/kunman/middleware/account/crud/contract"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID                   *uint32
	EntID                *uuid.UUID
	GoodID               *uuid.UUID
	DelegatedStakingID   *uuid.UUID
	CoinTypeID           *uuid.UUID
	AccountID            *uuid.UUID
	Address              *string
	Backup               *bool
	Active               *bool
	Locked               *bool
	LockedBy             *basetypes.AccountLockedBy
	Blocked              *bool
	ContractOperatorType *accounttypes.ContractOperatorType
	TransactionID        *uuid.UUID
	Conds                *contractcrud.Conds
	Offset               int32
	Limit                int32
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

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.GoodID = &_id
		return nil
	}
}

func WithDelegatedStakingID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid delegatedstakingid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.DelegatedStakingID = &_id
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

func WithBackup(backup *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Backup = backup
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

func WithContractOperatorType(contractOperatorType *accounttypes.ContractOperatorType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if contractOperatorType == nil {
			if must {
				return fmt.Errorf("invalid contractoperatortype")
			}
			return nil
		}
		switch *contractOperatorType {
		case accounttypes.ContractOperatorType_ContractOwner:
		case accounttypes.ContractOperatorType_ContractCalculator:
		default:
			return fmt.Errorf("invalid contractoperatortype")
		}
		h.ContractOperatorType = contractOperatorType
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error { //nolint
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &contractcrud.Conds{}
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
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{Op: conds.GetGoodID().GetOp(), Val: id}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{Op: conds.GetCoinTypeID().GetOp(), Val: id}
		}
		if conds.CoinTypeIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetCoinTypeIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.CoinTypeIDs = &cruder.Cond{Op: conds.GetCoinTypeIDs().GetOp(), Val: ids}
		}
		if conds.AccountID != nil {
			id, err := uuid.Parse(conds.GetAccountID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AccountID = &cruder.Cond{Op: conds.GetAccountID().GetOp(), Val: id}
		}
		if conds.DelegatedStakingID != nil {
			id, err := uuid.Parse(conds.GetDelegatedStakingID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.DelegatedStakingID = &cruder.Cond{Op: conds.GetDelegatedStakingID().GetOp(), Val: id}
		}
		if conds.Backup != nil {
			h.Conds.Backup = &cruder.Cond{
				Op:  conds.GetBackup().GetOp(),
				Val: conds.GetBackup().GetValue(),
			}
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
		if conds.ContractOperatorType != nil {
			h.Conds.ContractOperatorType = &cruder.Cond{
				Op:  conds.GetContractOperatorType().GetOp(),
				Val: accounttypes.ContractOperatorType(conds.GetContractOperatorType().GetValue()),
			}
		}
		if conds.DelegatedStakingIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetDelegatedStakingIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.DelegatedStakingIDs = &cruder.Cond{Op: conds.GetDelegatedStakingIDs().GetOp(), Val: ids}
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
