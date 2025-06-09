package appfee

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"
)

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdateAppFee(ctx context.Context) (*npool.AppFee, error) {
	if err := h.CheckAppGood(ctx); err != nil {
		return nil, err
	}

	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkAppFee(ctx); err != nil {
		return nil, err
	}

	feeHandler, err := appfeemw.NewHandler(
		ctx,
		appfeemw.WithAppID(h.AppID, true),
		appfeemw.WithGoodID(h.GoodID, true),
		appfeemw.WithAppGoodID(h.AppGoodID, true),
		appfeemw.WithProductPage(h.ProductPage, true),
		appfeemw.WithName(h.Name, true),
		appfeemw.WithBanner(h.Banner, true),
		appfeemw.WithUnitValue(h.UnitValue, true),
		appfeemw.WithMinOrderDurationSeconds(h.MinOrderDurationSeconds, true),
		appfeemw.WithCancelMode(h.CancelMode, true),
	)
	if err != nil {
		return nil, err
	}

	if err := feeHandler.UpdateFee(ctx); err != nil {
		return nil, err
	}
	return h.GetAppFee(ctx)
}
