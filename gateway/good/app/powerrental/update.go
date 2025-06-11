package powerrental

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
	apppowerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/app/powerrental"
)

// TODO: check start mode with power rental start mode

type updateHandler struct {
	*checkHandler
}

func (h *Handler) UpdatePowerRental(ctx context.Context) (*npool.AppPowerRental, error) {
	handler := &updateHandler{
		checkHandler: &checkHandler{
			Handler: h,
		},
	}
	if err := handler.checkPowerRental(ctx); err != nil {
		return nil, err
	}

	prHandler, err := apppowerrentalmw.NewHandler(
		ctx,
		apppowerrentalmw.WithID(h.ID, true),
		apppowerrentalmw.WithEntID(h.EntID, true),
		apppowerrentalmw.WithAppGoodID(h.AppGoodID, true),
		apppowerrentalmw.WithPurchasable(h.Purchasable, false),
		apppowerrentalmw.WithEnableProductPage(h.EnableProductPage, false),
		apppowerrentalmw.WithProductPage(h.ProductPage, false),
		apppowerrentalmw.WithOnline(h.Online, false),
		apppowerrentalmw.WithVisible(h.Visible, false),
		apppowerrentalmw.WithName(h.Name, false),
		apppowerrentalmw.WithDisplayIndex(h.DisplayIndex, false),
		apppowerrentalmw.WithBanner(h.Banner, false),
		apppowerrentalmw.WithServiceStartAt(h.ServiceStartAt, false),
		apppowerrentalmw.WithCancelMode(h.CancelMode, false),
		apppowerrentalmw.WithCancelableBeforeStartSeconds(h.CancelableBeforeStartSeconds, false),
		apppowerrentalmw.WithEnableSetCommission(h.EnableSetCommission, false),
		apppowerrentalmw.WithMinOrderAmount(h.MinOrderAmount, false),
		apppowerrentalmw.WithMaxOrderAmount(h.MaxOrderAmount, false),
		apppowerrentalmw.WithMaxUserAmount(h.MaxUserAmount, false),
		apppowerrentalmw.WithMinOrderDurationSeconds(h.MinOrderDurationSeconds, false),
		apppowerrentalmw.WithMaxOrderDurationSeconds(h.MaxOrderDurationSeconds, false),
		apppowerrentalmw.WithUnitPrice(h.UnitPrice, false),
		apppowerrentalmw.WithSaleStartAt(h.SaleStartAt, false),
		apppowerrentalmw.WithSaleEndAt(h.SaleEndAt, false),
		apppowerrentalmw.WithSaleMode(h.SaleMode, false),
		apppowerrentalmw.WithFixedDuration(h.FixedDuration, false),
		apppowerrentalmw.WithPackageWithRequireds(h.PackageWithRequireds, false),
		apppowerrentalmw.WithStartMode(h.StartMode, false),
	)
	if err != nil {
		return nil, err
	}

	if err := prHandler.UpdatePowerRental(ctx); err != nil {
		return nil, err
	}
	return h.GetPowerRental(ctx)
}
