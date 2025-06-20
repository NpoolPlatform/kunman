package appconfig

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"

	config1 "github.com/NpoolPlatform/kunman/gateway/order/app/config"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AdminGetAppConfigs(ctx context.Context, in *npool.AdminGetAppConfigsRequest) (*npool.AdminGetAppConfigsResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithAppID(in.TargetAppID, false),
		config1.WithOffset(in.Offset),
		config1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppConfigs",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppConfigsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetAppConfigs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppConfigs",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppConfigsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminGetAppConfigsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
