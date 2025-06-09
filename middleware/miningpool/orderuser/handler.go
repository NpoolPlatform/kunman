package orderuser

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/miningpool/middleware/v1/orderuser"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	orderusercrud "github.com/NpoolPlatform/kunman/middleware/miningpool/crud/orderuser"
	"github.com/NpoolPlatform/kunman/middleware/miningpool/gooduser"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	orderusercrud.Req
	CoinTypeID     *string
	Proportion     *decimal.Decimal
	RevenueAddress *string
	AutoPay        *bool
	Reqs           []*orderusercrud.Req
	Conds          *orderusercrud.Conds
	Offset         int32
	Limit          int32
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

func WithGoodUserID(gooduserid *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if gooduserid == nil {
			if must {
				return wlog.Errorf("invalid gooduserid")
			}
			return nil
		}
		guH, err := gooduser.NewHandler(ctx, gooduser.WithEntID(gooduserid, true))
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := guH.ExistGoodUser(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid gooduserid")
		}
		h.GoodUserID = guH.EntID
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
		id, err := uuid.Parse(*cointypeid)
		if err != nil {
			return wlog.WrapError(err)
		}

		_id := id.String()

		h.CoinTypeID = &_id
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

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return wlog.Errorf("invalid name")
			}
			return nil
		}
		h.Name = name
		return nil
	}
}

func WithProportion(proportion *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if proportion == nil {
			if must {
				return wlog.Errorf("invalid proportion")
			}
			return nil
		}
		_proportion, err := decimal.NewFromString(*proportion)
		if err != nil {
			return wlog.Errorf("invalid proportion,err: %v", err)
		}
		h.Proportion = &_proportion
		return nil
	}
}

func WithRevenueAddress(revenueaddress *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if revenueaddress == nil {
			if must {
				return wlog.Errorf("invalid revenueaddress")
			}
			return nil
		}
		h.RevenueAddress = revenueaddress
		return nil
	}
}

func WithReadPageLink(readpagelink *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if readpagelink == nil {
			if must {
				return wlog.Errorf("invalid readpagelink")
			}
			return nil
		}
		h.ReadPageLink = readpagelink
		return nil
	}
}

func WithAutoPay(autopay *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if autopay == nil {
			if must {
				return wlog.Errorf("invalid autopay")
			}
			return nil
		}
		h.AutoPay = autopay
		return nil
	}
}

//nolint:gocognit
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &orderusercrud.Conds{}
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
		if conds.Name != nil {
			h.Conds.Name = &cruder.Cond{
				Op:  conds.GetName().GetOp(),
				Val: conds.GetName().GetValue(),
			}
		}
		if conds.GoodUserID != nil {
			id, err := uuid.Parse(conds.GetGoodUserID().GetValue())
			if err != nil {
				return wlog.WrapError(err)
			}
			h.Conds.GoodUserID = &cruder.Cond{
				Op:  conds.GetGoodUserID().GetOp(),
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
