package config

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/commission/config"
	commissionconfigcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/app/commission/config"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID              *uint32
	EntID           *uuid.UUID
	AppID           *uuid.UUID
	ThresholdAmount *decimal.Decimal
	AmountOrPercent *decimal.Decimal
	StartAt         *uint32
	Invites         *uint32
	SettleType      *types.SettleType
	Disabled        *bool
	Level           *uint32
	Conds           *commissionconfigcrud.Conds
	Offset          int32
	Limit           int32
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
				return wlog.Errorf("invalid id")
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
				return wlog.Errorf("invalid entid")
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
				return wlog.Errorf("invalid appid")
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

//nolint:dupl
func WithThresholdAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid thresholdamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid thresholdamount")
		}
		h.ThresholdAmount = &_amount
		return nil
	}
}

func WithSettleType(settleType *types.SettleType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if settleType == nil {
			if must {
				return wlog.Errorf("invalid settletype")
			}
			return nil
		}
		switch *settleType {
		case types.SettleType_GoodOrderPayment:
		case types.SettleType_TechniqueServiceFee:
		default:
			return wlog.Errorf("invalid settletype")
		}
		h.SettleType = settleType
		return nil
	}
}

//nolint:dupl
func WithAmountOrPercent(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid amountorpercent")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if _amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid amountorpercent")
		}
		h.AmountOrPercent = &_amount
		return nil
	}
}

func WithStartAt(startAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if startAt == nil {
			if must {
				return wlog.Errorf("invalid startat")
			}
		}
		h.StartAt = startAt
		return nil
	}
}

func WithInvites(value *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid invites")
			}
		}
		h.Invites = value
		return nil
	}
}

func WithDisabled(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid disabled")
			}
		}
		h.Disabled = value
		return nil
	}
}

func WithLevel(value *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid level")
			}
		}
		h.Level = value
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &commissionconfigcrud.Conds{}
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
		if conds.SettleType != nil {
			h.Conds.SettleType = &cruder.Cond{
				Op:  conds.GetSettleType().GetOp(),
				Val: types.SettleType(conds.GetSettleType().GetValue()),
			}
		}
		if conds.StartAt != nil {
			h.Conds.StartAt = &cruder.Cond{
				Op:  conds.GetStartAt().GetOp(),
				Val: conds.GetStartAt().GetValue(),
			}
		}
		if conds.EndAt != nil {
			h.Conds.EndAt = &cruder.Cond{
				Op:  conds.GetEndAt().GetOp(),
				Val: conds.GetEndAt().GetValue(),
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
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.Disabled != nil {
			h.Conds.Disabled = &cruder.Cond{
				Op:  conds.GetDisabled().GetOp(),
				Val: conds.GetDisabled().GetValue(),
			}
		}
		if conds.Level != nil {
			h.Conds.Level = &cruder.Cond{
				Op:  conds.GetLevel().GetOp(),
				Val: conds.GetLevel().GetValue(),
			}
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: conds.GetID().GetValue(),
			}
		}
		return nil
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}
