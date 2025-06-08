package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	powerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/powerrental"
	powerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
)

func (h *Handler) UpdatePowerRental(ctx context.Context) (*npool.PowerRental, error) {
	handler := checkHandler{
		Handler: h,
	}
	if err := handler.checkPowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	if err := powerrentalmwcli.UpdatePowerRental(ctx, &powerrentalmwpb.PowerRentalReq{
		ID:                   h.ID,
		EntID:                h.EntID,
		GoodID:               h.GoodID,
		DeviceTypeID:         h.DeviceTypeID,
		VendorLocationID:     h.VendorLocationID,
		UnitPrice:            h.UnitPrice,
		QuantityUnit:         h.QuantityUnit,
		QuantityUnitAmount:   h.QuantityUnitAmount,
		DeliveryAt:           h.DeliveryAt,
		UnitLockDeposit:      h.UnitLockDeposit,
		DurationDisplayType:  h.DurationDisplayType,
		GoodType:             h.GoodType,
		Name:                 h.Name,
		ServiceStartAt:       h.ServiceStartAt,
		StartMode:            h.StartMode,
		TestOnly:             h.TestOnly,
		BenefitIntervalHours: h.BenefitIntervalHours,
		Purchasable:          h.Purchasable,
		Online:               h.Online,
		StockMode:            h.StockMode,
		State:                h.State,
		Total:                h.Total,
		MiningGoodStocks:     h.MiningGoodStocks,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetPowerRental(ctx)
}
