package paypal

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	fiatmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/fiat"
	appsubscriptionmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/subscription"
	fiatmw "github.com/NpoolPlatform/kunman/middleware/chain/fiat"
	appsubscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"

	"github.com/shopspring/decimal"
)

type appGoodHandler struct {
	appSubscription *appsubscriptionmwpb.Subscription

	fiat *fiatmwpb.Fiat
}

func (cli *PaymentClient) GetAppGood(ctx context.Context) error {
	handler := &appGoodHandler{}

	subscriptionHandler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithAppGoodID(cli.AppGoodID, true),
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	appSubscription, err := subscriptionHandler.GetSubscription(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}

	if appSubscription.PriceFiatID != "" {
		fiatHandler, err := fiatmw.NewHandler(
			ctx,
			fiatmw.WithEntID(&appSubscription.PriceFiatID, false),
		)
		if err != nil {
			return wlog.WrapError(err)
		}

		fiat, err := fiatHandler.GetFiat(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.fiat = fiat
	}
	handler.appSubscription = appSubscription
	cli.appGoodHandler = handler

	return nil
}

func (h *appGoodHandler) ProductID() string {
	return h.appSubscription.ProductID
}

func (h *appGoodHandler) Name() string {
	return h.appSubscription.AppGoodName
}

func (h *appGoodHandler) Description() string {
	return h.appSubscription.AppGoodName
}

func (h *appGoodHandler) PriceUnit() string {
	if h.fiat != nil {
		return h.fiat.Unit
	}
	return "USD"
}

func (h *appGoodHandler) TrialUnits() uint32 {
	return h.appSubscription.TrialUnits
}

func (h *appGoodHandler) TrialPrice() string {
	price, err := decimal.NewFromString(h.appSubscription.TrialFiatPrice)
	if err == nil {
		return price.Round(2).String()
	}
	return h.appSubscription.TrialUSDPrice
}

func (h *appGoodHandler) Price() string {
	price, err := decimal.NewFromString(h.appSubscription.FiatPrice)
	if err == nil {
		return price.Round(2).String()
	}
	return h.appSubscription.USDPrice
}

func (h *appGoodHandler) DurationUnits() uint32 {
	return h.appSubscription.DurationUnits
}

func (h *appGoodHandler) DurationUnit() (string, error) {
	switch h.appSubscription.DurationDisplayType {
	case types.GoodDurationType_GoodDurationByDay:
		return "DAY", nil
	case types.GoodDurationType_GoodDurationByWeek:
		return "WEEK", nil
	case types.GoodDurationType_GoodDurationByMonth:
		return "MONTH", nil
	case types.GoodDurationType_GoodDurationByYear:
		return "YEAR", nil
	}
	return "", wlog.Errorf("invalid duration")
}
