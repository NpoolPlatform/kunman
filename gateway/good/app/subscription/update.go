package subscription

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription"
	appsubscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"
)

// TODO: check start mode with power rental start mode

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateSubscription(ctx context.Context) (*npool.AppSubscription, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSubscription(ctx); err != nil {
		return nil, err
	}

	prHandler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithID(h.ID, true),
		appsubscriptionmw.WithEntID(h.EntID, true),
		appsubscriptionmw.WithAppGoodID(h.AppGoodID, true),
		appsubscriptionmw.WithName(h.Name, false),
		appsubscriptionmw.WithBanner(h.Banner, false),
		// appsubscriptionmw.WithEnableSetCommission(h.EnableSetCommission, false),
	)
	if err != nil {
		return nil, err
	}

	if err := prHandler.UpdateSubscription(ctx); err != nil {
		return nil, err
	}
	return h.GetSubscription(ctx)
}
