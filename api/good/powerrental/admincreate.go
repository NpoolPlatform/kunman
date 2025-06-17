package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/kunman/gateway/good/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/powerrental"
)

func (s *Server) AdminCreatePowerRental(ctx context.Context, in *npool.AdminCreatePowerRentalRequest) (*npool.AdminCreatePowerRentalResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,

		powerrental1.WithDeviceTypeID(&in.DeviceTypeID, true),
		powerrental1.WithVendorLocationID(&in.VendorLocationID, true),
		powerrental1.WithUnitPrice(&in.UnitPrice, true),
		powerrental1.WithQuantityUnit(&in.QuantityUnit, true),
		powerrental1.WithQuantityUnitAmount(&in.QuantityUnitAmount, true),
		powerrental1.WithDeliveryAt(in.DeliveryAt, false),
		powerrental1.WithUnitLockDeposit(in.UnitLockDeposit, false),
		powerrental1.WithDurationDisplayType(&in.DurationDisplayType, true),

		powerrental1.WithGoodType(&in.GoodType, true),
		powerrental1.WithName(&in.Name, true),
		powerrental1.WithServiceStartAt(in.ServiceStartAt, false),
		powerrental1.WithStartMode(&in.StartMode, true),
		powerrental1.WithTestOnly(in.TestOnly, false),
		powerrental1.WithBenefitIntervalHours(in.BenefitIntervalHours, false),
		powerrental1.WithPurchasable(in.Purchasable, false),
		powerrental1.WithOnline(in.Online, false),
		powerrental1.WithStockMode(&in.StockMode, true),

		powerrental1.WithTotal(&in.Total, true),
		powerrental1.WithMiningGoodStocks(in.MiningGoodStocks, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreatePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreatePowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreatePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreatePowerRentalResponse{
		Info: info,
	}, nil
}
