package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/poster"
)

func (s *Server) DeletePoster(ctx context.Context, in *npool.DeletePosterRequest) (*npool.DeletePosterResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithID(&in.ID, true),
		poster1.WithEntID(&in.EntID, true),
		poster1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePoster",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeletePoster(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePoster",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeletePosterResponse{
		Info: info,
	}, nil
}
