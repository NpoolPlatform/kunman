package powerrental

import (
	"context"

	apppowerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental"
	apppowerrentalmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/app/powerrental"
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
	if err := apppowerrentalmwcli.UpdatePowerRental(ctx, &apppowerrentalmwpb.PowerRentalReq{
		ID:                           h.ID,
		EntID:                        h.EntID,
		AppGoodID:                    h.AppGoodID,
		Purchasable:                  h.Purchasable,
		EnableProductPage:            h.EnableProductPage,
		ProductPage:                  h.ProductPage,
		Online:                       h.Online,
		Visible:                      h.Visible,
		Name:                         h.Name,
		DisplayIndex:                 h.DisplayIndex,
		Banner:                       h.Banner,
		ServiceStartAt:               h.ServiceStartAt,
		CancelMode:                   h.CancelMode,
		CancelableBeforeStartSeconds: h.CancelableBeforeStartSeconds,
		EnableSetCommission:          h.EnableSetCommission,
		MinOrderAmount:               h.MinOrderAmount,
		MaxOrderAmount:               h.MaxOrderAmount,
		MaxUserAmount:                h.MaxUserAmount,
		MinOrderDurationSeconds:      h.MinOrderDurationSeconds,
		MaxOrderDurationSeconds:      h.MaxOrderDurationSeconds,
		UnitPrice:                    h.UnitPrice,
		SaleStartAt:                  h.SaleStartAt,
		SaleEndAt:                    h.SaleEndAt,
		SaleMode:                     h.SaleMode,
		FixedDuration:                h.FixedDuration,
		PackageWithRequireds:         h.PackageWithRequireds,
		StartMode:                    h.StartMode,
	}); err != nil {
		return nil, err
	}
	return h.GetPowerRental(ctx)
}
