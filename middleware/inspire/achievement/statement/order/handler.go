//nolint:dupl
package orderstatement

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order"
	orderpaymentstatementmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/statement/order/payment"
	orderstatementcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/achievement/statement/order"
	orderpaymentstatementcrud "github.com/NpoolPlatform/kunman/middleware/inspire/crud/achievement/statement/order/payment"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	orderstatementcrud.Req
	PaymentStatementReqs []*orderpaymentstatementcrud.Req
	OrderStatementConds  *orderstatementcrud.Conds
	Offset               int32
	Limit                int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		OrderStatementConds: &orderstatementcrud.Conds{},
	}
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

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.UserID = &_id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = &_id
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodID = &_id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = &_id
		return nil
	}
}

func WithOrderUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderuserid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderUserID = &_id
		return nil
	}
}

func WithDirectContributorID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderuserid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.DirectContributorID = &_id
		return nil
	}
}

func WithGoodCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid goodcointypeid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodCoinTypeID = &_id
		return nil
	}
}

func WithUnits(value *string, must bool) func(context.Context, *Handler) error {
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
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("units is less than or equal to 0")
		}
		h.Units = &amount
		return nil
	}
}

func WithGoodValueUSD(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid goodvalueusd")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid goodvalueusd")
		}
		h.GoodValueUSD = &amount
		return nil
	}
}

func WithPaymentAmountUSD(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid paymentamountusd")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid paymentamountusd")
		}
		h.PaymentAmountUSD = &amount
		return nil
	}
}

func WithCommissionAmountUSD(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid commissionamountusd")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return wlog.WrapError(err)
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return wlog.Errorf("invalid commissionamountusd")
		}
		h.CommissionAmountUSD = &amount
		return nil
	}
}

func WithAppConfigID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid appconfigid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppConfigID = &_id
		return nil
	}
}

func WithCommissionConfigID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid commissionconfigid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.CommissionConfigID = &_id
		return nil
	}
}

func WithCommissionConfigType(value *types.CommissionConfigType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return wlog.Errorf("invalid commissionconfigtype")
			}
			return nil
		}
		switch *value {
		case types.CommissionConfigType_AppCommissionConfig:
		case types.CommissionConfigType_AppGoodCommissionConfig:
		case types.CommissionConfigType_LegacyCommissionConfig:
		case types.CommissionConfigType_WithoutCommissionConfig:
		default:
			return wlog.Errorf("invalid commissionconfigtype")
		}
		h.CommissionConfigType = value
		return nil
	}
}

func WithPaymentStatements(reqs []*orderpaymentstatementmwpb.StatementReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			_req := &orderpaymentstatementcrud.Req{}
			if req.EntID != nil {
				id, err := uuid.Parse(req.GetEntID())
				if err != nil {
					return wlog.WrapError(err)
				}
				_req.EntID = &id
			}

			id1, err := uuid.Parse(req.GetPaymentCoinTypeID())
			if err != nil {
				return wlog.WrapError(err)
			}
			_req.PaymentCoinTypeID = &id1

			amount1, err := decimal.NewFromString(req.GetAmount())
			if err != nil {
				return wlog.WrapError(err)
			}
			_req.Amount = &amount1

			amount2, err := decimal.NewFromString(req.GetCommissionAmount())
			if err != nil {
				return wlog.WrapError(err)
			}
			_req.CommissionAmount = &amount2

			h.PaymentStatementReqs = append(h.PaymentStatementReqs, _req)
		}
		return nil
	}
}

//nolint:funlen
func (h *Handler) withOrderStatementConds(conds *npool.Conds) error {
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
	}
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
	}
	if conds.UserID != nil {
		id, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
	}
	if conds.UserIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.UserIDs.GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderStatementConds.UserIDs = &cruder.Cond{Op: conds.GetUserIDs().GetOp(), Val: ids}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.GoodID = &cruder.Cond{Op: conds.GetGoodID().GetOp(), Val: id}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.AppGoodID = &cruder.Cond{Op: conds.GetAppGoodID().GetOp(), Val: id}
	}
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.OrderID = &cruder.Cond{Op: conds.GetOrderID().GetOp(), Val: id}
	}
	if conds.OrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.OrderIDs.GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderStatementConds.OrderIDs = &cruder.Cond{Op: conds.GetOrderIDs().GetOp(), Val: ids}
	}
	if conds.OrderUserID != nil {
		id, err := uuid.Parse(conds.GetOrderUserID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.OrderUserID = &cruder.Cond{Op: conds.GetOrderUserID().GetOp(), Val: id}
	}
	if conds.GoodCoinTypeID != nil {
		id, err := uuid.Parse(conds.GetGoodCoinTypeID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.GoodCoinTypeID = &cruder.Cond{Op: conds.GetGoodCoinTypeID().GetOp(), Val: id}
	}
	if conds.AppConfigID != nil {
		id, err := uuid.Parse(conds.GetAppConfigID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.AppConfigID = &cruder.Cond{Op: conds.GetAppConfigID().GetOp(), Val: id}
	}
	if conds.CommissionConfigID != nil {
		id, err := uuid.Parse(conds.GetCommissionConfigID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderStatementConds.CommissionConfigID = &cruder.Cond{Op: conds.GetCommissionConfigID().GetOp(), Val: id}
	}
	if conds.CommissionConfigType != nil {
		h.OrderStatementConds.CommissionConfigType = &cruder.Cond{
			Op:  conds.GetCommissionConfigType().GetOp(),
			Val: types.CommissionConfigType(conds.GetCommissionConfigType().GetValue()),
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		return h.withOrderStatementConds(conds)
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
