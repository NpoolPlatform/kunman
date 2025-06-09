package coin

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	mpbasetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/coin"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	coincrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/coin"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/pool"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	coincrud.Req
	Conds  *coincrud.Conds
	Offset int32
	Limit  int32
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

func WithCoinType(cointype *basetypes.CoinType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if cointype == nil {
			if must {
				return wlog.Errorf("invalid cointype")
			}
			return nil
		}
		if *cointype == basetypes.CoinType_DefaultCoinType {
			return wlog.Errorf("invalid cointype,not allow be default type")
		}
		h.CoinType = cointype
		return nil
	}
}

func WithCoinTypeID(cointypeid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if cointypeid == nil {
			if must {
				return wlog.Errorf("invalid cointypeid")
			}
			return nil
		}
		_id, err := uuid.Parse(*cointypeid)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithPoolID(poolid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if poolid == nil {
			if must {
				return wlog.Errorf("invalid poolid")
			}
			return nil
		}
		poolH, err := pool.NewHandler(ctx, pool.WithEntID(poolid, true))
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := poolH.ExistPool(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid poolid")
		}
		h.PoolID = poolH.EntID
		return nil
	}
}

func WithFeeRatio(feeratio *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if feeratio == nil {
			if must {
				return wlog.Errorf("invalid feeratio")
			}
			return nil
		}
		_feeratio, err := decimal.NewFromString(*feeratio)
		if err != nil {
			return wlog.WrapError(err)
		}

		if _feeratio.Sign() <= 0 || _feeratio.GreaterThan(decimal.NewFromInt(1)) {
			return wlog.Errorf("invalid feeratio")
		}

		h.FeeRatio = &_feeratio
		return nil
	}
}

func WithFixedRevenueAble(fixedrevenueable *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if fixedrevenueable == nil {
			if must {
				return wlog.Errorf("invalid fixedrevenueable")
			}
			return nil
		}
		h.FixedRevenueAble = fixedrevenueable
		return nil
	}
}

func WithLeastTransferAmount(leastTransferAmount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if leastTransferAmount == nil {
			if must {
				return wlog.Errorf("invalid leasttransferamount")
			}
			return nil
		}
		_leastTransferAmount, err := decimal.NewFromString(*leastTransferAmount)
		if err != nil {
			return wlog.WrapError(err)
		}
		if _leastTransferAmount.Sign() <= 0 {
			return wlog.Errorf("invalid leasttransferamount")
		}
		h.LeastTransferAmount = &_leastTransferAmount
		return nil
	}
}

func WithBenefitIntervalSeconds(benefitintervalseconds *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if benefitintervalseconds == nil {
			if must {
				return wlog.Errorf("invalid benefitintervalseconds")
			}
			return nil
		}
		if *benefitintervalseconds == 0 {
			return wlog.Errorf("invalid benefitintervalseconds")
		}
		h.BenefitIntervalSeconds = benefitintervalseconds
		return nil
	}
}

func WithRemark(remark *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if remark == nil {
			if must {
				return wlog.Errorf("invalid remark")
			}
			return nil
		}
		h.Remark = remark
		return nil
	}
}

//nolint:gocognit
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &coincrud.Conds{}
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
		if conds.PoolID != nil {
			id, err := uuid.Parse(conds.GetPoolID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.PoolID = &cruder.Cond{
				Op:  conds.GetPoolID().GetOp(),
				Val: id,
			}
		}
		if conds.CoinTypeIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetCoinTypeIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return wlog.WrapError(err)
				}
				ids = append(ids, _id)
			}

			h.Conds.CoinTypeIDs = &cruder.Cond{
				Op:  conds.GetCoinTypeIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.CoinTypeID = &cruder.Cond{
				Op:  conds.GetCoinTypeID().GetOp(),
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
		if conds.CoinType != nil {
			h.Conds.CoinType = &cruder.Cond{
				Op:  conds.GetCoinType().GetOp(),
				Val: basetypes.CoinType(conds.GetCoinType().GetValue()),
			}
		}
		if conds.MiningPoolType != nil {
			h.Conds.MiningPoolType = &cruder.Cond{
				Op:  conds.GetMiningPoolType().GetOp(),
				Val: mpbasetypes.MiningPoolType(conds.GetMiningPoolType().GetValue()),
			}
		}
		if conds.FixedRevenueAble != nil {
			h.Conds.FixedRevenueAble = &cruder.Cond{
				Op:  conds.GetFixedRevenueAble().GetOp(),
				Val: conds.GetFixedRevenueAble().GetValue(),
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
