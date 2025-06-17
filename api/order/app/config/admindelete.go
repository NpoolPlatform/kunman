package appconfig

import (
	"context"

	config1 "github.com/NpoolPlatform/kunman/gateway/order/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/app/config"
)

func (s *Server) AdminDeleteAppConfig(ctx context.Context, in *npool.AdminDeleteAppConfigRequest) (*npool.AdminDeleteAppConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(&in.ID, true),
		config1.WithEntID(&in.EntID, true),
		config1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminDeleteAppConfigResponse{
		Info: info,
	}, nil
}
