package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"
)

func (s *Server) DeleteTopMostGood(ctx context.Context, in *npool.DeleteTopMostGoodRequest) (*npool.DeleteTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(&in.ID, true),
		topmostgood1.WithEntID(&in.EntID, true),
		topmostgood1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTopMostGoodResponse{
		Info: info,
	}, nil
}
