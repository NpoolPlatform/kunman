package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/kunman/gateway/good/device/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/device/poster"
)

func (s *Server) AdminCreatePoster(ctx context.Context, in *npool.AdminCreatePosterRequest) (*npool.AdminCreatePosterResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithDeviceTypeID(&in.DeviceTypeID, true),
		poster1.WithPoster(&in.Poster, true),
		poster1.WithIndex(in.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreatePoster(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreatePosterResponse{
		Info: info,
	}, nil
}
