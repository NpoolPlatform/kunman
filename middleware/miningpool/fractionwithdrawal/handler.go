package fractionwithdrawal

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/miningpool/v1"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/fractionwithdrawal"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	fractionwithdrawalcrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/fractionwithdrawal"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/orderuser"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	fractionwithdrawalcrud.Req
	Reqs   []*fractionwithdrawalcrud.Req
	Conds  *fractionwithdrawalcrud.Conds
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

func WithAppID(appid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appid == nil {
			if must {
				return wlog.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*appid)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(userid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userid == nil {
			if must {
				return wlog.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*userid)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.UserID = &_id
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

func WithOrderUserID(orderuserid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if orderuserid == nil {
			if must {
				return wlog.Errorf("invalid orderuserid")
			}
			return nil
		}

		ouH, err := orderuser.NewHandler(ctx, orderuser.WithEntID(orderuserid, true))
		if err != nil {
			return wlog.WrapError(err)
		}
		exists, err := ouH.ExistOrderUser(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}

		if !exists {
			return wlog.Errorf("invalid orderuserid")
		}

		h.OrderUserID = ouH.EntID
		return nil
	}
}

func WithFractionWithdrawState(withdrawstate *basetypes.FractionWithdrawState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawstate == nil {
			if must {
				return wlog.Errorf("invalid withdrawstate")
			}
			return nil
		}
		if *withdrawstate == basetypes.FractionWithdrawState_DefaultFractionWithdrawState {
			return wlog.Errorf("invalid withdrawstate,not allow be default type")
		}
		h.FractionWithdrawState = withdrawstate
		return nil
	}
}

func WithWithdrawAt(withdrawtime *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if withdrawtime == nil {
			if must {
				return wlog.Errorf("invalid withdrawtime")
			}
			return nil
		}
		h.WithdrawAt = withdrawtime
		return nil
	}
}

func WithPromisePayAt(promisepayat *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if promisepayat == nil {
			if must {
				return wlog.Errorf("invalid promisepayat")
			}
			return nil
		}
		h.PromisePayAt = promisepayat
		return nil
	}
}

func WithMsg(msg *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if msg == nil {
			if must {
				return wlog.Errorf("invalid msg")
			}
			return nil
		}
		h.Msg = msg
		return nil
	}
}

//nolint:gocognit
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &fractionwithdrawalcrud.Conds{}
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
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
				Val: id,
			}
		}
		if conds.OrderUserID != nil {
			id, err := uuid.Parse(conds.GetOrderUserID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.OrderUserID = &cruder.Cond{
				Op:  conds.GetOrderUserID().GetOp(),
				Val: id,
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
		if conds.FractionWithdrawState != nil {
			h.Conds.FractionWithdrawState = &cruder.Cond{
				Op:  conds.GetFractionWithdrawState().GetOp(),
				Val: basetypes.FractionWithdrawState(conds.GetFractionWithdrawState().GetValue()),
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
