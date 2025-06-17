package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/poster"
)

func (s *Server) AdminDeletePoster(ctx context.Context, in *npool.AdminDeletePosterRequest) (*npool.AdminDeletePosterResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithID(&in.ID, true),
		poster1.WithEntID(&in.EntID, true),
		poster1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeletePoster",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeletePoster(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeletePoster",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeletePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeletePosterResponse{
		Info: info,
	}, nil
}
