package appcoin

import (
	"context"
	"fmt"

	appmw "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	constant "github.com/NpoolPlatform/kunman/pkg/const"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                       *uint32
	EntID                    *string
	AppID                    *string
	CoinTypeID               *string
	Name                     *string
	DisplayNames             []string
	Logo                     *string
	ForPay                   *bool
	WithdrawAutoReviewAmount *string
	MarketValue              *string
	SettlePercent            *uint32
	SettleTips               []string
	Setter                   *string
	ProductPage              *string
	DailyRewardAmount        *string
	Disabled                 *bool
	Display                  *bool
	DisplayIndex             *uint32
	MaxAmountPerWithdraw     *string
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
		_, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = id
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

		handler, err := appmw.NewHandler(
			ctx,
			appmw.WithEntID(id, true),
		)
		if err != nil {
			return err
		}

		_app, err := handler.GetApp(ctx)
		if err != nil {
			return err
		}
		if _app == nil {
			return fmt.Errorf("invalid app")
		}
		h.AppID = id
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

		handler, err := coinmw.NewHandler(
			ctx,
			coinmw.WithEntID(id, true),
		)
		if err != nil {
			return err
		}

		_coin, err := handler.GetCoin(ctx)
		if err != nil {
			return err
		}
		if _coin == nil {
			return fmt.Errorf("invalid coin")
		}
		h.CoinTypeID = id
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
				return fmt.Errorf("invalid withdrawautoreviewamount")
			}
			return nil
		}
		_, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.WithdrawAutoReviewAmount = amount
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
		_, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.MarketValue = amount
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
		_, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.Setter = id
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
		_, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.DailyRewardAmount = amount
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
		_, err := decimal.NewFromString(*amount)
		if err != nil {
			return err
		}
		h.MaxAmountPerWithdraw = amount
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
