package devicetype

import (
	"context"

	devicetype1 "github.com/NpoolPlatform/kunman/gateway/good/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/device"
)

func (s *Server) GetDeviceTypes(ctx context.Context, in *npool.GetDeviceTypesRequest) (*npool.GetDeviceTypesResponse, error) {
	handler, err := devicetype1.NewHandler(
		ctx,
		devicetype1.WithOffset(in.Offset),
		devicetype1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceTypes",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceTypesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDeviceTypes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceTypes",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceTypesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDeviceTypesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
