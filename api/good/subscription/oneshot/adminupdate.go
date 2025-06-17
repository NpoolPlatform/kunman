package oneshot

import (
	"context"

	oneshot1 "github.com/NpoolPlatform/kunman/gateway/good/subscription/oneshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/subscription/oneshot"
)

func (s *Server) AdminUpdateOneShot(ctx context.Context, in *npool.AdminUpdateOneShotRequest) (*npool.AdminUpdateOneShotResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithID(&in.ID, true),
		oneshot1.WithEntID(&in.EntID, true),
		oneshot1.WithGoodID(&in.GoodID, true),
		oneshot1.WithName(in.Name, false),
		oneshot1.WithQuota(in.Quota, false),
		oneshot1.WithUSDPrice(in.USDPrice, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateOneShot(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateOneShotResponse{
		Info: info,
	}, nil
}
