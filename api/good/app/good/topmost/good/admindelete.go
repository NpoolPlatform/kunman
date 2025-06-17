package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
)

func (s *Server) AdminDeleteTopMostGood(ctx context.Context, in *npool.AdminDeleteTopMostGoodRequest) (*npool.AdminDeleteTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(&in.ID, true),
		topmostgood1.WithEntID(&in.EntID, true),
		topmostgood1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteTopMostGoodResponse{
		Info: info,
	}, nil
}
