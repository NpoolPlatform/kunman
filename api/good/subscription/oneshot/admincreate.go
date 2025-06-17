package oneshot

import (
	"context"

	oneshot1 "github.com/NpoolPlatform/kunman/gateway/good/subscription/oneshot"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/subscription/oneshot"
)

func (s *Server) AdminCreateOneShot(ctx context.Context, in *npool.AdminCreateOneShotRequest) (*npool.AdminCreateOneShotResponse, error) {
	handler, err := oneshot1.NewHandler(
		ctx,
		oneshot1.WithName(&in.Name, true),
		oneshot1.WithQuota(&in.Quota, false),
		oneshot1.WithUSDPrice(&in.USDPrice, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateOneShot(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateOneShot",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateOneShotResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateOneShotResponse{
		Info: info,
	}, nil
}
