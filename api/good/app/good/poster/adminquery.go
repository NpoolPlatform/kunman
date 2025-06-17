package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/poster"
)

func (s *Server) AdminGetPosters(ctx context.Context, in *npool.AdminGetPostersRequest) (*npool.AdminGetPostersResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithAppID(&in.TargetAppID, true),
		poster1.WithOffset(in.Offset),
		poster1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetPosters",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetPostersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPosters(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetPosters",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetPostersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetPostersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
