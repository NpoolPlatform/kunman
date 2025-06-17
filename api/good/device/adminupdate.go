package devicetype

import (
	"context"

	devicetype1 "github.com/NpoolPlatform/kunman/gateway/good/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/device"
)

func (s *Server) AdminUpdateDeviceType(ctx context.Context, in *npool.AdminUpdateDeviceTypeRequest) (*npool.AdminUpdateDeviceTypeResponse, error) {
	handler, err := devicetype1.NewHandler(
		ctx,
		devicetype1.WithID(&in.ID, true),
		devicetype1.WithEntID(&in.EntID, true),
		devicetype1.WithType(in.DeviceType, false),
		devicetype1.WithManufacturerID(in.ManufacturerID, false),
		devicetype1.WithPowerConsumption(in.PowerConsumption, false),
		devicetype1.WithShipmentAt(in.ShipmentAt, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDeviceType(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateDeviceTypeResponse{
		Info: info,
	}, nil
}
