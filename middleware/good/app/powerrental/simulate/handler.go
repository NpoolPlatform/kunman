package appsimulatepowerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental/simulate"
	apppowerrental1 "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
	appgoodbasecrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/good/goodbase"
	appsimulatepowerrentalcrud "github.com/NpoolPlatform/kunman/middleware/good/crud/app/powerrental/simulate"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Here we actually only have simulate order for PowerRental good

type Handler struct {
	ID *uint32
	appsimulatepowerrentalcrud.Req
	AppSimulatePowerRentalConds *appsimulatepowerrentalcrud.Conds
	AppGoodBaseConds            *appgoodbasecrud.Conds
	Offset                      int32
	Limit                       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		AppSimulatePowerRentalConds: &appsimulatepowerrentalcrud.Conds{},
		AppGoodBaseConds:            &appgoodbasecrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
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

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appgoodid")
			}
			return nil
		}
		handler, err := apppowerrental1.NewHandler(
			ctx,
			apppowerrental1.WithAppGoodID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistPowerRental(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid appgoodid")
		}
		h.AppGoodID = handler.AppGoodID
		return nil
	}
}

func WithOrderUnits(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid units")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return wlog.Errorf("units is less than or equal to 0")
		}
		h.OrderUnits = &amount
		return nil
	}
}

func WithOrderDurationSeconds(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			if must {
				return wlog.Errorf("invalid duration")
			}
			return nil
		}
		if *duration <= uint32(0) {
			return wlog.Errorf("duration is less than or equal to 0")
		}
		h.OrderDurationSeconds = duration
		return nil
	}
}

func (h *Handler) withAppGoodBaseConds(conds *npool.Conds) error {
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func (h *Handler) withAppSimulatePowerRentalConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.AppSimulatePowerRentalConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppSimulatePowerRentalConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppSimulatePowerRentalConds.AppGoodID = &cruder.Cond{
			Op:  conds.GetAppGoodID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withAppGoodBaseConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withAppSimulatePowerRentalConds(conds)
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
