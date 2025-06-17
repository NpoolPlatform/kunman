package oneshot

import (
	"context"

	oneshot1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription/oneshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"
)

func (s *Server) AdminGetAppOneShots(ctx context.Context, in *npool.AdminGetAppOneShotsRequest) (*npool.AdminGetAppOneShotsResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithAppID(&in.TargetAppID, true),
		oneshot1.WithOffset(in.Offset),
		oneshot1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppOneShots",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppOneShotsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetOneShots(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetAppOneShots",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetAppOneShotsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminGetAppOneShotsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
