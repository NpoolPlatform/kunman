package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/powerrental"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"
)

func (h *Handler) UpdatePowerRental(ctx context.Context) (*npool.PowerRental, error) {
	handler := checkHandler{
		Handler: h,
	}
	if err := handler.checkPowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}

	prHandler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithID(h.ID, true),
		powerrentalmw.WithEntID(h.EntID, true),
		powerrentalmw.WithGoodID(h.GoodID, true),
		powerrentalmw.WithDeviceTypeID(h.DeviceTypeID, false),
		powerrentalmw.WithVendorLocationID(h.VendorLocationID, false),
		powerrentalmw.WithUnitPrice(h.UnitPrice, false),
		powerrentalmw.WithQuantityUnit(h.QuantityUnit, false),
		powerrentalmw.WithQuantityUnitAmount(h.QuantityUnitAmount, false),
		powerrentalmw.WithDeliveryAt(h.DeliveryAt, false),
		powerrentalmw.WithUnitLockDeposit(h.UnitLockDeposit, false),
		powerrentalmw.WithDurationDisplayType(h.DurationDisplayType, false),
		powerrentalmw.WithGoodType(h.GoodType, false),
		powerrentalmw.WithName(h.Name, false),
		powerrentalmw.WithServiceStartAt(h.ServiceStartAt, false),
		powerrentalmw.WithStartMode(h.StartMode, false),
		powerrentalmw.WithTestOnly(h.TestOnly, false),
		powerrentalmw.WithBenefitIntervalHours(h.BenefitIntervalHours, false),
		powerrentalmw.WithPurchasable(h.Purchasable, false),
		powerrentalmw.WithOnline(h.Online, false),
		powerrentalmw.WithStockMode(h.StockMode, false),
		powerrentalmw.WithTotal(h.Total, false),
		powerrentalmw.WithStocks(h.MiningGoodStocks, false),
	)
	if err != nil {
		return nil, err
	}

	if err := prHandler.UpdatePowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetPowerRental(ctx)
}
