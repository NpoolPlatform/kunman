package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/kunman/gateway/good/device/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/device/poster"
)

func (s *Server) GetPosters(ctx context.Context, in *npool.GetPostersRequest) (*npool.GetPostersResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithDeviceTypeID(in.DeviceTypeID, false),
		poster1.WithOffset(in.Offset),
		poster1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPosters",
			"In", in,
			"Error", err,
		)
		return &npool.GetPostersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPosters(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPosters",
			"In", in,
			"Error", err,
		)
		return &npool.GetPostersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPostersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
