package subscription

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription"
	appsubscriptionmw "github.com/NpoolPlatform/kunman/middleware/good/app/subscription"
)

type deleteHandler struct {
	*checkHandler
}

func (h *Handler) DeleteSubscription(ctx context.Context) (*npool.AppSubscription, error) {
	handler := &deleteHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkSubscription(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	info, err := h.GetSubscription(ctx)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if info == nil {
		return nil, wlog.Errorf("invalid subscription")
	}

	prHandler, err := appsubscriptionmw.NewHandler(
		ctx,
		appsubscriptionmw.WithID(h.ID, true),
		appsubscriptionmw.WithEntID(h.EntID, true),
		appsubscriptionmw.WithAppGoodID(h.AppGoodID, true),
	)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if err := prHandler.DeleteSubscription(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return info, nil
}
