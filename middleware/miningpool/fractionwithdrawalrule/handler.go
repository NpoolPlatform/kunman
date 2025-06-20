package fractionwithdrawalrule

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawalrule"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/coin"
	fractionwithdrawalrulecrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/fractionwithdrawalrule"
	poolcrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/pool"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID                    *uint32
	EntID                 *uuid.UUID
	PoolCoinTypeID        *uuid.UUID
	WithdrawInterval      *uint32
	PayoutThreshold       *decimal.Decimal
	LeastWithdrawalAmount *decimal.Decimal
	WithdrawFee           *decimal.Decimal
	Reqs                  []*fractionwithdrawalrulecrud.Req
	Conds                 *fractionwithdrawalrulecrud.Conds
	PoolConds             *poolcrud.Conds
	Offset                int32
	Limit                 int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid id")
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
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

func WithPoolCoinTypeID(poolcointypeid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if poolcointypeid == nil {
			if must {
				return wlog.Errorf("invalid poolcointypeid")
			}
			return nil
		}

		coinH, err := coin.NewHandler(ctx, coin.WithEntID(poolcointypeid, true))
		if err != nil {
			return wlog.WrapError(err)
		}

		exist, err := coinH.ExistCoin(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid poolcointypeid")
		}
		h.PoolCoinTypeID = coinH.EntID
		return nil
	}
}

func WithWithdrawInterval(withdrawinterval *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawinterval == nil {
			if must {
				return wlog.Errorf("invalid withdrawinterval")
			}
			return nil
		}
		h.WithdrawInterval = withdrawinterval
		return nil
	}
}

func WithLeastWithdrawalAmount(leastwithdrawalamount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if leastwithdrawalamount == nil {
			if must {
				return wlog.Errorf("invalid leastwithdrawalamount")
			}
			return nil
		}
		_leastwithdrawalamount, err := decimal.NewFromString(*leastwithdrawalamount)
		if err != nil {
			return wlog.Errorf("invalid leastwithdrawalamount,err: %v", err)
		}
		if _leastwithdrawalamount.Sign() <= 0 {
			return wlog.Errorf("invalid leastwithdrawalamount")
		}
		h.LeastWithdrawalAmount = &_leastwithdrawalamount
		return nil
	}
}

func WithPayoutThreshold(payoutthreshold *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if payoutthreshold == nil {
			if must {
				return wlog.Errorf("invalid payoutthreshold")
			}
			return nil
		}
		_payoutthreshold, err := decimal.NewFromString(*payoutthreshold)
		if err != nil {
			return wlog.Errorf("invalid payoutthreshold,err: %v", err)
		}
		if _payoutthreshold.Sign() <= 0 {
			return wlog.Errorf("invalid payoutthreshold")
		}
		h.PayoutThreshold = &_payoutthreshold
		return nil
	}
}

func WithWithdrawFee(withdrawrate *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawrate == nil {
			if must {
				return wlog.Errorf("invalid withdrawrate")
			}
			return nil
		}
		_withdrawrate, err := decimal.NewFromString(*withdrawrate)
		if err != nil {
			return wlog.Errorf("invalid withdrawrate,err: %v", err)
		}
		if _withdrawrate.Sign() < 0 {
			return wlog.Errorf("invalid withdrawrate")
		}
		h.WithdrawFee = &_withdrawrate
		return nil
	}
}

//nolint:gocognit
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &fractionwithdrawalrulecrud.Conds{}
		h.PoolConds = &poolcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return wlog.WrapError(err)
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.PoolCoinTypeID != nil {
			id, err := uuid.Parse(conds.GetPoolCoinTypeID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.PoolCoinTypeID = &cruder.Cond{
				Op:  conds.GetPoolCoinTypeID().GetOp(),
				Val: id,
			}
		}
		if conds.PoolID != nil {
			id, err := uuid.Parse(conds.GetPoolID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.PoolConds.EntID = &cruder.Cond{
				Op:  conds.GetPoolID().GetOp(),
				Val: id,
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
