package appfee

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/fee"
	appfeemw "github.com/NpoolPlatform/kunman/middleware/good/app/fee"

	"github.com/google/uuid"
)

func (h *Handler) CreateAppFee(ctx context.Context) (*npool.AppFee, error) {
	if h.AppGoodID == nil {
		h.AppGoodID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := appfeemw.NewHandler(
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

	if err := handler.CreateFee(ctx); err != nil {
		return nil, err
	}
	return h.GetAppFee(ctx)
}
