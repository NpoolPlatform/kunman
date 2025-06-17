package description

import (
	"context"

	description1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/description"
)

func (s *Server) DeleteDescription(ctx context.Context, in *npool.DeleteDescriptionRequest) (*npool.DeleteDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(&in.ID, true),
		description1.WithEntID(&in.EntID, true),
		description1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDescription",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDescription",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDescriptionResponse{
		Info: info,
	}, nil
}
