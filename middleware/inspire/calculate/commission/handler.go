package commission

import (
	"context"

	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	achievementusermwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/achievement/user"
	appcommissionconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/commission/config"
	appconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/config"
	appgoodcommissionconfigmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/good/commission/config"
	commissionmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/commission"
	registrationmwpb "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/invitation/registration"

	"github.com/shopspring/decimal"
)

type Handler struct {
	SettleType               types.SettleType
	SettleAmountType         types.SettleAmountType
	Inviters                 []*registrationmwpb.Registration
	AppConfig                *appconfigmwpb.AppConfig
	Commissions              []*commissionmwpb.Commission
	AppCommissionConfigs     []*appcommissionconfigmwpb.AppCommissionConfig
	AppGoodCommissionConfigs []*appgoodcommissionconfigmwpb.AppGoodCommissionConfig
	AchievementUsers         map[string]*achievementusermwpb.AchievementUser
	PaymentAmount            decimal.Decimal
	PaymentAmountUSD         decimal.Decimal
	PaymentCoinUSDCurrency   decimal.Decimal
	GoodValueUSD             decimal.Decimal
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

func WithSettleType(settleType types.SettleType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SettleType = settleType
		return nil
	}
}

func WithSettleAmountType(settleAmount types.SettleAmountType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SettleAmountType = settleAmount
		return nil
	}
}

func WithInviters(inviters []*registrationmwpb.Registration) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Inviters = inviters
		return nil
	}
}

func WithAppConfig(value *appconfigmwpb.AppConfig) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppConfig = value
		return nil
	}
}

func WithCommissions(commissions []*commissionmwpb.Commission) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Commissions = commissions
		return nil
	}
}

func WithAppCommissionConfigs(commissions []*appcommissionconfigmwpb.AppCommissionConfig) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppCommissionConfigs = commissions
		return nil
	}
}

func WithAppGoodCommissionConfigs(commissions []*appgoodcommissionconfigmwpb.AppGoodCommissionConfig) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodCommissionConfigs = commissions
		return nil
	}
}

func WithAchievementUsers(values map[string]*achievementusermwpb.AchievementUser) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AchievementUsers = values
		return nil
	}
}

func WithPaymentAmount(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.PaymentAmount = amount
		return nil
	}
}

func WithPaymentAmountUSD(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.PaymentAmountUSD = amount
		return nil
	}
}

func WithPaymentCoinUSDCurrency(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.PaymentCoinUSDCurrency = amount
		return nil
	}
}

func WithGoodValueUSD(value string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		amount, err := decimal.NewFromString(value)
		if err != nil {
			return err
		}
		h.GoodValueUSD = amount
		return nil
	}
}
