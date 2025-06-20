package coinusedfor

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/kunman/message/basetypes/chain/v1"
	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin/usedfor"
	coinusedforcrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/coin/usedfor"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID         *uint32
	EntID      *uuid.UUID
	CoinTypeID *uuid.UUID
	UsedFor    *types.CoinUsedFor
	Priority   *uint32
	Conds      *coinusedforcrud.Conds
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

func WithUsedFor(usedFor *types.CoinUsedFor, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if usedFor == nil {
			if must {
				return fmt.Errorf("invalid usedfor")
			}
			return nil
		}
		switch *usedFor {
		case types.CoinUsedFor_CoinUsedForCouponCash:
		case types.CoinUsedFor_CoinUsedForGoodFee:
		default:
			return fmt.Errorf("invalid usedfor")
		}
		h.UsedFor = usedFor
		return nil
	}
}

func WithPriority(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Priority = n
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &coinusedforcrud.Conds{}
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
		if conds.CoinTypeIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetCoinTypeIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.CoinTypeIDs = &cruder.Cond{
				Op:  conds.GetCoinTypeIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.UsedFor != nil {
			h.Conds.UsedFor = &cruder.Cond{
				Op:  cruder.EQ,
				Val: types.CoinUsedFor(conds.GetUsedFor().GetValue()),
			}
		}
		if conds.UsedFors != nil {
			usedFors := []types.CoinUsedFor{}
			for _, usedFor := range conds.GetUsedFors().GetValue() {
				usedFors = append(usedFors, types.CoinUsedFor(usedFor))
			}
			h.Conds.UsedFors = &cruder.Cond{
				Op:  cruder.IN,
				Val: usedFors,
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
