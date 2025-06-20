package history

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin/reward/history"
	historycrud "github.com/NpoolPlatform/kunman/middleware/good/crud/good/coin/reward/history"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	historycrud.Req
	HistoryConds *historycrud.Conds
	Offset       int32
	Limit        int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		HistoryConds: &historycrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func WithGoodID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = &id
		return nil
	}
}

func WithCoinTypeID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid cointypeid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.CoinTypeID = &id
		return nil
	}
}

func WithRewardDate(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return wlog.Errorf("invalid rewarddate")
			}
			return nil
		}
		h.RewardDate = n
		return nil
	}
}

func WithTID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid tid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.TID = &_id
		return nil
	}
}

func WithAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid amount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.Amount = &amount
		return nil
	}
}

func WithUnitAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid unitamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.UnitAmount = &amount
		return nil
	}
}

func WithUnitNetAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid unitnetamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.UnitNetAmount = &amount
		return nil
	}
}

//nolint:gocyclo
func (h *Handler) withHistoryConds(conds *npool.Conds) error {
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.HistoryConds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
	}
	if conds.ID != nil {
		h.HistoryConds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return err
		}
		h.HistoryConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			ids = append(ids, _id)
		}
		h.HistoryConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.CoinTypeID != nil {
		id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
		if err != nil {
			return err
		}
		h.HistoryConds.CoinTypeID = &cruder.Cond{
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
		h.HistoryConds.CoinTypeIDs = &cruder.Cond{
			Op:  conds.GetCoinTypeIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.RewardDate != nil {
		h.HistoryConds.RewardDate = &cruder.Cond{
			Op:  conds.GetRewardDate().GetOp(),
			Val: conds.GetRewardDate().GetValue(),
		}
	}
	if conds.StartAt != nil {
		h.HistoryConds.StartAt = &cruder.Cond{
			Op:  conds.GetStartAt().GetOp(),
			Val: conds.GetStartAt().GetValue(),
		}
	}
	if conds.EndAt != nil {
		h.HistoryConds.EndAt = &cruder.Cond{
			Op:  conds.GetEndAt().GetOp(),
			Val: conds.GetEndAt().GetValue(),
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		return h.withHistoryConds(conds)
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
