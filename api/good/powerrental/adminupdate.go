package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/kunman/gateway/good/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/powerrental"
)

func (s *Server) AdminUpdatePowerRental(ctx context.Context, in *npool.AdminUpdatePowerRentalRequest) (*npool.AdminUpdatePowerRentalResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(&in.ID, true),
		powerrental1.WithEntID(&in.EntID, true),
		powerrental1.WithGoodID(&in.GoodID, true),

		powerrental1.WithDeviceTypeID(in.DeviceTypeID, false),
		powerrental1.WithVendorLocationID(in.VendorLocationID, false),
		powerrental1.WithUnitPrice(in.UnitPrice, false),
		powerrental1.WithQuantityUnit(in.QuantityUnit, false),
		powerrental1.WithQuantityUnitAmount(in.QuantityUnitAmount, false),
		powerrental1.WithDeliveryAt(in.DeliveryAt, false),
		powerrental1.WithUnitLockDeposit(in.UnitLockDeposit, false),
		powerrental1.WithDurationDisplayType(in.DurationDisplayType, false),

		powerrental1.WithGoodType(in.GoodType, false),
		powerrental1.WithName(in.Name, false),
		powerrental1.WithServiceStartAt(in.ServiceStartAt, false),
		powerrental1.WithStartMode(in.StartMode, false),
		powerrental1.WithTestOnly(in.TestOnly, false),
		powerrental1.WithBenefitIntervalHours(in.BenefitIntervalHours, false),
		powerrental1.WithPurchasable(in.Purchasable, false),
		powerrental1.WithOnline(in.Online, false),
		powerrental1.WithStockMode(in.StockMode, false),
		powerrental1.WithState(in.State, false),

		powerrental1.WithTotal(in.Total, false),
		powerrental1.WithMiningGoodStocks(in.MiningGoodStocks, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdatePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdatePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdatePowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdatePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdatePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdatePowerRentalResponse{
		Info: info,
	}, nil
}
