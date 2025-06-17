package oneshot

import (
	"context"

	oneshot1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription/oneshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"
)

func (s *Server) GetAppOneShot(ctx context.Context, in *npool.GetAppOneShotRequest) (*npool.GetAppOneShotResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetOneShot(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppOneShotResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppOneShots(ctx context.Context, in *npool.GetAppOneShotsRequest) (*npool.GetAppOneShotsResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithAppID(&in.AppID, true),
		oneshot1.WithOffset(in.Offset),
		oneshot1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppOneShots",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppOneShotsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetOneShots(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppOneShots",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppOneShotsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppOneShotsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
