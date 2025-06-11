package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/kunman/framework/wlog"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/powerrental"
	powerrentalmw "github.com/NpoolPlatform/kunman/middleware/good/powerrental"

	"github.com/google/uuid"
)

func (h *Handler) CreatePowerRental(ctx context.Context) (*npool.PowerRental, error) {
	if h.GoodID == nil {
		h.GoodID = func() *string { s := uuid.NewString(); return &s }()
	}

	handler, err := powerrentalmw.NewHandler(
		ctx,
		powerrentalmw.WithGoodID(h.GoodID, true),
		powerrentalmw.WithDeviceTypeID(h.DeviceTypeID, true),
		powerrentalmw.WithVendorLocationID(h.VendorLocationID, true),
		powerrentalmw.WithUnitPrice(h.UnitPrice, true),
		powerrentalmw.WithQuantityUnit(h.QuantityUnit, true),
		powerrentalmw.WithQuantityUnitAmount(h.QuantityUnitAmount, true),
		powerrentalmw.WithDeliveryAt(h.DeliveryAt, true),
		powerrentalmw.WithUnitLockDeposit(h.UnitLockDeposit, true),
		powerrentalmw.WithDurationDisplayType(h.DurationDisplayType, true),
		powerrentalmw.WithGoodType(h.GoodType, true),
		powerrentalmw.WithName(h.Name, true),
		powerrentalmw.WithServiceStartAt(h.ServiceStartAt, true),
		powerrentalmw.WithStartMode(h.StartMode, true),
		powerrentalmw.WithTestOnly(h.TestOnly, true),
		powerrentalmw.WithBenefitIntervalHours(h.BenefitIntervalHours, true),
		powerrentalmw.WithPurchasable(h.Purchasable, true),
		powerrentalmw.WithOnline(h.Online, true),
		powerrentalmw.WithStockMode(h.StockMode, true),
		powerrentalmw.WithTotal(h.Total, true),
		powerrentalmw.WithStocks(h.MiningGoodStocks, true),
	)
	if err != nil {
		return nil, err
	}

	if err := handler.CreatePowerRental(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.GetPowerRental(ctx)
}
