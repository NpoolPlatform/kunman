package statement

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/good/ledger/statement"
	crud "github.com/NpoolPlatform/kunman/middleware/ledger/crud/good/ledger/statement"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	crud.Req
	Rollback *bool
	Reqs     []*crud.Req
	Conds    *crud.Conds
	Limit    int32
	Offset   int32
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

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid good id")
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

// nolint
func WithTotalAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid totalamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid totalamount")
		}
		h.TotalAmount = &_amount
		return nil
	}
}

// nolint
func WithUnsoldAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid unsold amount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid unsoldamount")
		}
		h.UnsoldAmount = &_amount
		return nil
	}
}

// nolint
func WithTechniqueServiceFeeAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid techniqueservicefeeamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid techniqueservicefeeamount")
		}
		h.TechniqueServiceFeeAmount = &_amount
		return nil
	}
}

func WithBenefitDate(date *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if date == nil {
			if must {
				return fmt.Errorf("invalid benefit date")
			}
			return nil
		}
		h.BenefitDate = date
		return nil
	}
}

func WithRollback(rollback *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if rollback == nil {
			if must {
				return fmt.Errorf("invalid rollback")
			}
			return nil
		}
		h.Rollback = rollback
		return nil
	}
}

// nolint
func WithReqs(reqs []*npool.GoodStatementReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*crud.Req{}
		for _, req := range reqs {
			if must {
				if req.GoodID == nil {
					return fmt.Errorf("invalid good id")
				}
				if req.CoinTypeID == nil {
					return fmt.Errorf("invalid coin type id")
				}
				if req.TotalAmount == nil {
					return fmt.Errorf("invalid total amount")
				}
				if req.UnsoldAmount == nil {
					return fmt.Errorf("invalid unsold amount")
				}
				if req.TechniqueServiceFeeAmount == nil {
					return fmt.Errorf("invalid technique service fee amount")
				}
				if req.BenefitDate == nil {
					return fmt.Errorf("invalid benefit date")
				}
			}
			_req := &crud.Req{}
			if req.ID != nil {
				_req.ID = req.ID
			}
			if req.EntID != nil {
				_id, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
				_req.EntID = &_id
			}
			if req.GoodID != nil {
				_id, err := uuid.Parse(*req.GoodID)
				if err != nil {
					return err
				}
				_req.GoodID = &_id
			}
			if req.CoinTypeID != nil {
				_id, err := uuid.Parse(*req.CoinTypeID)
				if err != nil {
					return err
				}
				_req.CoinTypeID = &_id
			}
			if req.TotalAmount != nil {
				amount, err := decimal.NewFromString(*req.TotalAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("invalid totalamount")
				}
				_req.TotalAmount = &amount
			}
			if req.UnsoldAmount != nil {
				amount, err := decimal.NewFromString(*req.UnsoldAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("invalid unsoldamount")
				}
				_req.UnsoldAmount = &amount
			}
			if req.TechniqueServiceFeeAmount != nil {
				amount, err := decimal.NewFromString(*req.TechniqueServiceFeeAmount)
				if err != nil {
					return err
				}
				if amount.Cmp(decimal.NewFromInt(0)) < 0 {
					return fmt.Errorf("invalid techniqueservicefeeamount")
				}
				_req.TechniqueServiceFeeAmount = &amount
			}
			if req.BenefitDate != nil {
				if *req.BenefitDate == 0 {
					return fmt.Errorf("invalid benefitdate")
				}
				_req.BenefitDate = req.BenefitDate
			}
			if req.Rollback != nil {
				h.Rollback = req.Rollback
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
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
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{
				Op:  conds.GetGoodID().GetOp(),
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
		if conds.Amount != nil {
			amount, err := decimal.NewFromString(conds.GetAmount().GetValue())
			if err != nil {
				return err
			}
			h.Conds.Amount = &cruder.Cond{
				Op:  conds.GetAmount().GetOp(),
				Val: amount,
			}
		}
		if conds.BenefitDate != nil {
			h.Conds.BenefitDate = &cruder.Cond{
				Op:  conds.GetBenefitDate().GetOp(),
				Val: conds.GetBenefitDate().GetValue(),
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
