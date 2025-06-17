package description

import (
	"context"

	description1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/description"
)

func (s *Server) GetDescriptions(ctx context.Context, in *npool.GetDescriptionsRequest) (*npool.GetDescriptionsResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithAppID(&in.AppID, true),
		description1.WithAppGoodID(in.AppGoodID, false),
		description1.WithOffset(in.Offset),
		description1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDescriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetDescriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDescriptions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDescriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetDescriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDescriptionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
