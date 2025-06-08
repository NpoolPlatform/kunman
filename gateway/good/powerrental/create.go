package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	powerrentalmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/powerrental"
	powerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"

	"github.com/google/uuid"
)

func (h *Handler) CreatePowerRental(ctx context.Context) (*npool.PowerRental, error) {
	if h.GoodID == nil {
		h.GoodID = func() *string { s := uuid.NewString(); return &s }()
	}
	if err := powerrentalmwcli.CreatePowerRental(ctx, &powerrentalmwpb.PowerRentalReq{
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
		Total:                h.Total,
		MiningGoodStocks:     h.MiningGoodStocks,
	}); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetPowerRental(ctx)
}
