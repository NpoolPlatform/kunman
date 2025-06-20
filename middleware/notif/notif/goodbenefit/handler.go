package goodbenefit

import (
	"context"
	"fmt"

	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
	crud "github.com/NpoolPlatform/kunman/middleware/notif/crud/notif/goodbenefit"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID          *uint32
	EntID       *uuid.UUID
	GoodID      *uuid.UUID
	GoodType    *goodtypes.GoodType
	GoodName    *string
	CoinTypeID  *uuid.UUID
	Amount      *decimal.Decimal
	State       *basetypes.Result
	Message     *string
	BenefitDate *uint32
	TxID        *uuid.UUID
	Generated   *bool
	DeletedAt   *uint32
	Conds       *crud.Conds
	Offset      int32
	Limit       int32
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

func WithGoodType(e *goodtypes.GoodType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid goodtype")
			}
			return nil
		}
		switch *e {
		case goodtypes.GoodType_PowerRental:
		case goodtypes.GoodType_LegacyPowerRental:
		default:
			return fmt.Errorf("invalid goodtype")
		}
		h.GoodType = e
		return nil
	}
}

func WithGoodName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid goodname")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid good name")
		}
		h.GoodName = name
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

func WithAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid amount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.Amount = &_amount
		return nil
	}
}

func WithState(state *basetypes.Result, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid state")
			}
			return nil
		}
		switch *state {
		case basetypes.Result_Fail:
		case basetypes.Result_Success:
		default:
			return fmt.Errorf("invalid state %s", *state)
		}

		h.State = state
		return nil
	}
}

func WithMessage(message *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if message == nil {
			if must {
				return fmt.Errorf("invalid message")
			}
			return nil
		}
		if *message == "" {
			return fmt.Errorf("invalid message")
		}
		h.Message = message
		return nil
	}
}

func WithGenerated(generated *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if generated == nil {
			if must {
				return fmt.Errorf("invalid generated")
			}
			return nil
		}
		h.Generated = generated
		return nil
	}
}

func WithBenefitDate(_date *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _date == nil {
			if must {
				return fmt.Errorf("invalid benefitdate")
			}
			return nil
		}
		h.BenefitDate = _date
		return nil
	}
}

func WithTxID(txID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if txID == nil {
			if must {
				return fmt.Errorf("invalid txid")
			}
			return nil
		}
		_txID, err := uuid.Parse(*txID)
		if err != nil {
			return err
		}
		h.TxID = &_txID
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
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
			goodID, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{Op: conds.GetGoodID().GetOp(), Val: goodID}
		}
		if conds.CoinTypeID != nil {
			goodID, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{Op: conds.GetCoinTypeID().GetOp(), Val: goodID}
		}
		if conds.Generated != nil {
			h.Conds.Generated = &cruder.Cond{Op: conds.GetGenerated().GetOp(), Val: conds.GetGenerated().GetValue()}
		}
		if conds.BenefitDateStart != nil {
			h.Conds.BenefitDateStart = &cruder.Cond{Op: conds.GetBenefitDateStart().GetOp(), Val: conds.GetBenefitDateStart().GetValue()}
		}
		if conds.BenefitDateEnd != nil {
			h.Conds.BenefitDateEnd = &cruder.Cond{Op: conds.GetBenefitDateEnd().GetOp(), Val: conds.GetBenefitDateEnd().GetValue()}
		}
		if conds.GoodType != nil {
			h.Conds.GoodType = &cruder.Cond{Op: conds.GetGoodType().GetOp(), Val: goodtypes.GoodType(conds.GetGoodType().GetValue())}
		}
		if conds.GoodTypes != nil {
			_types := []goodtypes.GoodType{}
			for _, _type := range conds.GetGoodTypes().GetValue() {
				_types = append(_types, goodtypes.GoodType(_type))
			}
			h.Conds.GoodTypes = &cruder.Cond{Op: conds.GetGoodTypes().GetOp(), Val: _types}
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
