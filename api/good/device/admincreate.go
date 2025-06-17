package devicetype

import (
	"context"

	devicetype1 "github.com/NpoolPlatform/kunman/gateway/good/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/device"
)

func (s *Server) AdminCreateDeviceType(ctx context.Context, in *npool.AdminCreateDeviceTypeRequest) (*npool.AdminCreateDeviceTypeResponse, error) {
	handler, err := devicetype1.NewHandler(
		ctx,
		devicetype1.WithType(&in.DeviceType, true),
		devicetype1.WithManufacturerID(&in.ManufacturerID, true),
		devicetype1.WithPowerConsumption(&in.PowerConsumption, true),
		devicetype1.WithShipmentAt(&in.ShipmentAt, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDeviceType(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateDeviceTypeResponse{
		Info: info,
	}, nil
}
