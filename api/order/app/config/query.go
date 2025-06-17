package appconfig

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/app/config"
	config1 "github.com/NpoolPlatform/kunman/gateway/order/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAppConfig(ctx context.Context, in *npool.GetAppConfigRequest) (*npool.GetAppConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppConfigResponse{
		Info: info,
	}, nil
}
