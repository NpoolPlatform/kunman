package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/kunman/gateway/good/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"
)

func (s *Server) AdminDeleteTopMost(ctx context.Context, in *npool.AdminDeleteTopMostRequest) (*npool.AdminDeleteTopMostResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithID(&in.ID, true),
		topmost1.WithEntID(&in.EntID, true),
		topmost1.WithAppID(&in.TargetAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteTopMostResponse{
		Info: info,
	}, nil
}
