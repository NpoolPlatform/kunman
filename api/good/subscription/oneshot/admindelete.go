package oneshot

import (
	"context"

	oneshot1 "github.com/NpoolPlatform/kunman/gateway/good/subscription/oneshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/subscription/oneshot"
)

func (s *Server) AdminDeleteOneShot(ctx context.Context, in *npool.AdminDeleteOneShotRequest) (*npool.AdminDeleteOneShotResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithID(&in.ID, true),
		oneshot1.WithEntID(&in.EntID, true),
		oneshot1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteOneShot(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteOneShotResponse{
		Info: info,
	}, nil
}
