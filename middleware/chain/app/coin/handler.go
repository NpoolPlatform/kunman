package appcoin

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	coincrud "github.com/NpoolPlatform/kunman/middleware/chain/crud/app/coin"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                       *uint32
	EntID                    *uuid.UUID
	AppID                    *uuid.UUID
	CoinTypeID               *uuid.UUID
	Name                     *string
	DisplayNames             []string
	Logo                     *string
	ForPay                   *bool
	WithdrawAutoReviewAmount *decimal.Decimal
	MarketValue              *decimal.Decimal
	SettlePercent            *uint32
	SettleTips               []string
	Setter                   *uuid.UUID
	ProductPage              *string
	DailyRewardAmount        *decimal.Decimal
	Disabled                 *bool
	Display                  *bool
	DisplayIndex             *uint32
	MaxAmountPerWithdraw     *decimal.Decimal
	Conds                    *coincrud.Conds
	Offset                   int32
	Limit                    int32
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
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

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid coinname")
		}
		h.Name = name
		return nil
	}
}

func WithDisplayNames(names []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.DisplayNames = names
		return nil
	}
}

func WithLogo(logo *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Logo = logo
		return nil
	}
}

func WithForPay(forPay *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ForPay = forPay
		return nil
	}
}

func WithWithdrawAutoReviewAmount(amount *string, must bool) func(context.Context, *Handler) error {
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
		h.WithdrawAutoReviewAmount = &_amount
		return nil
	}
}

func WithMarketValue(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid marketvalue")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.MarketValue = &_amount
		return nil
	}
}

func WithSettlePercent(percent *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if percent == nil {
			if must {
				return fmt.Errorf("invalid settlepercent")
			}
			return nil
		}
		if *percent == 0 {
			return fmt.Errorf("invalid percent")
		}
		h.SettlePercent = percent
		return nil
	}
}

func WithSettleTips(tips []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SettleTips = tips
		return nil
	}
}

func WithSetter(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid setter")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.Setter = &_id
		return nil
	}
}

func WithProductPage(page *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ProductPage = page
		return nil
	}
}

func WithDisabled(disabled *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Disabled = disabled
		return nil
	}
}

func WithDailyRewardAmount(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid dailyrewardamount")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.DailyRewardAmount = &_amount
		return nil
	}
}

func WithDisplay(display *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Display = display
		return nil
	}
}

func WithDisplayIndex(index *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.DisplayIndex = index
		return nil
	}
}

func WithMaxAmountPerWithdraw(amount *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if amount == nil {
			if must {
				return fmt.Errorf("invalid maxamountperwithdraw")
			}
			return nil
		}
		_amount, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.MaxAmountPerWithdraw = &_amount
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &coincrud.Conds{}
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
		if conds.ForPay != nil {
			h.Conds.ForPay = &cruder.Cond{
				Op:  conds.GetForPay().GetOp(),
				Val: conds.GetForPay().GetValue(),
			}
		}
		if conds.Disabled != nil {
			h.Conds.Disabled = &cruder.Cond{
				Op:  conds.GetDisabled().GetOp(),
				Val: conds.GetDisabled().GetValue(),
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
		if conds.CoinTypeIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetCoinTypeIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.CoinTypeIDs = &cruder.Cond{
				Op:  conds.GetCoinTypeIDs().GetOp(),
				Val: ids,
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
