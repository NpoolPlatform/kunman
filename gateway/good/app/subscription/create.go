package subscription

import (
	"context"

	paypal "github.com/NpoolPlatform/kunman/mal/payment/paypal"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription"
	appsubscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"

	"github.com/google/uuid"
)

// TODO: check start mode with power rental start mode
func (h *Handler) CreateSubscription(ctx context.Context) (*npool.AppSubscription, error) {
	if h.EntID == nil {
		h.EntID = func() *string { s := uuid.NewString(); return &s }()
	}
	if h.AppGoodID == nil {
		h.AppGoodID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithEntID(h.EntID, true),
		appsubscriptionmw.WithAppID(h.AppID, true),
		appsubscriptionmw.WithGoodID(h.GoodID, true),
		appsubscriptionmw.WithAppGoodID(h.AppGoodID, true),
		appsubscriptionmw.WithName(h.Name, true),
		appsubscriptionmw.WithBanner(h.Banner, true),
		// appsubscriptionmw.WithEnableSetCommission(h.EnableSetCommission, true),
		appsubscriptionmw.WithUSDPrice(h.USDPrice, true),
		appsubscriptionmw.WithProductID(h.ProductID, false),
		appsubscriptionmw.WithTrialUnits(h.TrialUnits, false),
		appsubscriptionmw.WithTrialUSDPrice(h.TrialUSDPrice, false),
		appsubscriptionmw.WithPriceFiatID(h.PriceFiatID, false),
		appsubscriptionmw.WithFiatPrice(h.FiatPrice, false),
		appsubscriptionmw.WithTrialFiatPrice(h.TrialFiatPrice, false),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreateSubscription(ctx); err != nil {
		return nil, err
	}

	info, err := h.GetSubscription(ctx)
	if err != nil {
		return nil, err
	}

	cli, err := paypal.NewPaymentClient(
		ctx,
		paypal.WithAppGoodID(*h.AppGoodID),
	)
	if err != nil {
		return nil, err
	}

	resp, err := cli.CreatePlan(ctx)
	if err != nil {
		return nil, err
	}

	handler, err = appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithID(&info.ID, true),
		appsubscriptionmw.WithPlanID(&resp.ID, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.UpdateSubscription(ctx); err != nil {
		return nil, err
	}

	return h.GetSubscription(ctx)
}
