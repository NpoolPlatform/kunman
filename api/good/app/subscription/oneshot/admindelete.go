package oneshot

import (
	"context"

	oneshot1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription/oneshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"
)

func (s *Server) AdminDeleteAppOneShot(ctx context.Context, in *npool.AdminDeleteAppOneShotRequest) (*npool.AdminDeleteAppOneShotResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithID(&in.ID, true),
		oneshot1.WithEntID(&in.EntID, true),
		oneshot1.WithAppID(&in.TargetAppID, true),
		oneshot1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteOneShot(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteAppOneShotResponse{
		Info: info,
	}, nil
}
