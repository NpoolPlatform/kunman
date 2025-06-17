package oneshot

import (
	"context"

	oneshot1 "github.com/NpoolPlatform/kunman/gateway/good/subscription/oneshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/subscription/oneshot"
)

func (s *Server) GetOneShot(ctx context.Context, in *npool.GetOneShotRequest) (*npool.GetOneShotResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.GetOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetOneShot(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.GetOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetOneShotResponse{
		Info: info,
	}, nil
}

func (s *Server) GetOneShots(ctx context.Context, in *npool.GetOneShotsRequest) (*npool.GetOneShotsResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithOffset(in.Offset),
		oneshot1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOneShots",
			"In", in,
			"Error", err,
		)
		return &npool.GetOneShotsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.GetOneShots(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOneShots",
			"In", in,
			"Error", err,
		)
		return &npool.GetOneShotsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetOneShotsResponse{
		Infos: infos,
	}, nil
}
