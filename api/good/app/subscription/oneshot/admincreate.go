package oneshot

import (
	"context"

	oneshot1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription/oneshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"
)

func (s *Server) AdminCreateAppOneShot(ctx context.Context, in *npool.AdminCreateAppOneShotRequest) (*npool.AdminCreateAppOneShotResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithAppID(&in.TargetAppID, true),
		oneshot1.WithGoodID(&in.GoodID, true),

		oneshot1.WithName(&in.Name, true),
		oneshot1.WithBanner(in.Banner, false),

		oneshot1.WithEnableSetCommission(in.EnableSetCommission, false),
		oneshot1.WithUSDPrice(&in.USDPrice, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateOneShot(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateAppOneShotResponse{
		Info: info,
	}, nil
}
