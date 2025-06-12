package config

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	appconfigcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/app/config"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"

	"github.com/google/uuid"
)

type Handler struct {
	ID               *uint32
	EntID            *uuid.UUID
	AppID            *uuid.UUID
	SettleMode       *types.SettleMode
	SettleAmountType *types.SettleAmountType
	SettleInterval   *types.SettleInterval
	CommissionType   *types.CommissionType
	SettleBenefit    *bool
	StartAt          *uint32
	MaxLevel         *uint32
	Conds            *appconfigcrud.Conds
	Offset           int32
	Limit            int32
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
			return wlog.WrapError(err)
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
			return wlog.WrapError(err)
		}
		h.AppID = &_id
		return nil
	}
}

func WithSettleMode(settleMode *types.SettleMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if settleMode == nil {
			if must {
				return wlog.Errorf("invalid settlemode")
			}
			return nil
		}
		switch *settleMode {
		case types.SettleMode_SettleWithGoodValue:
		case types.SettleMode_SettleWithPaymentAmount:
		default:
			return wlog.Errorf("invalid settlemode")
		}
		h.SettleMode = settleMode
		return nil
	}
}

func WithSettleAmountType(settleAmount *types.SettleAmountType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if settleAmount == nil {
			if must {
				return wlog.Errorf("invalid settleamounttype")
			}
			return nil
		}
		switch *settleAmount {
		case types.SettleAmountType_SettleByPercent:
		case types.SettleAmountType_SettleByAmount:
		default:
			return wlog.Errorf("invalid settleamount")
		}
		h.SettleAmountType = settleAmount
		return nil
	}
}

func WithSettleInterval(settleInterval *types.SettleInterval, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if settleInterval == nil {
			if must {
				return wlog.Errorf("invalid settleinterval")
			}
			return nil
		}
		switch *settleInterval {
		case types.SettleInterval_SettleAggregate:
		case types.SettleInterval_SettleMonthly:
		case types.SettleInterval_SettleYearly:
		case types.SettleInterval_SettleEveryOrder:
		default:
			return wlog.Errorf("invalid settlemode")
		}
		h.SettleInterval = settleInterval
		return nil
	}
}

func WithCommissionType(commissionType *types.CommissionType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if commissionType == nil {
			if must {
				return wlog.Errorf("invalid commissiontype")
			}
			return nil
		}
		switch *commissionType {
		case types.CommissionType_LayeredCommission:
		case types.CommissionType_DirectCommission:
		case types.CommissionType_LegacyCommission:
		case types.CommissionType_WithoutCommission:
		default:
			return wlog.Errorf("invalid commissiontype")
		}
		h.CommissionType = commissionType
		return nil
	}
}

func WithSettleBenefit(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid settlebenefit")
			}
		}
		h.SettleBenefit = value
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

func WithMaxLevel(value *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid MaxLevel")
			}
		}
		h.MaxLevel = value
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &appconfigcrud.Conds{}
		if conds == nil {
			return nil
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
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.SettleMode != nil {
			h.Conds.SettleMode = &cruder.Cond{
				Op:  conds.GetSettleMode().GetOp(),
				Val: types.SettleMode(conds.GetSettleMode().GetValue()),
			}
		}
		if conds.SettleAmountType != nil {
			h.Conds.SettleAmountType = &cruder.Cond{
				Op:  conds.GetSettleAmountType().GetOp(),
				Val: types.SettleAmountType(conds.GetSettleAmountType().GetValue()),
			}
		}
		if conds.SettleInterval != nil {
			h.Conds.SettleInterval = &cruder.Cond{
				Op:  conds.GetSettleInterval().GetOp(),
				Val: types.SettleInterval(conds.GetSettleInterval().GetValue()),
			}
		}
		if conds.CommissionType != nil {
			h.Conds.CommissionType = &cruder.Cond{
				Op:  conds.GetCommissionType().GetOp(),
				Val: types.CommissionType(conds.GetCommissionType().GetValue()),
			}
		}
		if conds.SettleBenefit != nil {
			h.Conds.SettleBenefit = &cruder.Cond{
				Op:  conds.GetSettleBenefit().GetOp(),
				Val: conds.GetSettleBenefit().GetValue(),
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
					return wlog.WrapError(err)
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{
				Op:  conds.GetEntIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: conds.GetID().GetValue(),
			}
		}
		if conds.MaxLevel != nil {
			h.Conds.MaxLevel = &cruder.Cond{
				Op:  conds.GetMaxLevel().GetOp(),
				Val: conds.GetMaxLevel().GetValue(),
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
