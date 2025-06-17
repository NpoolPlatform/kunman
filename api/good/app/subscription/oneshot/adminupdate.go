package oneshot

import (
	"context"

	oneshot1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription/oneshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"
)

func (s *Server) AdminUpdateAppOneShot(ctx context.Context, in *npool.AdminUpdateAppOneShotRequest) (*npool.AdminUpdateAppOneShotResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithID(&in.ID, true),
		oneshot1.WithEntID(&in.EntID, true),
		oneshot1.WithAppID(&in.TargetAppID, true),
		oneshot1.WithAppGoodID(&in.AppGoodID, true),

		oneshot1.WithName(in.Name, false),
		oneshot1.WithBanner(in.Banner, false),

		oneshot1.WithEnableSetCommission(in.EnableSetCommission, false),
		oneshot1.WithUSDPrice(in.USDPrice, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAppOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAppOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateOneShot(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateAppOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateAppOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateAppOneShotResponse{
		Info: info,
	}, nil
}
